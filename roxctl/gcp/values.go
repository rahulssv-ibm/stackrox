package gcp

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Values represent a collection of values that are injected during the GCP
// Marketplace deployment process.
type Values struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`

	MainImage       string `yaml:"main-image"`
	ScannerImage    string `yaml:"scanner-image"`
	MonitoringImage string `yaml:"monitoring-image"`

	License            string `yaml:"license"`
	StackroxIOUsername string `yaml:"stackrox-io-username"`
	StackroxIOPassword string `yaml:"stackrox-io-password"`

	Network string `yaml:"network"`

	PVCName         string `yaml:"pvc-name"`
	PVCStorageclass string `yaml:"pvc-storageclass"`
	PVCSize         uint32 `yaml:"pvc-size"`

	ServiceAccount string `yaml:"svcacct"`
}

func loadValues(filename string) (*Values, error) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var values Values
	if err := yaml.UnmarshalStrict(raw, &values); err != nil {
		return nil, err
	}
	return &values, nil
}
