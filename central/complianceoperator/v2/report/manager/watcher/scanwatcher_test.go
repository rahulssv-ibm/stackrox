package watcher

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/pkg/errors"
	coIntegrationMocks "github.com/stackrox/rox/central/complianceoperator/v2/integration/datastore/mocks"
	"github.com/stackrox/rox/central/complianceoperator/v2/scans/datastore/mocks"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/concurrency"
	"github.com/stackrox/rox/pkg/protocompat"
	"github.com/stackrox/rox/pkg/queue"
	"github.com/stackrox/rox/pkg/set"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

type testEvent func(*testing.T, ScanWatcher)

func handleScan(id string) func(*testing.T, ScanWatcher) {
	return func(t *testing.T, scanWatcher ScanWatcher) {
		err := scanWatcher.PushScan(&storage.ComplianceOperatorScanV2{
			Id: id,
		})
		require.NoError(t, err)
	}
}

func handleScanWithAnnotation(id, checkCount string) func(*testing.T, ScanWatcher) {
	return func(t *testing.T, scanWatcher ScanWatcher) {
		err := scanWatcher.PushScan(&storage.ComplianceOperatorScanV2{
			Id:          id,
			Annotations: map[string]string{checkCountAnnotationKey: checkCount},
		})
		require.NoError(t, err)
	}
}

func handleResult(id string) func(*testing.T, ScanWatcher) {
	return func(t *testing.T, scanWatcher ScanWatcher) {
		err := scanWatcher.PushCheckResult(&storage.ComplianceOperatorCheckResultV2{
			CheckId: id,
		})
		require.NoError(t, err)
	}
}

func TestScanWatcher(t *testing.T) {
	cases := map[string]struct {
		events          []testEvent
		assertScanID    string
		assertResultIDs []string
	}{
		"scan ready -> result -> result": {
			events: []testEvent{
				handleScanWithAnnotation("id-1", "2"),
				handleResult("id-1"),
				handleResult("id-2"),
			},
			assertScanID:    "id-1",
			assertResultIDs: []string{"id-1", "id-2"},
		},
		"scan -> result -> result -> scan ready": {
			events: []testEvent{
				handleScan("id-1"),
				handleResult("id-1"),
				handleResult("id-2"),
				handleScanWithAnnotation("id-1", "2"),
			},
			assertScanID:    "id-1",
			assertResultIDs: []string{"id-1", "id-2"},
		},
		"scan -> result -> scan ready -> result": {
			events: []testEvent{
				handleScan("id-1"),
				handleResult("id-1"),
				handleScanWithAnnotation("id-1", "2"),
				handleResult("id-2"),
			},
			assertScanID:    "id-1",
			assertResultIDs: []string{"id-1", "id-2"},
		},
		"result -> result -> scan ready": {
			events: []testEvent{
				handleResult("id-1"),
				handleResult("id-2"),
				handleScanWithAnnotation("id-1", "2"),
			},
			assertScanID:    "id-1",
			assertResultIDs: []string{"id-1", "id-2"},
		},
	}
	for tName, tCase := range cases {
		t.Run(tName, func(t *testing.T) {
			watcherID := "id"
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			resultQueue := queue.NewQueue[*ScanWatcherResults]()
			scanWatcher := NewScanWatcher(ctx, watcherID, resultQueue)
			for _, event := range tCase.events {
				event(t, scanWatcher)
			}
			require.Eventually(t, func() bool {
				return resultQueue.Len() > 0
			}, 200*time.Millisecond, 10*time.Millisecond)
			result := resultQueue.Pull()
			require.NotNil(t, result)
			assert.Equal(t, tCase.assertScanID, result.Scan.GetId())
			for _, checkID := range tCase.assertResultIDs {
				found := false
				for checkResult := range result.CheckResults {
					if checkID == checkResult {
						found = true
						break
					}
				}
				assert.Equal(t, true, found)
			}
		})
	}
}

