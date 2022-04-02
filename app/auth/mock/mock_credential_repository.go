// Code generated by MockGen. DO NOT EDIT.
// Source: ./domain/auth/credential_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	auth "github.com/ansidev/togo/domain/auth"
	user "github.com/ansidev/togo/domain/user"
	gomock "github.com/golang/mock/gomock"
)

// MockICredRepository is a mock of ICredRepository interface.
type MockICredRepository struct {
	ctrl     *gomock.Controller
	recorder *MockICredRepositoryMockRecorder
}

// MockICredRepositoryMockRecorder is the mock recorder for MockICredRepository.
type MockICredRepositoryMockRecorder struct {
	mock *MockICredRepository
}

// NewMockICredRepository creates a new mock instance.
func NewMockICredRepository(ctrl *gomock.Controller) *MockICredRepository {
	mock := &MockICredRepository{ctrl: ctrl}
	mock.recorder = &MockICredRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICredRepository) EXPECT() *MockICredRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockICredRepository) Get(token string) (auth.AuthenticationCredential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", token)
	ret0, _ := ret[0].(auth.AuthenticationCredential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockICredRepositoryMockRecorder) Get(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockICredRepository)(nil).Get), token)
}

// Save mocks base method.
func (m *MockICredRepository) Save(userModel user.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", userModel)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockICredRepositoryMockRecorder) Save(userModel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockICredRepository)(nil).Save), userModel)
}