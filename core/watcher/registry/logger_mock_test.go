// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/core/watcher/registry (interfaces: Logger)
//
// Generated by this command:
//
//	mockgen -package registry -destination logger_mock_test.go github.com/juju/juju/core/watcher/registry Logger
//

// Package registry is a generated GoMock package.
package registry

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// IsTraceEnabled mocks base method.
func (m *MockLogger) IsTraceEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsTraceEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsTraceEnabled indicates an expected call of IsTraceEnabled.
func (mr *MockLoggerMockRecorder) IsTraceEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsTraceEnabled", reflect.TypeOf((*MockLogger)(nil).IsTraceEnabled))
}

// Tracef mocks base method.
func (m *MockLogger) Tracef(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Tracef", varargs...)
}

// Tracef indicates an expected call of Tracef.
func (mr *MockLoggerMockRecorder) Tracef(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tracef", reflect.TypeOf((*MockLogger)(nil).Tracef), varargs...)
}