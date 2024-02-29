package schema

import (
	"reflect"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

var (
	// CreateTableSignatureIntegrationsStmt holds the create statement for table `signature_integrations`.
	CreateTableSignatureIntegrationsStmt = &postgres.CreateStmts{
		GormModel: (*SignatureIntegrations)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// SignatureIntegrationsSchema is the go schema for table `signature_integrations`.
	SignatureIntegrationsSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.SignatureIntegration)(nil)), "signature_integrations")
		return schema
	}()
)

const (
	SignatureIntegrationsTableName = "signature_integrations"
)

// SignatureIntegrations holds the Gorm model for Postgres table `signature_integrations`.
type SignatureIntegrations struct {
	Id         string `gorm:"column:id;type:varchar;primaryKey"`
	Name       string `gorm:"column:name;type:varchar;unique"`
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}
