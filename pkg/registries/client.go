package registries

import (
	"context"
	"sync"
	"time"

	"bitbucket.org/stack-rox/apollo/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/clientconn"
	"bitbucket.org/stack-rox/apollo/pkg/logging"
)

var (
	log = logging.New("registries")
)

const (
	updateInterval = 15 * time.Second
)

// A Client checks for new registry integrations.
type Client struct {
	updateTicker *time.Ticker

	registries []ImageRegistry
	lock       sync.RWMutex

	clusterID       string
	centralEndpoint string

	done chan struct{}
}

// NewRegistriesClient returns a new client of the registries API
func NewRegistriesClient(centralEndpoint string, clusterID string) *Client {
	return &Client{
		updateTicker:    time.NewTicker(updateInterval),
		clusterID:       clusterID,
		centralEndpoint: centralEndpoint,
		done:            make(chan struct{}),
	}
}

// Start runs the client
func (c *Client) Start() {
	for {
		select {
		case <-c.updateTicker.C:
			c.doUpdate()
		case <-c.done:
			return
		}
	}
}

func (c *Client) doUpdate() {
	conn, err := clientconn.GRPCConnection(c.centralEndpoint)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cli := v1.NewRegistryServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := cli.GetRegistries(ctx, &v1.GetRegistriesRequest{Cluster: c.clusterID})
	if err != nil {
		log.Errorf("Error checking registries: %s", err)
		return
	}
	c.replaceRegistries(resp)
}
func (c *Client) replaceRegistries(resp *v1.GetRegistriesResponse) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.registries = nil
	for _, registry := range resp.GetRegistries() {
		s, err := CreateRegistry(registry)
		if err != nil {
			log.Errorf("Could not instantiate registry %v: %s", registry, err)
			continue
		}
		c.registries = append(c.registries, s)
	}
}

// Stop stops polling for new registries.
func (c *Client) Stop() {
	c.done <- struct{}{}
}

// Registries returns the currently-defined set of image registries.
func (c *Client) Registries() []ImageRegistry {
	if c == nil {
		return nil
	}

	c.lock.RLock()
	defer c.lock.RUnlock()
	registries := make([]ImageRegistry, len(c.registries))
	for i, s := range c.registries {
		registries[i] = s
	}
	return registries
}
