// Package features helps enable or disable features.
package features

import (
	"fmt"
	"strings"
)

// A FeatureFlag is a product behavior that can be enabled or disabled using an
// environment variable.
type FeatureFlag interface {
	Name() string
	EnvVar() string
	Enabled() bool
	Default() bool
	Stage() string
}

var (
	// Flags contains all defined FeatureFlags by name.
	Flags = make(map[string]FeatureFlag)

	allPerStage = map[string]FeatureFlag{
		devPreviewString:  registerFeature("All dev-preview features", "ROX_ALL_DEV_PREVIEW"),
		techPreviewString: registerFeature("All tech-preview features", "ROX_ALL_TECH_PREVIEW", techPreview),
	}
)

// registerFeature registers and returns a new feature flag, configured with the
// provided options.
func registerFeature(name, envVar string, options ...option) FeatureFlag {
	if !strings.HasPrefix(envVar, "ROX_") {
		panic(fmt.Sprintf("invalid env var: %s, must start with ROX_", envVar))
	}
	f := &feature{
		name:   name,
		envVar: envVar,
	}
	for _, o := range options {
		o(f)
	}
	Flags[f.envVar] = f
	return f
}
