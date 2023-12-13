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
	baseTable = "alerts"
	storeName = "Alert"
)

var (
	log            = logging.LoggerForModule()
	schema         = pkgSchema.AlertsSchema
	targetResource = resources.Alert
)

type storeType = storage.Alert

// Store is the interface to interact with the storage for storage.Alert
type Store interface {
	Upsert(ctx context.Context, obj *storeType) error
	UpsertMany(ctx context.Context, objs []*storeType) error
	Delete(ctx context.Context, id string) error
	DeleteByQuery(ctx context.Context, q *v1.Query) ([]string, error)
	DeleteMany(ctx context.Context, identifiers []string) error

	Count(ctx context.Context) (int, error)
	Exists(ctx context.Context, id string) (bool, error)

	Get(ctx context.Context, id string) (*storeType, bool, error)
	GetByQuery(ctx context.Context, query *v1.Query) ([]*storeType, error)
	GetMany(ctx context.Context, identifiers []string) ([]*storeType, []int, error)
	GetIDs(ctx context.Context) ([]string, error)

	Walk(ctx context.Context, fn func(obj *storeType) error) error
}

// New returns a new Store instance using the provided sql instance.
func New(db postgres.DB) Store {
	return pgSearch.NewGenericStore[storeType, *storeType](
		db,
		schema,
		pkGetter,
		insertIntoAlerts,
		copyFromAlerts,
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
		return errors.Wrapf(sac.ErrResourceAccessDenied, "modifying alerts with IDs [%s] was denied", strings.Join(deniedIDs, ", "))
	}
	return nil
}

func insertIntoAlerts(batch *pgx.Batch, obj *storage.Alert) error {

	serialized, marshalErr := obj.Marshal()
	if marshalErr != nil {
		return marshalErr
	}

	values := []interface{}{
		// parent primary keys start
		pgutils.NilOrUUID(obj.GetId()),
		obj.GetPolicy().GetId(),
		obj.GetPolicy().GetName(),
		obj.GetPolicy().GetDescription(),
		obj.GetPolicy().GetDisabled(),
		obj.GetPolicy().GetCategories(),
		obj.GetPolicy().GetSeverity(),
		obj.GetPolicy().GetEnforcementActions(),
		pgutils.NilOrTime(obj.GetPolicy().GetLastUpdated()),
		obj.GetPolicy().GetSORTName(),
		obj.GetPolicy().GetSORTLifecycleStage(),
		obj.GetPolicy().GetSORTEnforcement(),
		obj.GetLifecycleStage(),
		pgutils.NilOrUUID(obj.GetClusterId()),
		obj.GetClusterName(),
		obj.GetNamespace(),
		pgutils.NilOrUUID(obj.GetNamespaceId()),
		pgutils.NilOrUUID(obj.GetDeployment().GetId()),
		obj.GetDeployment().GetName(),
		obj.GetDeployment().GetInactive(),
		obj.GetImage().GetId(),
		obj.GetImage().GetName().GetRegistry(),
		obj.GetImage().GetName().GetRemote(),
		obj.GetImage().GetName().GetTag(),
		obj.GetImage().GetName().GetFullName(),
		obj.GetResource().GetResourceType(),
		obj.GetResource().GetName(),
		obj.GetEnforcement().GetAction(),
		pgutils.NilOrTime(obj.GetTime()),
		obj.GetState(),
		serialized,
	}

	finalStr := "INSERT INTO alerts (Id, Policy_Id, Policy_Name, Policy_Description, Policy_Disabled, Policy_Categories, Policy_Severity, Policy_EnforcementActions, Policy_LastUpdated, Policy_SORTName, Policy_SORTLifecycleStage, Policy_SORTEnforcement, LifecycleStage, ClusterId, ClusterName, Namespace, NamespaceId, Deployment_Id, Deployment_Name, Deployment_Inactive, Image_Id, Image_Name_Registry, Image_Name_Remote, Image_Name_Tag, Image_Name_FullName, Resource_ResourceType, Resource_Name, Enforcement_Action, Time, State, serialized) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31) ON CONFLICT(Id) DO UPDATE SET Id = EXCLUDED.Id, Policy_Id = EXCLUDED.Policy_Id, Policy_Name = EXCLUDED.Policy_Name, Policy_Description = EXCLUDED.Policy_Description, Policy_Disabled = EXCLUDED.Policy_Disabled, Policy_Categories = EXCLUDED.Policy_Categories, Policy_Severity = EXCLUDED.Policy_Severity, Policy_EnforcementActions = EXCLUDED.Policy_EnforcementActions, Policy_LastUpdated = EXCLUDED.Policy_LastUpdated, Policy_SORTName = EXCLUDED.Policy_SORTName, Policy_SORTLifecycleStage = EXCLUDED.Policy_SORTLifecycleStage, Policy_SORTEnforcement = EXCLUDED.Policy_SORTEnforcement, LifecycleStage = EXCLUDED.LifecycleStage, ClusterId = EXCLUDED.ClusterId, ClusterName = EXCLUDED.ClusterName, Namespace = EXCLUDED.Namespace, NamespaceId = EXCLUDED.NamespaceId, Deployment_Id = EXCLUDED.Deployment_Id, Deployment_Name = EXCLUDED.Deployment_Name, Deployment_Inactive = EXCLUDED.Deployment_Inactive, Image_Id = EXCLUDED.Image_Id, Image_Name_Registry = EXCLUDED.Image_Name_Registry, Image_Name_Remote = EXCLUDED.Image_Name_Remote, Image_Name_Tag = EXCLUDED.Image_Name_Tag, Image_Name_FullName = EXCLUDED.Image_Name_FullName, Resource_ResourceType = EXCLUDED.Resource_ResourceType, Resource_Name = EXCLUDED.Resource_Name, Enforcement_Action = EXCLUDED.Enforcement_Action, Time = EXCLUDED.Time, State = EXCLUDED.State, serialized = EXCLUDED.serialized"
	batch.Queue(finalStr, values...)

	return nil
}

