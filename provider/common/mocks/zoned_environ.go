// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/provider/common (interfaces: ZonedEnviron)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	constraints "github.com/juju/juju/core/constraints"
	instance "github.com/juju/juju/core/instance"
	network "github.com/juju/juju/core/network"
	environs "github.com/juju/juju/environs"
	config "github.com/juju/juju/environs/config"
	context "github.com/juju/juju/environs/envcontext"
	instances "github.com/juju/juju/environs/instances"
	storage "github.com/juju/juju/internal/storage"
	version "github.com/juju/version/v2"
	gomock "go.uber.org/mock/gomock"
)

// MockZonedEnviron is a mock of ZonedEnviron interface.
type MockZonedEnviron struct {
	ctrl     *gomock.Controller
	recorder *MockZonedEnvironMockRecorder
}

// MockZonedEnvironMockRecorder is the mock recorder for MockZonedEnviron.
type MockZonedEnvironMockRecorder struct {
	mock *MockZonedEnviron
}

// NewMockZonedEnviron creates a new mock instance.
func NewMockZonedEnviron(ctrl *gomock.Controller) *MockZonedEnviron {
	mock := &MockZonedEnviron{ctrl: ctrl}
	mock.recorder = &MockZonedEnvironMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockZonedEnviron) EXPECT() *MockZonedEnvironMockRecorder {
	return m.recorder
}

// AdoptResources mocks base method.
func (m *MockZonedEnviron) AdoptResources(arg0 context.ProviderCallContext, arg1 string, arg2 version.Number) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdoptResources", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AdoptResources indicates an expected call of AdoptResources.
func (mr *MockZonedEnvironMockRecorder) AdoptResources(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdoptResources", reflect.TypeOf((*MockZonedEnviron)(nil).AdoptResources), arg0, arg1, arg2)
}

