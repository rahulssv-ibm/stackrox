package m21tom22

import (
	"github.com/gogo/protobuf/proto"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/migrator/migrations"
	"github.com/stackrox/rox/migrator/types"
	bolt "go.etcd.io/bbolt"
)

var clustersBucket = []byte("clusters")

var migration = types.Migration{
	StartingSeqNum: 21,
	VersionAfter:   storage.Version{SeqNum: 22},
	Run: func(databases *types.Databases) error {
		return rewrite(databases.BoltDB)
	},
}

func rewrite(db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(clustersBucket)
		if bucket == nil {
			return nil
		}
		return bucket.ForEach(func(k, v []byte) error {
			var cluster storage.Cluster
			err := proto.Unmarshal(v, &cluster)
			if err != nil {
				return err
			}
			if cluster.TolerationsConfig != nil {
				return nil
			}
			cluster.TolerationsConfig = &storage.TolerationsConfig{
				Disabled: true,
			}
			data, err := proto.Marshal(&cluster)
			if err != nil {
				return err
			}
			return bucket.Put(k, data)

		})
	})
}

func init() {
	migrations.MustRegisterMigration(migration)
}
