package datastore

import (
	"context"

	"github.com/stackrox/rox/central/nodecomponentedge/index"
	sacFilters "github.com/stackrox/rox/central/nodecomponentedge/sac"
	"github.com/stackrox/rox/central/nodecomponentedge/search"
	"github.com/stackrox/rox/central/nodecomponentedge/store"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/dackbox/graph"
	searchPkg "github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/search/filtered"
)

type datastoreImpl struct {
	storage       store.Store
	indexer       index.Indexer
	searcher      search.Searcher
	graphProvider graph.Provider
}

func (ds *datastoreImpl) Search(ctx context.Context, q *v1.Query) ([]searchPkg.Result, error) {
	return ds.searcher.Search(ctx, q)
}

func (ds *datastoreImpl) SearchEdges(ctx context.Context, q *v1.Query) ([]*v1.SearchResult, error) {
	return ds.searcher.SearchEdges(ctx, q)
}

func (ds *datastoreImpl) SearchRawEdges(ctx context.Context, q *v1.Query) ([]*storage.NodeComponentEdge, error) {
	edges, err := ds.searcher.SearchRawEdges(ctx, q)
	if err != nil {
		return nil, err
	}
	return edges, nil
}

func (ds *datastoreImpl) Count(ctx context.Context) (int, error) {
	results, err := ds.searcher.Search(ctx, searchPkg.EmptyQuery())
	if err != nil {
		return 0, err
	}
	return len(results), nil
}

func (ds *datastoreImpl) Get(ctx context.Context, id string) (*storage.NodeComponentEdge, bool, error) {
	filteredIDs, err := ds.filterReadable(ctx, []string{id})
	if err != nil || len(filteredIDs) != 1 {
		return nil, false, err
	}

	edge, found, err := ds.storage.Get(id)
	if err != nil || !found {
		return nil, false, err
	}
	return edge, true, nil
}

func (ds *datastoreImpl) Exists(ctx context.Context, id string) (bool, error) {
	filteredIDs, err := ds.filterReadable(ctx, []string{id})
	if err != nil || len(filteredIDs) != 1 {
		return false, err
	}

	found, err := ds.storage.Exists(id)
	if err != nil || !found {
		return false, err
	}
	return true, nil
}

func (ds *datastoreImpl) GetBatch(ctx context.Context, ids []string) ([]*storage.NodeComponentEdge, error) {
	filteredIDs, err := ds.filterReadable(ctx, ids)
	if err != nil {
		return nil, err
	}

	edges, _, err := ds.storage.GetBatch(filteredIDs)
	if err != nil {
		return nil, err
	}
	return edges, nil
}

func (ds *datastoreImpl) filterReadable(ctx context.Context, ids []string) ([]string, error) {
	var filteredIDs []string
	var err error
	graph.Context(ctx, ds.graphProvider, func(graphContext context.Context) {
		filteredIDs, err = filtered.ApplySACFilter(graphContext, ids, sacFilters.GetSACFilter())
	})
	return filteredIDs, err
}
