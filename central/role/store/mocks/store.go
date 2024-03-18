// Code generated by MockGen. DO NOT EDIT.
// Source: store.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/store.go -source store.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	search "github.com/stackrox/rox/pkg/search"
	gomock "go.uber.org/mock/gomock"
)

// MockPermissionSetStore is a mock of PermissionSetStore interface.
type MockPermissionSetStore struct {
	ctrl     *gomock.Controller
	recorder *MockPermissionSetStoreMockRecorder
}

// MockPermissionSetStoreMockRecorder is the mock recorder for MockPermissionSetStore.
type MockPermissionSetStoreMockRecorder struct {
	mock *MockPermissionSetStore
}

// NewMockPermissionSetStore creates a new mock instance.
func NewMockPermissionSetStore(ctrl *gomock.Controller) *MockPermissionSetStore {
	mock := &MockPermissionSetStore{ctrl: ctrl}
	mock.recorder = &MockPermissionSetStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPermissionSetStore) EXPECT() *MockPermissionSetStoreMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockPermissionSetStore) Count(ctx context.Context, q *v1.Query) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, q)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockPermissionSetStoreMockRecorder) Count(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockPermissionSetStore)(nil).Count), ctx, q)
}

// Delete mocks base method.
func (m *MockPermissionSetStore) Delete(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPermissionSetStoreMockRecorder) Delete(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPermissionSetStore)(nil).Delete), ctx, id)
}

// Get mocks base method.
func (m *MockPermissionSetStore) Get(ctx context.Context, id string) (*storage.PermissionSet, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*storage.PermissionSet)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockPermissionSetStoreMockRecorder) Get(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPermissionSetStore)(nil).Get), ctx, id)
}

// Search mocks base method.
func (m *MockPermissionSetStore) Search(ctx context.Context, q *v1.Query) ([]search.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, q)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockPermissionSetStoreMockRecorder) Search(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockPermissionSetStore)(nil).Search), ctx, q)
}

// Upsert mocks base method.
func (m *MockPermissionSetStore) Upsert(ctx context.Context, obj *storage.PermissionSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", ctx, obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockPermissionSetStoreMockRecorder) Upsert(ctx, obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockPermissionSetStore)(nil).Upsert), ctx, obj)
}

// UpsertMany mocks base method.
func (m *MockPermissionSetStore) UpsertMany(ctx context.Context, obj []*storage.PermissionSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertMany", ctx, obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertMany indicates an expected call of UpsertMany.
func (mr *MockPermissionSetStoreMockRecorder) UpsertMany(ctx, obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertMany", reflect.TypeOf((*MockPermissionSetStore)(nil).UpsertMany), ctx, obj)
}

// Walk mocks base method.
func (m *MockPermissionSetStore) Walk(ctx context.Context, fn func(*storage.PermissionSet) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Walk", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// Walk indicates an expected call of Walk.
func (mr *MockPermissionSetStoreMockRecorder) Walk(ctx, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Walk", reflect.TypeOf((*MockPermissionSetStore)(nil).Walk), ctx, fn)
}

// MockSimpleAccessScopeStore is a mock of SimpleAccessScopeStore interface.
type MockSimpleAccessScopeStore struct {
	ctrl     *gomock.Controller
	recorder *MockSimpleAccessScopeStoreMockRecorder
}

// MockSimpleAccessScopeStoreMockRecorder is the mock recorder for MockSimpleAccessScopeStore.
type MockSimpleAccessScopeStoreMockRecorder struct {
	mock *MockSimpleAccessScopeStore
}

// NewMockSimpleAccessScopeStore creates a new mock instance.
func NewMockSimpleAccessScopeStore(ctrl *gomock.Controller) *MockSimpleAccessScopeStore {
	mock := &MockSimpleAccessScopeStore{ctrl: ctrl}
	mock.recorder = &MockSimpleAccessScopeStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSimpleAccessScopeStore) EXPECT() *MockSimpleAccessScopeStoreMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockSimpleAccessScopeStore) Count(ctx context.Context, q *v1.Query) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, q)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockSimpleAccessScopeStoreMockRecorder) Count(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockSimpleAccessScopeStore)(nil).Count), ctx, q)
}

// Delete mocks base method.
func (m *MockSimpleAccessScopeStore) Delete(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockSimpleAccessScopeStoreMockRecorder) Delete(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSimpleAccessScopeStore)(nil).Delete), ctx, id)
}

