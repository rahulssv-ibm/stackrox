// Code generated by pg-bindings generator. DO NOT EDIT.

package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/stackrox/rox/central/metrics"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/logging"
	ops "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
	pkgSchema "github.com/stackrox/rox/pkg/postgres/schema"
	"github.com/stackrox/rox/pkg/sac/resources"
	pgSearch "github.com/stackrox/rox/pkg/search/postgres"
	"gorm.io/gorm"
)

const (
	baseTable = "active_components"
	storeName = "ActiveComponent"
)

var (
	log            = logging.LoggerForModule()
	schema         = pkgSchema.ActiveComponentsSchema
	targetResource = resources.Deployment
)

type storeType = storage.ActiveComponent

// Store is the interface to interact with the storage for storage.ActiveComponent
type Store interface {
	Upsert(ctx context.Context, obj *storeType) error
	UpsertMany(ctx context.Context, objs []*storeType) error
	Delete(ctx context.Context, id string) error
	DeleteByQuery(ctx context.Context, q *v1.Query) error
	DeleteMany(ctx context.Context, identifiers []string) error

	Count(ctx context.Context) (int, error)
	Exists(ctx context.Context, id string) (bool, error)

	Get(ctx context.Context, id string) (*storeType, bool, error)
	GetByQuery(ctx context.Context, query *v1.Query) ([]*storeType, error)
	GetMany(ctx context.Context, identifiers []string) ([]*storeType, []int, error)
	GetIDs(ctx context.Context) ([]string, error)

	Walk(ctx context.Context, fn func(obj *storeType) error) error
	WalkByQuery(ctx context.Context, query *v1.Query, fn func(obj *storeType) error) error
}

// New returns a new Store instance using the provided sql instance.
func New(db postgres.DB) Store {
	return pgSearch.NewGenericStore[storeType, *storeType](
		db,
		schema,
		pkGetter,
		insertIntoActiveComponents,
		copyFromActiveComponents,
		metricsSetAcquireDBConnDuration,
		metricsSetPostgresOperationDurationTime,
		pgSearch.GloballyScopedUpsertChecker[storeType, *storeType](targetResource),
		targetResource,
	)
}

// region Helper functions

func pkGetter(obj *storeType) string {
	return obj.GetId()
}

func metricsSetPostgresOperationDurationTime(start time.Time, op ops.Op) {
	metrics.SetPostgresOperationDurationTime(start, op, storeName)
}

func metricsSetAcquireDBConnDuration(start time.Time, op ops.Op) {
	metrics.SetAcquireDBConnDuration(start, op, storeName)
}

func insertIntoActiveComponents(batch *pgx.Batch, obj *storage.ActiveComponent) error {

	serialized, marshalErr := obj.Marshal()
	if marshalErr != nil {
		return marshalErr
	}

	values := []interface{}{
		// parent primary keys start
		obj.GetId(),
		pgutils.NilOrUUID(obj.GetDeploymentId()),
		obj.GetComponentId(),
		serialized,
	}

	finalStr := "INSERT INTO active_components (Id, DeploymentId, ComponentId, serialized) VALUES($1, $2, $3, $4) ON CONFLICT(Id) DO UPDATE SET Id = EXCLUDED.Id, DeploymentId = EXCLUDED.DeploymentId, ComponentId = EXCLUDED.ComponentId, serialized = EXCLUDED.serialized"
	batch.Queue(finalStr, values...)

	var query string

	for childIndex, child := range obj.GetActiveContextsSlice() {
		if err := insertIntoActiveComponentsActiveContextsSlices(batch, child, obj.GetId(), childIndex); err != nil {
			return err
		}
	}

	query = "delete from active_components_active_contexts_slices where active_components_Id = $1 AND idx >= $2"
	batch.Queue(query, obj.GetId(), len(obj.GetActiveContextsSlice()))
	return nil
}

