// Code generated by MockGen. DO NOT EDIT.
// Source: ./contracts.go

// Package user_repo is a generated GoMock package.
package user_repo

import (
	context "context"
	reflect "reflect"

	internal "github.com/SamStalschus/secrets-api/internal"
	gomock "github.com/golang/mock/gomock"
)

// MockIRepository is a mock of IRepository interface.
type MockIRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRepositoryMockRecorder
}

// MockIRepositoryMockRecorder is the mock recorder for MockIRepository.
type MockIRepositoryMockRecorder struct {
	mock *MockIRepository
}

// NewMockIRepository creates a new mock instance.
func NewMockIRepository(ctrl *gomock.Controller) *MockIRepository {
	mock := &MockIRepository{ctrl: ctrl}
	mock.recorder = &MockIRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRepository) EXPECT() *MockIRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockIRepository) CreateUser(ctx context.Context, user *internal.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockIRepositoryMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIRepository)(nil).CreateUser), ctx, user)
}

// FindUserByEmail mocks base method.
func (m *MockIRepository) FindUserByEmail(ctx context.Context, email string) (*internal.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", ctx, email)
	ret0, _ := ret[0].(*internal.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockIRepositoryMockRecorder) FindUserByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockIRepository)(nil).FindUserByEmail), ctx, email)
}

// FindUserByID mocks base method.
func (m *MockIRepository) FindUserByID(ctx context.Context, id string) (*internal.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByID", ctx, id)
	ret0, _ := ret[0].(*internal.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByID indicates an expected call of FindUserByID.
func (mr *MockIRepositoryMockRecorder) FindUserByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByID", reflect.TypeOf((*MockIRepository)(nil).FindUserByID), ctx, id)
}

// FindWithPasswordByEmail mocks base method.
func (m *MockIRepository) FindWithPasswordByEmail(ctx context.Context, email string) (*internal.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindWithPasswordByEmail", ctx, email)
	ret0, _ := ret[0].(*internal.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindWithPasswordByEmail indicates an expected call of FindWithPasswordByEmail.
func (mr *MockIRepositoryMockRecorder) FindWithPasswordByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindWithPasswordByEmail", reflect.TypeOf((*MockIRepository)(nil).FindWithPasswordByEmail), ctx, email)
}
