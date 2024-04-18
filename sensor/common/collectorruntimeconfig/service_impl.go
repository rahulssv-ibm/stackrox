package collectorruntimeconfig

import (
	"context"
	// "sync/atomic"

	// "github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	// "github.com/pkg/errors"
	"github.com/stackrox/rox/generated/internalapi/central"
	// "github.com/stackrox/rox/generated/internalapi/compliance"
	"github.com/stackrox/rox/generated/internalapi/sensor"
	// "github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/centralsensor"
	"github.com/stackrox/rox/pkg/grpc/authz/idcheck"
	// "github.com/stackrox/rox/pkg/k8sutil"
	// "github.com/stackrox/rox/pkg/sync"
	"github.com/stackrox/rox/sensor/common"
	"github.com/stackrox/rox/sensor/common/message"
	// "github.com/stackrox/rox/sensor/common/orchestrator"
	"google.golang.org/grpc"
)

// CollectorService is the struct that manages the compliance results and audit log events
type serviceImpl struct {
	sensor.UnimplementedCollectorServiceServer

	collectorC <-chan common.MessageToCollectorWithAddress
}

func (s *serviceImpl) Notify(e common.SensorComponentEvent) {
	// switch e {
	// case common.SensorComponentEventCentralReachable:
	//	s.offlineMode.Store(false)
	// case common.SensorComponentEventOfflineMode:
	//	s.offlineMode.Store(true)
	//}
}

func (s *serviceImpl) Start() error {
	return nil
}

func (s *serviceImpl) Stop(_ error) {}

func (s *serviceImpl) Capabilities() []centralsensor.SensorCapability {
	return nil
}

func (s *serviceImpl) ProcessMessage(_ *central.MsgToSensor) error {
	return nil
}

func (s *serviceImpl) ResponsesC() <-chan *message.ExpiringMessage {
	return nil
}

// type connectionManager struct {
//	connectionLock sync.RWMutex
//	connectionMap  map[string]sensor.CollectorService_CommunicateServer
//}
//
// func newConnectionManager() *connectionManager {
//	return &connectionManager{
//		connectionMap: make(map[string]sensor.CollectorService_CommunicateServer),
//	}
//}
//
// func (c *connectionManager) add(node string, connection sensor.CollectorService_CommunicateServer) {
//	c.connectionLock.Lock()
//	defer c.connectionLock.Unlock()
//
//	c.connectionMap[node] = connection
//}
//
// func (c *connectionManager) remove(node string) {
//	c.connectionLock.Lock()
//	defer c.connectionLock.Unlock()
//
//	delete(c.connectionMap, node)
//}
//
// func (c *connectionManager) forEach(fn func(node string, server sensor.CollectorService_CommunicateServer)) {
//	c.connectionLock.RLock()
//	defer c.connectionLock.RUnlock()
//
//	for node, server := range c.connectionMap {
//		fn(node, server)
//	}
//}

//// GetScrapeConfig returns the scrape configuration for the given node name and scrape ID.
// func (s *serviceImpl) GetScrapeConfig(_ context.Context, nodeName string) (*sensor.MsgToCollector_ScrapeConfig, error) {
//	nodeScrapeConfig, err := s.orchestrator.GetNodeScrapeConfig(nodeName)
//	if err != nil {
//		return nil, err
//	}
//
//	rt, _ := k8sutil.ParseContainerRuntimeString(nodeScrapeConfig.ContainerRuntimeVersion)
//
//	return &sensor.MsgToCollector_ScrapeConfig{
//		ContainerRuntime: rt,
//		IsMasterNode:     nodeScrapeConfig.IsMasterNode,
//	}, nil
//}

// func (s *serviceImpl) startSendingLoop() {
//	//	for msg := range s.collectorC {
//	//		if msg.Broadcast {
//	//			s.connectionManager.forEach(func(node string, server sensor.CollectorService_CommunicateServer) {
//	//				err := server.Send(msg.Msg)
//	//				if err != nil {
//	//					log.Errorf("Error sending broadcast MessageToCollectorWithAddress to node %q: %v", node, err)
//	//					return
//	//				}
//	//			})
//	//		} else {
//	//			con, ok := s.connectionManager.connectionMap[msg.Hostname]
//	//			if !ok {
//	//				log.Errorf("Unable to find connection to compliance: %q", msg.Hostname)
//	//				return
//	//			}
//	//			err := con.Send(msg.Msg)
//	//			if err != nil {
//	//				log.Errorf("Error sending MessageToCollectorWithAddress to node %q: %v", msg.Hostname, err)
//	//				return
//	//			}
//	//		}
//	//	}
//}

