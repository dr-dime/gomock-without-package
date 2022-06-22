// Code generated by MockGen. DO NOT EDIT.
// Source: p1/p1.go

// Package mock_p1 is a generated GoMock package.
package mock_p1

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockP1 is a mock of P1 interface.
type MockP1 struct {
	ctrl     *gomock.Controller
	recorder *MockP1MockRecorder
}

// MockP1MockRecorder is the mock recorder for MockP1.
type MockP1MockRecorder struct {
	mock *MockP1
}

// NewMockP1 creates a new mock instance.
func NewMockP1(ctrl *gomock.Controller) *MockP1 {
	mock := &MockP1{ctrl: ctrl}
	mock.recorder = &MockP1MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockP1) EXPECT() *MockP1MockRecorder {
	return m.recorder
}

// M2 mocks base method.
func (m *MockP1) M2() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "M2")
}

// M2 indicates an expected call of M2.
func (mr *MockP1MockRecorder) M2() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "M2", reflect.TypeOf((*MockP1)(nil).M2))
}