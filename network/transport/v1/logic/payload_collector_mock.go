// Code generated by MockGen. DO NOT EDIT.
// Source: network/transport/v1/logic/payload_collector.go

// Package logic is a generated GoMock package.
package logic

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	hash "github.com/nuts-foundation/nuts-node/crypto/hash"
)

// MockmissingPayloadCollector is a mock of missingPayloadCollector interface.
type MockmissingPayloadCollector struct {
	ctrl     *gomock.Controller
	recorder *MockmissingPayloadCollectorMockRecorder
}

// MockmissingPayloadCollectorMockRecorder is the mock recorder for MockmissingPayloadCollector.
type MockmissingPayloadCollectorMockRecorder struct {
	mock *MockmissingPayloadCollector
}

// NewMockmissingPayloadCollector creates a new mock instance.
func NewMockmissingPayloadCollector(ctrl *gomock.Controller) *MockmissingPayloadCollector {
	mock := &MockmissingPayloadCollector{ctrl: ctrl}
	mock.recorder = &MockmissingPayloadCollectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockmissingPayloadCollector) EXPECT() *MockmissingPayloadCollectorMockRecorder {
	return m.recorder
}

// findAndQueryMissingPayloads mocks base method.
func (m *MockmissingPayloadCollector) findAndQueryMissingPayloads() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "findAndQueryMissingPayloads")
	ret0, _ := ret[0].(error)
	return ret0
}

// findAndQueryMissingPayloads indicates an expected call of findAndQueryMissingPayloads.
func (mr *MockmissingPayloadCollectorMockRecorder) findAndQueryMissingPayloads() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "findAndQueryMissingPayloads", reflect.TypeOf((*MockmissingPayloadCollector)(nil).findAndQueryMissingPayloads))
}

// findMissingPayloads mocks base method.
func (m *MockmissingPayloadCollector) findMissingPayloads() ([]hash.SHA256Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "findMissingPayloads")
	ret0, _ := ret[0].([]hash.SHA256Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// findMissingPayloads indicates an expected call of findMissingPayloads.
func (mr *MockmissingPayloadCollectorMockRecorder) findMissingPayloads() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "findMissingPayloads", reflect.TypeOf((*MockmissingPayloadCollector)(nil).findMissingPayloads))
}
