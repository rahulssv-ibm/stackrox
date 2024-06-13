//go:build compliance

package tests

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/ComplianceAsCode/compliance-operator/pkg/apis"
	complianceoperatorv1 "github.com/ComplianceAsCode/compliance-operator/pkg/apis/compliance/v1alpha1"
	v1 "github.com/stackrox/rox/generated/api/v1"
	v2 "github.com/stackrox/rox/generated/api/v2"
	"github.com/stackrox/rox/pkg/retry"
	"github.com/stackrox/rox/pkg/testutils/centralgrpc"
	"github.com/stackrox/rox/pkg/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	extscheme "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/scheme"
	apiMetaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	cached "k8s.io/client-go/discovery/cached"
	"k8s.io/client-go/dynamic"
	cgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/restmapper"
	dynclient "sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	coNamespace = "openshift-compliance"
	scanConfig1 = v2.ComplianceScanConfiguration{
		ScanName: "testConfig1",
		ScanConfig: &v2.BaseComplianceScanConfigurationSettings{
			Description: "test123",
			OneTimeScan: false,
			Profiles:    []string{"ocp4-e8"},
			ScanSchedule: &v2.Schedule{
				Hour:         2,
				Minute:       0,
				IntervalType: v2.Schedule_WEEKLY,
				Interval: &v2.Schedule_DaysOfWeek_{
					DaysOfWeek: &v2.Schedule_DaysOfWeek{
						Days: []int32{4, 3},
					},
				},
			},
		},
	}
	modifiedScanConfig1 = v2.ComplianceScanConfiguration{
		ScanName: "testConfig1",
		ScanConfig: &v2.BaseComplianceScanConfigurationSettings{
			Description: "test456",
			OneTimeScan: false,
			Profiles:    []string{"rhcos4-e8"},
			ScanSchedule: &v2.Schedule{
				Hour:         2,
				Minute:       0,
				IntervalType: v2.Schedule_WEEKLY,
				Interval: &v2.Schedule_DaysOfWeek_{
					DaysOfWeek: &v2.Schedule_DaysOfWeek{
						Days: []int32{1, 5},
					},
				},
			},
		},
	}
)

func createDynamicClient(t *testing.T) dynclient.Client {
	restCfg := getConfig(t)
	k8sClient := createK8sClient(t)

	k8sScheme := runtime.NewScheme()

	err := cgoscheme.AddToScheme(k8sScheme)
	require.NoError(t, err, "error adding Kubernetes Scheme to client")

	err = extscheme.AddToScheme(k8sScheme)
	require.NoError(t, err, "error adding Kubernetes Scheme to client")

	cachedClientDiscovery := cached.NewMemCacheClient(k8sClient.Discovery())
	restMapper := restmapper.NewDeferredDiscoveryRESTMapper(cachedClientDiscovery)
	restMapper.Reset()

	client, err := dynclient.New(
		restCfg,
		dynclient.Options{
			Scheme:         k8sScheme,
			Mapper:         restMapper,
			WarningHandler: dynclient.WarningHandlerOptions{SuppressWarnings: true},
		},
	)
	require.NoError(t, err, "failed to create dynamic client")

	// Add all the Compliance Operator schemes to the client so we can use
	// it for dealing with the Compliance Operator directly.
	err = apis.AddToScheme(client.Scheme())
	require.NoError(t, err, "failed to add Compliance Operator schemes to client")
	return client
}

func getGVR(api apiMetaV1.APIResource) schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    api.Group,
		Version:  api.Version,
		Resource: api.Name,
	}
}

// AssertResourceDoesExist asserts whether the given resource exits in the cluster
func AssertResourceDoesExist(ctx context.Context, t *testing.T, client *dynamic.DynamicClient, resourceName string, namespace string, api apiMetaV1.APIResource) *unstructured.Unstructured {
	var cli dynamic.ResourceInterface
	var obj *unstructured.Unstructured
	var err error
	require.Eventually(t, func() bool {
		cli = client.Resource(getGVR(api))
		if namespace != "" {
			cli = client.Resource(getGVR(api)).Namespace(namespace)
		}
		obj, err = cli.Get(ctx, resourceName, apiMetaV1.GetOptions{})
		return err == nil
	}, 30*time.Second, 10*time.Millisecond)
	return obj
}

