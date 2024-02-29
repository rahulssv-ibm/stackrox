// This file was originally generated with
// //go:generate cp ../../../central/cve/dackbox/crud.go cve/crud.go
package dackbox

import (
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/dackbox/crud"
	"github.com/stackrox/rox/pkg/dbhelper"
	"github.com/stackrox/rox/pkg/protocompat"
)

var (
	// Bucket stores the child image vulnerabilities.
	Bucket = []byte("image_vuln")

	// BucketHandler is the bucket's handler.
	BucketHandler = &dbhelper.BucketHandler{BucketPrefix: Bucket}

	// Reader reads storage.CVEs directly from the store.
	Reader = crud.NewReader(
		crud.WithAllocFunction(Alloc),
	)

	// Upserter writes storage.CVEs directly to the store.
	Upserter = crud.NewUpserter(
		crud.WithKeyFunction(KeyFunc),
		crud.AddToIndex(),
	)

	// Deleter deletes vulns from the store.
	Deleter = crud.NewDeleter(
		crud.Shared(),
		crud.RemoveFromIndex(),
	)
)

// KeyFunc returns the key for a CVE.
func KeyFunc(msg protocompat.Message) []byte {
	unPrefixed := []byte(msg.(interface{ GetId() string }).GetId())
	return dbhelper.GetBucketKey(Bucket, unPrefixed)
}

// Alloc allocates a CVE.
func Alloc() protocompat.Message {
	return &storage.CVE{}
}