func insertIntoActiveComponentsActiveContextsSlices(batch *pgx.Batch, obj *storage.ActiveComponent_ActiveContext, activeComponentID string, idx int) error {

	values := []interface{}{
		// parent primary keys start
		activeComponentID,
		idx,
		obj.GetContainerName(),
		obj.GetImageId(),
	}

	finalStr := "INSERT INTO active_components_active_contexts_slices (active_components_Id, idx, ContainerName, ImageId) VALUES($1, $2, $3, $4) ON CONFLICT(active_components_Id, idx) DO UPDATE SET active_components_Id = EXCLUDED.active_components_Id, idx = EXCLUDED.idx, ContainerName = EXCLUDED.ContainerName, ImageId = EXCLUDED.ImageId"
	batch.Queue(finalStr, values...)

	return nil
}

func copyFromActiveComponents(ctx context.Context, s pgSearch.Deleter, tx *postgres.Tx, objs ...*storage.ActiveComponent) error {
	batchSize := pgSearch.MaxBatchSize
	if len(objs) < batchSize {
		batchSize = len(objs)
	}
	inputRows := make([][]interface{}, 0, batchSize)

	// This is a copy so first we must delete the rows and re-add them
	// Which is essentially the desired behaviour of an upsert.
	deletes := make([]string, 0, batchSize)

	copyCols := []string{
		"id",
		"deploymentid",
		"componentid",
		"serialized",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj "+
			"in the loop is not used as it only consists of the parent ID and the index.  Putting this here as a stop gap "+
			"to simply use the object.  %s", obj)

		serialized, marshalErr := obj.Marshal()
		if marshalErr != nil {
			return marshalErr
		}

		inputRows = append(inputRows, []interface{}{
			obj.GetId(),
			pgutils.NilOrUUID(obj.GetDeploymentId()),
			obj.GetComponentId(),
			serialized,
		})

		// Add the ID to be deleted.
		deletes = append(deletes, obj.GetId())

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			if err := s.DeleteMany(ctx, deletes); err != nil {
				return err
			}
			// clear the inserts and vals for the next batch
			deletes = deletes[:0]

			if _, err := tx.CopyFrom(ctx, pgx.Identifier{"active_components"}, copyCols, pgx.CopyFromRows(inputRows)); err != nil {
				return err
			}
			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	for idx, obj := range objs {
		_ = idx // idx may or may not be used depending on how nested we are, so avoid compile-time errors.

		if err := copyFromActiveComponentsActiveContextsSlices(ctx, s, tx, obj.GetId(), obj.GetActiveContextsSlice()...); err != nil {
			return err
		}
	}

	return nil
}

func copyFromActiveComponentsActiveContextsSlices(ctx context.Context, s pgSearch.Deleter, tx *postgres.Tx, activeComponentID string, objs ...*storage.ActiveComponent_ActiveContext) error {
	batchSize := pgSearch.MaxBatchSize
	if len(objs) < batchSize {
		batchSize = len(objs)
	}
	inputRows := make([][]interface{}, 0, batchSize)

	copyCols := []string{
		"active_components_id",
		"idx",
		"containername",
		"imageid",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj "+
			"in the loop is not used as it only consists of the parent ID and the index.  Putting this here as a stop gap "+
			"to simply use the object.  %s", obj)

		inputRows = append(inputRows, []interface{}{
			activeComponentID,
			idx,
			obj.GetContainerName(),
			obj.GetImageId(),
		})

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			if _, err := tx.CopyFrom(ctx, pgx.Identifier{"active_components_active_contexts_slices"}, copyCols, pgx.CopyFromRows(inputRows)); err != nil {
				return err
			}
			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	return nil
}

// endregion Helper functions

// region Used for testing

// CreateTableAndNewStore returns a new Store instance for testing.
func CreateTableAndNewStore(ctx context.Context, db postgres.DB, gormDB *gorm.DB) Store {
	pkgSchema.ApplySchemaForTable(ctx, gormDB, baseTable)
	return New(db)
}

// Destroy drops the tables associated with the target object type.
func Destroy(ctx context.Context, db postgres.DB) {
	dropTableActiveComponents(ctx, db)
}

func dropTableActiveComponents(ctx context.Context, db postgres.DB) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS active_components CASCADE")
	dropTableActiveComponentsActiveContextsSlices(ctx, db)

}

func dropTableActiveComponentsActiveContextsSlices(ctx context.Context, db postgres.DB) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS active_components_active_contexts_slices CASCADE")

}

// endregion Used for testing