func waitForComplianceSuiteToComplete(t *testing.T, suiteName string, interval, timeout time.Duration) {
	client := createDynamicClient(t)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	timer := time.NewTimer(timeout)
	defer timer.Stop()

	log.Info("Waiting for ComplianceSuite to reach DONE phase")
	for {
		select {
		case <-ticker.C:
			var suite complianceoperatorv1.ComplianceSuite
			err := client.Get(context.TODO(), types.NamespacedName{Name: suiteName, Namespace: "openshift-compliance"}, &suite)
			require.NoError(t, err, "failed to get ComplianceSuite %s", suiteName)

			if suite.Status.Phase == "DONE" {
				log.Infof("ComplianceSuite %s reached DONE phase", suiteName)
				return
			}
			log.Infof("ComplianceSuite %s is in %s phase", suiteName, suite.Status.Phase)
		case <-timer.C:
			t.Fatalf("Timed out waiting for ComplianceSuite to complete")
		}
	}
}

func TestCentralSendsScanConfiguration(t *testing.T) {
	// TODO
	// connect to central
	// create config
	// send config
	// connect central to sensor
	// check that scan configuration arrived at sensors

	ctx := context.Background()

	// k8sClient := createDynamicClient(t)

	conn := centralgrpc.GRPCConnectionToCentral(t)
	service := v2.NewComplianceScanConfigurationServiceClient(conn)

	// Get the clusters
	resp := getIntegrations(t)
	assert.Len(t, resp.Integrations, 1, "failed to assert there is only a single compliance integration")

	// Get the profiles for the cluster
	// clusterID := resp.Integrations[0].GetId()

	/*
		expectedScanConfig := &central.ApplyComplianceScanConfigRequest{
			ScanRequest: &central.ApplyComplianceScanConfigRequest_UpdateScan{
				UpdateScan: &central.ApplyComplianceScanConfigRequest_UpdateScheduledScan{
					ScanSettings: &central.ApplyComplianceScanConfigRequest_BaseScanSettings{
						ScanName: "testConfig1",
						Profiles: []string{"ocp4-e8"},
					},
					Cron: "0 1 * * *",
				},
			},
		}
	*/

	req, err := service.CreateComplianceScanConfiguration(ctx, &scanConfig1)
	assert.NoError(t, err)

	query := &v2.RawQuery{Query: ""}
	scanConfigs, err := service.ListComplianceScanConfigurations(ctx, query)
	assert.NoError(t, err)
	assert.Equal(t, len(scanConfigs.GetConfigurations()), 1)

	assert.Equal(t, req, scanConfigs.GetConfigurations()[0])

	conn.Close()

	req, err = service.CreateComplianceScanConfiguration(ctx, &modifiedScanConfig1)
	assert.NoError(t, err)

	conn.Connect()

	query = &v2.RawQuery{Query: ""}
	scanConfigs, err = service.ListComplianceScanConfigurations(ctx, query)
	assert.NoError(t, err)
	assert.Equal(t, len(scanConfigs.GetConfigurations()), 1)
	assert.Equal(t, req, scanConfigs.GetConfigurations()[0])

	conn.Close()

	reqDelete := &v2.ResourceByID{
		Id: "testConfig1",
	}
	_, err = service.DeleteComplianceScanConfiguration(ctx, reqDelete)
	assert.NoError(t, err)

	conn.Connect()

	query = &v2.RawQuery{Query: ""}
	scanConfigs, err = service.ListComplianceScanConfigurations(ctx, query)
	assert.NoError(t, err)
	assert.Equal(t, len(scanConfigs.GetConfigurations()), 0)
}

