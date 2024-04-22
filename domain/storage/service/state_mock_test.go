// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/storage/service (interfaces: StoragePoolState)
//
// Generated by this command:
//
//	mockgen -package service -destination state_mock_test.go github.com/juju/juju/domain/storage/service StoragePoolState
//

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	storage "github.com/juju/juju/domain/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockStoragePoolState is a mock of StoragePoolState interface.
type MockStoragePoolState struct {
	ctrl     *gomock.Controller
	recorder *MockStoragePoolStateMockRecorder
}

// MockStoragePoolStateMockRecorder is the mock recorder for MockStoragePoolState.
type MockStoragePoolStateMockRecorder struct {
	mock *MockStoragePoolState
}

// NewMockStoragePoolState creates a new mock instance.
func NewMockStoragePoolState(ctrl *gomock.Controller) *MockStoragePoolState {
	mock := &MockStoragePoolState{ctrl: ctrl}
	mock.recorder = &MockStoragePoolStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStoragePoolState) EXPECT() *MockStoragePoolStateMockRecorder {
	return m.recorder
}

// CreateStoragePool mocks base method.
func (m *MockStoragePoolState) CreateStoragePool(arg0 context.Context, arg1 storage.StoragePoolDetails) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStoragePool", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateStoragePool indicates an expected call of CreateStoragePool.
func (mr *MockStoragePoolStateMockRecorder) CreateStoragePool(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStoragePool", reflect.TypeOf((*MockStoragePoolState)(nil).CreateStoragePool), arg0, arg1)
}

// DeleteStoragePool mocks base method.
func (m *MockStoragePoolState) DeleteStoragePool(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStoragePool", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStoragePool indicates an expected call of DeleteStoragePool.
func (mr *MockStoragePoolStateMockRecorder) DeleteStoragePool(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStoragePool", reflect.TypeOf((*MockStoragePoolState)(nil).DeleteStoragePool), arg0, arg1)
}

// GetStoragePoolByName mocks base method.
func (m *MockStoragePoolState) GetStoragePoolByName(arg0 context.Context, arg1 string) (storage.StoragePoolDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStoragePoolByName", arg0, arg1)
	ret0, _ := ret[0].(storage.StoragePoolDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStoragePoolByName indicates an expected call of GetStoragePoolByName.
func (mr *MockStoragePoolStateMockRecorder) GetStoragePoolByName(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStoragePoolByName", reflect.TypeOf((*MockStoragePoolState)(nil).GetStoragePoolByName), arg0, arg1)
}

// ListStoragePools mocks base method.
func (m *MockStoragePoolState) ListStoragePools(arg0 context.Context, arg1 storage.Names, arg2 storage.Providers) ([]storage.StoragePoolDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStoragePools", arg0, arg1, arg2)
	ret0, _ := ret[0].([]storage.StoragePoolDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStoragePools indicates an expected call of ListStoragePools.
func (mr *MockStoragePoolStateMockRecorder) ListStoragePools(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStoragePools", reflect.TypeOf((*MockStoragePoolState)(nil).ListStoragePools), arg0, arg1, arg2)
}

// ReplaceStoragePool mocks base method.
func (m *MockStoragePoolState) ReplaceStoragePool(arg0 context.Context, arg1 storage.StoragePoolDetails) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplaceStoragePool", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReplaceStoragePool indicates an expected call of ReplaceStoragePool.
func (mr *MockStoragePoolStateMockRecorder) ReplaceStoragePool(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplaceStoragePool", reflect.TypeOf((*MockStoragePoolState)(nil).ReplaceStoragePool), arg0, arg1)
}