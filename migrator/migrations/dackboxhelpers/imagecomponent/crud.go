// This file was originally generated with
// //go:generate cp ../../../central/imagecomponent/dackbox/crud.go imagecomponent/crud.go
package dackbox

import (
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/dackbox/crud"
	"github.com/stackrox/rox/pkg/dbhelper"
	"github.com/stackrox/rox/pkg/protocompat"
)

var (
	// Bucket stores the child image components.
	Bucket = []byte("image_component")

	// BucketHandler is the bucket's handler.
	BucketHandler = &dbhelper.BucketHandler{BucketPrefix: Bucket}

	// Reader reads storage.ImageComponents from the store.
	Reader = crud.NewReader(
		crud.WithAllocFunction(Alloc),
	)

	// Upserter writes components to the store.
	Upserter = crud.NewUpserter(
		crud.WithKeyFunction(KeyFunc),
		crud.AddToIndex(),
	)

	// Deleter deletes components to the store.
	Deleter = crud.NewDeleter(
		crud.Shared(),
		crud.RemoveFromIndex(),
	)
)

// KeyFunc returns the key for an ImageComponent.
func KeyFunc(msg protocompat.Message) []byte {
	unPrefixed := []byte(msg.(interface{ GetId() string }).GetId())
	return dbhelper.GetBucketKey(Bucket, unPrefixed)
}

// Alloc allocates an ImageComponent.
func Alloc() protocompat.Message {
	return &storage.ImageComponent{}
}
