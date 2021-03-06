// Code generated by MockGen. DO NOT EDIT.
// Source: service/service.go

// Package service is a generated GoMock package.
package service

import (
	model "Desktop/todo-backend/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CreateTodo mocks base method.
func (m *MockService) CreateTodo(todo model.SendTodoElements) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTodo", todo)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTodo indicates an expected call of CreateTodo.
func (mr *MockServiceMockRecorder) CreateTodo(todo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTodo", reflect.TypeOf((*MockService)(nil).CreateTodo), todo)
}

// GetTodoElements mocks base method.
func (m *MockService) GetTodoElements() ([]model.TodoElements, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTodoElements")
	ret0, _ := ret[0].([]model.TodoElements)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTodoElements indicates an expected call of GetTodoElements.
func (mr *MockServiceMockRecorder) GetTodoElements() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTodoElements", reflect.TypeOf((*MockService)(nil).GetTodoElements))
}
