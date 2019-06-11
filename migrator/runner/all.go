package runner

import (
	// Import these packages to trigger the registration.
	_ "github.com/stackrox/rox/migrator/migrations/m_0_to_m_1_create_version_bucket"
	_ "github.com/stackrox/rox/migrator/migrations/m_10_to_11_processwhitelist_cluster_namespace"
	_ "github.com/stackrox/rox/migrator/migrations/m_1_to_2_alert_violation"
	_ "github.com/stackrox/rox/migrator/migrations/m_2_to_3_network_flows_in_badger"
	_ "github.com/stackrox/rox/migrator/migrations/m_3_to_4_cluster_spec_to_status"
	_ "github.com/stackrox/rox/migrator/migrations/m_4_to_5_auth_provider_default_group"
	_ "github.com/stackrox/rox/migrator/migrations/m_5_to_6_collection_method"
	_ "github.com/stackrox/rox/migrator/migrations/m_6_to_7_image_update_time"
	_ "github.com/stackrox/rox/migrator/migrations/m_7_to_8_sac_cluster_namespace"
	_ "github.com/stackrox/rox/migrator/migrations/m_8_to_m_9_list_alert_cluster"
	_ "github.com/stackrox/rox/migrator/migrations/m_9_to_m_10_image_component"
)
