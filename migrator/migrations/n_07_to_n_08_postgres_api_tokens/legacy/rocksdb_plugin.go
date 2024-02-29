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
	bucket = []byte("apiTokens")
)

type Store interface {
	UpsertMany(ctx context.Context, objs []*storage.TokenMetadata) error
	Walk(ctx context.Context, fn func(obj *storage.TokenMetadata) error) error
}

type storeImpl struct {
	crud db.Crud
}

func alloc() protocompat.Message {
	return &storage.TokenMetadata{}
}

func keyFunc(msg protocompat.Message) []byte {
	return []byte(msg.(*storage.TokenMetadata).GetId())
}

// New returns a new Store instance using the provided rocksdb instance.
func New(db *rocksdb.RocksDB) (Store, error) {
	baseCRUD := generic.NewCRUD(db, bucket, keyFunc, alloc, false)
	return &storeImpl{crud: baseCRUD}, nil
}

// UpsertMany batches objects into the DB
func (b *storeImpl) UpsertMany(_ context.Context, objs []*storage.TokenMetadata) error {
	msgs := make([]protocompat.Message, 0, len(objs))
	for _, o := range objs {
		msgs = append(msgs, o)
	}

	return b.crud.UpsertMany(msgs)
}

// Walk iterates over all of the objects in the store and applies the closure
func (b *storeImpl) Walk(_ context.Context, fn func(obj *storage.TokenMetadata) error) error {
	return b.crud.Walk(func(msg protocompat.Message) error {
		return fn(msg.(*storage.TokenMetadata))
	})
}
