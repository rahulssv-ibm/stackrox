package detector

import (
	"context"
	"errors"
	"time"

	"github.com/cenkalti/backoff/v3"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/internalapi/central"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/booleanpolicy/augmentedobjs"
	"github.com/stackrox/rox/pkg/concurrency"
	"github.com/stackrox/rox/pkg/env"
	"github.com/stackrox/rox/pkg/expiringcache"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/images/types"
	"github.com/stackrox/rox/pkg/protoutils"
	"github.com/stackrox/rox/pkg/set"
	"github.com/stackrox/rox/pkg/sync"
	"github.com/stackrox/rox/sensor/common/clusterid"
	"github.com/stackrox/rox/sensor/common/detector/metrics"
	"github.com/stackrox/rox/sensor/common/imagecacheutils"
	"github.com/stackrox/rox/sensor/common/registry"
	"github.com/stackrox/rox/sensor/common/scan"
	"github.com/stackrox/rox/sensor/common/store"
	"google.golang.org/grpc/status"
)

var (
	scanTimeout = env.ScanTimeout.DurationSetting()
)

type scanResult struct {
	context                context.Context
	action                 central.ResourceAction
	deployment             *storage.Deployment
	images                 []*storage.Image
	networkPoliciesApplied *augmentedobjs.NetworkPoliciesApplied
}

type imageChanResult struct {
	image        *storage.Image
	containerIdx int
}

type enricher struct {
	imageSvc       v1.ImageServiceClient
	scanResultChan chan scanResult

	serviceAccountStore store.ServiceAccountStore
	localScan           *scan.LocalScan
	imageCache          expiringcache.Cache
	stopSig             concurrency.Signal
	regStore            *registry.Store
}

type cacheValue struct {
	lock      sync.RWMutex
	signal    concurrency.Signal
	image     *storage.Image
	localScan *scan.LocalScan
	regStore  *registry.Store
}

func (c *cacheValue) waitAndGet() *storage.Image {
	// We need to wait before locking
	c.signal.Wait()
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.image
}

func scanImage(ctx context.Context, svc v1.ImageServiceClient, req *scanImageRequest, _ *scan.LocalScan) (*v1.ScanImageInternalResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, scanTimeout)
	defer cancel()

	internalReq := &v1.ScanImageInternalRequest{
		Image: req.containerImage,
	}
	if features.SourcedAutogeneratedIntegrations.Enabled() {
		internalReq.Source = &v1.ScanImageInternalRequest_Source{
			ClusterId:        req.clusterID,
			Namespace:        req.namespace,
			ImagePullSecrets: req.pullSecrets,
		}
	}

	return svc.ScanImageInternal(ctx, internalReq)
}

func scanImageLocal(ctx context.Context, svc v1.ImageServiceClient, req *scanImageRequest, localScan *scan.LocalScan) (*v1.ScanImageInternalResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, scanTimeout)
	defer cancel()

	img, err := localScan.EnrichLocalImageInNamespace(ctx, svc, req.containerImage, req.namespace, "", false)

	return &v1.ScanImageInternalResponse{
		Image: img,
	}, err
}

type scanFunc func(ctx context.Context, svc v1.ImageServiceClient, req *scanImageRequest, localScan *scan.LocalScan) (*v1.ScanImageInternalResponse, error)