// func (s *serviceImpl) RunScrape(msg *sensor.MsgToCollector) int {
//	var count int
//
//	s.connectionManager.forEach(func(node string, server sensor.CollectorService_CommunicateServer) {
//		err := server.Send(msg)
//		if err != nil {
//			log.Errorf("Error sending compliance request to node %q: %v", node, err)
//			return
//		}
//		count++
//	})
//	return count
//}

func (s *serviceImpl) Communicate(server sensor.CollectorService_CommunicateServer) error {
	// incomingMD := metautils.ExtractIncoming(server.Context())
	// hostname := incomingMD.Get("rox-compliance-nodename")
	// if hostname == "" {
	//	return errors.New("compliance did not transmit a hostname in initial metadata")
	//}

	// log.Infof("Received connection from %q", hostname)

	// s.connectionManager.add(hostname, server)
	// defer s.connectionManager.remove(hostname)

	// conf, err := s.GetScrapeConfig(server.Context(), hostname)
	// if err != nil {
	//	log.Errorf("getting scrape config for %q: %v", hostname, err)
	//	conf = &sensor.MsgToCollector_ScrapeConfig{
	//		ContainerRuntime: storage.ContainerRuntime_UNKNOWN_CONTAINER_RUNTIME,
	//	}
	//}

	// err = server.Send(&sensor.MsgToCollector{
	//	Msg: &sensor.MsgToCollector_Config{
	//		Config: conf,
	//	},
	// })
	// if err != nil {
	//	return errors.Wrapf(err, "sending config to %q", hostname)
	//}

	//// Set up this node to start collecting audit log events if it's a master node. It may send a message if the feature is already enabled
	// if conf.GetIsMasterNode() {
	//	log.Infof("Adding node %s to list of eligible compliance nodes for audit log collection because it is on a master node", hostname)
	//	s.auditLogCollectionManager.AddEligibleCollectorNode(hostname, server)
	//	defer s.auditLogCollectionManager.RemoveEligibleCollectorNode(hostname)
	//}

	// go s.startSendingLoop()

	// for {
	//	msg, err := server.Recv()
	//	if err != nil {
	//		log.Errorf("receiving from compliance %q: %v", hostname, err)
	//		return err
	//	}
	//	switch t := msg.Msg.(type) {
	//	case *sensor.MsgFromCollector_Return:
	//		log.Infof("Received compliance return from %q", msg.GetNode())
	//		s.output <- t.Return
	//	case *sensor.MsgFromCollector_AuditEvents:
	//		// if we are offline we do not send more audit logs to the manager nor the detector.
	//		// Upon reconnection Central will sync the last state and Sensor will request to Collector to start
	//		// sending the audit logs based on that state.
	//		if s.offlineMode.Load() {
	//			continue
	//		}
	//		s.auditEvents <- t.AuditEvents
	//		s.auditLogCollectionManager.AuditMessagesChan() <- msg
	//	case *sensor.MsgFromCollector_NodeInventory:
	//		s.nodeInventories <- t.NodeInventory
	//	}
	//}
	return nil
}

// RegisterServiceServer registers this service with the given gRPC Server.
func (s *serviceImpl) RegisterServiceServer(grpcServer *grpc.Server) {
	sensor.RegisterCollectorServiceServer(grpcServer, s)
}

// RegisterServiceHandler implements the APIService interface, but the agent does not accept calls over the gRPC gateway
func (s *serviceImpl) RegisterServiceHandler(context.Context, *runtime.ServeMux, *grpc.ClientConn) error {
	return nil
}

// AuthFuncOverride specifies the auth criteria for this API.
func (s *serviceImpl) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, idcheck.CollectorOnly().Authorized(ctx, fullMethodName)
}