// AllInstances mocks base method.
func (m *MockZonedEnviron) AllInstances(arg0 context.ProviderCallContext) ([]instances.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllInstances", arg0)
	ret0, _ := ret[0].([]instances.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllInstances indicates an expected call of AllInstances.
func (mr *MockZonedEnvironMockRecorder) AllInstances(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllInstances", reflect.TypeOf((*MockZonedEnviron)(nil).AllInstances), arg0)
}

// AllRunningInstances mocks base method.
func (m *MockZonedEnviron) AllRunningInstances(arg0 context.ProviderCallContext) ([]instances.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllRunningInstances", arg0)
	ret0, _ := ret[0].([]instances.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllRunningInstances indicates an expected call of AllRunningInstances.
func (mr *MockZonedEnvironMockRecorder) AllRunningInstances(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllRunningInstances", reflect.TypeOf((*MockZonedEnviron)(nil).AllRunningInstances), arg0)
}

// AvailabilityZones mocks base method.
func (m *MockZonedEnviron) AvailabilityZones(arg0 context.ProviderCallContext) (network.AvailabilityZones, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailabilityZones", arg0)
	ret0, _ := ret[0].(network.AvailabilityZones)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AvailabilityZones indicates an expected call of AvailabilityZones.
func (mr *MockZonedEnvironMockRecorder) AvailabilityZones(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailabilityZones", reflect.TypeOf((*MockZonedEnviron)(nil).AvailabilityZones), arg0)
}

// Bootstrap mocks base method.
func (m *MockZonedEnviron) Bootstrap(arg0 environs.BootstrapContext, arg1 context.ProviderCallContext, arg2 environs.BootstrapParams) (*environs.BootstrapResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bootstrap", arg0, arg1, arg2)
	ret0, _ := ret[0].(*environs.BootstrapResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Bootstrap indicates an expected call of Bootstrap.
func (mr *MockZonedEnvironMockRecorder) Bootstrap(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bootstrap", reflect.TypeOf((*MockZonedEnviron)(nil).Bootstrap), arg0, arg1, arg2)
}

// Config mocks base method.
func (m *MockZonedEnviron) Config() *config.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*config.Config)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MockZonedEnvironMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockZonedEnviron)(nil).Config))
}

// ConstraintsValidator mocks base method.
func (m *MockZonedEnviron) ConstraintsValidator(arg0 context.ProviderCallContext) (constraints.Validator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConstraintsValidator", arg0)
	ret0, _ := ret[0].(constraints.Validator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConstraintsValidator indicates an expected call of ConstraintsValidator.
func (mr *MockZonedEnvironMockRecorder) ConstraintsValidator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConstraintsValidator", reflect.TypeOf((*MockZonedEnviron)(nil).ConstraintsValidator), arg0)
}

// ControllerInstances mocks base method.
func (m *MockZonedEnviron) ControllerInstances(arg0 context.ProviderCallContext, arg1 string) ([]instance.Id, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerInstances", arg0, arg1)
	ret0, _ := ret[0].([]instance.Id)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerInstances indicates an expected call of ControllerInstances.
func (mr *MockZonedEnvironMockRecorder) ControllerInstances(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerInstances", reflect.TypeOf((*MockZonedEnviron)(nil).ControllerInstances), arg0, arg1)
}

// Create mocks base method.
func (m *MockZonedEnviron) Create(arg0 context.ProviderCallContext, arg1 environs.CreateParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockZonedEnvironMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockZonedEnviron)(nil).Create), arg0, arg1)
}

// DeriveAvailabilityZones mocks base method.
func (m *MockZonedEnviron) DeriveAvailabilityZones(arg0 context.ProviderCallContext, arg1 environs.StartInstanceParams) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeriveAvailabilityZones", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeriveAvailabilityZones indicates an expected call of DeriveAvailabilityZones.
func (mr *MockZonedEnvironMockRecorder) DeriveAvailabilityZones(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeriveAvailabilityZones", reflect.TypeOf((*MockZonedEnviron)(nil).DeriveAvailabilityZones), arg0, arg1)
}

// Destroy mocks base method.
func (m *MockZonedEnviron) Destroy(arg0 context.ProviderCallContext) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy.
func (mr *MockZonedEnvironMockRecorder) Destroy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockZonedEnviron)(nil).Destroy), arg0)
}

