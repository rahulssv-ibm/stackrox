package mappings

import (
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/search"
)

// OptionsMap defines the search options for Vulnerabilities stored in images.
var OptionsMap = search.Walk(v1.SearchCategory_COMPONENT_VULN_EDGE, "componentcveedge", (*storage.ComponentCVEEdge)(nil))
