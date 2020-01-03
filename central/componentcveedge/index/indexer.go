// Code generated by blevebindings generator. DO NOT EDIT.

package index

import (
	bleve "github.com/blevesearch/bleve"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	blevehelper "github.com/stackrox/rox/pkg/blevehelper"
	search "github.com/stackrox/rox/pkg/search"
	blevesearch "github.com/stackrox/rox/pkg/search/blevesearch"
)

type Indexer interface {
	AddComponentCVEEdge(componentcveedge *storage.ComponentCVEEdge) error
	AddComponentCVEEdges(componentcveedges []*storage.ComponentCVEEdge) error
	DeleteComponentCVEEdge(id string) error
	DeleteComponentCVEEdges(ids []string) error
	ResetIndex() error
	Search(q *v1.Query, opts ...blevesearch.SearchOption) ([]search.Result, error)
}

func New(index bleve.Index) Indexer {
	wrapper, err := blevehelper.NewBleveWrapper(index, resourceName)
	if err != nil {
		panic(err)
	}
	return &indexerImpl{index: wrapper}
}
