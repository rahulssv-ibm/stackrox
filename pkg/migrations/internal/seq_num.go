// Code generated by make bootstrap_migration generator. DO NOT EDIT.
package internal

var (
	// CurrentDBVersionSeqNum is the current DB version number.
	// This must be incremented every time we write a migration.
	// It is a shared constant between central and the migrator binary.
	CurrentDBVersionSeqNum = 189

	// LastRocksDBVersionSeqNum is the sequence number for the last RocksDB version.
	LastRocksDBVersionSeqNum = 112

	// LastRocksDBToPostgresVersionSeqNum is the sequence number for the last Rocks -> Postgres migration (n migration)
	LastRocksDBToPostgresVersionSeqNum = 57
)
