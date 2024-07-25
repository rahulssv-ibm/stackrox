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
	// Flags contains all defined FeatureFlags.
	Flags = []FeatureFlag{}
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
	Flags = append(Flags, f)
	return f
}

// FindFlagByVariable searches for the flag with the given variable, and returns
// the flag if found, and nil otherwise.
func FindFlagByVariable(envVar string) FeatureFlag {
	for _, flag := range Flags {
		if flag.EnvVar() == envVar {
			return flag
		}
	}
	return nil
}