// ACS API test suite for integration testing for the Compliance Operator.
func TestComplianceV2Integration(t *testing.T) {
	resp := getIntegrations(t)
	assert.Len(t, resp.Integrations, 1, "failed to assert there is only a single compliance integration")
	assert.Equal(t, resp.Integrations[0].ClusterName, "remote", "failed to find integration for cluster called \"remote\"")
	assert.Equal(t, resp.Integrations[0].Namespace, "openshift-compliance", "failed to find integration for \"openshift-compliance\" namespace")
}

func TestComplianceV2ProfileGet(t *testing.T) {
	conn := centralgrpc.GRPCConnectionToCentral(t)
	client := v2.NewComplianceProfileServiceClient(conn)

	// Get the clusters
	resp := getIntegrations(t)
	assert.Len(t, resp.Integrations, 1, "failed to assert there is only a single compliance integration")

	// Get the profiles for the cluster
	clusterID := resp.Integrations[0].ClusterId
	profileList, err := client.ListComplianceProfiles(context.TODO(), &v2.ProfilesForClusterRequest{ClusterId: clusterID})
	assert.Greater(t, len(profileList.Profiles), 0, "failed to assert the cluster has profiles")

	// Now take the ID from one of the cluster profiles to get the specific profile.
	profile, err := client.GetComplianceProfile(context.TODO(), &v2.ResourceByID{Id: profileList.Profiles[0].Id})
	if err != nil {
		t.Fatal(err)
	}
	assert.Greater(t, len(profile.Rules), 0, "failed to verify the selected profile contains any rules")
}

func TestComplianceV2ProfileGetSummaries(t *testing.T) {
	conn := centralgrpc.GRPCConnectionToCentral(t)
	client := v2.NewComplianceProfileServiceClient(conn)

	// Get the clusters
	resp := getIntegrations(t)
	assert.Len(t, resp.Integrations, 1, "failed to assert there is only a single compliance integration")

	// Get the profiles for the cluster
	clusterID := resp.Integrations[0].ClusterId
	profileSummaries, err := client.ListProfileSummaries(context.TODO(), &v2.ClustersProfileSummaryRequest{ClusterIds: []string{clusterID}})
	assert.NoError(t, err)
	assert.Greater(t, len(profileSummaries.Profiles), 0, "failed to assert the cluster has profiles")
}

// Helper to get the integrations as the cluster id is needed in many API calls
func getIntegrations(t *testing.T) *v2.ListComplianceIntegrationsResponse {
	conn := centralgrpc.GRPCConnectionToCentral(t)
	client := v2.NewComplianceIntegrationServiceClient(conn)

	q := &v2.RawQuery{Query: ""}
	resp, err := client.ListComplianceIntegrations(context.TODO(), q)
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(t, resp.Integrations, 1, "failed to assert there is only a single compliance integration")

	return resp
}