func TestScanWatcherCancel(t *testing.T) {
	watcherID := "id"
	ctx, cancel := context.WithCancel(context.Background())
	resultQueue := queue.NewQueue[*ScanWatcherResults]()
	scanWatcher := NewScanWatcher(ctx, watcherID, resultQueue)
	handleScan("id-1")(t, scanWatcher)
	handleResult("id-1")(t, scanWatcher)
	cancel()
	select {
	case <-scanWatcher.Finished().Done():
	case <-time.After(100 * time.Millisecond):
		t.Error("timeout waiting for the watcher to stop")
	}
	assert.Equal(t, 0, resultQueue.Len())
}

func TestScanWatcherStop(t *testing.T) {
	watcherID := "id"
	resultQueue := queue.NewQueue[*ScanWatcherResults]()
	scanWatcher := NewScanWatcher(context.Background(), watcherID, resultQueue)
	handleScan("id-1")(t, scanWatcher)
	handleResult("id-1")(t, scanWatcher)
	scanWatcher.Stop()
	select {
	case <-scanWatcher.Finished().Done():
	case <-time.After(100 * time.Millisecond):
		t.Error("timeout waiting for the watcher to stop")
	}
	assert.Equal(t, 0, resultQueue.Len())
}

func TestScanWatcherTimeout(t *testing.T) {
	resultQueue := queue.NewQueue[*ScanWatcherResults]()
	ctx, cancel := context.WithCancel(context.Background())
	finishedSignal := concurrency.NewSignal()
	scanWatcher := &scanWatcherImpl{
		ctx:          ctx,
		cancel:       cancel,
		watcherID:    "id",
		scanC:        make(chan *storage.ComplianceOperatorScanV2),
		resultC:      make(chan *storage.ComplianceOperatorCheckResultV2),
		stopped:      &finishedSignal,
		readyQueue:   resultQueue,
		checkResults: set.NewStringSet(),
	}
	timeoutC := make(chan time.Time)
	go scanWatcher.run(timeoutC)
	handleScan("id-1")(t, scanWatcher)
	handleResult("id-1")(t, scanWatcher)
	// We signal the timeout
	timeoutC <- time.Now()
	select {
	case <-scanWatcher.Finished().Done():
	case <-time.After(100 * time.Millisecond):
		t.Error("timeout waiting for the watcher to stop")
	}
	// We should have a result in the queue with an error
	require.Equal(t, 1, resultQueue.Len())
	result := resultQueue.Pull()
	assert.Error(t, result.Error)
}

func TestGetIDFromScan(t *testing.T) {
	scan := &storage.ComplianceOperatorScanV2{}
	_, err := GetWatcherIDFromScan(scan)
	assert.Error(t, err)
	scan.ClusterId = "cluster-1"
	_, err = GetWatcherIDFromScan(scan)
	assert.Error(t, err)
	scan.Id = "scan-1"
	id, err := GetWatcherIDFromScan(scan)
	assert.NoError(t, err)
	assert.Equal(t, "", id)
	scan.LastStartedTime = protocompat.TimestampNow()
	id, err = GetWatcherIDFromScan(scan)
	assert.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("%s:%s:%s", scan.ClusterId, scan.Id, scan.LastStartedTime), id)
}

