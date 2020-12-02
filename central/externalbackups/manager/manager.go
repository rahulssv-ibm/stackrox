package manager

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/stackrox/rox/central/externalbackups/plugins"
	"github.com/stackrox/rox/central/externalbackups/plugins/types"
	"github.com/stackrox/rox/central/externalbackups/scheduler"
	"github.com/stackrox/rox/central/integrationhealth/reporter"
	"github.com/stackrox/rox/central/role/resources"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/concurrency"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/protoconv/schedule"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/sync"
)

var (
	log = logging.LoggerForModule()
)

// Manager implements the interface for external backups
type Manager interface {
	Upsert(ctx context.Context, backup *storage.ExternalBackup) error
	Test(ctx context.Context, backup *storage.ExternalBackup) error
	Remove(ctx context.Context, id string)

	Backup(ctx context.Context, id string) error
}

// New returns a new external backup manager
func New() Manager {
	return &managerImpl{
		scheduler:            scheduler.New(),
		idsToExternalBackups: make(map[string]types.ExternalBackup),
	}
}

var (
	externalBkpSAC = sac.ForResource(resources.BackupPlugins)
)

type managerImpl struct {
	scheduler scheduler.Scheduler

	lock       sync.Mutex
	inProgress concurrency.Flag

	idsToExternalBackups map[string]types.ExternalBackup
}

func renderExternalBackupFromProto(backup *storage.ExternalBackup) (types.ExternalBackup, error) {
	creator, ok := plugins.Registry[backup.GetType()]
	if !ok {
		return nil, fmt.Errorf("external backup with type %q is not implemented", backup.GetType())
	}

	backupInterface, err := creator(backup, reporter.Singleton())
	if err != nil {
		return nil, err
	}
	return backupInterface, nil
}

func (m *managerImpl) Upsert(ctx context.Context, backup *storage.ExternalBackup) error {
	if ok, err := externalBkpSAC.WriteAllowed(ctx); err != nil {
		return err
	} else if !ok {
		return errors.New("permission denied")
	}

	backupInterface, err := renderExternalBackupFromProto(backup)
	if err != nil {
		return err
	}

	cronTab, err := schedule.ConvertToCronTab(backup.GetSchedule())
	if err != nil {
		return err
	}

	if err := m.scheduler.UpsertBackup(backup.GetId(), cronTab, backupInterface); err != nil {
		return err
	}

	m.lock.Lock()
	defer m.lock.Unlock()
	m.idsToExternalBackups[backup.GetId()] = backupInterface

	return nil
}

func (m *managerImpl) Test(ctx context.Context, backup *storage.ExternalBackup) error {
	if ok, err := externalBkpSAC.WriteAllowed(ctx); err != nil {
		return err
	} else if !ok {
		return errors.New("permission denied")
	}

	backupInterface, err := renderExternalBackupFromProto(backup)
	if err != nil {
		return err
	}
	return backupInterface.Test()
}

func (m *managerImpl) getBackup(id string) (types.ExternalBackup, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	backup, ok := m.idsToExternalBackups[id]
	if !ok {
		return nil, fmt.Errorf("backup with id %q does not exist", id)
	}
	return backup, nil
}

func (m *managerImpl) Backup(ctx context.Context, id string) error {
	if ok, err := externalBkpSAC.WriteAllowed(ctx); err != nil {
		return err
	} else if !ok {
		return errors.New("permission denied")
	}

	backup, err := m.getBackup(id)
	if err != nil {
		return err
	}

	if m.inProgress.TestAndSet(true) {
		return errors.New("backup already in progress")
	}

	defer m.inProgress.Set(false)

	if err := m.scheduler.RunBackup(backup); err != nil {
		return err
	}
	return nil
}

func (m *managerImpl) Remove(ctx context.Context, id string) {
	if ok, err := externalBkpSAC.WriteAllowed(ctx); err != nil || !ok {
		return
	}

	m.scheduler.RemoveBackup(id)

	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.idsToExternalBackups, id)
}