func TestComplianceV2CreateGetScanConfigurations(t *testing.T) {
	ctx := context.Background()
	conn := centralgrpc.GRPCConnectionToCentral(t)
	service := v2.NewComplianceScanConfigurationServiceClient(conn)
	serviceCluster := v1.NewClustersServiceClient(conn)
	clusters, err := serviceCluster.GetClusters(ctx, &v1.GetClustersRequest{})
	assert.NoError(t, err)
	clusterID := clusters.GetClusters()[0].GetId()
	testName := fmt.Sprintf("test-%s", uuid.NewV4().String())
	req := &v2.ComplianceScanConfiguration{
		ScanName: testName,
		Id:       "",
		Clusters: []string{clusterID},
		ScanConfig: &v2.BaseComplianceScanConfigurationSettings{
			OneTimeScan: false,
			Profiles:    []string{"rhcos4-e8"},
			Description: "test config",
			ScanSchedule: &v2.Schedule{
				IntervalType: 1,
				Hour:         15,
				Minute:       0,
				Interval: &v2.Schedule_DaysOfWeek_{
					DaysOfWeek: &v2.Schedule_DaysOfWeek{
						Days: []int32{1, 2, 3, 4, 5, 6},
					},
				},
			},
		},
	}

	resp, err := service.CreateComplianceScanConfiguration(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, req.GetScanName(), resp.GetScanName())

	query := &v2.RawQuery{Query: ""}
	scanConfigs, err := service.ListComplianceScanConfigurations(ctx, query)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(scanConfigs.GetConfigurations()), 1)
	assert.GreaterOrEqual(t, scanConfigs.TotalCount, int32(1))

	configs := scanConfigs.GetConfigurations()
	scanconfigID := getscanConfigID(testName, configs)
	defer deleteScanConfig(ctx, scanconfigID, service)

	serviceResult := v2.NewComplianceResultsServiceClient(conn)
	query = &v2.RawQuery{Query: ""}
	err = retry.WithRetry(func() error {
		results, err := serviceResult.GetComplianceScanResults(ctx, query)
		if err != nil {
			return err
		}

		resultsList := results.GetScanResults()
		for i := 0; i < len(resultsList); i++ {
			if resultsList[i].GetScanName() == testName {
				return nil
			}
		}
		return errors.New("scan result not found")
	}, retry.BetweenAttempts(func(previousAttemptNumber int) {
		time.Sleep(60 * time.Second)
	}), retry.Tries(10))
	assert.NoError(t, err)

	// Create a different scan configuration with the same profile
	duplicateTestName := fmt.Sprintf("test-%s", uuid.NewV4().String())
	duplicateProfileReq := &v2.ComplianceScanConfiguration{
		ScanName: duplicateTestName,
		Id:       "",
		Clusters: []string{clusterID},
		ScanConfig: &v2.BaseComplianceScanConfigurationSettings{
			OneTimeScan: false,
			Profiles:    []string{"rhcos4-e8"},
			Description: "test config with duplicate profile",
			ScanSchedule: &v2.Schedule{
				IntervalType: 1,
				Hour:         15,
				Minute:       0,
				Interval: &v2.Schedule_DaysOfWeek_{
					DaysOfWeek: &v2.Schedule_DaysOfWeek{
						Days: []int32{1, 2, 3, 4, 5, 6},
					},
				},
			},
		},
	}

	// Verify that the duplicate profile was not created and the error message is correct
	_, err = service.CreateComplianceScanConfiguration(ctx, duplicateProfileReq)
	assert.Contains(t, err.Error(), "already uses profile")

	query = &v2.RawQuery{Query: ""}
	scanConfigs, err = service.ListComplianceScanConfigurations(ctx, query)
	assert.NoError(t, err)
	assert.Equal(t, len(scanConfigs.GetConfigurations()), 1)

	// Create a scan configuration with invalid profiles configuration
	// contains both rhcos4-high and ocp4-e8 profiles. This is going
	// to fail validation, so we don't need to worry about running a larger
	// profile (e.g., rhcos4-high), since it won't increase test times.
	invalidProfileTestName := fmt.Sprintf("test-%s", uuid.NewV4().String())
	invalidProfileReq := &v2.ComplianceScanConfiguration{
		ScanName: invalidProfileTestName,
		Id:       "",
		Clusters: []string{clusterID},
		ScanConfig: &v2.BaseComplianceScanConfigurationSettings{
			OneTimeScan: false,
			Profiles:    []string{"rhcos4-high", "ocp4-cis-node"},
			Description: "test config with invalid profiles",
			ScanSchedule: &v2.Schedule{
				IntervalType: 1,
				Hour:         15,
				Minute:       0,
				Interval: &v2.Schedule_DaysOfWeek_{
					DaysOfWeek: &v2.Schedule_DaysOfWeek{
						Days: []int32{1, 2, 3, 4, 5, 6},
					},
				},
			},
		},
	}

	// Verify that the invalid scan configuration was not created and the error message is correct
	_, err = service.CreateComplianceScanConfiguration(ctx, invalidProfileReq)
	if err == nil {
		t.Fatal("expected error creating scan configuration with invalid profiles")
	}
	assert.Contains(t, err.Error(), "profiles must have the same product")

	query = &v2.RawQuery{Query: ""}
	scanConfigs, err = service.ListComplianceScanConfigurations(ctx, query)
	assert.NoError(t, err)
	assert.Equal(t, len(scanConfigs.GetConfigurations()), 1)
}

