// Code generated by make bootstrap_migration generator. DO NOT EDIT.
package runner

import (
	// Postgres -> Postgres migrations
	_ "github.com/stackrox/rox/migrator/migrations/m_168_to_m_169_postgres_remove_clustercve_permission"
	_ "github.com/stackrox/rox/migrator/migrations/m_169_to_m_170_collections_sac_resource_migration"
	_ "github.com/stackrox/rox/migrator/migrations/m_170_to_m_171_create_policy_categories_and_edges"
	_ "github.com/stackrox/rox/migrator/migrations/m_171_to_m_172_move_scope_to_collection_in_report_configurations"
	_ "github.com/stackrox/rox/migrator/migrations/m_172_to_m_173_network_flows_partition"
	_ "github.com/stackrox/rox/migrator/migrations/m_173_to_m_174_group_unique_constraint"
	_ "github.com/stackrox/rox/migrator/migrations/m_174_to_m_175_enable_search_on_api_tokens"
	_ "github.com/stackrox/rox/migrator/migrations/m_175_to_m_176_create_notification_schedule_table"
	_ "github.com/stackrox/rox/migrator/migrations/m_176_to_m_177_network_baselines_cidr"
	_ "github.com/stackrox/rox/migrator/migrations/m_177_to_m_178_group_permissions"
	_ "github.com/stackrox/rox/migrator/migrations/m_178_to_m_179_embedded_collections_search_label"
	_ "github.com/stackrox/rox/migrator/migrations/m_179_to_m_180_openshift_policy_exclusions"
	_ "github.com/stackrox/rox/migrator/migrations/m_180_to_m_181_move_to_blobstore"
	_ "github.com/stackrox/rox/migrator/migrations/m_181_to_m_182_group_role_permission_with_access_one"
	_ "github.com/stackrox/rox/migrator/migrations/m_182_to_m_183_remove_default_scope_manager_role"
	_ "github.com/stackrox/rox/migrator/migrations/m_183_to_m_184_move_declarative_config_health"
	_ "github.com/stackrox/rox/migrator/migrations/m_184_to_m_185_remove_policy_vulnerability_report_resources"
	_ "github.com/stackrox/rox/migrator/migrations/m_185_to_m_186_more_policy_migrations"
	_ "github.com/stackrox/rox/migrator/migrations/m_186_to_m_187_add_blob_search"
	_ "github.com/stackrox/rox/migrator/migrations/m_187_to_m_188_remove_reports_without_migrated_collections"
	_ "github.com/stackrox/rox/migrator/migrations/m_188_to_m_189_logimbues_add_time_column"
	_ "github.com/stackrox/rox/migrator/migrations/m_189_to_m_190_vulnerability_requests_add_name"
	_ "github.com/stackrox/rox/migrator/migrations/m_190_to_m_191_plop_add_closed_time_column"
	_ "github.com/stackrox/rox/migrator/migrations/m_191_to_m_192_vulnerability_requests_searchable_scope"
	_ "github.com/stackrox/rox/migrator/migrations/m_192_to_m_193_report_config_migrations"
	_ "github.com/stackrox/rox/migrator/migrations/m_193_to_m_194_policy_updates_for_4_3"
	_ "github.com/stackrox/rox/migrator/migrations/m_194_to_m_195_vuln_request_global_scope"
	_ "github.com/stackrox/rox/migrator/migrations/m_195_to_m_196_vuln_request_users"
	_ "github.com/stackrox/rox/migrator/migrations/m_196_to_m_197_vulnerability_requests_for_legacy"
	_ "github.com/stackrox/rox/migrator/migrations/m_197_to_m_198_add_oidc_claim_mappings"
	_ "github.com/stackrox/rox/migrator/migrations/m_198_to_m_199_policy_description_and_criteria_updates"
	_ "github.com/stackrox/rox/migrator/migrations/m_199_to_m_200_policy_updates_for_4_5"
	_ "github.com/stackrox/rox/migrator/migrations/m_200_to_m_201_compliance_v2_for_4_5"
	_ "github.com/stackrox/rox/migrator/migrations/m_201_to_m_202_vuln_request_v1_to_v2"
	_ "github.com/stackrox/rox/migrator/migrations/m_202_to_m_203_vuln_requests_for_suppressed_cves"
	_ "github.com/stackrox/rox/migrator/migrations/m_203_to_m_204_openshift_policy_exclusions_for_4_5"
	_ "github.com/stackrox/rox/migrator/migrations/m_204_to_m_205_clusters_platform_type_and_k8_version"
	_ "github.com/stackrox/rox/migrator/migrations/m_205_to_m_206_remove_bad_gorm_index"
)
