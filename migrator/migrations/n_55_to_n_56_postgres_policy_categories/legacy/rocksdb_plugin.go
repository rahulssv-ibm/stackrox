package legacy

import (
	"context"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/db"
	"github.com/stackrox/rox/pkg/protocompat"
	"github.com/stackrox/rox/pkg/rocksdb"
	generic "github.com/stackrox/rox/pkg/rocksdb/crud"
)

var (
	bucket = []byte("policy_categories")
)

type Store interface {
	UpsertMany(ctx context.Context, objs []*storage.PolicyCategory) error
	Walk(ctx context.Context, fn func(obj *storage.PolicyCategory) error) error
}

type storeImpl struct {
	crud db.Crud
}

func alloc() protocompat.Message {
	return &storage.PolicyCategory{}
}

func keyFunc(msg protocompat.Message) []byte {
	return []byte(msg.(*storage.PolicyCategory).GetId())
}
func uniqKeyFunc(msg protocompat.Message) []byte {
	return []byte(msg.(*storage.PolicyCategory).GetName())
}

// New returns a new Store instance using the provided rocksdb instance.
func New(db *rocksdb.RocksDB) (Store, error) {
	baseCRUD := generic.NewUniqueKeyCRUD(db, bucket, keyFunc, alloc, uniqKeyFunc, false)
	return &storeImpl{crud: baseCRUD}, nil
}

// UpsertMany batches objects into the DB
func (b *storeImpl) UpsertMany(_ context.Context, objs []*storage.PolicyCategory) error {
	msgs := make([]protocompat.Message, 0, len(objs))
	for _, o := range objs {
		msgs = append(msgs, o)
	}

	return b.crud.UpsertMany(msgs)
}

// Walk iterates over all of the objects in the store and applies the closure
func (b *storeImpl) Walk(_ context.Context, fn func(obj *storage.PolicyCategory) error) error {
	return b.crud.Walk(func(msg protocompat.Message) error {
		return fn(msg.(*storage.PolicyCategory))
	})
}
