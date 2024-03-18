package store

import (
	"context"

	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/search"
)

// PermissionSetStore provides storage functionality for permission sets.
//
//go:generate mockgen-wrapper
type PermissionSetStore interface {
	Get(ctx context.Context, id string) (*storage.PermissionSet, bool, error)
	Count(ctx context.Context, q *v1.Query) (int, error)
	Search(ctx context.Context, q *v1.Query) ([]search.Result, error)
	Upsert(ctx context.Context, obj *storage.PermissionSet) error
	UpsertMany(ctx context.Context, obj []*storage.PermissionSet) error
	Delete(ctx context.Context, id string) error
	Walk(ctx context.Context, fn func(obj *storage.PermissionSet) error) error
}

// SimpleAccessScopeStore provides storage functionality for simple access scopes.
//
//go:generate mockgen-wrapper
type SimpleAccessScopeStore interface {
	Get(ctx context.Context, id string) (*storage.SimpleAccessScope, bool, error)
	Count(ctx context.Context, q *v1.Query) (int, error)
	Search(ctx context.Context, q *v1.Query) ([]search.Result, error)
	Exists(ctx context.Context, id string) (bool, error)
	Upsert(ctx context.Context, obj *storage.SimpleAccessScope) error
	UpsertMany(ctx context.Context, obj []*storage.SimpleAccessScope) error
	Delete(ctx context.Context, id string) error
	Walk(ctx context.Context, fn func(obj *storage.SimpleAccessScope) error) error
}

// RoleStore provides storage functionality for roles.
//
//go:generate mockgen-wrapper
type RoleStore interface {
	Get(ctx context.Context, id string) (*storage.Role, bool, error)
	Count(ctx context.Context, q *v1.Query) (int, error)
	Search(ctx context.Context, q *v1.Query) ([]search.Result, error)
	Upsert(ctx context.Context, obj *storage.Role) error
	UpsertMany(ctx context.Context, obj []*storage.Role) error
	Delete(ctx context.Context, id string) error
	Walk(ctx context.Context, fn func(obj *storage.Role) error) error
}
