package enrichment

import (
	"fmt"
	"sync"
	"time"

	"bitbucket.org/stack-rox/apollo/central/db"
	"bitbucket.org/stack-rox/apollo/central/risk"
	"bitbucket.org/stack-rox/apollo/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/logging"
	"bitbucket.org/stack-rox/apollo/pkg/sources"
	"github.com/karlseguin/ccache"
	"golang.org/x/time/rate"
)

const (
	imageDataExpiration = 10 * time.Minute

	maxCacheSize = 500
	itemsToPrune = 100
)

var (
	logger = logging.LoggerForModule()
)

// Enricher enriches images with data from registries and scanners.
type Enricher struct {
	storage interface {
		db.DeploymentStorage
		db.ImageStorage
		db.ImageIntegrationStorage
		db.MultiplierStorage
		db.AlertStorage
	}

	imageIntegrationMutex sync.Mutex
	imageIntegrations     map[string]*sources.ImageIntegration

	metadataLimiter *rate.Limiter
	metadataCache   *ccache.Cache

	scanLimiter *rate.Limiter
	scanCache   *ccache.Cache

	scorerMutex sync.Mutex
	scorer      *risk.Scorer
}

// New creates and returns a new Enricher.
func New(storage db.Storage, scorer *risk.Scorer) (*Enricher, error) {
	e := &Enricher{
		storage:         storage,
		scorer:          scorer,
		metadataLimiter: rate.NewLimiter(rate.Every(5*time.Second), 3),
		metadataCache:   ccache.New(ccache.Configure().MaxSize(maxCacheSize).ItemsToPrune(itemsToPrune)),
		scanLimiter:     rate.NewLimiter(rate.Every(5*time.Second), 3),
		scanCache:       ccache.New(ccache.Configure().MaxSize(maxCacheSize).ItemsToPrune(itemsToPrune)),
	}
	if err := e.initializeImageIntegrations(); err != nil {
		return nil, err
	}
	if err := e.initializeMultipliers(); err != nil {
		return nil, err
	}
	return e, nil
}

func (e *Enricher) initializeImageIntegrations() error {
	protoImageIntegrations, err := e.storage.GetImageIntegrations(&v1.GetImageIntegrationsRequest{})
	if err != nil {
		return err
	}
	e.imageIntegrations = make(map[string]*sources.ImageIntegration, len(protoImageIntegrations))
	for _, protoImageIntegration := range protoImageIntegrations {
		integration, err := sources.NewImageIntegration(protoImageIntegration)
		if err != nil {
			return fmt.Errorf("error generating an image integration from a persisted image integration: %s", err)
		}
		e.imageIntegrations[protoImageIntegration.GetId()] = integration
	}
	return nil
}

// UpdateImageIntegration updates the enricher's map of active image integratinos
func (e *Enricher) UpdateImageIntegration(integration *sources.ImageIntegration) {
	e.imageIntegrationMutex.Lock()
	defer e.imageIntegrationMutex.Unlock()
	e.imageIntegrations[integration.GetId()] = integration
}

// RemoveImageIntegration removes a image integration from the enricher's map of active image integrations
func (e *Enricher) RemoveImageIntegration(id string) {
	e.imageIntegrationMutex.Lock()
	defer e.imageIntegrationMutex.Unlock()
	delete(e.imageIntegrations, id)
}

func (e *Enricher) initializeMultipliers() error {
	protoMultipliers, err := e.storage.GetMultipliers()
	if err != nil {
		return err
	}
	for _, mult := range protoMultipliers {
		e.scorer.UpdateUserDefinedMultiplier(mult)
	}
	return nil
}

// Enrich enriches a deployment with data from registries and scanners.
func (e *Enricher) Enrich(deployment *v1.Deployment) (enriched bool, err error) {
	for _, c := range deployment.GetContainers() {
		if updated, err := e.enrichImage(c.GetImage()); err != nil {
			return false, err
		} else if updated {
			enriched = true
		}
	}

	if enriched {
		err = e.storage.UpdateDeployment(deployment)
	}

	return
}

// EnrichWithImageIntegration takes in a deployment and integration
func (e *Enricher) EnrichWithImageIntegration(deployment *v1.Deployment, integration *sources.ImageIntegration) bool {
	e.imageIntegrationMutex.Lock()
	defer e.imageIntegrationMutex.Unlock()
	var wasUpdated bool
	// TODO(cgorman) These may have a real ordering that we need to adhere to
	for _, category := range integration.GetCategories() {
		switch category {
		case v1.ImageIntegrationCategory_REGISTRY:
			updated := e.enrichWithRegistry(deployment, integration.Registry)
			if !wasUpdated {
				wasUpdated = updated
			}
		case v1.ImageIntegrationCategory_SCANNER:
			updated := e.enrichWithScanner(deployment, integration.Scanner)
			if !wasUpdated {
				wasUpdated = updated
			}
		}
	}
	return wasUpdated
}

func (e *Enricher) enrichImage(image *v1.Image) (bool, error) {
	updatedMetadata, err := e.enrichWithMetadata(image)
	if err != nil {
		return false, err
	}
	updatedScan, err := e.enrichWithScan(image)
	if err != nil {
		return false, err
	}
	if image.GetName().GetSha() != "" && (updatedMetadata || updatedScan) {
		// Store image in the database
		return true, e.storage.UpdateImage(image)
	}
	return false, nil
}
