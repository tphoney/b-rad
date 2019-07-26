// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/{{github}}/client (interfaces: Client)

// Package mocks is a generated GoMock package.
package mocks

import (
	types "github.com/{{github}}/types"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// {{child}} mocks base method
func (m *MockClient) {{child}}(arg0, arg1, arg2 int64) (*types.{{child}}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "{{child}}", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.{{child}})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// {{child}} indicates an expected call of {{child}}
func (mr *MockClientMockRecorder) {{child}}(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "{{child}}", reflect.TypeOf((*MockClient)(nil).{{child}}), arg0, arg1, arg2)
}

// {{child}}Create mocks base method
func (m *MockClient) {{child}}Create(arg0, arg1 int64, arg2 *types.{{child}}) (*types.{{child}}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "{{child}}Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.{{child}})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// {{child}}Create indicates an expected call of {{child}}Create
func (mr *MockClientMockRecorder) {{child}}Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "{{child}}Create", reflect.TypeOf((*MockClient)(nil).{{child}}Create), arg0, arg1, arg2)
}

// {{child}}Delete mocks base method
func (m *MockClient) {{child}}Delete(arg0, arg1, arg2 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "{{child}}Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// {{child}}Delete indicates an expected call of {{child}}Delete
func (mr *MockClientMockRecorder) {{child}}Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "{{child}}Delete", reflect.TypeOf((*MockClient)(nil).{{child}}Delete), arg0, arg1, arg2)
}

// {{child}}List mocks base method
func (m *MockClient) {{child}}List(arg0, arg1 int64) ([]*types.{{child}}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "{{child}}List", arg0, arg1)
	ret0, _ := ret[0].([]*types.{{child}})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// {{child}}List indicates an expected call of {{child}}List
func (mr *MockClientMockRecorder) {{child}}List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "{{child}}List", reflect.TypeOf((*MockClient)(nil).{{child}}List), arg0, arg1)
}

// {{child}}Update mocks base method
func (m *MockClient) {{child}}Update(arg0, arg1, arg2 int64, arg3 *types.{{child}}Input) (*types.{{child}}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "{{child}}Update", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*types.{{child}})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// {{child}}Update indicates an expected call of {{child}}Update
func (mr *MockClientMockRecorder) {{child}}Update(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "{{child}}Update", reflect.TypeOf((*MockClient)(nil).{{child}}Update), arg0, arg1, arg2, arg3)
}

// {{parent}} mocks base method
func (m *MockClient) {{parent}}(arg0, arg1 int64) (*types.{{parent}}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "{{parent}}", arg0, arg1)
	ret0, _ := ret[0].(*types.{{parent}})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// {{parent}} indicates an expected call of {{parent}}
func (mr *MockClientMockRecorder) {{parent}}(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "{{parent}}", reflect.TypeOf((*MockClient)(nil).{{parent}}), arg0, arg1)
}

// {{parent}}Create mocks base method
func (m *MockClient) {{parent}}Create(arg0 int64, arg1 *types.{{parent}}) (*types.{{parent}}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "{{parent}}Create", arg0, arg1)
	ret0, _ := ret[0].(*types.{{parent}})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// {{parent}}Create indicates an expected call of {{parent}}Create
func (mr *MockClientMockRecorder) {{parent}}Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "{{parent}}Create", reflect.TypeOf((*MockClient)(nil).{{parent}}Create), arg0, arg1)
}

// {{parent}}Delete mocks base method
func (m *MockClient) {{parent}}Delete(arg0, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "{{parent}}Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// {{parent}}Delete indicates an expected call of {{parent}}Delete
func (mr *MockClientMockRecorder) {{parent}}Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "{{parent}}Delete", reflect.TypeOf((*MockClient)(nil).{{parent}}Delete), arg0, arg1)
}

// {{parent}}List mocks base method
func (m *MockClient) {{parent}}List(arg0 int64) ([]*types.{{parent}}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "{{parent}}List", arg0)
	ret0, _ := ret[0].([]*types.{{parent}})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// {{parent}}List indicates an expected call of {{parent}}List
func (mr *MockClientMockRecorder) {{parent}}List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "{{parent}}List", reflect.TypeOf((*MockClient)(nil).{{parent}}List), arg0)
}

// {{parent}}Update mocks base method
func (m *MockClient) {{parent}}Update(arg0, arg1 int64, arg2 *types.{{parent}}Input) (*types.{{parent}}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "{{parent}}Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.{{parent}})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// {{parent}}Update indicates an expected call of {{parent}}Update
func (mr *MockClientMockRecorder) {{parent}}Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "{{parent}}Update", reflect.TypeOf((*MockClient)(nil).{{parent}}Update), arg0, arg1, arg2)
}

// Login mocks base method
func (m *MockClient) Login(arg0, arg1 string) (*types.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(*types.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login
func (mr *MockClientMockRecorder) Login(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockClient)(nil).Login), arg0, arg1)
}

