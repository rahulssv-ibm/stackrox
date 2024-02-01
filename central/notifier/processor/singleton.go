package processor

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/stackrox/rox/central/integrationhealth/reporter"
	"github.com/stackrox/rox/central/notifier/datastore"
	encConfigDatastore "github.com/stackrox/rox/central/notifier/encconfig/datastore"
	notifierUtils "github.com/stackrox/rox/central/notifiers/utils"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/declarativeconfig"
	"github.com/stackrox/rox/pkg/env"
	"github.com/stackrox/rox/pkg/errorhelpers"
	"github.com/stackrox/rox/pkg/notifier"
	"github.com/stackrox/rox/pkg/notifiers"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/sac/resources"
	"github.com/stackrox/rox/pkg/sync"
	"github.com/stackrox/rox/pkg/utils"
)

const (
	// When we fail to notify on an alert, retry every hour for 4 hours, and only retry up to 100 alerts
	retryAlertsEvery = 5 * time.Minute
	retryAlertsFor   = 1 * time.Hour
)

var (
	once sync.Once

	ns   notifier.Set
	loop notifier.Loop
	pr   notifier.Processor

	// Create a context that can access and update notifiers and namespaces for initialization.
	notifierCtx = declarativeconfig.WithModifyDeclarativeOrImperative(
		sac.WithGlobalAccessScopeChecker(context.Background(),
			sac.AllowFixedScopes(
				sac.AccessModeScopeKeys(storage.Access_READ_ACCESS, storage.Access_READ_WRITE_ACCESS),
				sac.ResourceScopeKeys(resources.Integration, resources.Namespace))))
)

func initialize() {
	// Keep track of the notifiers in use.
	ns = notifier.NewNotifierSet(retryAlertsFor)

	// When alerts are generated, we will want to notify.
	pr = New(ns, reporter.Singleton())

	notifierDatastore := datastore.Singleton()
	protoNotifiers, err := notifierDatastore.GetNotifiers(notifierCtx)
	if err != nil {
		log.Panicf("unable to fetch notifiers: %v", err)
	}

	if env.EncNotifierCreds.BooleanSetting() {
		currentKey, activeIndex, err := notifierUtils.GetActiveNotifierEncryptionKey()
		if err != nil {
			utils.Should(errors.Wrap(err, "Error reading encryption key, notifiers will be unable to send notifications"))
		}

		encConfigDataStore := encConfigDatastore.Singleton()
		encConfig, err := encConfigDataStore.GetConfig()
		if err != nil {
			utils.Should(errors.Wrap(err, "Error getting notifier encryption config"))
		}
		if encConfig == nil {
			// This will be true when secured notifiers feature is enabled for the first time as the config will not exist yet in the db
			encConfig = &storage.NotifierEncConfig{ActiveKeyIndex: 0}
			err = encConfigDataStore.UpsertConfig(encConfig)
			if err != nil {
				utils.Should(errors.Wrap(err, "Error initializing notifier encryption config %d"))
			}
		}
		storedKeyIndex := int(encConfig.GetActiveKeyIndex())

		var newKeyAdded bool
		var prevKey string
		if activeIndex != storedKeyIndex {
			newKeyAdded = true
			prevKey, err = notifierUtils.GetNotifierEncryptionKeyAtIndex(storedKeyIndex)
			if err != nil {
				utils.Should(errors.Wrap(err, "Error reading previous encryption key, notifiers will be unable to send notifications"))
			}
		}
		if err := secureNotifiers(notifierDatastore, protoNotifiers, currentKey, prevKey, newKeyAdded); err != nil {
			if newKeyAdded {
				// If we did a rekey, then update the stored key index
				encConfig = &storage.NotifierEncConfig{ActiveKeyIndex: int32(activeIndex)}
				err = encConfigDataStore.UpsertConfig(encConfig)
				if err != nil {
					utils.Should(errors.Wrapf(err, "Error updating notifier encryption config's stored key index to %d", activeIndex))
				}
			}
		} else {
			utils.Should(err)
		}
	}

	// Create actionable notifiers from the loaded protos.
	for _, protoNotifier := range protoNotifiers {
		notifier, err := notifiers.CreateNotifier(protoNotifier)
		if err != nil {
			utils.Should(errors.Wrapf(err, "error creating notifier with %v (%v) and type %v", protoNotifier.GetId(), protoNotifier.GetName(), protoNotifier.GetType()))
			continue
		}
		pr.UpdateNotifier(notifierCtx, notifier)
	}

	// When alerts have failed, we will want to retry the notifications.
	loop = notifier.NewLoop(ns, retryAlertsEvery)
	loop.Start(notifierCtx)
}

// secureNotifiers encrypts/rekeys notifiers
func secureNotifiers(notifierDatastore datastore.DataStore, protoNotifiers []*storage.Notifier,
	currentKey string, prevKey string, needsRekey bool) error {
	var notifiersToUpsert []*storage.Notifier

	errorList := errorhelpers.NewErrorList("Securing notifiers: ")
	for _, protoNotifier := range protoNotifiers {
		secured, err := notifierUtils.IsNotifierSecured(protoNotifier)
		if err != nil {
			errorList.AddError(errors.Wrapf(err, "Error checking if the notifier %s is secured, notifications to this notifier will fail",
				protoNotifier.GetId()))
			continue
		}
		if !secured {
			// If notifier is not secured, then we just need to secure it using the latest key and continue. No need to rekey.
			err := notifierUtils.SecureNotifier(protoNotifier, currentKey)
			if err != nil {
				// Don't send out error from crypto lib
				errorList.AddError(fmt.Errorf("error securing notifier %s, notifications to this notifier will fail", protoNotifier.GetId()))
				continue
			}
			notifiersToUpsert = append(notifiersToUpsert, protoNotifier)
			continue
		}

		if needsRekey {
			// If a notifier is already secured and needsRekey = true i.e (storedKeyIndex != activeKeyIndex)
			// then we need to decrypt using old key and encrypt using the active key
			err = notifierUtils.RekeyNotifier(protoNotifier, prevKey, currentKey)
			if err != nil {
				errorList.AddError(fmt.Errorf("error rekeying notifier %s, notifications to this notifier will fail", protoNotifier.GetId()))
				continue
			}
			notifiersToUpsert = append(notifiersToUpsert, protoNotifier)
		}
	}

	err := notifierDatastore.UpsertManyNotifiers(notifierCtx, notifiersToUpsert)
	if err != nil {
		errorList.AddError(errors.Wrap(err, "Error upserting secured notifiers, several notifiers will be unable to send notifications"))
	}
	return errorList.ToError()
}

// Singleton provides the interface for processing notifications.
func Singleton() notifier.Processor {
	once.Do(initialize)
	return pr
}
