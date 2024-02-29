package schema

import (
	"reflect"

	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/search"
)

var (
	// CreateTablePolicyCategoriesStmt holds the create statement for table `policy_categories`.
	CreateTablePolicyCategoriesStmt = &postgres.CreateStmts{
		GormModel: (*PolicyCategories)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// PolicyCategoriesSchema is the go schema for table `policy_categories`.
	PolicyCategoriesSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.PolicyCategory)(nil)), "policy_categories")
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_POLICY_CATEGORIES, "policycategory", (*storage.PolicyCategory)(nil)))
		return schema
	}()
)

const (
	PolicyCategoriesTableName = "policy_categories"
)

// PolicyCategories holds the Gorm model for Postgres table `policy_categories`.
type PolicyCategories struct {
	Id         string `gorm:"column:id;type:varchar;primaryKey"`
	Name       string `gorm:"column:name;type:varchar;unique"`
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}
