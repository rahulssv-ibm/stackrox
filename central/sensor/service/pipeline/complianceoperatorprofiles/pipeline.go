package complianceoperatorprofiles

import (
	"context"

	"github.com/pkg/errors"
	"github.com/stackrox/rox/central/complianceoperator/manager"
	"github.com/stackrox/rox/central/complianceoperator/profiles/datastore"
	countMetrics "github.com/stackrox/rox/central/metrics"
	"github.com/stackrox/rox/central/sensor/service/common"
	"github.com/stackrox/rox/central/sensor/service/pipeline"
	"github.com/stackrox/rox/central/sensor/service/pipeline/reconciliation"
	"github.com/stackrox/rox/generated/internalapi/central"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/centralsensor"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
	"github.com/stackrox/rox/pkg/set"
)

var (
	_ pipeline.Fragment = (*pipelineImpl)(nil)
)

// GetPipeline returns an instantiation of this particular pipeline
func GetPipeline() pipeline.Fragment {
	return NewPipeline(datastore.Singleton(), manager.Singleton())
}

// NewPipeline returns a new instance of Pipeline.
func NewPipeline(datastore datastore.DataStore, manager manager.Manager) pipeline.Fragment {
	return &pipelineImpl{
		datastore: datastore,
		manager:   manager,
	}
}

type pipelineImpl struct {
	datastore datastore.DataStore
	manager   manager.Manager
}

func (s *pipelineImpl) Capabilities() []centralsensor.CentralCapability {
	return nil
}

func (s *pipelineImpl) Reconcile(ctx context.Context, clusterID string, storeMap *reconciliation.StoreMap) error {
	existingIDs := set.NewStringSet()
	walkFn := func() error {
		existingIDs.Clear()
		return s.datastore.Walk(ctx, func(profile *storage.ComplianceOperatorProfile) error {
			if profile.GetClusterId() == clusterID {
				existingIDs.Add(profile.GetId())
			}
			return nil
		})
	}
	if err := pgutils.RetryIfPostgres(walkFn); err != nil {
		return err
	}
	store := storeMap.Get((*central.SensorEvent_ComplianceOperatorProfile)(nil))
	return reconciliation.Perform(store, existingIDs, "complianceoperatorprofiles", func(id string) error {
		return s.datastore.Delete(ctx, id)
	})
}

func (s *pipelineImpl) Match(msg *central.MsgFromSensor) bool {
	if features.ComplianceEnhancements.Enabled() {
		return msg.GetEvent().GetComplianceOperatorProfileV2() != nil || msg.GetEvent().GetComplianceOperatorProfile() != nil
	}
	return msg.GetEvent().GetComplianceOperatorProfile() != nil
}

// Run runs the pipeline template on the input and returns the output.
func (s *pipelineImpl) Run(ctx context.Context, clusterID string, msg *central.MsgFromSensor, _ common.MessageInjector) error {
	defer countMetrics.IncrementResourceProcessedCounter(pipeline.ActionToOperation(msg.GetEvent().GetAction()), metrics.ComplianceOperatorProfile)

	event := msg.GetEvent()
	// If a sensor sends in a v1 compliance message we will still process it the v1 way in the event
	// a sensor is not updated or does not have the flag on.
	switch event.Resource.(type) {
	case *central.SensorEvent_ComplianceOperatorProfile:
		return s.processComplianceProfile(ctx, event, clusterID)
	case *central.SensorEvent_ComplianceOperatorProfileV2:
		if !features.ComplianceEnhancements.Enabled() {
			return errors.New("Next gen compliance is disabled.  Message unexpected.")
		}
		return s.processComplianceProfileV2(ctx, event, clusterID)
	}

	return errors.Errorf("unexpected message %t.", event.Resource)
}

func (s *pipelineImpl) OnFinish(_ string) {}

func (s *pipelineImpl) processComplianceProfile(_ context.Context, event *central.SensorEvent, clusterID string) error {
	profile := event.GetComplianceOperatorProfile()
	profile.ClusterId = clusterID

	switch event.GetAction() {
	case central.ResourceAction_REMOVE_RESOURCE:
		return s.manager.DeleteProfile(profile)
	default:
		return s.manager.AddProfile(profile)
	}
}

func (s *pipelineImpl) processComplianceProfileV2(ctx context.Context, event *central.SensorEvent, clusterID string) error {
	if !features.ComplianceEnhancements.Enabled() {
		return errors.New("Next gen compliance is disabled.  Message unexpected.")
	}
}
