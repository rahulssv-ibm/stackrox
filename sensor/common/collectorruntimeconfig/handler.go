package collectorruntimeconfig

import (
	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/internalapi/central"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/centralsensor"
	"github.com/stackrox/rox/pkg/concurrency"

	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/sync"
	"github.com/stackrox/rox/sensor/common"
	"github.com/stackrox/rox/sensor/common/message"
)

var (
	log = logging.LoggerForModule()
)

// Store is a store for network graph external sources.
//
//go:generate mockgen-wrapper
type Store interface {
	CollectorConfigValueStream() concurrency.ReadOnlyValueStream[*storage.CollectorConfig]
}

// Handler forwards the collector runtime config received from Central to Collectors.
type Handler interface {
	common.SensorComponent
}

type handlerImpl struct {
	stopSig   concurrency.Signal
	updateSig concurrency.Signal

	collectorConfig            *storage.CollectorConfig
	collectorConfigProtoStream *concurrency.ValueStream[*storage.CollectorConfig]

	lock sync.Mutex
}

func (h *handlerImpl) Start() error {
	go h.run()
	return nil
}

func (h *handlerImpl) Stop(_ error) {
	h.stopSig.Signal()
}

func (h *handlerImpl) Notify(common.SensorComponentEvent) {}

func (h *handlerImpl) Capabilities() []centralsensor.SensorCapability {
	// return []centralsensor.SensorCapability{centralsensor.NetworkGraphExternalSrcsCap}
	return nil
}

func getCollectorConfig(msg *central.MsgToSensor) *storage.CollectorConfig {
	if clusterConfig := msg.GetClusterConfig(); clusterConfig != nil {
		if config := clusterConfig.GetConfig(); config != nil {
			if collectorConfig := config.GetCollectorConfig(); collectorConfig != nil {
				return collectorConfig
			}
		}
	}

	return nil
}

func (h *handlerImpl) ProcessMessage(msg *central.MsgToSensor) error {
	collectorConfig := getCollectorConfig(msg)
	if collectorConfig == nil {
		return nil
	}

	log.Infof("Got collector config %+v", collectorConfig)
	select {
	case <-h.stopSig.Done():
		return errors.New("could not process external network entities request")
	default:
		h.lock.Lock()
		defer h.lock.Unlock()

		h.collectorConfig = collectorConfig
		h.updateSig.Signal()
		return nil
	}
}

func (h *handlerImpl) ResponsesC() <-chan *message.ExpiringMessage {
	return nil
}

func (h *handlerImpl) run() {
	for {
		select {
		case <-h.updateSig.Done():
			h.pushConfigToValueStream()
		case <-h.stopSig.Done():
			return
		}
	}
}

func (h *handlerImpl) pushConfigToValueStream() {
	h.lock.Lock()
	defer h.lock.Unlock()

	defer h.updateSig.Reset()

	h.collectorConfigProtoStream.Push(h.collectorConfig)
}

func (h *handlerImpl) CollectorConfigValueStream() concurrency.ReadOnlyValueStream[*storage.CollectorConfig] {
	return h.collectorConfigProtoStream
}
