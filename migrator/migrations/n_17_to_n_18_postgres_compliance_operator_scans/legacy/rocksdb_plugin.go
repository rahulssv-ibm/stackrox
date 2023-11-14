// Code generated by rocksdb-bindings generator. DO NOT EDIT.
package legacy

import (
	"context"

	proto "github.com/CrowdStrike/csproto"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/db"
	"github.com/stackrox/rox/pkg/rocksdb"
	generic "github.com/stackrox/rox/pkg/rocksdb/crud"
)

var (
	bucket = []byte("complianceoperatorscans")
)

type Store interface {
	UpsertMany(ctx context.Context, objs []*storage.ComplianceOperatorScan) error
	Walk(ctx context.Context, fn func(obj *storage.ComplianceOperatorScan) error) error
}

type storeImpl struct {
	crud db.Crud
}

func alloc() proto.Message {
	return &storage.ComplianceOperatorScan{}
}

func keyFunc(msg proto.Message) []byte {
	return []byte(msg.(*storage.ComplianceOperatorScan).GetId())
}

// New returns a new Store instance using the provided rocksdb instance.
func New(db *rocksdb.RocksDB) (Store, error) {
	baseCRUD := generic.NewCRUD(db, bucket, keyFunc, alloc, false)
	return &storeImpl{crud: baseCRUD}, nil
}

// UpsertMany batches objects into the DB
func (b *storeImpl) UpsertMany(_ context.Context, objs []*storage.ComplianceOperatorScan) error {
	msgs := make([]proto.Message, 0, len(objs))
	for _, o := range objs {
		msgs = append(msgs, o)
	}

	return b.crud.UpsertMany(msgs)
}

// Walk iterates over all of the objects in the store and applies the closure
func (b *storeImpl) Walk(_ context.Context, fn func(obj *storage.ComplianceOperatorScan) error) error {
	return b.crud.Walk(func(msg proto.Message) error {
		return fn(msg.(*storage.ComplianceOperatorScan))
	})
}
