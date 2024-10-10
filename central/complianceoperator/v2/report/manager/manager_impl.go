package manager

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/pkg/errors"
	profileDatastore "github.com/stackrox/rox/central/complianceoperator/v2/profiles/datastore"
	snapshotDS "github.com/stackrox/rox/central/complianceoperator/v2/report/datastore"
	reportGen "github.com/stackrox/rox/central/complianceoperator/v2/report/manager/complianceReportgenerator"
	"github.com/stackrox/rox/central/complianceoperator/v2/report/manager/watcher"
	scanConfigurationDS "github.com/stackrox/rox/central/complianceoperator/v2/scanconfigurations/datastore"
	scanDS "github.com/stackrox/rox/central/complianceoperator/v2/scans/datastore"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/concurrency"
	"github.com/stackrox/rox/pkg/env"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/protocompat"
	"github.com/stackrox/rox/pkg/queue"
	"github.com/stackrox/rox/pkg/sync"
	"github.com/stackrox/rox/pkg/uuid"
	"golang.org/x/sync/semaphore"
)

var (
	log         = logging.LoggerForModule()
	maxRequests = 100
)

type reportRequest struct {
	scanConfig *storage.ComplianceOperatorScanConfigurationV2
	ctx        context.Context
	snapshotID string
}

type managerImpl struct {
	datastore         scanConfigurationDS.DataStore
	scanDataStore     scanDS.DataStore
	profileDataStore  profileDatastore.DataStore
	snapshotDataStore snapshotDS.DataStore

	runningReportConfigs map[string]*reportRequest
	// channel for report job requests
	reportRequests chan *reportRequest
	stopper        concurrency.Stopper
	// isStarted will make sure only one start routine runs for an instance of manager
	isStarted atomic.Bool

	// isStopped will prevent manager from being re-started once it is stopped
	isStopped atomic.Bool

	// Mutex to synchronize access to runningReportConfigs map
	mu             sync.Mutex
	concurrencySem *semaphore.Weighted
	reportGen      reportGen.ComplianceReportGenerator

	watchingScansLock sync.Mutex
	// watchingScans a map holding the ScanWatchers
	watchingScans map[string]watcher.ScanWatcher
	// readyQueue holds the scan that are ready to be reported
	readyQueue *queue.Queue[*watcher.ScanWatcherResults]

	watchingScanConfigsLock sync.Mutex
	// watchingScanConfigs a map holding the ScanConfigWatchers
	watchingScanConfigs map[string]watcher.ScanConfigWatcher
	// scanConfigReadyQueue holds the scan configurations that are ready to be reported
	scanConfigReadyQueue *queue.Queue[*watcher.ScanConfigWatcherResults]
}

func New(scanConfigDS scanConfigurationDS.DataStore, scanDataStore scanDS.DataStore, profileDataStore profileDatastore.DataStore, snapshotDatastore snapshotDS.DataStore, reportGen reportGen.ComplianceReportGenerator) Manager {
	return &managerImpl{
		datastore:            scanConfigDS,
		scanDataStore:        scanDataStore,
		profileDataStore:     profileDataStore,
		snapshotDataStore:    snapshotDatastore,
		stopper:              concurrency.NewStopper(),
		runningReportConfigs: make(map[string]*reportRequest, maxRequests),
		reportRequests:       make(chan *reportRequest, maxRequests),
		concurrencySem:       semaphore.NewWeighted(int64(env.ReportExecutionMaxConcurrency.IntegerSetting())),
		reportGen:            reportGen,
		watchingScans:        make(map[string]watcher.ScanWatcher),
		readyQueue:           queue.NewQueue[*watcher.ScanWatcherResults](),
		watchingScanConfigs:  make(map[string]watcher.ScanConfigWatcher),
		scanConfigReadyQueue: queue.NewQueue[*watcher.ScanConfigWatcherResults](),
	}
}

