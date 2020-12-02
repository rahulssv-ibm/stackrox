package datastore

import (
	"context"

	"github.com/stackrox/rox/central/integrationhealth/store/rocksdb"
	"github.com/stackrox/rox/central/role/resources"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	imageSAC    = sac.ForResource(resources.ImageIntegration)
	notifierSAC = sac.ForResource(resources.Notifier)
	backupSAC   = sac.ForResource(resources.BackupPlugins)
)

type datastoreImpl struct {
	store rocksdb.Store
}

func (ds *datastoreImpl) GetRegistriesAndScanners(ctx context.Context) ([]*storage.IntegrationHealth, error) {
	if ok, err := imageSAC.ReadAllowed(ctx); err != nil {
		return nil,
			status.Errorf(codes.Internal, "Failed to retrieve health for registries and scanners: %v", err)
	} else if !ok {
		return nil, nil
	}
	return ds.getIntegrationsOfType(storage.IntegrationHealth_IMAGE_INTEGRATION)
}

func (ds *datastoreImpl) GetNotifierPlugins(ctx context.Context) ([]*storage.IntegrationHealth, error) {
	if ok, err := notifierSAC.ReadAllowed(ctx); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to retrieve health for notifiers: %v", err)
	} else if !ok {
		return nil, nil
	}
	return ds.getIntegrationsOfType(storage.IntegrationHealth_NOTIFIER)
}

func (ds *datastoreImpl) GetBackupPlugins(ctx context.Context) ([]*storage.IntegrationHealth, error) {
	if ok, err := backupSAC.ReadAllowed(ctx); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to retrieve health for backup plugins: %v", err)
	} else if !ok {
		return nil, nil
	}
	return ds.getIntegrationsOfType(storage.IntegrationHealth_BACKUP)
}

func (ds *datastoreImpl) UpdateIntegrationHealth(ctx context.Context, integrationHealth *storage.IntegrationHealth) error {
	if ok, err := writeAllowed(ctx, integrationHealth.GetType()); err != nil {
		return status.Errorf(codes.Internal, "Failed to update health for integration %s: %v",
			integrationHealth.Id, err)
	} else if !ok {
		return nil
	}
	return ds.store.Upsert(integrationHealth)
}

func (ds *datastoreImpl) RemoveIntegrationHealth(ctx context.Context, id string) error {
	currentHealth, exists, err := ds.GetIntegrationHealth(ctx, id)
	if err != nil {
		return status.Errorf(codes.Internal, "unable to find integration health for integration %s", id)
	}
	if !exists {
		return nil
	}
	if ok, err := writeAllowed(ctx, currentHealth.GetType()); err != nil {
		return status.Errorf(codes.Internal, "Failed to remove health for integration %s: %v", id, err)
	} else if !ok {
		return nil
	}
	return ds.store.Delete(id)
}

func (ds *datastoreImpl) GetIntegrationHealth(ctx context.Context, id string) (*storage.IntegrationHealth, bool, error) {
	health, found, err := ds.store.Get(id)
	if !found || err != nil {
		return nil, false, err
	}
	if ok, err := readAllowed(ctx, health.GetType()); err != nil {
		return nil, false, status.Errorf(codes.Internal, "Failed to get health for integration %s: %v", id, err)
	} else if !ok {
		return nil, false, nil
	}
	return health, found, err
}

func writeAllowed(ctx context.Context, typ storage.IntegrationHealth_Type) (bool, error) {
	switch typ {
	case storage.IntegrationHealth_IMAGE_INTEGRATION:
		return imageSAC.WriteAllowed(ctx)
	case storage.IntegrationHealth_NOTIFIER:
		return notifierSAC.WriteAllowed(ctx)
	case storage.IntegrationHealth_BACKUP:
		return backupSAC.WriteAllowed(ctx)
	default:
		return false, utils.Should(status.Error(codes.Internal, "Unknown integration type"))
	}
}

func readAllowed(ctx context.Context, typ storage.IntegrationHealth_Type) (bool, error) {
	switch typ {
	case storage.IntegrationHealth_IMAGE_INTEGRATION:
		return imageSAC.ReadAllowed(ctx)
	case storage.IntegrationHealth_NOTIFIER:
		return notifierSAC.ReadAllowed(ctx)
	case storage.IntegrationHealth_BACKUP:
		return backupSAC.ReadAllowed(ctx)
	default:
		return false, utils.Should(status.Error(codes.Internal, "Unknown integration type"))
	}
}

func (ds *datastoreImpl) getIntegrationsOfType(integrationType storage.IntegrationHealth_Type) ([]*storage.IntegrationHealth, error) {
	var integrationHealth []*storage.IntegrationHealth
	err := ds.store.Walk(func(obj *storage.IntegrationHealth) error {
		if obj.GetType() == integrationType {
			integrationHealth = append(integrationHealth, obj)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return integrationHealth, nil
}
