// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/controller/crossmodelsecrets (interfaces: SecretsState)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	secrets "github.com/juju/juju/core/secrets"
)

// MockSecretsState is a mock of SecretsState interface.
type MockSecretsState struct {
	ctrl     *gomock.Controller
	recorder *MockSecretsStateMockRecorder
}

// MockSecretsStateMockRecorder is the mock recorder for MockSecretsState.
type MockSecretsStateMockRecorder struct {
	mock *MockSecretsState
}

// NewMockSecretsState creates a new mock instance.
func NewMockSecretsState(ctrl *gomock.Controller) *MockSecretsState {
	mock := &MockSecretsState{ctrl: ctrl}
	mock.recorder = &MockSecretsStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretsState) EXPECT() *MockSecretsStateMockRecorder {
	return m.recorder
}

// GetSecret mocks base method.
func (m *MockSecretsState) GetSecret(arg0 *secrets.URI) (*secrets.SecretMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", arg0)
	ret0, _ := ret[0].(*secrets.SecretMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecret indicates an expected call of GetSecret.
func (mr *MockSecretsStateMockRecorder) GetSecret(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockSecretsState)(nil).GetSecret), arg0)
}

// GetSecretValue mocks base method.
func (m *MockSecretsState) GetSecretValue(arg0 *secrets.URI, arg1 int) (secrets.SecretValue, *secrets.ValueRef, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretValue", arg0, arg1)
	ret0, _ := ret[0].(secrets.SecretValue)
	ret1, _ := ret[1].(*secrets.ValueRef)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSecretValue indicates an expected call of GetSecretValue.
func (mr *MockSecretsStateMockRecorder) GetSecretValue(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretValue", reflect.TypeOf((*MockSecretsState)(nil).GetSecretValue), arg0, arg1)
}