// Code generated by MockGen. DO NOT EDIT.
// Source: ./contracts.go

// Package auth is a generated GoMock package.
package auth

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockProvider is a mock of Provider interface.
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider.
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance.
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// CheckPassword mocks base method.
func (m *MockProvider) CheckPassword(password, providedPassword string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckPassword", password, providedPassword)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckPassword indicates an expected call of CheckPassword.
func (mr *MockProviderMockRecorder) CheckPassword(password, providedPassword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPassword", reflect.TypeOf((*MockProvider)(nil).CheckPassword), password, providedPassword)
}

// EncryptPassword mocks base method.
func (m *MockProvider) EncryptPassword(password string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EncryptPassword", password)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EncryptPassword indicates an expected call of EncryptPassword.
func (mr *MockProviderMockRecorder) EncryptPassword(password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncryptPassword", reflect.TypeOf((*MockProvider)(nil).EncryptPassword), password)
}

// NewJwt mocks base method.
func (m *MockProvider) NewJwt(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewJwt", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewJwt indicates an expected call of NewJwt.
func (mr *MockProviderMockRecorder) NewJwt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewJwt", reflect.TypeOf((*MockProvider)(nil).NewJwt), arg0)
}

// ValidateJwt mocks base method.
func (m *MockProvider) ValidateJwt(token string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateJwt", token)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateJwt indicates an expected call of ValidateJwt.
func (mr *MockProviderMockRecorder) ValidateJwt(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateJwt", reflect.TypeOf((*MockProvider)(nil).ValidateJwt), token)
}
