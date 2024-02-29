package m187tom188

import (
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/migrator/migrations"
	"github.com/stackrox/rox/migrator/types"
	"github.com/stackrox/rox/pkg/logging"
)

const (
	startSeqNum = 187
)

var (
	migration = types.Migration{
		StartingSeqNum: startSeqNum,
		VersionAfter:   &storage.Version{SeqNum: int32(startSeqNum + 1)},
		Run:            migrate,
	}

	log = logging.LoggerForModule()
)

func init() {
	migrations.MustRegisterMigration(migration)
}
