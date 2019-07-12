package enrichment

import (
	"time"

	"github.com/stackrox/rox/central/imageintegration"
	"github.com/stackrox/rox/pkg/expiringcache"
	"github.com/stackrox/rox/pkg/images/enricher"
	"github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/sync"
)

var (
	once sync.Once

	ie enricher.ImageEnricher
	en Enricher

	scanCacheOnce sync.Once
	scanCache     expiringcache.Cache

	metadataCacheOnce sync.Once
	metadataCache     expiringcache.Cache

	imageCacheExpiryDuration = 4 * time.Hour
)

func initialize() {
	ie = enricher.New(imageintegration.Set(), metrics.CentralSubsystem, ImageMetadataCacheSingleton(), ImageScanCacheSingleton())
	en = New(ie)
}

// Singleton provides the singleton Enricher to use.
func Singleton() Enricher {
	once.Do(initialize)
	return en
}

// ImageEnricherSingleton provides the singleton ImageEnricher to use.
func ImageEnricherSingleton() enricher.ImageEnricher {
	once.Do(initialize)
	return ie
}

// ImageScanCacheSingleton returns the cache for image scans
func ImageScanCacheSingleton() expiringcache.Cache {
	scanCacheOnce.Do(func() {
		scanCache = expiringcache.NewExpiringCache(imageCacheExpiryDuration)
	})
	return scanCache
}

// ImageMetadataCacheSingleton returns the cache for image metadata
func ImageMetadataCacheSingleton() expiringcache.Cache {
	metadataCacheOnce.Do(func() {
		metadataCache = expiringcache.NewExpiringCache(imageCacheExpiryDuration)
	})
	return metadataCache
}
