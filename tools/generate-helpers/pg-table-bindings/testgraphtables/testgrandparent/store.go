// Code generated by pg-bindings generator. DO NOT EDIT.

package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/stackrox/rox/central/metrics"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/grpc/authn"
	"github.com/stackrox/rox/pkg/logging"
	ops "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/postgres"
	pkgSchema "github.com/stackrox/rox/pkg/postgres/schema"
	"github.com/stackrox/rox/pkg/sac/resources"
	pgSearch "github.com/stackrox/rox/pkg/search/postgres"
	"gorm.io/gorm"
)

const (
	baseTable = "test_grandparents"
	storeName = "TestGrandparent"

	// using copyFrom, we may not even want to batch.  It would probably be simpler
	// to deal with failures if we just sent it all.  Something to think about as we
	// proceed and move into more e2e and larger performance testing
	batchSize = 10000
)

var (
	log            = logging.LoggerForModule()
	schema         = pkgSchema.TestGrandparentsSchema
	targetResource = resources.Namespace
)

type storeType = storage.TestGrandparent

// Store is the interface to interact with the storage for storage.TestGrandparent
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
}

// New returns a new Store instance using the provided sql instance.
func New(db postgres.DB) Store {
	return pgSearch.NewGenericStore[storeType, *storeType](
		db,
		schema,
		pkGetter,
		insertIntoTestGrandparents,
		copyFromTestGrandparents,
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

func insertIntoTestGrandparents(ctx context.Context, batch *pgx.Batch, obj *storage.TestGrandparent) error {

	serialized, marshalErr := obj.Marshal()
	if marshalErr != nil {
		return marshalErr
	}

	ctxIdentity := authn.IdentityFromContextOrNil(ctx)
	if ctxIdentity == nil {
		return nil
	}

	values := []interface{}{
		// parent primary keys start
		obj.GetId(),
		obj.GetVal(),
		obj.GetPriority(),
		obj.GetRiskScore(),
		serialized,
	}

	finalStr := "INSERT INTO test_grandparents (Id, Val, Priority, RiskScore, serialized) VALUES($1, $2, $3, $4, $5) ON CONFLICT(Id) DO UPDATE SET Id = EXCLUDED.Id, Val = EXCLUDED.Val, Priority = EXCLUDED.Priority, RiskScore = EXCLUDED.RiskScore, serialized = EXCLUDED.serialized"
	batch.Queue(finalStr, values...)

	var query string

	for childIndex, child := range obj.GetEmbedded() {
		if err := insertIntoTestGrandparentsEmbeddeds(ctx, batch, child, obj.GetId(), childIndex); err != nil {
			return err
		}
	}

	query = "delete from test_grandparents_embeddeds where test_grandparents_Id = $1 AND idx >= $2"
	batch.Queue(query, obj.GetId(), len(obj.GetEmbedded()))
	return nil
}

func insertIntoTestGrandparentsEmbeddeds(ctx context.Context, batch *pgx.Batch, obj *storage.TestGrandparent_Embedded, testGrandparentID string, idx int) error {

	ctxIdentity := authn.IdentityFromContextOrNil(ctx)
	if ctxIdentity == nil {
		return nil
	}

	values := []interface{}{
		// parent primary keys start
		testGrandparentID,
		idx,
		obj.GetVal(),
	}

	finalStr := "INSERT INTO test_grandparents_embeddeds (test_grandparents_Id, idx, Val) VALUES($1, $2, $3) ON CONFLICT(test_grandparents_Id, idx) DO UPDATE SET test_grandparents_Id = EXCLUDED.test_grandparents_Id, idx = EXCLUDED.idx, Val = EXCLUDED.Val"
	batch.Queue(finalStr, values...)

	var query string

	for childIndex, child := range obj.GetEmbedded2() {
		if err := insertIntoTestGrandparentsEmbeddedsEmbedded2(ctx, batch, child, testGrandparentID, idx, childIndex); err != nil {
			return err
		}
	}

	query = "delete from test_grandparents_embeddeds_embedded2 where test_grandparents_Id = $1 AND test_grandparents_embeddeds_idx = $2 AND idx >= $3"
	batch.Queue(query, testGrandparentID, idx, len(obj.GetEmbedded2()))
	return nil
}

func insertIntoTestGrandparentsEmbeddedsEmbedded2(ctx context.Context, batch *pgx.Batch, obj *storage.TestGrandparent_Embedded_Embedded2, testGrandparentID string, testGrandparentEmbeddedIdx int, idx int) error {

	ctxIdentity := authn.IdentityFromContextOrNil(ctx)
	if ctxIdentity == nil {
		return nil
	}

	values := []interface{}{
		// parent primary keys start
		testGrandparentID,
		testGrandparentEmbeddedIdx,
		idx,
		obj.GetVal(),
	}

	finalStr := "INSERT INTO test_grandparents_embeddeds_embedded2 (test_grandparents_Id, test_grandparents_embeddeds_idx, idx, Val) VALUES($1, $2, $3, $4) ON CONFLICT(test_grandparents_Id, test_grandparents_embeddeds_idx, idx) DO UPDATE SET test_grandparents_Id = EXCLUDED.test_grandparents_Id, test_grandparents_embeddeds_idx = EXCLUDED.test_grandparents_embeddeds_idx, idx = EXCLUDED.idx, Val = EXCLUDED.Val"
	batch.Queue(finalStr, values...)

	return nil
}

func copyFromTestGrandparents(ctx context.Context, s pgSearch.Deleter, tx *postgres.Tx, objs ...*storage.TestGrandparent) error {
	inputRows := make([][]interface{}, 0, batchSize)

	// This is a copy so first we must delete the rows and re-add them
	// Which is essentially the desired behaviour of an upsert.
	deletes := make([]string, 0, batchSize)

	copyCols := []string{
		"id",
		"val",
		"priority",
		"riskscore",
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
			obj.GetVal(),
			obj.GetPriority(),
			obj.GetRiskScore(),
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

			if _, err := tx.CopyFrom(ctx, pgx.Identifier{"test_grandparents"}, copyCols, pgx.CopyFromRows(inputRows)); err != nil {
				return err
			}
			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	for idx, obj := range objs {
		_ = idx // idx may or may not be used depending on how nested we are, so avoid compile-time errors.

		if err := copyFromTestGrandparentsEmbeddeds(ctx, s, tx, obj.GetId(), obj.GetEmbedded()...); err != nil {
			return err
		}
	}

	return nil
}

func copyFromTestGrandparentsEmbeddeds(ctx context.Context, s pgSearch.Deleter, tx *postgres.Tx, testGrandparentID string, objs ...*storage.TestGrandparent_Embedded) error {
	inputRows := make([][]interface{}, 0, batchSize)

	copyCols := []string{
		"test_grandparents_id",
		"idx",
		"val",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj "+
			"in the loop is not used as it only consists of the parent ID and the index.  Putting this here as a stop gap "+
			"to simply use the object.  %s", obj)

		inputRows = append(inputRows, []interface{}{
			testGrandparentID,
			idx,
			obj.GetVal(),
		})

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			if _, err := tx.CopyFrom(ctx, pgx.Identifier{"test_grandparents_embeddeds"}, copyCols, pgx.CopyFromRows(inputRows)); err != nil {
				return err
			}
			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	for idx, obj := range objs {
		_ = idx // idx may or may not be used depending on how nested we are, so avoid compile-time errors.

		if err := copyFromTestGrandparentsEmbeddedsEmbedded2(ctx, s, tx, testGrandparentID, idx, obj.GetEmbedded2()...); err != nil {
			return err
		}
	}

	return nil
}

func copyFromTestGrandparentsEmbeddedsEmbedded2(ctx context.Context, s pgSearch.Deleter, tx *postgres.Tx, testGrandparentID string, testGrandparentEmbeddedIdx int, objs ...*storage.TestGrandparent_Embedded_Embedded2) error {
	inputRows := make([][]interface{}, 0, batchSize)

	copyCols := []string{
		"test_grandparents_id",
		"test_grandparents_embeddeds_idx",
		"idx",
		"val",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj "+
			"in the loop is not used as it only consists of the parent ID and the index.  Putting this here as a stop gap "+
			"to simply use the object.  %s", obj)

		inputRows = append(inputRows, []interface{}{
			testGrandparentID,
			testGrandparentEmbeddedIdx,
			idx,
			obj.GetVal(),
		})

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			if _, err := tx.CopyFrom(ctx, pgx.Identifier{"test_grandparents_embeddeds_embedded2"}, copyCols, pgx.CopyFromRows(inputRows)); err != nil {
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
	dropTableTestGrandparents(ctx, db)
}

func dropTableTestGrandparents(ctx context.Context, db postgres.DB) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS test_grandparents CASCADE")
	dropTableTestGrandparentsEmbeddeds(ctx, db)

}

func dropTableTestGrandparentsEmbeddeds(ctx context.Context, db postgres.DB) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS test_grandparents_embeddeds CASCADE")
	dropTableTestGrandparentsEmbeddedsEmbedded2(ctx, db)

}

func dropTableTestGrandparentsEmbeddedsEmbedded2(ctx context.Context, db postgres.DB) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS test_grandparents_embeddeds_embedded2 CASCADE")

}

// endregion Used for testing
