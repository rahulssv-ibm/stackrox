//go:build test_all

package k8sutil

import (
	"strings"
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stretchr/testify/assert"
)

var (
	crioCGroup = `
11:freezer:/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod61223a2c_dead_11e9_8887_42010a000004.slice/crio-6f6b9de202f6ec9939bcbecf0beebe1f56c6f521922d4073467136b536c75f53.scope
10:perf_event:/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod61223a2c_dead_11e9_8887_42010a000004.slice/crio-6f6b9de202f6ec9939bcbecf0beebe1f56c6f521922d4073467136b536c75f53.scope
9:cpuset:/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod61223a2c_dead_11e9_8887_42010a000004.slice/crio-6f6b9de202f6ec9939bcbecf0beebe1f56c6f521922d4073467136b536c75f53.scope
8:devices:/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod61223a2c_dead_11e9_8887_42010a000004.slice/crio-6f6b9de202f6ec9939bcbecf0beebe1f56c6f521922d4073467136b536c75f53.scope
7:cpuacct,cpu:/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod61223a2c_dead_11e9_8887_42010a000004.slice/crio-6f6b9de202f6ec9939bcbecf0beebe1f56c6f521922d4073467136b536c75f53.scope
6:memory:/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod61223a2c_dead_11e9_8887_42010a000004.slice/crio-6f6b9de202f6ec9939bcbecf0beebe1f56c6f521922d4073467136b536c75f53.scope
5:blkio:/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod61223a2c_dead_11e9_8887_42010a000004.slice/crio-6f6b9de202f6ec9939bcbecf0beebe1f56c6f521922d4073467136b536c75f53.scope
4:pids:/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod61223a2c_dead_11e9_8887_42010a000004.slice/crio-6f6b9de202f6ec9939bcbecf0beebe1f56c6f521922d4073467136b536c75f53.scope
3:hugetlb:/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod61223a2c_dead_11e9_8887_42010a000004.slice/crio-6f6b9de202f6ec9939bcbecf0beebe1f56c6f521922d4073467136b536c75f53.scope
2:net_prio,net_cls:/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod61223a2c_dead_11e9_8887_42010a000004.slice/crio-6f6b9de202f6ec9939bcbecf0beebe1f56c6f521922d4073467136b536c75f53.scope
1:name=systemd:/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod61223a2c_dead_11e9_8887_42010a000004.slice/crio-6f6b9de202f6ec9939bcbecf0beebe1f56c6f521922d4073467136b536c75f53.scope
`

	unknownRuntimeCGroup = `
13:name=systemd:/kubepods/kubepods/burstable/podd99bfe2c-dfce-11e9-8dc5-025000000001/4eba512587b2599c2cadaf8111327b95347366ee787387ba7dbb74e80793a1cd
12:pids:/kubepods/kubepods/burstable/podd99bfe2c-dfce-11e9-8dc5-025000000001/4eba512587b2599c2cadaf8111327b95347366ee787387ba7dbb74e80793a1cd
11:hugetlb:/kubepods/kubepods/burstable/podd99bfe2c-dfce-11e9-8dc5-025000000001/4eba512587b2599c2cadaf8111327b95347366ee787387ba7dbb74e80793a1cd
10:net_prio:/kubepods/kubepods/burstable/podd99bfe2c-dfce-11e9-8dc5-025000000001/4eba512587b2599c2cadaf8111327b95347366ee787387ba7dbb74e80793a1cd
9:perf_event:/kubepods/kubepods/burstable/podd99bfe2c-dfce-11e9-8dc5-025000000001/4eba512587b2599c2cadaf8111327b95347366ee787387ba7dbb74e80793a1cd
8:net_cls:/kubepods/kubepods/burstable/podd99bfe2c-dfce-11e9-8dc5-025000000001/4eba512587b2599c2cadaf8111327b95347366ee787387ba7dbb74e80793a1cd
7:freezer:/kubepods/kubepods/burstable/podd99bfe2c-dfce-11e9-8dc5-025000000001/4eba512587b2599c2cadaf8111327b95347366ee787387ba7dbb74e80793a1cd
6:devices:/kubepods/kubepods/burstable/podd99bfe2c-dfce-11e9-8dc5-025000000001/4eba512587b2599c2cadaf8111327b95347366ee787387ba7dbb74e80793a1cd
5:memory:/kubepods/kubepods/burstable/podd99bfe2c-dfce-11e9-8dc5-025000000001/4eba512587b2599c2cadaf8111327b95347366ee787387ba7dbb74e80793a1cd
4:blkio:/kubepods/kubepods/burstable/podd99bfe2c-dfce-11e9-8dc5-025000000001/4eba512587b2599c2cadaf8111327b95347366ee787387ba7dbb74e80793a1cd
3:cpuacct:/kubepods/kubepods/burstable/podd99bfe2c-dfce-11e9-8dc5-025000000001/4eba512587b2599c2cadaf8111327b95347366ee787387ba7dbb74e80793a1cd
2:cpu:/kubepods/kubepods/burstable/podd99bfe2c-dfce-11e9-8dc5-025000000001/4eba512587b2599c2cadaf8111327b95347366ee787387ba7dbb74e80793a1cd
1:cpuset:/kubepods/kubepods/burstable/podd99bfe2c-dfce-11e9-8dc5-025000000001/4eba512587b2599c2cadaf8111327b95347366ee787387ba7dbb74e80793a1cd
`
)

func TestInferContainerRuntime_CRIO(t *testing.T) {
	rt, err := inferContainerRuntimeFromCGroupFile(strings.NewReader(crioCGroup))
	assert.NoError(t, err)
	assert.Equal(t, storage.ContainerRuntime_CRIO_CONTAINER_RUNTIME, rt)
}

func TestInferContainerRuntime_Unknown(t *testing.T) {
	rt, err := inferContainerRuntimeFromCGroupFile(strings.NewReader(unknownRuntimeCGroup))
	assert.Error(t, err)
	assert.Equal(t, storage.ContainerRuntime_UNKNOWN_CONTAINER_RUNTIME, rt)
}
