// Code generated by MockGen. DO NOT EDIT.
// Source: ./contracts.go

// Package apierr is a generated GoMock package.
package apierr

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

// BadRequest mocks base method.
func (m *MockProvider) BadRequest(message string, err error) *Message {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BadRequest", message, err)
	ret0, _ := ret[0].(*Message)
	return ret0
}

// BadRequest indicates an expected call of BadRequest.
func (mr *MockProviderMockRecorder) BadRequest(message, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BadRequest", reflect.TypeOf((*MockProvider)(nil).BadRequest), message, err)
}

// InternalServerError mocks base method.
func (m *MockProvider) InternalServerError(err error) *Message {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InternalServerError", err)
	ret0, _ := ret[0].(*Message)
	return ret0
}

// InternalServerError indicates an expected call of InternalServerError.
func (mr *MockProviderMockRecorder) InternalServerError(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InternalServerError", reflect.TypeOf((*MockProvider)(nil).InternalServerError), err)
}

// Unauthorized mocks base method.
func (m *MockProvider) Unauthorized(message string) *Message {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unauthorized", message)
	ret0, _ := ret[0].(*Message)
	return ret0
}

// Unauthorized indicates an expected call of Unauthorized.
func (mr *MockProviderMockRecorder) Unauthorized(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unauthorized", reflect.TypeOf((*MockProvider)(nil).Unauthorized), message)
}
