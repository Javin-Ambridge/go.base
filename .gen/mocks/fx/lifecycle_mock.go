// Code generated by MockGen. DO NOT EDIT.
// Source: go.uber.org/fx (interfaces: Lifecycle)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	fx "go.uber.org/fx"
	reflect "reflect"
)

// MockLifecycle is a mock of Lifecycle interface
type MockLifecycle struct {
	ctrl     *gomock.Controller
	recorder *MockLifecycleMockRecorder
}

// MockLifecycleMockRecorder is the mock recorder for MockLifecycle
type MockLifecycleMockRecorder struct {
	mock *MockLifecycle
}

// NewMockLifecycle creates a new mock instance
func NewMockLifecycle(ctrl *gomock.Controller) *MockLifecycle {
	mock := &MockLifecycle{ctrl: ctrl}
	mock.recorder = &MockLifecycleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLifecycle) EXPECT() *MockLifecycleMockRecorder {
	return m.recorder
}

// Append mocks base method
func (m *MockLifecycle) Append(arg0 fx.Hook) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Append", arg0)
}

// Append indicates an expected call of Append
func (mr *MockLifecycleMockRecorder) Append(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Append", reflect.TypeOf((*MockLifecycle)(nil).Append), arg0)
}
