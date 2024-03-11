// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/worker/bootstrap (interfaces: ControllerConfigService,FlagService,ObjectStoreGetter,SystemState,HTTPClient,CredentialService,CloudService,StorageService,ApplicationService,SpaceService)
//
// Generated by this command:
//
//	mockgen -package bootstrap -destination bootstrap_mock_test.go github.com/juju/juju/internal/worker/bootstrap ControllerConfigService,FlagService,ObjectStoreGetter,SystemState,HTTPClient,CredentialService,CloudService,StorageService,ApplicationService,SpaceService
//

// Package bootstrap is a generated GoMock package.
package bootstrap

import (
	context "context"
	http "net/http"
	reflect "reflect"

	charm "github.com/juju/charm/v13"
	cloud "github.com/juju/juju/cloud"
	controller "github.com/juju/juju/controller"
	network "github.com/juju/juju/core/network"
	objectstore "github.com/juju/juju/core/objectstore"
	service "github.com/juju/juju/domain/application/service"
	credential "github.com/juju/juju/domain/credential"
	service0 "github.com/juju/juju/domain/storage/service"
	bootstrap "github.com/juju/juju/internal/bootstrap"
	services "github.com/juju/juju/internal/charm/services"
	storage "github.com/juju/juju/internal/storage"
	state "github.com/juju/juju/state"
	binarystorage "github.com/juju/juju/state/binarystorage"
	gomock "go.uber.org/mock/gomock"
)

// MockControllerConfigService is a mock of ControllerConfigService interface.
type MockControllerConfigService struct {
	ctrl     *gomock.Controller
	recorder *MockControllerConfigServiceMockRecorder
}

// MockControllerConfigServiceMockRecorder is the mock recorder for MockControllerConfigService.
type MockControllerConfigServiceMockRecorder struct {
	mock *MockControllerConfigService
}

// NewMockControllerConfigService creates a new mock instance.
func NewMockControllerConfigService(ctrl *gomock.Controller) *MockControllerConfigService {
	mock := &MockControllerConfigService{ctrl: ctrl}
	mock.recorder = &MockControllerConfigServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockControllerConfigService) EXPECT() *MockControllerConfigServiceMockRecorder {
	return m.recorder
}

// ControllerConfig mocks base method.
func (m *MockControllerConfigService) ControllerConfig(arg0 context.Context) (controller.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig", arg0)
	ret0, _ := ret[0].(controller.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerConfig indicates an expected call of ControllerConfig.
func (mr *MockControllerConfigServiceMockRecorder) ControllerConfig(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockControllerConfigService)(nil).ControllerConfig), arg0)
}

// MockFlagService is a mock of FlagService interface.
type MockFlagService struct {
	ctrl     *gomock.Controller
	recorder *MockFlagServiceMockRecorder
}

// MockFlagServiceMockRecorder is the mock recorder for MockFlagService.
type MockFlagServiceMockRecorder struct {
	mock *MockFlagService
}

// NewMockFlagService creates a new mock instance.
func NewMockFlagService(ctrl *gomock.Controller) *MockFlagService {
	mock := &MockFlagService{ctrl: ctrl}
	mock.recorder = &MockFlagServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFlagService) EXPECT() *MockFlagServiceMockRecorder {
	return m.recorder
}

