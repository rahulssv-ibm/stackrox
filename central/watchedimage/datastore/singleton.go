package datastore

import (
	"github.com/stackrox/rox/central/globaldb"
	"github.com/stackrox/rox/central/watchedimage/datastore/internal/store/rocksdb"
	"github.com/stackrox/rox/pkg/sync"
	"github.com/stackrox/rox/pkg/utils"
)

var (
	instance DataStore
	once     sync.Once
)

// Singleton returns the instance of DataStore to use.
func Singleton() DataStore {
	once.Do(func() {
		store, err := rocksdb.New(globaldb.GetRocksDB())
		utils.Must(err)
		instance = New(store)
	})
	return instance
}
