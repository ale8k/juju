// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/modelconfig (interfaces: SecretBackendService)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/service_mock.go github.com/juju/juju/apiserver/facades/client/modelconfig SecretBackendService
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockSecretBackendService is a mock of SecretBackendService interface.
type MockSecretBackendService struct {
	ctrl     *gomock.Controller
	recorder *MockSecretBackendServiceMockRecorder
}

// MockSecretBackendServiceMockRecorder is the mock recorder for MockSecretBackendService.
type MockSecretBackendServiceMockRecorder struct {
	mock *MockSecretBackendService
}

// NewMockSecretBackendService creates a new mock instance.
func NewMockSecretBackendService(ctrl *gomock.Controller) *MockSecretBackendService {
	mock := &MockSecretBackendService{ctrl: ctrl}
	mock.recorder = &MockSecretBackendServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretBackendService) EXPECT() *MockSecretBackendServiceMockRecorder {
	return m.recorder
}

// PingSecretBackend mocks base method.
func (m *MockSecretBackendService) PingSecretBackend(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PingSecretBackend", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PingSecretBackend indicates an expected call of PingSecretBackend.
func (mr *MockSecretBackendServiceMockRecorder) PingSecretBackend(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PingSecretBackend", reflect.TypeOf((*MockSecretBackendService)(nil).PingSecretBackend), arg0, arg1)
}