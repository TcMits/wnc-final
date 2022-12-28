// Source: interfaces.go

// Package task is a generated GoMock package.
package task

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIExecuteTask is a mock of IExecuteTask interface.
type MockIExecuteTask[Template any] struct {
	ctrl     *gomock.Controller
	recorder *MockIExecuteTaskMockRecorder[Template]
}

// MockIExecuteTaskMockRecorder is the mock recorder for MockIExecuteTask.
type MockIExecuteTaskMockRecorder[Template any] struct {
	mock *MockIExecuteTask[Template]
}

// NewMockIExecuteTask creates a new mock instance.
func NewMockIExecuteTask[Template any](ctrl *gomock.Controller) *MockIExecuteTask[Template] {
	mock := &MockIExecuteTask[Template]{ctrl: ctrl}
	mock.recorder = &MockIExecuteTaskMockRecorder[Template]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIExecuteTask[Template]) EXPECT() *MockIExecuteTaskMockRecorder[Template] {
	return m.recorder
}

// ExecuteTask mocks base method.
func (m *MockIExecuteTask[Template]) ExecuteTask(arg0 context.Context, arg1 Template) error {
	ret := m.ctrl.Call(m, "ExecuteTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecuteTask indicates an expected call of ExecuteTask.
func (mr *MockIExecuteTaskMockRecorder[Template]) ExecuteTask(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteTask", reflect.TypeOf((*MockIExecuteTask[Template])(nil).ExecuteTask), arg0, arg1)
}