// Member mocks base method
func (m *MockClient) Member(arg0 int64, arg1 string) (*types.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Member", arg0, arg1)
	ret0, _ := ret[0].(*types.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Member indicates an expected call of Member
func (mr *MockClientMockRecorder) Member(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Member", reflect.TypeOf((*MockClient)(nil).Member), arg0, arg1)
}

// MemberCreate mocks base method
func (m *MockClient) MemberCreate(arg0 *types.MembershipInput) (*types.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MemberCreate", arg0)
	ret0, _ := ret[0].(*types.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MemberCreate indicates an expected call of MemberCreate
func (mr *MockClientMockRecorder) MemberCreate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MemberCreate", reflect.TypeOf((*MockClient)(nil).MemberCreate), arg0)
}

// MemberDelete mocks base method
func (m *MockClient) MemberDelete(arg0 int64, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MemberDelete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MemberDelete indicates an expected call of MemberDelete
func (mr *MockClientMockRecorder) MemberDelete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MemberDelete", reflect.TypeOf((*MockClient)(nil).MemberDelete), arg0, arg1)
}

// MemberList mocks base method
func (m *MockClient) MemberList(arg0 int64) ([]*types.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MemberList", arg0)
	ret0, _ := ret[0].([]*types.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MemberList indicates an expected call of MemberList
func (mr *MockClientMockRecorder) MemberList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MemberList", reflect.TypeOf((*MockClient)(nil).MemberList), arg0)
}

// MemberUpdate mocks base method
func (m *MockClient) MemberUpdate(arg0 *types.MembershipInput) (*types.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MemberUpdate", arg0)
	ret0, _ := ret[0].(*types.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MemberUpdate indicates an expected call of MemberUpdate
func (mr *MockClientMockRecorder) MemberUpdate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MemberUpdate", reflect.TypeOf((*MockClient)(nil).MemberUpdate), arg0)
}

// Project mocks base method
func (m *MockClient) Project(arg0 int64) (*types.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Project", arg0)
	ret0, _ := ret[0].(*types.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Project indicates an expected call of Project
func (mr *MockClientMockRecorder) Project(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Project", reflect.TypeOf((*MockClient)(nil).Project), arg0)
}

// ProjectCreate mocks base method
func (m *MockClient) ProjectCreate(arg0 *types.Project) (*types.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectCreate", arg0)
	ret0, _ := ret[0].(*types.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectCreate indicates an expected call of ProjectCreate
func (mr *MockClientMockRecorder) ProjectCreate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectCreate", reflect.TypeOf((*MockClient)(nil).ProjectCreate), arg0)
}

// ProjectDelete mocks base method
func (m *MockClient) ProjectDelete(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProjectDelete indicates an expected call of ProjectDelete
func (mr *MockClientMockRecorder) ProjectDelete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectDelete", reflect.TypeOf((*MockClient)(nil).ProjectDelete), arg0)
}

// ProjectList mocks base method
func (m *MockClient) ProjectList() ([]*types.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectList")
	ret0, _ := ret[0].([]*types.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectList indicates an expected call of ProjectList
func (mr *MockClientMockRecorder) ProjectList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectList", reflect.TypeOf((*MockClient)(nil).ProjectList))
}

// ProjectUpdate mocks base method
func (m *MockClient) ProjectUpdate(arg0 int64, arg1 *types.ProjectInput) (*types.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectUpdate", arg0, arg1)
	ret0, _ := ret[0].(*types.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectUpdate indicates an expected call of ProjectUpdate
func (mr *MockClientMockRecorder) ProjectUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectUpdate", reflect.TypeOf((*MockClient)(nil).ProjectUpdate), arg0, arg1)
}

// Register mocks base method
func (m *MockClient) Register(arg0, arg1 string) (*types.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(*types.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register
func (mr *MockClientMockRecorder) Register(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockClient)(nil).Register), arg0, arg1)
}

// Self mocks base method
func (m *MockClient) Self() (*types.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Self")
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Self indicates an expected call of Self
func (mr *MockClientMockRecorder) Self() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Self", reflect.TypeOf((*MockClient)(nil).Self))
}

// Token mocks base method
func (m *MockClient) Token() (*types.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Token")
	ret0, _ := ret[0].(*types.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Token indicates an expected call of Token
func (mr *MockClientMockRecorder) Token() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Token", reflect.TypeOf((*MockClient)(nil).Token))
}

// User mocks base method
func (m *MockClient) User(arg0 string) (*types.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "User", arg0)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// User indicates an expected call of User
func (mr *MockClientMockRecorder) User(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "User", reflect.TypeOf((*MockClient)(nil).User), arg0)
}

// UserCreate mocks base method
func (m *MockClient) UserCreate(arg0 *types.User) (*types.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserCreate", arg0)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserCreate indicates an expected call of UserCreate
func (mr *MockClientMockRecorder) UserCreate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserCreate", reflect.TypeOf((*MockClient)(nil).UserCreate), arg0)
}

// UserDelete mocks base method
func (m *MockClient) UserDelete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UserDelete indicates an expected call of UserDelete
func (mr *MockClientMockRecorder) UserDelete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserDelete", reflect.TypeOf((*MockClient)(nil).UserDelete), arg0)
}

// UserList mocks base method
func (m *MockClient) UserList() ([]*types.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserList")
	ret0, _ := ret[0].([]*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserList indicates an expected call of UserList
func (mr *MockClientMockRecorder) UserList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserList", reflect.TypeOf((*MockClient)(nil).UserList))
}

// UserUpdate mocks base method
func (m *MockClient) UserUpdate(arg0 string, arg1 *types.UserInput) (*types.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserUpdate", arg0, arg1)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserUpdate indicates an expected call of UserUpdate
func (mr *MockClientMockRecorder) UserUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserUpdate", reflect.TypeOf((*MockClient)(nil).UserUpdate), arg0, arg1)
}