// Exists mocks base method.
func (m *MockSimpleAccessScopeStore) Exists(ctx context.Context, id string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockSimpleAccessScopeStoreMockRecorder) Exists(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockSimpleAccessScopeStore)(nil).Exists), ctx, id)
}

// Get mocks base method.
func (m *MockSimpleAccessScopeStore) Get(ctx context.Context, id string) (*storage.SimpleAccessScope, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*storage.SimpleAccessScope)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockSimpleAccessScopeStoreMockRecorder) Get(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSimpleAccessScopeStore)(nil).Get), ctx, id)
}

// Search mocks base method.
func (m *MockSimpleAccessScopeStore) Search(ctx context.Context, q *v1.Query) ([]search.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, q)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockSimpleAccessScopeStoreMockRecorder) Search(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockSimpleAccessScopeStore)(nil).Search), ctx, q)
}

// Upsert mocks base method.
func (m *MockSimpleAccessScopeStore) Upsert(ctx context.Context, obj *storage.SimpleAccessScope) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", ctx, obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockSimpleAccessScopeStoreMockRecorder) Upsert(ctx, obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockSimpleAccessScopeStore)(nil).Upsert), ctx, obj)
}

// UpsertMany mocks base method.
func (m *MockSimpleAccessScopeStore) UpsertMany(ctx context.Context, obj []*storage.SimpleAccessScope) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertMany", ctx, obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertMany indicates an expected call of UpsertMany.
func (mr *MockSimpleAccessScopeStoreMockRecorder) UpsertMany(ctx, obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertMany", reflect.TypeOf((*MockSimpleAccessScopeStore)(nil).UpsertMany), ctx, obj)
}

// Walk mocks base method.
func (m *MockSimpleAccessScopeStore) Walk(ctx context.Context, fn func(*storage.SimpleAccessScope) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Walk", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// Walk indicates an expected call of Walk.
func (mr *MockSimpleAccessScopeStoreMockRecorder) Walk(ctx, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Walk", reflect.TypeOf((*MockSimpleAccessScopeStore)(nil).Walk), ctx, fn)
}

// MockRoleStore is a mock of RoleStore interface.
type MockRoleStore struct {
	ctrl     *gomock.Controller
	recorder *MockRoleStoreMockRecorder
}

// MockRoleStoreMockRecorder is the mock recorder for MockRoleStore.
type MockRoleStoreMockRecorder struct {
	mock *MockRoleStore
}

// NewMockRoleStore creates a new mock instance.
func NewMockRoleStore(ctrl *gomock.Controller) *MockRoleStore {
	mock := &MockRoleStore{ctrl: ctrl}
	mock.recorder = &MockRoleStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoleStore) EXPECT() *MockRoleStoreMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockRoleStore) Count(ctx context.Context, q *v1.Query) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, q)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockRoleStoreMockRecorder) Count(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockRoleStore)(nil).Count), ctx, q)
}

// Delete mocks base method.
func (m *MockRoleStore) Delete(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRoleStoreMockRecorder) Delete(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRoleStore)(nil).Delete), ctx, id)
}

// Get mocks base method.
func (m *MockRoleStore) Get(ctx context.Context, id string) (*storage.Role, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*storage.Role)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockRoleStoreMockRecorder) Get(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRoleStore)(nil).Get), ctx, id)
}

// Search mocks base method.
func (m *MockRoleStore) Search(ctx context.Context, q *v1.Query) ([]search.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, q)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockRoleStoreMockRecorder) Search(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockRoleStore)(nil).Search), ctx, q)
}

// Upsert mocks base method.
func (m *MockRoleStore) Upsert(ctx context.Context, obj *storage.Role) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", ctx, obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockRoleStoreMockRecorder) Upsert(ctx, obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockRoleStore)(nil).Upsert), ctx, obj)
}

// UpsertMany mocks base method.
func (m *MockRoleStore) UpsertMany(ctx context.Context, obj []*storage.Role) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertMany", ctx, obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertMany indicates an expected call of UpsertMany.
func (mr *MockRoleStoreMockRecorder) UpsertMany(ctx, obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertMany", reflect.TypeOf((*MockRoleStore)(nil).UpsertMany), ctx, obj)
}

// Walk mocks base method.
func (m *MockRoleStore) Walk(ctx context.Context, fn func(*storage.Role) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Walk", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// Walk indicates an expected call of Walk.
func (mr *MockRoleStoreMockRecorder) Walk(ctx, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Walk", reflect.TypeOf((*MockRoleStore)(nil).Walk), ctx, fn)
}
