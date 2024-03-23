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
	"github.com/stackrox/rox/pkg/protocompat"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/sac/resources"
	"github.com/stackrox/rox/pkg/search"
	pgSearch "github.com/stackrox/rox/pkg/search/postgres"
	"gorm.io/gorm"
)

const (
	baseTable = "clusters"
	storeName = "Cluster"
)

var (
	log            = logging.LoggerForModule()
	schema         = pkgSchema.ClustersSchema
	targetResource = resources.Cluster
)

type storeType = storage.Cluster

// Store is the interface to interact with the storage for storage.Cluster
type Store interface {
	Upsert(ctx context.Context, obj *storeType) error
	UpsertMany(ctx context.Context, objs []*storeType) error
	Delete(ctx context.Context, id string) error
	DeleteByQuery(ctx context.Context, q *v1.Query) ([]string, error)
	DeleteMany(ctx context.Context, identifiers []string) error

	Count(ctx context.Context, q *v1.Query) (int, error)
	Exists(ctx context.Context, id string) (bool, error)
	Search(ctx context.Context, q *v1.Query) ([]search.Result, error)

	Get(ctx context.Context, id string) (*storeType, bool, error)
	GetByQuery(ctx context.Context, query *v1.Query) ([]*storeType, error)
	GetMany(ctx context.Context, identifiers []string) ([]*storeType, []int, error)
	GetIDs(ctx context.Context) ([]string, error)

	Walk(ctx context.Context, fn func(obj *storeType) error) error
	WalkByQuery(ctx context.Context, query *v1.Query, fn func(obj *storeType) error) error
}

// New returns a new Store instance using the provided sql instance.
func New(db postgres.DB) Store {
	// Use of pgSearch.NewGenericStoreWithCache can be dangerous with high cardinality stores,
	// and be the source of memory pressure. Think twice about the need for in-memory caching
	// of the whole store.
	return pgSearch.NewGenericStoreWithCache[storeType, *storeType](
		db,
		schema,
		pkGetter,
		insertIntoClusters,
		nil,
		metricsSetAcquireDBConnDuration,
		metricsSetPostgresOperationDurationTime,
		metricsSetCacheOperationDurationTime,

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

func metricsSetCacheOperationDurationTime(start time.Time, op ops.Op) {
	metrics.SetCacheOperationDurationTime(start, op, storeName)
}

func isUpsertAllowed(ctx context.Context, objs ...*storeType) error {
	scopeChecker := sac.GlobalAccessScopeChecker(ctx).AccessMode(storage.Access_READ_WRITE_ACCESS).Resource(targetResource)
	if scopeChecker.IsAllowed() {
		return nil
	}
	var deniedIDs []string
	for _, obj := range objs {
		subScopeChecker := scopeChecker.ClusterID(obj.GetId())
		if !subScopeChecker.IsAllowed() {
			deniedIDs = append(deniedIDs, obj.GetId())
		}
	}
	if len(deniedIDs) != 0 {
		return errors.Wrapf(sac.ErrResourceAccessDenied, "modifying clusters with IDs [%s] was denied", strings.Join(deniedIDs, ", "))
	}
	return nil
}

func insertIntoClusters(batch *pgx.Batch, obj *storage.Cluster) error {

	serialized, marshalErr := protocompat.Marshal(obj)
	if marshalErr != nil {
		return marshalErr
	}

	values := []interface{}{
		// parent primary keys start
		pgutils.NilOrUUID(obj.GetId()),
		obj.GetName(),
		pgutils.EmptyOrMap(obj.GetLabels()),
		obj.GetStatus().GetProviderMetadata().GetCluster().GetType(),
		serialized,
	}

	finalStr := "INSERT INTO clusters (Id, Name, Labels, Status_ProviderMetadata_Cluster_Type, serialized) VALUES($1, $2, $3, $4, $5) ON CONFLICT(Id) DO UPDATE SET Id = EXCLUDED.Id, Name = EXCLUDED.Name, Labels = EXCLUDED.Labels, Status_ProviderMetadata_Cluster_Type = EXCLUDED.Status_ProviderMetadata_Cluster_Type, serialized = EXCLUDED.serialized"
	batch.Queue(finalStr, values...)

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
	dropTableClusters(ctx, db)
}

func dropTableClusters(ctx context.Context, db postgres.DB) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS clusters CASCADE")

}

// endregion Used for testing
