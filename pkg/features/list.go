package features

//lint:file-ignore U1000 we want to introduce this feature flag unused.

var (
	// SourcedAutogeneratedIntegrations enables adding a "source" to autogenerated integrations.
	// TODO(ROX-20353): if this is enabled by default, make sure to address sensor reconciliation of ImageIntegration resources.
	SourcedAutogeneratedIntegrations = registerUnchangeableFeature("Enable autogenerated integrations with cluster/namespace/secret source", "ROX_SOURCED_AUTOGENERATED_INTEGRATIONS", false)

	// VulnMgmtWorkloadCVEs enables APIs and UI pages for the VM Workload CVE enhancements
	VulnMgmtWorkloadCVEs = registerUnchangeableFeature("Vuln Mgmt Workload CVEs", "ROX_VULN_MGMT_WORKLOAD_CVES", true)

	// StoreEventHashes stores the hashes of successfully processed objects we receive from Sensor into the database
	StoreEventHashes = registerUnchangeableFeature("Store Event Hashes", "ROX_STORE_EVENT_HASHES", true)

	// PreventSensorRestartOnDisconnect enables a new behavior in Sensor where it avoids restarting when the gRPC connection with Central ends.
	PreventSensorRestartOnDisconnect = registerUnchangeableFeature("Prevent Sensor restart on disconnect", "ROX_PREVENT_SENSOR_RESTART_ON_DISCONNECT", true)

	// MoveInitBundlesUI is front-end only move from integrations to clusters route.
	MoveInitBundlesUI = registerUnchangeableFeature("Move init-bundles UI", "ROX_MOVE_INIT_BUNDLES_UI", true)

	// ComplianceEnhancements enables APIs and UI pages for Compliance 2.0
	ComplianceEnhancements = registerFeature("Compliance enhancements", "ROX_COMPLIANCE_ENHANCEMENTS", true)

	// AdministrationEvents enables APIs (including collection) and UI pages for administration events.
	AdministrationEvents = registerFeature("Enable administration events", "ROX_ADMINISTRATION_EVENTS", true)

	// PostgresDatastore defines if PostgresSQL should be used
	PostgresDatastore = registerUnchangeableFeature("Enable Postgres Datastore", "ROX_POSTGRES_DATASTORE", true)

	// ActiveVulnMgmt defines if the active vuln mgmt feature is enabled
	ActiveVulnMgmt = registerFeature("Enable Active Vulnerability Management", "ROX_ACTIVE_VULN_MGMT", false)

	// VulnReportingEnhancements enables APIs and UI pages for VM Reporting enhancements including downloadable reports
	VulnReportingEnhancements = registerFeature("Enable Vulnerability Reporting enhancements", "ROX_VULN_MGMT_REPORTING_ENHANCEMENTS", true)

	// UnifiedCVEDeferral enables APIs and UI pages for unified deferral workflow.
	UnifiedCVEDeferral = registerFeature("Enable new unified Vulnerability deferral workflow", "ROX_VULN_MGMT_UNIFIED_CVE_DEFERRAL", false)

	// ClusterAwareDeploymentCheck enables roxctl deployment check to check deployments on the cluster level.
	ClusterAwareDeploymentCheck = registerFeature("Enables cluster level check for the 'roxctl deployment check' command.", "ROX_CLUSTER_AWARE_DEPLOYMENT_CHECK", true)

	// WorkloadCVEsFixabilityFilters enables Workload CVE UI controls for fixability filters and default filters
	WorkloadCVEsFixabilityFilters = registerFeature("Enables Workload CVE fixability filters", "ROX_WORKLOAD_CVES_FIXABILITY_FILTERS", false)

	// SensorReconciliationOnReconnect enables sensors to support reconciliation when reconnecting
	SensorReconciliationOnReconnect = registerFeature("Enable Sensors to support reconciliation on reconnect", "ROX_SENSOR_RECONCILIATION", true)

	// AuthMachineToMachine allows to exchange ID tokens for Central tokens without requiring user interaction.
	AuthMachineToMachine = registerFeature("Enable Auth Machine to Machine functionalities", "ROX_AUTH_MACHINE_TO_MACHINE", true)

	// PolicyCriteriaModal enables a modal for selecting policy criteria when editing a policy
	PolicyCriteriaModal = registerFeature("Enable modal to select policy criteria when editing a policy", "ROX_POLICY_CRITERIA_MODAL", false)

	// SensorDeploymentBuildOptimization enables a performance improvement by skipping deployments processing when no dependency or spec changed
	SensorDeploymentBuildOptimization = registerFeature("Enables a performance improvement by skipping deployments processing when no dependency or spec changed", "ROX_DEPLOYMENT_BUILD_OPTIMIZATION", true)

	// DeploymentVolumeSearch enables search on the volume fields of deployments
	_ = registerFeature("Enables search on the volume fields of deployments", "ROX_DEPLOYMENT_VOLUME_SEARCH", true)

	// DeploymentSecretSearch enables search on the secret fields of deployments
	_ = registerFeature("Enables search on the secret fields of deployments", "ROX_DEPLOYMENT_SECRET_SEARCH", true)

	// DeploymentEnvvarSearch enables search on the environment variable fields of deployments
	_ = registerFeature("Enables search on the environment variable fields of deployments", "ROX_DEPLOYMENT_ENVVAR_SEARCH", true)

	// SecretFileSearch enables search on the file fields of secrets
	_ = registerFeature("Enables search on the file fields of secrets", "ROX_SECRET_FILE_SEARCH", true)

	// SensorCapturesIntermediateEvents enables sensor to capture intermediate events when it is disconnected from central
	SensorCapturesIntermediateEvents = registerFeature("Enables sensor to capture intermediate events when it is disconnected from central", "ROX_CAPTURE_INTERMEDIATE_EVENTS", true)

	// ScannerV4Support enables various capabilities associated with ACS + Clair consolidated scanner.
	ScannerV4Support = registerUnchangeableFeature("Enable Scanner V4 Support", "ROX_SCANNER_V4_SUPPORT", true)

	// ScannerV4 indicates Scanner V4 is installed and should be used as the default image scanner in Central/Sensor.
	ScannerV4 = registerFeature("Enables Scanner V4 runtime functionality", "ROX_SCANNER_V4", false)

	// CloudCredentials enables support for short-lived cloud credentials.
	CloudCredentials = registerFeature("Enable support for short-lived cloud credentials", "ROX_CLOUD_CREDENTIALS", true)

	// CloudSources enables support for cloud source integrations.
	CloudSources = registerFeature("Enable support for cloud source integrations", "ROX_CLOUD_SOURCES", true)

	// ComplianceHierachyControlData enables support for compliance hierarchy control data.
	ComplianceHierachyControlData = registerFeature("Enable support to display and import the compliance hierarchy control data", "ROX_COMPLIANCE_HIERARCHY_CONTROL_DATA", false)

	// VulnMgmtNodePlatformCVEs enables new APIs and UI for VM 2.0 Node and Platform CVEs
	VulnMgmtNodePlatformCVEs = registerFeature("Enables support for Node and Platform CVEs in VM 2.0", "ROX_VULN_MGMT_NODE_PLATFORM_CVES", false)

	// ComplianceReporting enables support for compliance reporting.
	ComplianceReporting = registerFeature("Enable support for V2 compliance reporting", "ROX_COMPLIANCE_REPORTING", false)

	// UnqualifiedSearchRegistries enables support for unqualified search registries and short name aliases.
	UnqualifiedSearchRegistries = registerFeature("Enable support for unqualified search registries and short name aliases", "ROX_UNQUALIFIED_SEARCH_REGISTRIES", false)
)
