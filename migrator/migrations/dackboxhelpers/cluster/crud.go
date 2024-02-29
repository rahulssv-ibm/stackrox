// This file was originally generated with
// //go:generate cp ../../../central/cluster/dackbox/crud.go cluster/crud.go
package dackbox

import (
	"github.com/stackrox/rox/pkg/dbhelper"
)

var (
	// Bucket is the prefix for stored clusters.
	Bucket = []byte("clusters")

	// BucketHandler is the bucket's handler.
	BucketHandler = &dbhelper.BucketHandler{BucketPrefix: Bucket}
)