func (m *managerImpl) SubmitReportRequest(ctx context.Context, scanConfig *storage.ComplianceOperatorScanConfigurationV2) error {
	if !features.ComplianceReporting.Enabled() {
		return nil
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.runningReportConfigs[scanConfig.GetId()]; ok {
		return errors.New(fmt.Sprintf("Report request for scan configuration %q already in process", scanConfig.GetScanConfigName()))
	}

	req := &reportRequest{
		scanConfig: scanConfig,
		ctx:        context.WithoutCancel(ctx),
	}
	log.Infof("Submitting report for scan config %s at %v for execution with req %v.", scanConfig.GetScanConfigName(), time.Now().Format(time.RFC822), *req)
	select {
	case m.reportRequests <- req:
		m.runningReportConfigs[scanConfig.GetId()] = req
	default:
		return errors.New(fmt.Sprintf("Error submitting report request for for scan configuration %q, request limit reached",
			scanConfig.GetScanConfigName()))
	}

	return nil
}

func (m *managerImpl) Start() {
	if m.isStopped.Load() {
		log.Error("Compliance report manager already stopped. It cannot be re-started once stopped.")
		return
	}
	swapped := m.isStarted.CompareAndSwap(false, true)
	if !swapped {
		log.Error("Compliance report manager already running")
		return
	}
	log.Info("Starting compliance report manager")
	go m.runReports()
	go m.handleReadyScan()
	go m.handleReadyScanConfig()
}

func (m *managerImpl) Stop() {
	if m.isStarted.Load() {
		log.Error("Compliance report manager not started")
		return
	}
	swapped := m.isStopped.CompareAndSwap(false, true)
	if !swapped {
		log.Error("Compliance report manager already stopped")
		return
	}
	logging.Info("Stopping compliance report manager")
	concurrency.WithLock(&m.watchingScansLock, func() {
		for _, scanWatcher := range m.watchingScans {
			scanWatcher.Stop()
		}
	})
	concurrency.WithLock(&m.watchingScanConfigsLock, func() {
		for _, scanConfigWatcher := range m.watchingScanConfigs {
			scanConfigWatcher.Stop()
		}
	})
	m.stopper.Client().Stop()
	err := m.stopper.Client().Stopped().Wait()
	if err != nil {
		logging.Errorf("Error stopping compliance report manager : %v", err)
	}
}

func (m *managerImpl) generateReport(req *reportRequest) {
	defer m.concurrencySem.Release(1)

	clusterIds := []string{}
	profiles := []string{}
	for _, cluster := range req.scanConfig.GetClusters() {
		clusterIds = append(clusterIds, cluster.GetClusterId())
	}

	for _, profile := range req.scanConfig.GetProfiles() {
		profiles = append(profiles, profile.GetProfileName())
	}

	repRequest := &reportGen.ComplianceReportRequest{
		ScanConfigName: req.scanConfig.GetScanConfigName(),
		ScanConfigID:   req.scanConfig.GetId(),
		Profiles:       profiles,
		ClusterIDs:     clusterIds,
		Notifiers:      req.scanConfig.GetNotifiers(),
		Ctx:            req.ctx,
	}
	m.reportGen.ProcessReportRequest(repRequest)
	logging.Infof("Executing report request for scan config %q", req.scanConfig.GetId())

	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.runningReportConfigs, req.scanConfig.GetId())

}

func (m *managerImpl) runReports() {
	defer m.stopper.Flow().ReportStopped()
	for {
		select {
		case <-m.stopper.Flow().StopRequested():
			logging.Info("Signal received to stop compliance report manager")
			return
		case req := <-m.reportRequests:
			if err := m.concurrencySem.Acquire(context.Background(), 1); err != nil {
				log.Errorf("Error acquiring semaphore to run new report: %v", err)
				continue
			}
			logging.Infof("Executing report %q at %v", req.scanConfig.GetId(), time.Now().Format(time.RFC822))
			var w watcher.ScanConfigWatcher
			concurrency.WithLock(&m.watchingScanConfigsLock, func() {
				w = m.watchingScanConfigs[req.scanConfig.GetId()]
			})
			reportType := storage.ComplianceOperatorReportStatus_ON_DEMAND
			notificationMethod := storage.ComplianceOperatorReportStatus_EMAIL
			snapshot := &storage.ComplianceOperatorReportSnapshotV2{
				ReportId:            uuid.NewV4().String(),
				ScanConfigurationId: req.scanConfig.GetId(),
				Name:                req.scanConfig.GetScanConfigName(),
				Description:         req.scanConfig.GetDescription(),
				ReportStatus: &storage.ComplianceOperatorReportStatus{
					RunState:                 storage.ComplianceOperatorReportStatus_WAITING,
					StartedAt:                nil,
					ReportRequestType:        reportType,
					ReportNotificationMethod: notificationMethod,
				},
				User: req.scanConfig.GetModifiedBy(),
			}
			if w == nil {
				log.Infof("No watcher running")
				snapshot.ReportStatus.RunState = storage.ComplianceOperatorReportStatus_PREPARING
				log.Infof("Creating snapshot %s", snapshot.GetReportId())
				if err := m.snapshotDataStore.UpsertSnapshot(req.ctx, snapshot); err != nil {
					log.Errorf("Unable to create snapshot: ", err)
					continue
				}
				req.snapshotID = snapshot.GetReportId()
				go m.generateReport(req)
				continue
			}
			log.Infof("Creating snapshot %s", snapshot.GetReportId())
			if err := m.snapshotDataStore.UpsertSnapshot(req.ctx, snapshot); err != nil {
				log.Errorf("Unable to create snapshot: ", err)
				continue
			}
			w.Subscribe(snapshot.GetReportId())
		}
	}
}

