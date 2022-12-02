// Code generated by pg-bindings generator. DO NOT EDIT.
package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	metrics "github.com/stackrox/rox/central/metrics"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	ops "github.com/stackrox/rox/pkg/metrics"
	search "github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/search/blevesearch"
	"github.com/stackrox/rox/pkg/search/postgres"
	"github.com/stackrox/rox/pkg/search/postgres/mapping"
)

func init() {
	mapping.RegisterCategoryToTable(v1.SearchCategory_PROCESS_LISTENING_ON_PORT, schema)
}

// NewIndexer returns new indexer for `storage.ProcessListeningOnPortStorage`.
func NewIndexer(db *pgxpool.Pool) *indexerImpl {
	return &indexerImpl{
		db: db,
	}
}

type indexerImpl struct {
	db *pgxpool.Pool
}

func (b *indexerImpl) Count(ctx context.Context, q *v1.Query, opts ...blevesearch.SearchOption) (int, error) {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Count, "ProcessListeningOnPortStorage")

	return postgres.RunCountRequest(ctx, v1.SearchCategory_PROCESS_LISTENING_ON_PORT, q, b.db)
}

func (b *indexerImpl) Search(ctx context.Context, q *v1.Query, opts ...blevesearch.SearchOption) ([]search.Result, error) {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Search, "ProcessListeningOnPortStorage")

	return postgres.RunSearchRequest(ctx, v1.SearchCategory_PROCESS_LISTENING_ON_PORT, q, b.db)
}

//// Stubs for satisfying interfaces

func (b *indexerImpl) AddProcessListeningOnPortStorage(deployment *storage.ProcessListeningOnPortStorage) error {
	return nil
}

func (b *indexerImpl) AddProcessListeningOnPortStorages(_ []*storage.ProcessListeningOnPortStorage) error {
	return nil
}

func (b *indexerImpl) DeleteProcessListeningOnPortStorage(id string) error {
	return nil
}

func (b *indexerImpl) DeleteProcessListeningOnPortStorages(_ []string) error {
	return nil
}

func (b *indexerImpl) MarkInitialIndexingComplete() error {
	return nil
}

func (b *indexerImpl) NeedsInitialIndexing() (bool, error) {
	return false, nil
}