// DestroyController mocks base method.
func (m *MockZonedEnviron) DestroyController(arg0 context.ProviderCallContext, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroyController", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyController indicates an expected call of DestroyController.
func (mr *MockZonedEnvironMockRecorder) DestroyController(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyController", reflect.TypeOf((*MockZonedEnviron)(nil).DestroyController), arg0, arg1)
}

// InstanceAvailabilityZoneNames mocks base method.
func (m *MockZonedEnviron) InstanceAvailabilityZoneNames(arg0 context.ProviderCallContext, arg1 []instance.Id) (map[instance.Id]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceAvailabilityZoneNames", arg0, arg1)
	ret0, _ := ret[0].(map[instance.Id]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstanceAvailabilityZoneNames indicates an expected call of InstanceAvailabilityZoneNames.
func (mr *MockZonedEnvironMockRecorder) InstanceAvailabilityZoneNames(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceAvailabilityZoneNames", reflect.TypeOf((*MockZonedEnviron)(nil).InstanceAvailabilityZoneNames), arg0, arg1)
}

// InstanceTypes mocks base method.
func (m *MockZonedEnviron) InstanceTypes(arg0 context.ProviderCallContext, arg1 constraints.Value) (instances.InstanceTypesWithCostMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceTypes", arg0, arg1)
	ret0, _ := ret[0].(instances.InstanceTypesWithCostMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstanceTypes indicates an expected call of InstanceTypes.
func (mr *MockZonedEnvironMockRecorder) InstanceTypes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceTypes", reflect.TypeOf((*MockZonedEnviron)(nil).InstanceTypes), arg0, arg1)
}

// Instances mocks base method.
func (m *MockZonedEnviron) Instances(arg0 context.ProviderCallContext, arg1 []instance.Id) ([]instances.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Instances", arg0, arg1)
	ret0, _ := ret[0].([]instances.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Instances indicates an expected call of Instances.
func (mr *MockZonedEnvironMockRecorder) Instances(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Instances", reflect.TypeOf((*MockZonedEnviron)(nil).Instances), arg0, arg1)
}

// PrecheckInstance mocks base method.
func (m *MockZonedEnviron) PrecheckInstance(arg0 context.ProviderCallContext, arg1 environs.PrecheckInstanceParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrecheckInstance", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PrecheckInstance indicates an expected call of PrecheckInstance.
func (mr *MockZonedEnvironMockRecorder) PrecheckInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrecheckInstance", reflect.TypeOf((*MockZonedEnviron)(nil).PrecheckInstance), arg0, arg1)
}

// PrepareForBootstrap mocks base method.
func (m *MockZonedEnviron) PrepareForBootstrap(arg0 environs.BootstrapContext, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareForBootstrap", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PrepareForBootstrap indicates an expected call of PrepareForBootstrap.
func (mr *MockZonedEnvironMockRecorder) PrepareForBootstrap(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareForBootstrap", reflect.TypeOf((*MockZonedEnviron)(nil).PrepareForBootstrap), arg0, arg1)
}

// Provider mocks base method.
func (m *MockZonedEnviron) Provider() environs.EnvironProvider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Provider")
	ret0, _ := ret[0].(environs.EnvironProvider)
	return ret0
}

// Provider indicates an expected call of Provider.
func (mr *MockZonedEnvironMockRecorder) Provider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Provider", reflect.TypeOf((*MockZonedEnviron)(nil).Provider))
}

// SetConfig mocks base method.
func (m *MockZonedEnviron) SetConfig(arg0 *config.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetConfig", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetConfig indicates an expected call of SetConfig.
func (mr *MockZonedEnvironMockRecorder) SetConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfig", reflect.TypeOf((*MockZonedEnviron)(nil).SetConfig), arg0)
}

// StartInstance mocks base method.
func (m *MockZonedEnviron) StartInstance(arg0 context.ProviderCallContext, arg1 environs.StartInstanceParams) (*environs.StartInstanceResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartInstance", arg0, arg1)
	ret0, _ := ret[0].(*environs.StartInstanceResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartInstance indicates an expected call of StartInstance.
func (mr *MockZonedEnvironMockRecorder) StartInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartInstance", reflect.TypeOf((*MockZonedEnviron)(nil).StartInstance), arg0, arg1)
}

// StopInstances mocks base method.
func (m *MockZonedEnviron) StopInstances(arg0 context.ProviderCallContext, arg1 ...instance.Id) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StopInstances", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopInstances indicates an expected call of StopInstances.
func (mr *MockZonedEnvironMockRecorder) StopInstances(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopInstances", reflect.TypeOf((*MockZonedEnviron)(nil).StopInstances), varargs...)
}

// StorageProvider mocks base method.
func (m *MockZonedEnviron) StorageProvider(arg0 storage.ProviderType) (storage.Provider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageProvider", arg0)
	ret0, _ := ret[0].(storage.Provider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageProvider indicates an expected call of StorageProvider.
func (mr *MockZonedEnvironMockRecorder) StorageProvider(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageProvider", reflect.TypeOf((*MockZonedEnviron)(nil).StorageProvider), arg0)
}

// StorageProviderTypes mocks base method.
func (m *MockZonedEnviron) StorageProviderTypes() ([]storage.ProviderType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageProviderTypes")
	ret0, _ := ret[0].([]storage.ProviderType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageProviderTypes indicates an expected call of StorageProviderTypes.
func (mr *MockZonedEnvironMockRecorder) StorageProviderTypes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageProviderTypes", reflect.TypeOf((*MockZonedEnviron)(nil).StorageProviderTypes))
}