func TestGetIDFromResult(t *testing.T) {
	timeNow := protocompat.TimestampNow()
	ctrl := gomock.NewController(t)
	ds := mocks.NewMockDataStore(ctrl)
	// Error querying the Scan DataStore
	ds.EXPECT().SearchScans(gomock.Any(), gomock.Any()).Times(1).
		DoAndReturn(func(_, _ any) ([]*storage.ComplianceOperatorScanV2, error) {
			return nil, errors.New("db error")
		})
	result := &storage.ComplianceOperatorCheckResultV2{}
	_, err := GetWatcherIDFromCheckResult(result, ds)
	assert.Error(t, err)
	// No Scan retrieved
	ds.EXPECT().SearchScans(gomock.Any(), gomock.Any()).Times(1).
		DoAndReturn(func(_, _ any) ([]*storage.ComplianceOperatorScanV2, error) {
			return nil, nil
		})
	_, err = GetWatcherIDFromCheckResult(result, ds)
	assert.Error(t, err)
	// Scan retrieved successfully
	ds.EXPECT().SearchScans(gomock.Any(), gomock.Any()).Times(4).
		DoAndReturn(func(_, _ any) ([]*storage.ComplianceOperatorScanV2, error) {
			return []*storage.ComplianceOperatorScanV2{
				{
					ClusterId:       "cluster-1",
					Id:              "scan-1",
					LastStartedTime: timeNow,
				},
			}, nil
		})
	// Empty annotation
	_, err = GetWatcherIDFromCheckResult(result, ds)
	assert.Error(t, err)
	// Invalid format in the annotation
	result.Annotations = map[string]string{
		lastScannedAnnotationKey: protocompat.TimestampNow().String(),
	}
	// The timestamp does not coincide with the Scan start time
	_, err = GetWatcherIDFromCheckResult(result, ds)
	assert.Error(t, err)
	result.Annotations = map[string]string{
		lastScannedAnnotationKey: protocompat.TimestampNow().AsTime().Format(time.RFC3339Nano),
	}
	// Successful execution
	_, err = GetWatcherIDFromCheckResult(result, ds)
	assert.Error(t, err)
	result.Annotations = map[string]string{
		lastScannedAnnotationKey: timeNow.AsTime().Format(time.RFC3339Nano),
	}
	id, err := GetWatcherIDFromCheckResult(result, ds)
	assert.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("cluster-1:scan-1:%s", timeNow.String()), id)
}

func TestIsComplianceOperatorHealthy(t *testing.T) {
	clusterID := "cluster-id"
	ctrl := gomock.NewController(t)
	ds := coIntegrationMocks.NewMockDataStore(ctrl)

	// DataStore error
	ds.EXPECT().GetComplianceIntegrationByCluster(gomock.Any(), gomock.Any()).Times(1).
		DoAndReturn(func(_, _ any) ([]*storage.ComplianceIntegration, error) {
			return []*storage.ComplianceIntegration{}, ComplianceOperatorIntegrationDataStoreError
		})
	assert.Error(t, IsComplianceOperatorHealthy(clusterID, ds))

	// No integrations retrieved
	ds.EXPECT().GetComplianceIntegrationByCluster(gomock.Any(), gomock.Any()).Times(1).
		DoAndReturn(func(_, _ any) ([]*storage.ComplianceIntegration, error) {
			return []*storage.ComplianceIntegration{}, ComplianceOperatorIntegrationZeroIngerationsError
		})
	assert.Error(t, IsComplianceOperatorHealthy(clusterID, ds))

	// Compliance Operator not installed
	ds.EXPECT().GetComplianceIntegrationByCluster(gomock.Any(), gomock.Any()).Times(1).
		DoAndReturn(func(_, _ any) ([]*storage.ComplianceIntegration, error) {
			return []*storage.ComplianceIntegration{
				{
					OperatorInstalled: false,
				},
			}, nil
		})
	err := IsComplianceOperatorHealthy(clusterID, ds)
	assert.Error(t, err)
	assert.Error(t, ComplianceOperatorNotInstalledError, err)

	// Minimum version error
	ds.EXPECT().GetComplianceIntegrationByCluster(gomock.Any(), gomock.Any()).Times(1).
		DoAndReturn(func(_, _ any) ([]*storage.ComplianceIntegration, error) {
			return []*storage.ComplianceIntegration{
				{
					OperatorInstalled: true,
					Version:           "v1.5.0",
				},
			}, nil
		})
	err = IsComplianceOperatorHealthy(clusterID, ds)
	assert.Error(t, err)
	assert.Equal(t, ComplianceOperatorVersionError, err)

	// Compliance Operator is healthy
	ds.EXPECT().GetComplianceIntegrationByCluster(gomock.Any(), gomock.Any()).Times(1).
		DoAndReturn(func(_, _ any) ([]*storage.ComplianceIntegration, error) {
			return []*storage.ComplianceIntegration{
				{
					OperatorInstalled: true,
					Version:           "v1.6.0",
				},
			}, nil
		})
	assert.NoError(t, IsComplianceOperatorHealthy(clusterID, ds))
}