func copyFromAlerts(ctx context.Context, s pgSearch.Deleter, tx *postgres.Tx, objs ...*storage.Alert) error {
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
		"policy_id",
		"policy_name",
		"policy_description",
		"policy_disabled",
		"policy_categories",
		"policy_severity",
		"policy_enforcementactions",
		"policy_lastupdated",
		"policy_sortname",
		"policy_sortlifecyclestage",
		"policy_sortenforcement",
		"lifecyclestage",
		"clusterid",
		"clustername",
		"namespace",
		"namespaceid",
		"deployment_id",
		"deployment_name",
		"deployment_inactive",
		"image_id",
		"image_name_registry",
		"image_name_remote",
		"image_name_tag",
		"image_name_fullname",
		"resource_resourcetype",
		"resource_name",
		"enforcement_action",
		"time",
		"state",
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
			obj.GetPolicy().GetId(),
			obj.GetPolicy().GetName(),
			obj.GetPolicy().GetDescription(),
			obj.GetPolicy().GetDisabled(),
			obj.GetPolicy().GetCategories(),
			obj.GetPolicy().GetSeverity(),
			obj.GetPolicy().GetEnforcementActions(),
			pgutils.NilOrTime(obj.GetPolicy().GetLastUpdated()),
			obj.GetPolicy().GetSORTName(),
			obj.GetPolicy().GetSORTLifecycleStage(),
			obj.GetPolicy().GetSORTEnforcement(),
			obj.GetLifecycleStage(),
			pgutils.NilOrUUID(obj.GetClusterId()),
			obj.GetClusterName(),
			obj.GetNamespace(),
			pgutils.NilOrUUID(obj.GetNamespaceId()),
			pgutils.NilOrUUID(obj.GetDeployment().GetId()),
			obj.GetDeployment().GetName(),
			obj.GetDeployment().GetInactive(),
			obj.GetImage().GetId(),
			obj.GetImage().GetName().GetRegistry(),
			obj.GetImage().GetName().GetRemote(),
			obj.GetImage().GetName().GetTag(),
			obj.GetImage().GetName().GetFullName(),
			obj.GetResource().GetResourceType(),
			obj.GetResource().GetName(),
			obj.GetEnforcement().GetAction(),
			pgutils.NilOrTime(obj.GetTime()),
			obj.GetState(),
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

			if _, err := tx.CopyFrom(ctx, pgx.Identifier{"alerts"}, copyCols, pgx.CopyFromRows(inputRows)); err != nil {
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
	dropTableAlerts(ctx, db)
}

func dropTableAlerts(ctx context.Context, db postgres.DB) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS alerts CASCADE")

}

// endregion Used for testing