func (c *cacheValue) scanWithRetries(ctx context.Context, svc v1.ImageServiceClient, req *scanImageRequest, scanFn scanFunc) (*v1.ScanImageInternalResponse, error) {
	eb := backoff.NewExponentialBackOff()
	eb.InitialInterval = 5 * time.Second
	eb.Multiplier = 2
	eb.MaxInterval = 4 * time.Minute
	eb.MaxElapsedTime = 0 // Never stop the backoff, leave that decision to the parent context.

	eb.Reset()

outer:
	for {
		// We want to get the time spent in backoff without including the time it took to scan the image.
		timeSpentInBackoffSoFar := eb.GetElapsedTime()
		scannedImage, err := scanFn(ctx, svc, req, c.localScan)
		if err != nil {
			for _, detail := range status.Convert(err).Details() {
				// If the client is effectively rate-limited, backoff and try again.
				if _, isTooManyParallelScans := detail.(*v1.ScanImageInternalResponseDetails_TooManyParallelScans); isTooManyParallelScans {
					time.Sleep(eb.NextBackOff())
					continue outer
				}
			}

			// If cluster local scan is rate-limited, backoff and try again
			if errors.Is(err, scan.ErrTooManyParallelScans) {
				dur := eb.NextBackOff()
				log.Debugf("local scan rate limited, backing off for %q: %q", dur, req.containerImage.GetName().GetFullName())
				time.Sleep(dur)
				continue outer
			}

			return nil, err
		}

		metrics.ObserveTimeSpentInExponentialBackoff(timeSpentInBackoffSoFar)

		return scannedImage, nil
	}
}

func (c *cacheValue) scanAndSet(ctx context.Context, svc v1.ImageServiceClient, req *scanImageRequest) {
	defer c.signal.Signal()

	// Ask Central to scan the image if the image is not local otherwise scan with local scanner
	scanImageFn := scanImage
	if c.regStore.IsLocal(req.containerImage.GetName()) {
		scanImageFn = scanImageLocal
		log.Debugf("Sending scan to local scanner for image %q", req.containerImage.GetName().GetFullName())
	} else {
		log.Debugf("Sending scan to central for image %q", req.containerImage.GetName().GetFullName())
	}

	scannedImage, err := c.scanWithRetries(ctx, svc, req, scanImageFn)
	// lock here to set the image
	c.lock.Lock()
	defer c.lock.Unlock()
	if err != nil {
		// Ignore the error and set the image to something basic,
		// so alerting can progress.
		log.Errorf("Scan request failed for image %q: %s", req.containerImage.GetName().GetFullName(), err)
		c.image = types.ToImage(req.containerImage)
		return
	}

	log.Debugf("Successful image scan for image %s: %d components returned by scanner", req.containerImage.GetName().GetFullName(), len(scannedImage.GetImage().GetScan().GetComponents()))
	c.image = scannedImage.GetImage()
}

func newEnricher(cache expiringcache.Cache, serviceAccountStore store.ServiceAccountStore, registryStore *registry.Store, localScan *scan.LocalScan) *enricher {
	return &enricher{
		scanResultChan:      make(chan scanResult),
		serviceAccountStore: serviceAccountStore,
		imageCache:          cache,
		stopSig:             concurrency.NewSignal(),
		localScan:           localScan,
		regStore:            registryStore,
	}
}

func (e *enricher) getImageFromCache(key string) (*storage.Image, bool) {
	value, _ := e.imageCache.Get(key).(*cacheValue)
	if value == nil {
		return nil, false
	}
	return value.waitAndGet(), true
}

