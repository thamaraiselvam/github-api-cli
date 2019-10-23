// Code generated by MockGen. DO NOT EDIT.
// Source: ../cmd/service/service.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	types "github.com/thamaraiselvam/git-api-cli/cmd/types"
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

// GetUser mocks base method
func (m *MockClient) GetUser() (types.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser")
	ret0, _ := ret[0].(types.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser
func (mr *MockClientMockRecorder) GetUser() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockClient)(nil).GetUser))
}

// GetFollowers mocks base method
func (m *MockClient) GetFollowers() (types.Followers, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFollowers")
	ret0, _ := ret[0].(types.Followers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFollowers indicates an expected call of GetFollowers
func (mr *MockClientMockRecorder) GetFollowers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFollowers", reflect.TypeOf((*MockClient)(nil).GetFollowers))
}

// GetRepos mocks base method
func (m *MockClient) GetRepos() ([]types.RepoInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepos")
	ret0, _ := ret[0].([]types.RepoInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepos indicates an expected call of GetRepos
func (mr *MockClientMockRecorder) GetRepos() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepos", reflect.TypeOf((*MockClient)(nil).GetRepos))
}
