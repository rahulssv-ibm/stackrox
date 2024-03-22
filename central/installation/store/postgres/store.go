// Code generated by pg-bindings generator. DO NOT EDIT.
package postgres

import (
	"context"
	"time"

	"github.com/stackrox/rox/central/metrics"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/logging"
	ops "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
	pkgSchema "github.com/stackrox/rox/pkg/postgres/schema"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/sac/resources"
	"github.com/stackrox/rox/pkg/sync"
)

const (
	baseTable = "installation_infos"

	getStmt    = "SELECT serialized FROM installation_infos LIMIT 1"
	deleteStmt = "DELETE FROM installation_infos"
)

var (
	log            = logging.LoggerForModule()
	schema         = pkgSchema.InstallationInfosSchema
	targetResource = resources.InstallationInfo
)

// Store is the interface to interact with the storage for storage.InstallationInfo
type Store interface {
	Get(ctx context.Context) (*storage.InstallationInfo, bool, error)
	Upsert(ctx context.Context, obj *storage.InstallationInfo) error
	Delete(ctx context.Context) error
}

type storeImpl struct {
	db    postgres.DB
	mutex sync.Mutex
}

// New returns a new Store instance using the provided sql instance.
func New(db postgres.DB) Store {
	return &storeImpl{
		db: db,
	}
}

func insertIntoInstallationInfos(ctx context.Context, tx *postgres.Tx, obj *storage.InstallationInfo) error {
	serialized, marshalErr := obj.MarshalVT()
	if marshalErr != nil {
		return marshalErr
	}

	values := []interface{}{
		// parent primary keys start
		serialized,
	}

	finalStr := "INSERT INTO installation_infos (serialized) VALUES($1)"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}
	return nil
}

// Upsert saves the current state of an object in storage.
func (s *storeImpl) Upsert(ctx context.Context, obj *storage.InstallationInfo) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Upsert, "InstallationInfo")

	scopeChecker := sac.GlobalAccessScopeChecker(ctx).AccessMode(storage.Access_READ_WRITE_ACCESS).Resource(targetResource)
	if !scopeChecker.IsAllowed() {
		return sac.ErrResourceAccessDenied
	}

	return pgutils.Retry(func() error {
		return s.retryableUpsert(ctx, obj)
	})
}

func (s *storeImpl) retryableUpsert(ctx context.Context, obj *storage.InstallationInfo) error {
	conn, release, err := s.acquireConn(ctx, ops.Get, "InstallationInfo")
	if err != nil {
		return err
	}
	defer release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}

	if _, err := tx.Exec(ctx, deleteStmt); err != nil {
		return err
	}

	if err := insertIntoInstallationInfos(ctx, tx, obj); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return err
		}
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}

// Get returns the object, if it exists from the store.
func (s *storeImpl) Get(ctx context.Context) (*storage.InstallationInfo, bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Get, "InstallationInfo")

	scopeChecker := sac.GlobalAccessScopeChecker(ctx).AccessMode(storage.Access_READ_ACCESS).Resource(targetResource)
	if !scopeChecker.IsAllowed() {
		return nil, false, nil
	}

	return pgutils.Retry3(func() (*storage.InstallationInfo, bool, error) {
		return s.retryableGet(ctx)
	})
}

func (s *storeImpl) retryableGet(ctx context.Context) (*storage.InstallationInfo, bool, error) {
	conn, release, err := s.acquireConn(ctx, ops.Get, "InstallationInfo")
	if err != nil {
		return nil, false, err
	}
	defer release()

	row := conn.QueryRow(ctx, getStmt)
	var data []byte
	if err := row.Scan(&data); err != nil {
		return nil, false, pgutils.ErrNilIfNoRows(err)
	}

	var msg storage.InstallationInfo
	if err := msg.Unmarshal(data); err != nil {
		return nil, false, err
	}
	return &msg, true, nil
}

func (s *storeImpl) acquireConn(ctx context.Context, op ops.Op, typ string) (*postgres.Conn, func(), error) {
	defer metrics.SetAcquireDBConnDuration(time.Now(), op, typ)
	conn, err := s.db.Acquire(ctx)
	if err != nil {
		return nil, nil, err
	}
	return conn, conn.Release, nil
}

// Delete removes the singleton from the store
func (s *storeImpl) Delete(ctx context.Context) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Remove, "InstallationInfo")

	scopeChecker := sac.GlobalAccessScopeChecker(ctx).AccessMode(storage.Access_READ_WRITE_ACCESS).Resource(targetResource)
	if !scopeChecker.IsAllowed() {
		return sac.ErrResourceAccessDenied
	}

	return pgutils.Retry(func() error {
		return s.retryableDelete(ctx)
	})
}

func (s *storeImpl) retryableDelete(ctx context.Context) error {
	conn, release, err := s.acquireConn(ctx, ops.Remove, "InstallationInfo")
	if err != nil {
		return err
	}
	defer release()

	if _, err := conn.Exec(ctx, deleteStmt); err != nil {
		return err
	}
	return nil
}

// Used for Testing

// Destroy drops the tables associated with the target object type.
func Destroy(ctx context.Context, db postgres.DB) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS installation_infos CASCADE")
}
