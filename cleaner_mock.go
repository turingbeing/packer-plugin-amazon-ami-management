// Code generated by MockGen. DO NOT EDIT.
// Source: cleaner.go

// Package main is a generated GoMock package.
package main

import (
	ec2 "github.com/aws/aws-sdk-go/service/ec2"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCleanable is a mock of Cleanable interface
type MockCleanable struct {
	ctrl     *gomock.Controller
	recorder *MockCleanableMockRecorder
}

// MockCleanableMockRecorder is the mock recorder for MockCleanable
type MockCleanableMockRecorder struct {
	mock *MockCleanable
}

// NewMockCleanable creates a new mock instance
func NewMockCleanable(ctrl *gomock.Controller) *MockCleanable {
	mock := &MockCleanable{ctrl: ctrl}
	mock.recorder = &MockCleanableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCleanable) EXPECT() *MockCleanableMockRecorder {
	return m.recorder
}

// RetrieveCandidateImages mocks base method
func (m *MockCleanable) RetrieveCandidateImages() ([]*ec2.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveCandidateImages")
	ret0, _ := ret[0].([]*ec2.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveCandidateImages indicates an expected call of RetrieveCandidateImages
func (mr *MockCleanableMockRecorder) RetrieveCandidateImages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveCandidateImages", reflect.TypeOf((*MockCleanable)(nil).RetrieveCandidateImages))
}

// DeleteImage mocks base method
func (m *MockCleanable) DeleteImage(arg0 *ec2.Image) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteImage", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteImage indicates an expected call of DeleteImage
func (mr *MockCleanableMockRecorder) DeleteImage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteImage", reflect.TypeOf((*MockCleanable)(nil).DeleteImage), arg0)
}

// IsUsed mocks base method
func (m *MockCleanable) IsUsed(arg0 *ec2.Image) *Used {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUsed", arg0)
	ret0, _ := ret[0].(*Used)
	return ret0
}

// IsUsed indicates an expected call of IsUsed
func (mr *MockCleanableMockRecorder) IsUsed(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUsed", reflect.TypeOf((*MockCleanable)(nil).IsUsed), arg0)
}
