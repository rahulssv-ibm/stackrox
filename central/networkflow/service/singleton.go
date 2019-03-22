package service

import (
	deploymentDataStore "github.com/stackrox/rox/central/deployment/datastore"
	networkFlowStoreSingleton "github.com/stackrox/rox/central/networkflow/store/singleton"
	"github.com/stackrox/rox/central/networkpolicies/graph"
	"github.com/stackrox/rox/pkg/sync"
)

var (
	once sync.Once

	as Service
)

func initialize() {
	as = New(networkFlowStoreSingleton.Singleton(), deploymentDataStore.Singleton(), graph.Singleton())
}

// Singleton provides the instance of the Service interface to register.
func Singleton() Service {
	once.Do(initialize)
	return as
}
