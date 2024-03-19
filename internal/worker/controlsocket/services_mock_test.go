// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/worker/controlsocket (interfaces: UserService,PermissionService)
//
// Generated by this command:
//
//	mockgen -package controlsocket -destination services_mock_test.go github.com/juju/juju/internal/worker/controlsocket UserService,PermissionService
//

// Package controlsocket is a generated GoMock package.
package controlsocket

import (
	context "context"
	reflect "reflect"

	permission "github.com/juju/juju/core/permission"
	user "github.com/juju/juju/core/user"
	service "github.com/juju/juju/domain/user/service"
	auth "github.com/juju/juju/internal/auth"
	gomock "go.uber.org/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// AddUser mocks base method.
func (m *MockUserService) AddUser(arg0 context.Context, arg1 service.AddUserArg) (user.UUID, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", arg0, arg1)
	ret0, _ := ret[0].(user.UUID)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddUser indicates an expected call of AddUser.
func (mr *MockUserServiceMockRecorder) AddUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockUserService)(nil).AddUser), arg0, arg1)
}

// GetUserByAuth mocks base method.
func (m *MockUserService) GetUserByAuth(arg0 context.Context, arg1 string, arg2 auth.Password) (user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByAuth", arg0, arg1, arg2)
	ret0, _ := ret[0].(user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByAuth indicates an expected call of GetUserByAuth.
func (mr *MockUserServiceMockRecorder) GetUserByAuth(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByAuth", reflect.TypeOf((*MockUserService)(nil).GetUserByAuth), arg0, arg1, arg2)
}

// GetUserByName mocks base method.
func (m *MockUserService) GetUserByName(arg0 context.Context, arg1 string) (user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByName", arg0, arg1)
	ret0, _ := ret[0].(user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByName indicates an expected call of GetUserByName.
func (mr *MockUserServiceMockRecorder) GetUserByName(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByName", reflect.TypeOf((*MockUserService)(nil).GetUserByName), arg0, arg1)
}

// RemoveUser mocks base method.
func (m *MockUserService) RemoveUser(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveUser indicates an expected call of RemoveUser.
func (mr *MockUserServiceMockRecorder) RemoveUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUser", reflect.TypeOf((*MockUserService)(nil).RemoveUser), arg0, arg1)
}

// MockPermissionService is a mock of PermissionService interface.
type MockPermissionService struct {
	ctrl     *gomock.Controller
	recorder *MockPermissionServiceMockRecorder
}

// MockPermissionServiceMockRecorder is the mock recorder for MockPermissionService.
type MockPermissionServiceMockRecorder struct {
	mock *MockPermissionService
}

// NewMockPermissionService creates a new mock instance.
func NewMockPermissionService(ctrl *gomock.Controller) *MockPermissionService {
	mock := &MockPermissionService{ctrl: ctrl}
	mock.recorder = &MockPermissionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPermissionService) EXPECT() *MockPermissionServiceMockRecorder {
	return m.recorder
}

// AddUserPermission mocks base method.
func (m *MockPermissionService) AddUserPermission(arg0 context.Context, arg1 string, arg2 permission.Access) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUserPermission", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUserPermission indicates an expected call of AddUserPermission.
func (mr *MockPermissionServiceMockRecorder) AddUserPermission(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserPermission", reflect.TypeOf((*MockPermissionService)(nil).AddUserPermission), arg0, arg1, arg2)
}