func (e *enricher) runScan(req *scanImageRequest) imageChanResult {
	// Cache key is either going to be image full name or image ID.
	// In case of image full name, we can skip. In case of image ID, we should make sure to check if the image's name
	// is equal / contained in the images `Names` field.
	key := imagecacheutils.GetImageCacheKey(req.containerImage)

	// If the container image says that the image is not pullable, don't even bother trying to scan
	if req.containerImage.GetNotPullable() {
		log.Warnf("Skipping image scan for image: %s. Not pullable", req.containerImage.GetName().GetFullName())
		return imageChanResult{
			image:        types.ToImage(req.containerImage),
			containerIdx: req.containerIdx,
		}
	}

	// forceEnrichImageWithSignatures will be set to true in case we have an image where a cached value exists for the
	// digest, but the name has not been added to the "Names" field. In this case, we will force a
	// re-scan of the image, which should only fetch & verify signatures (since we already have a scan
	// result associated, this should not matter.
	var forceEnrichImageWithSignatures bool

	img, ok := e.getImageFromCache(key)
	if ok {
		// If the container image name is already within the cached images names, we can short-circuit.
		if protoutils.SliceContains(req.containerImage.GetName(), img.GetNames()) {
			log.Debugf("Image scan loaded from cache: %s: Components: (%d)", req.containerImage.GetName().GetFullName(), len(img.GetScan().GetComponents()))
			return imageChanResult{
				image:        img,
				containerIdx: req.containerIdx,
			}
		}
		// We found an image that is already in cache (i.e. with the same digest), but the image name is different.
		// Ensuring we have a fully enriched image (especially regarding image signatures), we need to make sure to
		// scan this image once more. This should result in the signatures + signature verification being re-done.
		forceEnrichImageWithSignatures = true
	}

	newValue := &cacheValue{
		signal:    concurrency.NewSignal(),
		localScan: e.localScan,
		regStore:  e.regStore,
	}
	value := e.imageCache.GetOrSet(key, newValue).(*cacheValue)
	if forceEnrichImageWithSignatures || newValue == value {
		value.scanAndSet(concurrency.AsContext(&e.stopSig), e.imageSvc, req)
	}
	return imageChanResult{
		image:        value.waitAndGet(),
		containerIdx: req.containerIdx,
	}
}

type scanImageRequest struct {
	containerIdx         int
	containerImage       *storage.ContainerImage
	clusterID, namespace string
	pullSecrets          []string
}

func (e *enricher) runImageScanAsync(imageChan chan<- imageChanResult, req *scanImageRequest) {
	go func() {
		// unguarded send (push to channel outside a select) is allowed because the imageChan is a buffered channel of exact size
		imageChan <- e.runScan(req)
	}()
}

func (e *enricher) getImages(deployment *storage.Deployment) []*storage.Image {
	imageChan := make(chan imageChanResult, len(deployment.GetContainers()))

	var pullSecrets []string
	if features.SourcedAutogeneratedIntegrations.Enabled() {
		pullSecretsSet := set.NewStringSet(e.serviceAccountStore.GetImagePullSecrets(deployment.GetNamespace(), deployment.GetServiceAccount())...)
		pullSecretsSet.AddAll(deployment.GetImagePullSecrets()...)
		pullSecrets = pullSecretsSet.AsSlice()
	}
	for idx, container := range deployment.GetContainers() {
		e.runImageScanAsync(imageChan, &scanImageRequest{
			containerIdx:   idx,
			containerImage: container.GetImage(),
			clusterID:      clusterid.Get(),
			namespace:      deployment.GetNamespace(),
			pullSecrets:    pullSecrets,
		})
	}
	images := make([]*storage.Image, len(deployment.GetContainers()))
	for i := 0; i < len(deployment.GetContainers()); i++ {
		imgResult := <-imageChan

		// This will ensure that when we change the Name of the image
		// that it will not cause a potential race condition
		image := *imgResult.image.Clone()
		// Overwrite the image Name as a workaround to the fact that we fetch the image by ID
		// The ID may actually have many names that refer to it. e.g. busybox:latest and busybox:1.31 could have the
		// exact same ID
		image.Name = deployment.Containers[imgResult.containerIdx].GetImage().GetName()
		images[imgResult.containerIdx] = &image
	}
	return images
}

func (e *enricher) blockingScan(ctx context.Context, deployment *storage.Deployment, netpolApplied *augmentedobjs.NetworkPoliciesApplied, action central.ResourceAction) {
	select {
	case <-e.stopSig.Done():
		return
	case e.scanResultChan <- scanResult{
		context:                ctx,
		action:                 action,
		deployment:             deployment,
		images:                 e.getImages(deployment),
		networkPoliciesApplied: netpolApplied,
	}:
	}
}

func (e *enricher) outputChan() <-chan scanResult {
	return e.scanResultChan
}

func (e *enricher) stop() {
	e.stopSig.Signal()
}
