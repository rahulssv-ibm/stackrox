package mappings

import (
	imageMapping "github.com/stackrox/rox/central/image/index/mappings"
	processIndicatorMapping "github.com/stackrox/rox/central/processindicator/index/mappings"
	"github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/pkg/search"
)

// OptionsMap is exposed for e2e test
var OptionsMap = mergeMaps(map[search.FieldLabel]*v1.SearchField{
	search.Cluster:   search.NewStringField(v1.SearchCategory_DEPLOYMENTS, "deployment.cluster_name"),
	search.ClusterID: search.NewField(v1.SearchCategory_DEPLOYMENTS, "deployment.cluster_id", v1.SearchDataType_SEARCH_STRING, search.OptionHidden|search.OptionStore),
	search.Namespace: search.NewStringField(v1.SearchCategory_DEPLOYMENTS, "deployment.namespace"),
	search.Label:     search.NewMapField(v1.SearchCategory_DEPLOYMENTS, "deployment.labels"),

	search.CPUCoresLimit:    search.NewNumericField(v1.SearchCategory_DEPLOYMENTS, "deployment.containers.resources.cpu_cores_limit"),
	search.CPUCoresRequest:  search.NewField(v1.SearchCategory_DEPLOYMENTS, "deployment.containers.resources.cpu_cores_request", v1.SearchDataType_SEARCH_NUMERIC, search.OptionStore),
	search.DeploymentID:     search.NewField(v1.SearchCategory_DEPLOYMENTS, "deployment.id", v1.SearchDataType_SEARCH_STRING, search.OptionHidden|search.OptionStore),
	search.DeploymentName:   search.NewStringField(v1.SearchCategory_DEPLOYMENTS, "deployment.name"),
	search.DeploymentType:   search.NewStringField(v1.SearchCategory_DEPLOYMENTS, "deployment.type"),
	search.AddCapabilities:  search.NewStringField(v1.SearchCategory_DEPLOYMENTS, "deployment.containers.security_context.add_capabilities"),
	search.DropCapabilities: search.NewStringField(v1.SearchCategory_DEPLOYMENTS, "deployment.containers.security_context.drop_capabilities"),
	search.EnvironmentKey:   search.NewStringField(v1.SearchCategory_DEPLOYMENTS, "deployment.containers.config.env.key"),
	search.EnvironmentValue: search.NewStringField(v1.SearchCategory_DEPLOYMENTS, "deployment.containers.config.env.value"),
	search.ImagePullSecret:  search.NewStringField(v1.SearchCategory_DEPLOYMENTS, "deployment.image_pull_secrets"),
	search.MemoryLimit:      search.NewNumericField(v1.SearchCategory_DEPLOYMENTS, "deployment.containers.resources.memory_mb_limit"),
	search.MemoryRequest:    search.NewNumericField(v1.SearchCategory_DEPLOYMENTS, "deployment.containers.resources.memory_mb_request"),
	search.Privileged:       search.NewBoolField(v1.SearchCategory_DEPLOYMENTS, "deployment.containers.security_context.privileged"),

	search.SecretName:        search.NewStringField(v1.SearchCategory_DEPLOYMENTS, "deployment.containers.secrets.name"),
	search.SecretPath:        search.NewStringField(v1.SearchCategory_DEPLOYMENTS, "deployment.containers.secrets.path"),
	search.ServiceAccount:    search.NewStringField(v1.SearchCategory_DEPLOYMENTS, "deployment.service_account"),
	search.VolumeName:        search.NewField(v1.SearchCategory_DEPLOYMENTS, "deployment.containers.volumes.name", v1.SearchDataType_SEARCH_STRING, search.OptionStore),
	search.VolumeSource:      search.NewStringField(v1.SearchCategory_DEPLOYMENTS, "deployment.containers.volumes.source"),
	search.VolumeDestination: search.NewStringField(v1.SearchCategory_DEPLOYMENTS, "deployment.containers.volumes.destination"),
	search.VolumeReadonly:    search.NewBoolField(v1.SearchCategory_DEPLOYMENTS, "deployment.containers.volumes.read_only"),
	search.VolumeType:        search.NewStringField(v1.SearchCategory_DEPLOYMENTS, "deployment.containers.volumes.type"),

	"ImageRelationship": search.NewField(v1.SearchCategory_DEPLOYMENTS, "deployment.containers.image.name.sha", v1.SearchDataType_SEARCH_STRING, search.OptionHidden|search.OptionStore),
}, imageMapping.OptionsMap, processIndicatorMapping.OptionsMap)

func mergeMaps(m1 map[search.FieldLabel]*v1.SearchField, otherMaps ...map[search.FieldLabel]*v1.SearchField) map[search.FieldLabel]*v1.SearchField {
	finalMap := make(map[search.FieldLabel]*v1.SearchField)
	for _, m := range otherMaps {
		for k, v := range m {
			finalMap[k] = v
		}
	}
	// This ordering is actually important because the fields defined in this file have precendence over ones in other OptionsMaps
	for k, v := range m1 {
		finalMap[k] = v
	}

	return finalMap
}
