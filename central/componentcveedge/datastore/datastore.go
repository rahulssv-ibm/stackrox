package datastore

import (
	"context"

	"github.com/stackrox/rox/central/componentcveedge/index"
	"github.com/stackrox/rox/central/componentcveedge/search"
	"github.com/stackrox/rox/central/componentcveedge/store"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	searchPkg "github.com/stackrox/rox/pkg/search"
)

// DataStore is an intermediary to Component/CVE edge storage.
//go:generate mockgen-wrapper DataStore
type DataStore interface {
	Search(ctx context.Context, q *v1.Query) ([]searchPkg.Result, error)
	SearchEdges(ctx context.Context, q *v1.Query) ([]*v1.SearchResult, error)
	SearchRawEdges(ctx context.Context, q *v1.Query) ([]*storage.ComponentCVEEdge, error)

	Exists(ctx context.Context, id string) (bool, error)
	Get(ctx context.Context, id string) (*storage.ComponentCVEEdge, bool, error)
	Count(ctx context.Context) (int, error)
	GetBatch(ctx context.Context, id []string) ([]*storage.ComponentCVEEdge, error)

	Upsert(ctx context.Context, cve *storage.ComponentCVEEdge) error
	Delete(ctx context.Context, ids ...string) error
}

// New returns a new instance of a DataStore.
func New(storage store.Store, indexer index.Indexer, searcher search.Searcher) (DataStore, error) {
	ds := &datastoreImpl{
		storage:  storage,
		indexer:  indexer,
		searcher: searcher,
	}
	return ds, nil
}
