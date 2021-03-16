// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/worker/uniter/runner/jujuc (interfaces: ContextRelation)

// Package jujuc is a generated GoMock package.
package jujuc

import (
	gomock "github.com/golang/mock/gomock"
	params "github.com/juju/juju/apiserver/params"
	life "github.com/juju/juju/core/life"
	relation "github.com/juju/juju/core/relation"
	reflect "reflect"
)

// MockContextRelation is a mock of ContextRelation interface
type MockContextRelation struct {
	ctrl     *gomock.Controller
	recorder *MockContextRelationMockRecorder
}

// MockContextRelationMockRecorder is the mock recorder for MockContextRelation
type MockContextRelationMockRecorder struct {
	mock *MockContextRelation
}

// NewMockContextRelation creates a new mock instance
func NewMockContextRelation(ctrl *gomock.Controller) *MockContextRelation {
	mock := &MockContextRelation{ctrl: ctrl}
	mock.recorder = &MockContextRelationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContextRelation) EXPECT() *MockContextRelationMockRecorder {
	return m.recorder
}

// ApplicationSettings mocks base method
func (m *MockContextRelation) ApplicationSettings() (Settings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationSettings")
	ret0, _ := ret[0].(Settings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplicationSettings indicates an expected call of ApplicationSettings
func (mr *MockContextRelationMockRecorder) ApplicationSettings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationSettings", reflect.TypeOf((*MockContextRelation)(nil).ApplicationSettings))
}

// FakeId mocks base method
func (m *MockContextRelation) FakeId() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FakeId")
	ret0, _ := ret[0].(string)
	return ret0
}

// FakeId indicates an expected call of FakeId
func (mr *MockContextRelationMockRecorder) FakeId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FakeId", reflect.TypeOf((*MockContextRelation)(nil).FakeId))
}

// Id mocks base method
func (m *MockContextRelation) Id() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(int)
	return ret0
}

// Id indicates an expected call of Id
func (mr *MockContextRelationMockRecorder) Id() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockContextRelation)(nil).Id))
}

// Life mocks base method
func (m *MockContextRelation) Life() life.Value {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Life")
	ret0, _ := ret[0].(life.Value)
	return ret0
}

// Life indicates an expected call of Life
func (mr *MockContextRelationMockRecorder) Life() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Life", reflect.TypeOf((*MockContextRelation)(nil).Life))
}

// Name mocks base method
func (m *MockContextRelation) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockContextRelationMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockContextRelation)(nil).Name))
}

// ReadApplicationSettings mocks base method
func (m *MockContextRelation) ReadApplicationSettings(arg0 string) (params.Settings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadApplicationSettings", arg0)
	ret0, _ := ret[0].(params.Settings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadApplicationSettings indicates an expected call of ReadApplicationSettings
func (mr *MockContextRelationMockRecorder) ReadApplicationSettings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadApplicationSettings", reflect.TypeOf((*MockContextRelation)(nil).ReadApplicationSettings), arg0)
}

// ReadSettings mocks base method
func (m *MockContextRelation) ReadSettings(arg0 string) (params.Settings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadSettings", arg0)
	ret0, _ := ret[0].(params.Settings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadSettings indicates an expected call of ReadSettings
func (mr *MockContextRelationMockRecorder) ReadSettings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadSettings", reflect.TypeOf((*MockContextRelation)(nil).ReadSettings), arg0)
}

// RemoteApplicationName mocks base method
func (m *MockContextRelation) RemoteApplicationName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteApplicationName")
	ret0, _ := ret[0].(string)
	return ret0
}

// RemoteApplicationName indicates an expected call of RemoteApplicationName
func (mr *MockContextRelationMockRecorder) RemoteApplicationName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteApplicationName", reflect.TypeOf((*MockContextRelation)(nil).RemoteApplicationName))
}

// SetStatus mocks base method
func (m *MockContextRelation) SetStatus(arg0 relation.Status) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStatus", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStatus indicates an expected call of SetStatus
func (mr *MockContextRelationMockRecorder) SetStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatus", reflect.TypeOf((*MockContextRelation)(nil).SetStatus), arg0)
}

// Settings mocks base method
func (m *MockContextRelation) Settings() (Settings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Settings")
	ret0, _ := ret[0].(Settings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Settings indicates an expected call of Settings
func (mr *MockContextRelationMockRecorder) Settings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Settings", reflect.TypeOf((*MockContextRelation)(nil).Settings))
}

// Suspended mocks base method
func (m *MockContextRelation) Suspended() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Suspended")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Suspended indicates an expected call of Suspended
func (mr *MockContextRelationMockRecorder) Suspended() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Suspended", reflect.TypeOf((*MockContextRelation)(nil).Suspended))
}

// UnitNames mocks base method
func (m *MockContextRelation) UnitNames() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnitNames")
	ret0, _ := ret[0].([]string)
	return ret0
}

// UnitNames indicates an expected call of UnitNames
func (mr *MockContextRelationMockRecorder) UnitNames() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnitNames", reflect.TypeOf((*MockContextRelation)(nil).UnitNames))
}