// HandleScan starts a new ScanWatcher if needed and pushes the scan to it
func (m *managerImpl) HandleScan(ctx context.Context, scan *storage.ComplianceOperatorScanV2) error {
	if !features.ComplianceReporting.Enabled() {
		return nil
	}
	id, err := watcher.GetWatcherIDFromScan(scan)
	if err != nil {
		return err
	}
	if id == "" {
		return nil
	}
	concurrency.WithLock(&m.watchingScansLock, func() {
		var scanWatcher watcher.ScanWatcher
		var found bool
		if scanWatcher, found = m.watchingScans[id]; !found {
			scanWatcher = watcher.NewScanWatcher(ctx, id, m.readyQueue)
			m.watchingScans[id] = scanWatcher
		}
		err = scanWatcher.PushScan(scan)
	})

	if err != nil {
		return err
	}

	return nil
}

// HandleResult starts a new ScanWatcher if needed and pushes the checkResult to it
func (m *managerImpl) HandleResult(ctx context.Context, result *storage.ComplianceOperatorCheckResultV2) error {
	if !features.ComplianceReporting.Enabled() {
		return nil
	}
	id, err := watcher.GetWatcherIDFromCheckResult(ctx, result, m.scanDataStore)
	if err != nil {
		return err
	}
	if id == "" {
		return nil
	}
	concurrency.WithLock(&m.watchingScansLock, func() {
		var scanWatcher watcher.ScanWatcher
		var found bool
		if scanWatcher, found = m.watchingScans[id]; !found {
			scanWatcher = watcher.NewScanWatcher(ctx, id, m.readyQueue)
			m.watchingScans[id] = scanWatcher
		}
		err = scanWatcher.PushCheckResult(result)
	})
	return err
}

// handleReadyScan pulls scans that are ready to be reported
func (m *managerImpl) handleReadyScan() {
	if !features.ComplianceReporting.Enabled() {
		return
	}
	for {
		select {
		case <-m.stopper.Flow().StopRequested():
			return
		default:
			if scanResult := m.readyQueue.PullBlocking(m.stopper.LowLevel().GetStopRequestSignal()); scanResult != nil {
				log.Infof("Scan %s done with %d checks", scanResult.Scan.GetScanName(), len(scanResult.CheckResults))
				w, err := m.getScanConfigWatcher(scanResult.Ctx, scanResult.Scan, m.datastore, m.scanConfigReadyQueue, false, false)
				if err != nil {
					log.Errorf("Unable to create the ScanConfigWatcher: %v", err)
					continue
				}
				if err := w.PushScanResults(scanResult); err != nil {
					log.Errorf("Unable to push scan %s: %v", scanResult.Scan.GetScanName(), err)
				}
				concurrency.WithLock(&m.watchingScansLock, func() {
					delete(m.watchingScans, scanResult.WatcherID)
				})
			}
		}
	}
}

