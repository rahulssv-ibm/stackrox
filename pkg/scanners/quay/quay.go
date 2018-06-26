package quay

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"bitbucket.org/stack-rox/apollo/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/images"
	"bitbucket.org/stack-rox/apollo/pkg/logging"
	"bitbucket.org/stack-rox/apollo/pkg/registries"
	quayRegistry "bitbucket.org/stack-rox/apollo/pkg/registries/quay"
	"bitbucket.org/stack-rox/apollo/pkg/scanners"
	"bitbucket.org/stack-rox/apollo/pkg/urlfmt"
)

const (
	requestTimeout = 5 * time.Second
)

var (
	log = logging.LoggerForModule()
)

type quay struct {
	client *http.Client

	endpoint   string
	oauthToken string
	registry   registries.ImageRegistry

	protoImageIntegration *v1.ImageIntegration
}

func newScanner(protoImageIntegration *v1.ImageIntegration) (*quay, error) {
	quayConfig, ok := protoImageIntegration.IntegrationConfig.(*v1.ImageIntegration_Quay)
	if !ok {
		return nil, fmt.Errorf("Quay config must be specified")
	}
	config := quayConfig.Quay

	registry, err := quayRegistry.NewRegistryFromConfig(quayConfig.Quay, protoImageIntegration)
	if err != nil {
		return nil, err
	}

	endpoint, err := urlfmt.FormatURL(config.GetEndpoint(), true, false)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Timeout: requestTimeout,
	}
	scanner := &quay{
		client: client,

		registry:   registry,
		endpoint:   endpoint,
		oauthToken: config.GetOauthToken(),

		protoImageIntegration: protoImageIntegration,
	}
	return scanner, nil
}

func (q *quay) sendRequest(method string, values url.Values, pathSegments ...string) ([]byte, int, error) {
	fullURL, err := urlfmt.FullyQualifiedURL(q.endpoint, values, pathSegments...)
	if err != nil {
		return nil, -1, err
	}
	req, err := http.NewRequest(method, fullURL, nil)
	if err != nil {
		return nil, -1, err
	}
	if q.oauthToken != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", q.oauthToken))
	}
	resp, err := q.client.Do(req)
	if err != nil {
		return nil, -1, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return body, resp.StatusCode, nil
}

// Test initiates a test of the Quay Scanner which verifies that we have the proper scan permissions
func (q *quay) Test() error {
	return q.registry.Test()
}

// GetLastScan retrieves the most recent scan
func (q *quay) GetLastScan(image *v1.Image) (*v1.ImageScan, error) {
	if image == nil || image.GetName().GetRemote() == "" || image.GetName().GetTag() == "" {
		return nil, nil
	}

	values := url.Values{}
	values.Add("features", "true")
	values.Add("vulnerabilities", "true")
	digest := images.NewDigest(image.GetMetadata().GetRegistrySha()).Digest()
	body, status, err := q.sendRequest("GET", values, "api", "v1", "repository", image.GetName().GetRemote(), "manifest", digest, "security")
	if err != nil {
		return nil, err
	} else if status != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status code %d when retrieving image scan for %s", status, images.Wrapper{Image: image})
	}
	scan, err := parseImageScan(body)
	if err != nil {
		return nil, err
	}
	if scan.Data.Layer == nil {
		return nil, fmt.Errorf("Layer for image %s was not found", image.GetName().GetFullName())
	}
	return convertScanToImageScan(image, scan), nil
}

// Match decides if the image is contained within this scanner
func (q *quay) Match(image *v1.Image) bool {
	return q.registry.Match(image)
}

func (q *quay) Global() bool {
	return len(q.protoImageIntegration.GetClusters()) == 0
}

func init() {
	scanners.Registry["quay"] = func(integration *v1.ImageIntegration) (scanners.ImageScanner, error) {
		scan, err := newScanner(integration)
		return scan, err
	}
}