func TestComplianceV2DeleteComplianceScanConfigurations(t *testing.T) {
	ctx := context.Background()
	conn := centralgrpc.GRPCConnectionToCentral(t)
	service := v2.NewComplianceScanConfigurationServiceClient(conn)
	// Retrieve the results from the scan configuration once the scan is complete
	serviceCluster := v1.NewClustersServiceClient(conn)
	clusters, err := serviceCluster.GetClusters(ctx, &v1.GetClustersRequest{})
	assert.NoError(t, err)

	clusterID := clusters.GetClusters()[0].GetId()
	testName := fmt.Sprintf("test-%s", uuid.NewV4().String())
	req := &v2.ComplianceScanConfiguration{
		ScanName: testName,
		Id:       "",
		Clusters: []string{clusterID},
		ScanConfig: &v2.BaseComplianceScanConfigurationSettings{
			OneTimeScan: false,
			Profiles:    []string{"rhcos4-e8"},
			Description: "test config",
			ScanSchedule: &v2.Schedule{
				IntervalType: 1,
				Hour:         15,
				Minute:       0,
				Interval: &v2.Schedule_DaysOfWeek_{
					DaysOfWeek: &v2.Schedule_DaysOfWeek{
						Days: []int32{1, 2, 3, 4, 5, 6},
					},
				},
			},
		},
	}

	resp, err := service.CreateComplianceScanConfiguration(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, req.GetScanName(), resp.GetScanName())

	query := &v2.RawQuery{Query: ""}
	scanConfigs, err := service.ListComplianceScanConfigurations(ctx, query)
	configs := scanConfigs.GetConfigurations()
	scanconfigID := getscanConfigID(testName, configs)
	reqDelete := &v2.ResourceByID{
		Id: scanconfigID,
	}
	_, err = service.DeleteComplianceScanConfiguration(ctx, reqDelete)
	assert.NoError(t, err)

	// Verify scan configuration no longer exists
	scanConfigs, err = service.ListComplianceScanConfigurations(ctx, query)
	configs = scanConfigs.GetConfigurations()
	scanconfigID = getscanConfigID(testName, configs)
	assert.Empty(t, scanconfigID)
}

func TestComplianceV2ComplianceObjectMetadata(t *testing.T) {
	ctx := context.Background()
	conn := centralgrpc.GRPCConnectionToCentral(t)
	service := v2.NewComplianceScanConfigurationServiceClient(conn)
	serviceCluster := v1.NewClustersServiceClient(conn)
	clusters, err := serviceCluster.GetClusters(ctx, &v1.GetClustersRequest{})
	assert.NoError(t, err)
	clusterID := clusters.GetClusters()[0].GetId()
	testName := fmt.Sprintf("test-%s", uuid.NewV4().String())
	req := &v2.ComplianceScanConfiguration{
		ScanName: testName,
		Id:       "",
		Clusters: []string{clusterID},
		ScanConfig: &v2.BaseComplianceScanConfigurationSettings{
			OneTimeScan: false,
			Profiles:    []string{"rhcos4-e8"},
			Description: "test config",
			ScanSchedule: &v2.Schedule{
				IntervalType: 1,
				Hour:         15,
				Minute:       0,
				Interval: &v2.Schedule_DaysOfWeek_{
					DaysOfWeek: &v2.Schedule_DaysOfWeek{
						Days: []int32{1, 2, 3, 4, 5, 6},
					},
				},
			},
		},
	}

	resp, err := service.CreateComplianceScanConfiguration(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, req.GetScanName(), resp.GetScanName())

	query := &v2.RawQuery{Query: ""}
	scanConfigs, err := service.ListComplianceScanConfigurations(ctx, query)
	configs := scanConfigs.GetConfigurations()
	scanconfigID := getscanConfigID(testName, configs)
	defer deleteScanConfig(ctx, scanconfigID, service)

	// Ensure the ScanSetting and ScanSettingBinding have ACS metadata
	client := createDynamicClient(t)
	var scanSetting complianceoperatorv1.ScanSetting
	err = client.Get(context.TODO(), types.NamespacedName{Name: testName, Namespace: "openshift-compliance"}, &scanSetting)
	require.NoError(t, err, "failed to get ScanSetting %s", testName)

	assert.Contains(t, scanSetting.Labels, "app.kubernetes.io/name")
	assert.Equal(t, scanSetting.Labels["app.kubernetes.io/name"], "stackrox")
	assert.Contains(t, scanSetting.Annotations, "owner")
	assert.Equal(t, scanSetting.Annotations["owner"], "stackrox")

	var scanSettingBinding complianceoperatorv1.ScanSetting
	err = client.Get(context.TODO(), types.NamespacedName{Name: testName, Namespace: "openshift-compliance"}, &scanSettingBinding)
	require.NoError(t, err, "failed to get ScanSettingBinding %s", testName)
	assert.Contains(t, scanSettingBinding.Labels, "app.kubernetes.io/name")
	assert.Equal(t, scanSettingBinding.Labels["app.kubernetes.io/name"], "stackrox")
	assert.Contains(t, scanSettingBinding.Annotations, "owner")
	assert.Equal(t, scanSettingBinding.Annotations["owner"], "stackrox")
}

