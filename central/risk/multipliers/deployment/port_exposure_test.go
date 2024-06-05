package deployment

import (
	"context"
	"testing"

	"github.com/stackrox/rox/central/risk/multipliers"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/protocompat"
	"github.com/stretchr/testify/assert"
)

func TestPortExposureScore(t *testing.T) {
	portMultiplier := NewReachability()

	deployment := multipliers.GetMockDeployment()
	expectedScore := &storage.Risk_Result{
		Name: ReachabilityHeading,
		Factors: []*storage.Risk_Result_Factor{
			{Message: "Port 22 is exposed to external clients"},
			{Message: "Port 23 is exposed in the cluster"},
			{Message: "Port 24 is exposed on node interfaces"},
		},
		Score: 1.6,
	}
	score := portMultiplier.Score(context.Background(), deployment, nil)
	assert.True(t, protocompat.Equal(expectedScore, score))
}
