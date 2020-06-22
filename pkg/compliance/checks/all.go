package checks

import (
	// Make sure all checks from all standards are registered.
	_ "github.com/stackrox/rox/pkg/compliance/checks/docker"
	_ "github.com/stackrox/rox/pkg/compliance/checks/kubernetes"
	_ "github.com/stackrox/rox/pkg/compliance/checks/nist800-190"
)
