// This file was originally generated with
// //go:generate cp ../../../central/namespace/dackbox/crud.go namespace/crud.go
package dackbox

import (
	"github.com/stackrox/rox/pkg/dbhelper"
)

var (
	// Bucket is the prefix for stored namespaces.
	Bucket = []byte("namespaces")

	// BucketHandler is the bucket's handler.
	BucketHandler = &dbhelper.BucketHandler{BucketPrefix: Bucket}
)