// GetFlag mocks base method.
func (m *MockFlagService) GetFlag(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFlag", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFlag indicates an expected call of GetFlag.
func (mr *MockFlagServiceMockRecorder) GetFlag(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlag", reflect.TypeOf((*MockFlagService)(nil).GetFlag), arg0, arg1)
}

// SetFlag mocks base method.
func (m *MockFlagService) SetFlag(arg0 context.Context, arg1 string, arg2 bool, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetFlag", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetFlag indicates an expected call of SetFlag.
func (mr *MockFlagServiceMockRecorder) SetFlag(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFlag", reflect.TypeOf((*MockFlagService)(nil).SetFlag), arg0, arg1, arg2, arg3)
}

// MockObjectStoreGetter is a mock of ObjectStoreGetter interface.
type MockObjectStoreGetter struct {
	ctrl     *gomock.Controller
	recorder *MockObjectStoreGetterMockRecorder
}

// MockObjectStoreGetterMockRecorder is the mock recorder for MockObjectStoreGetter.
type MockObjectStoreGetterMockRecorder struct {
	mock *MockObjectStoreGetter
}

// NewMockObjectStoreGetter creates a new mock instance.
func NewMockObjectStoreGetter(ctrl *gomock.Controller) *MockObjectStoreGetter {
	mock := &MockObjectStoreGetter{ctrl: ctrl}
	mock.recorder = &MockObjectStoreGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockObjectStoreGetter) EXPECT() *MockObjectStoreGetterMockRecorder {
	return m.recorder
}

// GetObjectStore mocks base method.
func (m *MockObjectStoreGetter) GetObjectStore(arg0 context.Context, arg1 string) (objectstore.ObjectStore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectStore", arg0, arg1)
	ret0, _ := ret[0].(objectstore.ObjectStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectStore indicates an expected call of GetObjectStore.
func (mr *MockObjectStoreGetterMockRecorder) GetObjectStore(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectStore", reflect.TypeOf((*MockObjectStoreGetter)(nil).GetObjectStore), arg0, arg1)
}

// MockSystemState is a mock of SystemState interface.
type MockSystemState struct {
	ctrl     *gomock.Controller
	recorder *MockSystemStateMockRecorder
}

// MockSystemStateMockRecorder is the mock recorder for MockSystemState.
type MockSystemStateMockRecorder struct {
	mock *MockSystemState
}

// NewMockSystemState creates a new mock instance.
func NewMockSystemState(ctrl *gomock.Controller) *MockSystemState {
	mock := &MockSystemState{ctrl: ctrl}
	mock.recorder = &MockSystemStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSystemState) EXPECT() *MockSystemStateMockRecorder {
	return m.recorder
}

// AddApplication mocks base method.
func (m *MockSystemState) AddApplication(arg0 state.AddApplicationArgs, arg1 objectstore.ObjectStore) (bootstrap.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddApplication", arg0, arg1)
	ret0, _ := ret[0].(bootstrap.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddApplication indicates an expected call of AddApplication.
func (mr *MockSystemStateMockRecorder) AddApplication(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddApplication", reflect.TypeOf((*MockSystemState)(nil).AddApplication), arg0, arg1)
}

// ApplyOperation mocks base method.
func (m *MockSystemState) ApplyOperation(arg0 *state.UpdateUnitOperation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyOperation", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyOperation indicates an expected call of ApplyOperation.
func (mr *MockSystemStateMockRecorder) ApplyOperation(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyOperation", reflect.TypeOf((*MockSystemState)(nil).ApplyOperation), arg0)
}

// Charm mocks base method.
func (m *MockSystemState) Charm(arg0 string) (bootstrap.Charm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Charm", arg0)
	ret0, _ := ret[0].(bootstrap.Charm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Charm indicates an expected call of Charm.
func (mr *MockSystemStateMockRecorder) Charm(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Charm", reflect.TypeOf((*MockSystemState)(nil).Charm), arg0)
}

// CloudService mocks base method.
func (m *MockSystemState) CloudService(arg0 string) (bootstrap.CloudService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudService", arg0)
	ret0, _ := ret[0].(bootstrap.CloudService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudService indicates an expected call of CloudService.
func (mr *MockSystemStateMockRecorder) CloudService(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudService", reflect.TypeOf((*MockSystemState)(nil).CloudService), arg0)
}

// ControllerModelUUID mocks base method.
func (m *MockSystemState) ControllerModelUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerModelUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ControllerModelUUID indicates an expected call of ControllerModelUUID.
func (mr *MockSystemStateMockRecorder) ControllerModelUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerModelUUID", reflect.TypeOf((*MockSystemState)(nil).ControllerModelUUID))
}

// Machine mocks base method.
func (m *MockSystemState) Machine(arg0 string) (bootstrap.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Machine", arg0)
	ret0, _ := ret[0].(bootstrap.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Machine indicates an expected call of Machine.
func (mr *MockSystemStateMockRecorder) Machine(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Machine", reflect.TypeOf((*MockSystemState)(nil).Machine), arg0)
}

// Model mocks base method.
func (m *MockSystemState) Model() (bootstrap.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model")
	ret0, _ := ret[0].(bootstrap.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Model indicates an expected call of Model.
func (mr *MockSystemStateMockRecorder) Model() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockSystemState)(nil).Model))
}

// ModelUUID mocks base method.
func (m *MockSystemState) ModelUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ModelUUID indicates an expected call of ModelUUID.
func (mr *MockSystemStateMockRecorder) ModelUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelUUID", reflect.TypeOf((*MockSystemState)(nil).ModelUUID))
}

// PrepareCharmUpload mocks base method.
func (m *MockSystemState) PrepareCharmUpload(arg0 string) (services.UploadedCharm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareCharmUpload", arg0)
	ret0, _ := ret[0].(services.UploadedCharm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareCharmUpload indicates an expected call of PrepareCharmUpload.
func (mr *MockSystemStateMockRecorder) PrepareCharmUpload(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareCharmUpload", reflect.TypeOf((*MockSystemState)(nil).PrepareCharmUpload), arg0)
}

// PrepareLocalCharmUpload mocks base method.
func (m *MockSystemState) PrepareLocalCharmUpload(arg0 string) (*charm.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareLocalCharmUpload", arg0)
	ret0, _ := ret[0].(*charm.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareLocalCharmUpload indicates an expected call of PrepareLocalCharmUpload.
func (mr *MockSystemStateMockRecorder) PrepareLocalCharmUpload(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareLocalCharmUpload", reflect.TypeOf((*MockSystemState)(nil).PrepareLocalCharmUpload), arg0)
}

// SaveCloudService mocks base method.
func (m *MockSystemState) SaveCloudService(arg0 state.SaveCloudServiceArgs) (*state.CloudService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCloudService", arg0)
	ret0, _ := ret[0].(*state.CloudService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveCloudService indicates an expected call of SaveCloudService.
func (mr *MockSystemStateMockRecorder) SaveCloudService(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCloudService", reflect.TypeOf((*MockSystemState)(nil).SaveCloudService), arg0)
}

// SetAPIHostPorts mocks base method.
func (m *MockSystemState) SetAPIHostPorts(arg0 controller.Config, arg1, arg2 []network.SpaceHostPorts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAPIHostPorts", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetAPIHostPorts indicates an expected call of SetAPIHostPorts.
func (mr *MockSystemStateMockRecorder) SetAPIHostPorts(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAPIHostPorts", reflect.TypeOf((*MockSystemState)(nil).SetAPIHostPorts), arg0, arg1, arg2)
}

// ToolsStorage mocks base method.
func (m *MockSystemState) ToolsStorage(arg0 objectstore.ObjectStore) (binarystorage.StorageCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToolsStorage", arg0)
	ret0, _ := ret[0].(binarystorage.StorageCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToolsStorage indicates an expected call of ToolsStorage.
func (mr *MockSystemStateMockRecorder) ToolsStorage(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToolsStorage", reflect.TypeOf((*MockSystemState)(nil).ToolsStorage), arg0)
}

// Unit mocks base method.
func (m *MockSystemState) Unit(arg0 string) (bootstrap.Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unit", arg0)
	ret0, _ := ret[0].(bootstrap.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unit indicates an expected call of Unit.
func (mr *MockSystemStateMockRecorder) Unit(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unit", reflect.TypeOf((*MockSystemState)(nil).Unit), arg0)
}

// UpdateUploadedCharm mocks base method.
func (m *MockSystemState) UpdateUploadedCharm(arg0 state.CharmInfo) (services.UploadedCharm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUploadedCharm", arg0)
	ret0, _ := ret[0].(services.UploadedCharm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUploadedCharm indicates an expected call of UpdateUploadedCharm.
func (mr *MockSystemStateMockRecorder) UpdateUploadedCharm(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUploadedCharm", reflect.TypeOf((*MockSystemState)(nil).UpdateUploadedCharm), arg0)
}

// MockHTTPClient is a mock of HTTPClient interface.
type MockHTTPClient struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPClientMockRecorder
}

// MockHTTPClientMockRecorder is the mock recorder for MockHTTPClient.
type MockHTTPClientMockRecorder struct {
	mock *MockHTTPClient
}

// NewMockHTTPClient creates a new mock instance.
func NewMockHTTPClient(ctrl *gomock.Controller) *MockHTTPClient {
	mock := &MockHTTPClient{ctrl: ctrl}
	mock.recorder = &MockHTTPClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPClient) EXPECT() *MockHTTPClientMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockHTTPClient) Do(arg0 *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockHTTPClientMockRecorder) Do(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockHTTPClient)(nil).Do), arg0)
}

// MockCredentialService is a mock of CredentialService interface.
type MockCredentialService struct {
	ctrl     *gomock.Controller
	recorder *MockCredentialServiceMockRecorder
}

// MockCredentialServiceMockRecorder is the mock recorder for MockCredentialService.
type MockCredentialServiceMockRecorder struct {
	mock *MockCredentialService
}

// NewMockCredentialService creates a new mock instance.
func NewMockCredentialService(ctrl *gomock.Controller) *MockCredentialService {
	mock := &MockCredentialService{ctrl: ctrl}
	mock.recorder = &MockCredentialServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCredentialService) EXPECT() *MockCredentialServiceMockRecorder {
	return m.recorder
}

// CloudCredential mocks base method.
func (m *MockCredentialService) CloudCredential(arg0 context.Context, arg1 credential.ID) (cloud.Credential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudCredential", arg0, arg1)
	ret0, _ := ret[0].(cloud.Credential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudCredential indicates an expected call of CloudCredential.
func (mr *MockCredentialServiceMockRecorder) CloudCredential(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudCredential", reflect.TypeOf((*MockCredentialService)(nil).CloudCredential), arg0, arg1)
}

// MockCloudService is a mock of CloudService interface.
type MockCloudService struct {
	ctrl     *gomock.Controller
	recorder *MockCloudServiceMockRecorder
}

// MockCloudServiceMockRecorder is the mock recorder for MockCloudService.
type MockCloudServiceMockRecorder struct {
	mock *MockCloudService
}

// NewMockCloudService creates a new mock instance.
func NewMockCloudService(ctrl *gomock.Controller) *MockCloudService {
	mock := &MockCloudService{ctrl: ctrl}
	mock.recorder = &MockCloudServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudService) EXPECT() *MockCloudServiceMockRecorder {
	return m.recorder
}

// Cloud mocks base method.
func (m *MockCloudService) Cloud(arg0 context.Context, arg1 string) (*cloud.Cloud, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cloud", arg0, arg1)
	ret0, _ := ret[0].(*cloud.Cloud)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cloud indicates an expected call of Cloud.
func (mr *MockCloudServiceMockRecorder) Cloud(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cloud", reflect.TypeOf((*MockCloudService)(nil).Cloud), arg0, arg1)
}

// MockStorageService is a mock of StorageService interface.
type MockStorageService struct {
	ctrl     *gomock.Controller
	recorder *MockStorageServiceMockRecorder
}

// MockStorageServiceMockRecorder is the mock recorder for MockStorageService.
type MockStorageServiceMockRecorder struct {
	mock *MockStorageService
}

// NewMockStorageService creates a new mock instance.
func NewMockStorageService(ctrl *gomock.Controller) *MockStorageService {
	mock := &MockStorageService{ctrl: ctrl}
	mock.recorder = &MockStorageServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageService) EXPECT() *MockStorageServiceMockRecorder {
	return m.recorder
}

// CreateStoragePool mocks base method.
func (m *MockStorageService) CreateStoragePool(arg0 context.Context, arg1 string, arg2 storage.ProviderType, arg3 service0.PoolAttrs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStoragePool", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateStoragePool indicates an expected call of CreateStoragePool.
func (mr *MockStorageServiceMockRecorder) CreateStoragePool(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStoragePool", reflect.TypeOf((*MockStorageService)(nil).CreateStoragePool), arg0, arg1, arg2, arg3)
}

// MockApplicationService is a mock of ApplicationService interface.
type MockApplicationService struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationServiceMockRecorder
}

// MockApplicationServiceMockRecorder is the mock recorder for MockApplicationService.
type MockApplicationServiceMockRecorder struct {
	mock *MockApplicationService
}

// NewMockApplicationService creates a new mock instance.
func NewMockApplicationService(ctrl *gomock.Controller) *MockApplicationService {
	mock := &MockApplicationService{ctrl: ctrl}
	mock.recorder = &MockApplicationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationService) EXPECT() *MockApplicationServiceMockRecorder {
	return m.recorder
}

// CreateApplication mocks base method.
func (m *MockApplicationService) CreateApplication(arg0 context.Context, arg1 string, arg2 service.AddApplicationParams, arg3 ...service.AddUnitParams) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateApplication", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateApplication indicates an expected call of CreateApplication.
func (mr *MockApplicationServiceMockRecorder) CreateApplication(arg0, arg1, arg2 any, arg3 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateApplication", reflect.TypeOf((*MockApplicationService)(nil).CreateApplication), varargs...)
}

// MockSpaceService is a mock of SpaceService interface.
type MockSpaceService struct {
	ctrl     *gomock.Controller
	recorder *MockSpaceServiceMockRecorder
}

// MockSpaceServiceMockRecorder is the mock recorder for MockSpaceService.
type MockSpaceServiceMockRecorder struct {
	mock *MockSpaceService
}

// NewMockSpaceService creates a new mock instance.
func NewMockSpaceService(ctrl *gomock.Controller) *MockSpaceService {
	mock := &MockSpaceService{ctrl: ctrl}
	mock.recorder = &MockSpaceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpaceService) EXPECT() *MockSpaceServiceMockRecorder {
	return m.recorder
}

// GetAllSpaces mocks base method.
func (m *MockSpaceService) GetAllSpaces(arg0 context.Context) (network.SpaceInfos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSpaces", arg0)
	ret0, _ := ret[0].(network.SpaceInfos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSpaces indicates an expected call of GetAllSpaces.
func (mr *MockSpaceServiceMockRecorder) GetAllSpaces(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSpaces", reflect.TypeOf((*MockSpaceService)(nil).GetAllSpaces), arg0)
}

// Space mocks base method.
func (m *MockSpaceService) Space(arg0 context.Context, arg1 string) (*network.SpaceInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Space", arg0, arg1)
	ret0, _ := ret[0].(*network.SpaceInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Space indicates an expected call of Space.
func (mr *MockSpaceServiceMockRecorder) Space(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Space", reflect.TypeOf((*MockSpaceService)(nil).Space), arg0, arg1)
}

// SpaceByName mocks base method.
func (m *MockSpaceService) SpaceByName(arg0 context.Context, arg1 string) (*network.SpaceInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpaceByName", arg0, arg1)
	ret0, _ := ret[0].(*network.SpaceInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SpaceByName indicates an expected call of SpaceByName.
func (mr *MockSpaceServiceMockRecorder) SpaceByName(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpaceByName", reflect.TypeOf((*MockSpaceService)(nil).SpaceByName), arg0, arg1)
}
