// Code generated by MockGen. DO NOT EDIT.
// Source: ../git/interface/walk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	git "github.com/libgit2/git2go/v31"
)

// MockRevWalk is a mock of RevWalk interface.
type MockRevWalk struct {
	ctrl     *gomock.Controller
	recorder *MockRevWalkMockRecorder
}

// MockRevWalkMockRecorder is the mock recorder for MockRevWalk.
type MockRevWalkMockRecorder struct {
	mock *MockRevWalk
}

// NewMockRevWalk creates a new mock instance.
func NewMockRevWalk(ctrl *gomock.Controller) *MockRevWalk {
	mock := &MockRevWalk{ctrl: ctrl}
	mock.recorder = &MockRevWalkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRevWalk) EXPECT() *MockRevWalkMockRecorder {
	return m.recorder
}

// Iterate mocks base method.
func (m *MockRevWalk) Iterate(iterator git.RevWalkIterator) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iterate", iterator)
	ret0, _ := ret[0].(error)
	return ret0
}

// Iterate indicates an expected call of Iterate.
func (mr *MockRevWalkMockRecorder) Iterate(iterator interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iterate", reflect.TypeOf((*MockRevWalk)(nil).Iterate), iterator)
}

// PushHead mocks base method.
func (m *MockRevWalk) PushHead() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushHead")
	ret0, _ := ret[0].(error)
	return ret0
}

// PushHead indicates an expected call of PushHead.
func (mr *MockRevWalkMockRecorder) PushHead() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushHead", reflect.TypeOf((*MockRevWalk)(nil).PushHead))
}