package satoken

import (
	"bytes"
	"io/ioutil"
)

// LoadTokenFromFile loads the Kubernetes service account JWT token from the canonical file location and returns the
// token or an error.
func LoadTokenFromFile() (string, error) {
	contents, err := ioutil.ReadFile(ServiceAccountTokenJWTPath)
	if err != nil {
		return "", err
	}
	return string(bytes.TrimSpace(contents)), nil
}

// LoadNamespaceFromFile loads the Kubernetes service account namespace (which is the same as the pod namespace)
// from the canonical file location and returns the namespace or an error.
func LoadNamespaceFromFile() (string, error) {
	contents, err := ioutil.ReadFile(ServiceAccountTokenNamespacePath)
	if err != nil {
		return "", err
	}
	return string(bytes.TrimSpace(contents)), nil
}
