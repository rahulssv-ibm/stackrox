package manager

import (
	"context"

	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/internalapi/central"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/networkgraph"
	"github.com/stackrox/rox/pkg/timestamp"
)

//go:generate mockgen-wrapper
// The Manager manages network baselines.
// ALL writes to network baselines MUST go through the manager.
type Manager interface {
	// ProcessDeploymentCreate notifies the baseline manager of a deployment create.
	// The baseline manager then creates a baseline for this deployment if it does not already exist.
	// It must only be called by trusted code, since it assumes the caller has full access to modify
	// network baselines in the datastore.
	ProcessDeploymentCreate(deploymentID, deploymentName, clusterID, namespace string) error
	// ProcessDeploymentDelete notifies the baseline manager of a deployment delete.
	// The baseline manager then updates all the existing baselines that had an edge to this
	// delete deployment.
	// It must only be called by trusted code, since it assumes the caller has full access to modify
	// network baselines in the datastore.
	ProcessDeploymentDelete(deploymentID string) error
	// ProcessFlowUpdate notifies the baseline manager of a dump of a batch of network flows.
	// It must only be called by trusted code, since it assumes the caller has full access to modify
	// network baselines in the datastore.
	ProcessFlowUpdate(flows map[networkgraph.NetworkConnIndicator]timestamp.MicroTS) error

	// ProcessBaselineStatusUpdate processes a user-filed request to modify the baseline status.
	// The error it returns will be a status.Error.
	ProcessBaselineStatusUpdate(ctx context.Context, modifyRequest *v1.ModifyBaselineStatusForPeersRequest) error
	// ProcessNetworkPolicyUpdate is invoked when we there is a change to the network policies. Changed network
	// policy is passed in allow updating relevant baselines.
	ProcessNetworkPolicyUpdate(ctx context.Context, action central.ResourceAction, policy *storage.NetworkPolicy) error
	// ProcessBaselineLockUpdate updates a baseline's lock status. This locks the baseline if lockBaseline is true
	ProcessBaselineLockUpdate(ctx context.Context, deploymentID string, lockBaseline bool) error
}
