package service

import (
	"github.com/stackrox/rox/central/enrichment"
	"github.com/stackrox/rox/central/image/datastore"
	"github.com/stackrox/rox/pkg/sync"
)

var (
	once sync.Once

	as Service
)

func initialize() {
	as = New(datastore.Singleton(), enrichment.ImageEnricherSingleton(), enrichment.ImageMetadataCacheSingleton(), enrichment.ImageScanCacheSingleton())
}

// Singleton provides the instance of the Service interface to register.
func Singleton() Service {
	once.Do(initialize)
	return as
}
