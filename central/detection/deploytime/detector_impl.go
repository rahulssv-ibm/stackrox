package deploytime

import (
	"context"

	"github.com/pkg/errors"
	deploymentDataStore "github.com/stackrox/rox/central/deployment/datastore"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/detection"
	"github.com/stackrox/rox/pkg/detection/deploytime"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/utils"
)

var (
	log = logging.LoggerForModule()

	executorCtx = sac.WithAllAccess(context.Background())
)

type detectorImpl struct {
	policySet      detection.PolicySet
	deployments    deploymentDataStore.DataStore
	singleDetector deploytime.Detector
}

// UpsertPolicy adds or updates a policy in the set.
func (d *detectorImpl) PolicySet() detection.PolicySet {
	return d.policySet
}

// Detect runs detection on an deployment, returning any generated alerts.
func (d *detectorImpl) Detect(ctx deploytime.DetectionContext, deployment *storage.Deployment, images []*storage.Image) ([]*storage.Alert, error) {
	return d.singleDetector.Detect(ctx, deployment, images)
}

func (d *detectorImpl) AlertsForPolicy(policyID string) ([]*storage.Alert, error) {
	if features.BooleanPolicyLogic.Enabled() {
		return nil, utils.Should(errors.New("search-based policy evaluation is deprecated"))
	}
	exe := newAllDeploymentsExecutor(executorCtx, d.deployments)
	err := d.policySet.ForOne(policyID, exe)
	if err != nil {
		return nil, err
	}
	return exe.GetAlerts(), nil
}