// getScanConfigWatcher returns the ScanConfigWatcher of a given scan
func (m *managerImpl) getScanConfigWatcher(ctx context.Context, scan *storage.ComplianceOperatorScanV2, ds scanConfigurationDS.DataStore, queue *queue.Queue[*watcher.ScanConfigWatcherResults], onDemand bool, reportDownload bool) (watcher.ScanConfigWatcher, error) {
	sc, err := watcher.GetScanConfigFromScan(ctx, scan, ds)
	if err != nil {
		return nil, errors.Errorf("Unable to get scan config id: %v", err)
	}
	if sc == nil {
		return nil, errors.Errorf("ScanConfiguration not found for scan %s", scan.GetScanName())
	}
	var w watcher.ScanConfigWatcher
	var ok bool
	concurrency.WithLock(&m.watchingScanConfigsLock, func() {
		if w, ok = m.watchingScanConfigs[sc.GetId()]; !ok {
			log.Infof("Staring config watcher %s", sc.GetId())
			w = watcher.NewScanConfigWatcher(ctx, sc.GetId(), sc, m.scanDataStore, m.profileDataStore, queue)
			m.watchingScanConfigs[sc.GetId()] = w
		}
	})
	if !ok {
		reportType := storage.ComplianceOperatorReportStatus_SCHEDULED
		if onDemand {
			reportType = storage.ComplianceOperatorReportStatus_ON_DEMAND
		}
		notificationMethod := storage.ComplianceOperatorReportStatus_EMAIL
		if reportDownload {
			notificationMethod = storage.ComplianceOperatorReportStatus_DOWNLOAD
		}
		snapshot := &storage.ComplianceOperatorReportSnapshotV2{
			ReportId:            uuid.NewV4().String(),
			ScanConfigurationId: sc.GetId(),
			Name:                sc.GetScanConfigName(),
			Description:         sc.GetDescription(),
			ReportStatus: &storage.ComplianceOperatorReportStatus{
				RunState:                 storage.ComplianceOperatorReportStatus_WAITING,
				StartedAt:                scan.GetLastStartedTime(),
				ReportRequestType:        reportType,
				ReportNotificationMethod: notificationMethod,
			},
			User: sc.GetModifiedBy(),
		}
		log.Infof("Creating snapshot %s", snapshot.GetReportId())
		if err := m.snapshotDataStore.UpsertSnapshot(ctx, snapshot); err != nil {
			return nil, err
		}
		w.Subscribe(snapshot.GetReportId())
	}
	return w, nil
}

// handleReadyScanConfig pulls scan configs that are ready to be reported
func (m *managerImpl) handleReadyScanConfig() {
	if !features.ComplianceReporting.Enabled() {
		return
	}
	for {
		select {
		case <-m.stopper.Flow().StopRequested():
			return
		default:
			if result := m.scanConfigReadyQueue.PullBlocking(m.stopper.LowLevel().GetStopRequestSignal()); result != nil {
				log.Infof("Scan Config %s done with %d scans and %d reports", result.ScanConfig.GetScanConfigName(), len(result.ScanResults), len(result.ReportSnapshotIDs))
				concurrency.WithLock(&m.watchingScanConfigsLock, func() {
					delete(m.watchingScanConfigs, result.WatcherID)
				})
				for _, id := range result.ReportSnapshotIDs {
					snapshot, found, err := m.snapshotDataStore.GetSnapshot(result.Ctx, id)
					if err != nil {
						log.Errorf("Unable to retrieve snapshot %s: %v", id, err)
						continue
					}
					if !found {
						log.Errorf("Snapshot %s not found", id)
						continue
					}
					snapshot.ReportStatus.RunState = storage.ComplianceOperatorReportStatus_PREPARING
					snapshot.ReportStatus.CompletedAt = protocompat.TimestampNow()
					log.Infof("Updating snapshot %s", snapshot.GetReportId())
					if err := m.snapshotDataStore.UpsertSnapshot(result.Ctx, snapshot); err != nil {
						log.Errorf("Unable to update the snapshot %s: %v", snapshot.GetReportId(), err)
						continue
					}
					generateReportReq := &reportRequest{
						ctx:        result.Ctx,
						scanConfig: result.ScanConfig,
						snapshotID: snapshot.GetReportId(),
					}
					go m.generateReport(generateReportReq)
				}
			}
		}
	}
}
