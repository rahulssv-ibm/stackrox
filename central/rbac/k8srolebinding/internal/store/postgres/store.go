// Code generated by pg-bindings generator. DO NOT EDIT.

package postgres

import (
	"context"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/central/metrics"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/logging"
	ops "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
	pkgSchema "github.com/stackrox/rox/pkg/postgres/schema"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/sac/resources"
	pgSearch "github.com/stackrox/rox/pkg/search/postgres"
	"gorm.io/gorm"
)

const (
	baseTable = "role_bindings"
	storeName = "K8SRoleBinding"
)

var (
	log            = logging.LoggerForModule()
	schema         = pkgSchema.RoleBindingsSchema
	targetResource = resources.K8sRoleBinding
)

type storeType = storage.K8SRoleBinding

// Store is the interface to interact with the storage for storage.K8SRoleBinding
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
		insertIntoRoleBindings,
		copyFromRoleBindings,
		metricsSetAcquireDBConnDuration,
		metricsSetPostgresOperationDurationTime,
		isUpsertAllowed,
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
func isUpsertAllowed(ctx context.Context, objs ...*storeType) error {
	scopeChecker := sac.GlobalAccessScopeChecker(ctx).AccessMode(storage.Access_READ_WRITE_ACCESS).Resource(targetResource)
	if scopeChecker.IsAllowed() {
		return nil
	}
	var deniedIDs []string
	for _, obj := range objs {
		subScopeChecker := scopeChecker.ClusterID(obj.GetClusterId()).Namespace(obj.GetNamespace())
		if !subScopeChecker.IsAllowed() {
			deniedIDs = append(deniedIDs, obj.GetId())
		}
	}
	if len(deniedIDs) != 0 {
		return errors.Wrapf(sac.ErrResourceAccessDenied, "modifying k8SRoleBindings with IDs [%s] was denied", strings.Join(deniedIDs, ", "))
	}
	return nil
}

func insertIntoRoleBindings(batch *pgx.Batch, obj *storage.K8SRoleBinding) error {

	serialized, marshalErr := obj.Marshal()
	if marshalErr != nil {
		return marshalErr
	}

	values := []interface{}{
		// parent primary keys start
		pgutils.NilOrUUID(obj.GetId()),
		obj.GetName(),
		obj.GetNamespace(),
		pgutils.NilOrUUID(obj.GetClusterId()),
		obj.GetClusterName(),
		obj.GetClusterRole(),
		pgutils.EmptyOrMap(obj.GetLabels()),
		pgutils.EmptyOrMap(obj.GetAnnotations()),
		pgutils.NilOrUUID(obj.GetRoleId()),
		serialized,
	}

	finalStr := "INSERT INTO role_bindings (Id, Name, Namespace, ClusterId, ClusterName, ClusterRole, Labels, Annotations, RoleId, serialized) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) ON CONFLICT(Id) DO UPDATE SET Id = EXCLUDED.Id, Name = EXCLUDED.Name, Namespace = EXCLUDED.Namespace, ClusterId = EXCLUDED.ClusterId, ClusterName = EXCLUDED.ClusterName, ClusterRole = EXCLUDED.ClusterRole, Labels = EXCLUDED.Labels, Annotations = EXCLUDED.Annotations, RoleId = EXCLUDED.RoleId, serialized = EXCLUDED.serialized"
	batch.Queue(finalStr, values...)

	var query string

	for childIndex, child := range obj.GetSubjects() {
		if err := insertIntoRoleBindingsSubjects(batch, child, obj.GetId(), childIndex); err != nil {
			return err
		}
	}

	query = "delete from role_bindings_subjects where role_bindings_Id = $1 AND idx >= $2"
	batch.Queue(query, pgutils.NilOrUUID(obj.GetId()), len(obj.GetSubjects()))
	return nil
}

func insertIntoRoleBindingsSubjects(batch *pgx.Batch, obj *storage.Subject, roleBindingID string, idx int) error {

	values := []interface{}{
		// parent primary keys start
		pgutils.NilOrUUID(roleBindingID),
		idx,
		obj.GetKind(),
		obj.GetName(),
	}

	finalStr := "INSERT INTO role_bindings_subjects (role_bindings_Id, idx, Kind, Name) VALUES($1, $2, $3, $4) ON CONFLICT(role_bindings_Id, idx) DO UPDATE SET role_bindings_Id = EXCLUDED.role_bindings_Id, idx = EXCLUDED.idx, Kind = EXCLUDED.Kind, Name = EXCLUDED.Name"
	batch.Queue(finalStr, values...)

	return nil
}

func copyFromRoleBindings(ctx context.Context, s pgSearch.Deleter, tx *postgres.Tx, objs ...*storage.K8SRoleBinding) error {
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
		"name",
		"namespace",
		"clusterid",
		"clustername",
		"clusterrole",
		"labels",
		"annotations",
		"roleid",
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
			pgutils.NilOrUUID(obj.GetId()),
			obj.GetName(),
			obj.GetNamespace(),
			pgutils.NilOrUUID(obj.GetClusterId()),
			obj.GetClusterName(),
			obj.GetClusterRole(),
			pgutils.EmptyOrMap(obj.GetLabels()),
			pgutils.EmptyOrMap(obj.GetAnnotations()),
			pgutils.NilOrUUID(obj.GetRoleId()),
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

			if _, err := tx.CopyFrom(ctx, pgx.Identifier{"role_bindings"}, copyCols, pgx.CopyFromRows(inputRows)); err != nil {
				return err
			}
			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	for idx, obj := range objs {
		_ = idx // idx may or may not be used depending on how nested we are, so avoid compile-time errors.

		if err := copyFromRoleBindingsSubjects(ctx, s, tx, obj.GetId(), obj.GetSubjects()...); err != nil {
			return err
		}
	}

	return nil
}

func copyFromRoleBindingsSubjects(ctx context.Context, s pgSearch.Deleter, tx *postgres.Tx, roleBindingID string, objs ...*storage.Subject) error {
	batchSize := pgSearch.MaxBatchSize
	if len(objs) < batchSize {
		batchSize = len(objs)
	}
	inputRows := make([][]interface{}, 0, batchSize)

	copyCols := []string{
		"role_bindings_id",
		"idx",
		"kind",
		"name",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj "+
			"in the loop is not used as it only consists of the parent ID and the index.  Putting this here as a stop gap "+
			"to simply use the object.  %s", obj)

		inputRows = append(inputRows, []interface{}{
			pgutils.NilOrUUID(roleBindingID),
			idx,
			obj.GetKind(),
			obj.GetName(),
		})

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			if _, err := tx.CopyFrom(ctx, pgx.Identifier{"role_bindings_subjects"}, copyCols, pgx.CopyFromRows(inputRows)); err != nil {
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
	dropTableRoleBindings(ctx, db)
}

func dropTableRoleBindings(ctx context.Context, db postgres.DB) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS role_bindings CASCADE")
	dropTableRoleBindingsSubjects(ctx, db)

}

func dropTableRoleBindingsSubjects(ctx context.Context, db postgres.DB) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS role_bindings_subjects CASCADE")

}

// endregion Used for testing