func deleteScanConfig(ctx context.Context, scanID string, service v2.ComplianceScanConfigurationServiceClient) error {
	req := &v2.ResourceByID{
		Id: scanID,
	}
	_, err := service.DeleteComplianceScanConfiguration(ctx, req)
	return err
}

func getscanConfigID(configName string, scanConfigs []*v2.ComplianceScanConfigurationStatus) string {
	configID := ""
	for i := 0; i < len(scanConfigs); i++ {
		if scanConfigs[i].GetScanName() == configName {
			configID = scanConfigs[i].GetId()
		}

	}
	return configID
}

func TestComplianceV2ScheduleRescan(t *testing.T) {
	conn := centralgrpc.GRPCConnectionToCentral(t)
	client := v2.NewComplianceScanConfigurationServiceClient(conn)
	integrationClient := v2.NewComplianceIntegrationServiceClient(conn)
	resp, err := integrationClient.ListComplianceIntegrations(context.TODO(), &v2.RawQuery{Query: ""})
	if err != nil {
		t.Fatal(err)
	}
	clusterId := resp.Integrations[0].ClusterId

	scanConfigName := "e8-scan-schedule"
	sc := v2.ComplianceScanConfiguration{
		ScanName: scanConfigName,
		ScanConfig: &v2.BaseComplianceScanConfigurationSettings{
			OneTimeScan: false,
			Profiles:    []string{"ocp4-e8"},
			ScanSchedule: &v2.Schedule{
				IntervalType: 3,
				Hour:         0,
				Minute:       0,
			},
			Description: "Scan schedule for the Austrailian Essential Eight profile to run daily.",
		},
		Clusters: []string{clusterId},
	}
	scanConfig, err := client.CreateComplianceScanConfiguration(context.TODO(), &sc)
	if err != nil {
		t.Fatal(err)
	}

	defer client.DeleteComplianceScanConfiguration(context.TODO(), &v2.ResourceByID{Id: scanConfig.GetId()})

	waitForComplianceSuiteToComplete(t, scanConfig.ScanName, 2*time.Second, 5*time.Minute)

	// Invoke a rescan
	_, err = client.RunComplianceScanConfiguration(context.TODO(), &v2.ResourceByID{Id: scanConfig.GetId()})
	require.NoError(t, err, "failed to rerun scan schedule %s", scanConfigName)

	// Assert the scan is rerunning on the cluster using the Compliance Operator CRDs
	waitForComplianceSuiteToComplete(t, scanConfig.ScanName, 2*time.Second, 5*time.Minute)
}
