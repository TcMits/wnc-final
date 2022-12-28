// Source: interfaces.go

// Package sse is a generated GoMock package.
package sse

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockINotify is a mock of INotify interface.
type MockINotify struct {
	ctrl     *gomock.Controller
	recorder *MockINotifyMockRecorder
}

// MockINotifyMockRecorder is the mock recorder for MockINotify.
type MockINotifyMockRecorder struct {
	mock *MockINotify
}

// NewMockINotify creates a new mock instance.
func NewMockINotify(ctrl *gomock.Controller) *MockINotify {
	mock := &MockINotify{ctrl: ctrl}
	mock.recorder = &MockINotifyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockINotify) EXPECT() *MockINotifyMockRecorder {
	return m.recorder
}

// Notify mocks base method.
func (m *MockINotify) Notify(arg0 *MessagePayload) error {
	ret := m.ctrl.Call(m, "Notify", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Notify indicates an expected call of Notify.
func (mr *MockINotifyMockRecorder) Notify(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notify", reflect.TypeOf((*MockINotify)(nil).Notify), arg0)
}
