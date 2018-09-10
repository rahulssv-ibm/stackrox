package service

import (
	"sync"

	deploymentDataStore "github.com/stackrox/rox/central/deployment/datastore"
	"github.com/stackrox/rox/central/secret/datastore"
)

var (
	once sync.Once
	as   Service
)

func initialize() {
	as = New(datastore.Singleton(), deploymentDataStore.Singleton())
}

// Singleton provides the instance of the Service interface to register.
func Singleton() Service {
	once.Do(initialize)
	return as
}
