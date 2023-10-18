// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-application-networking-k8s/pkg/deploy/lattice (interfaces: AccessLogSubscriptionManager)

// Package lattice is a generated GoMock package.
package lattice

import (
	context "context"
	reflect "reflect"

	lattice "github.com/aws/aws-application-networking-k8s/pkg/model/lattice"
	gomock "github.com/golang/mock/gomock"
)

// MockAccessLogSubscriptionManager is a mock of AccessLogSubscriptionManager interface.
type MockAccessLogSubscriptionManager struct {
	ctrl     *gomock.Controller
	recorder *MockAccessLogSubscriptionManagerMockRecorder
}

// MockAccessLogSubscriptionManagerMockRecorder is the mock recorder for MockAccessLogSubscriptionManager.
type MockAccessLogSubscriptionManagerMockRecorder struct {
	mock *MockAccessLogSubscriptionManager
}

// NewMockAccessLogSubscriptionManager creates a new mock instance.
func NewMockAccessLogSubscriptionManager(ctrl *gomock.Controller) *MockAccessLogSubscriptionManager {
	mock := &MockAccessLogSubscriptionManager{ctrl: ctrl}
	mock.recorder = &MockAccessLogSubscriptionManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessLogSubscriptionManager) EXPECT() *MockAccessLogSubscriptionManagerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAccessLogSubscriptionManager) Create(arg0 context.Context, arg1 *lattice.AccessLogSubscription) (*lattice.AccessLogSubscriptionStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*lattice.AccessLogSubscriptionStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAccessLogSubscriptionManagerMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAccessLogSubscriptionManager)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockAccessLogSubscriptionManager) Delete(arg0 context.Context, arg1 *lattice.AccessLogSubscription) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAccessLogSubscriptionManagerMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAccessLogSubscriptionManager)(nil).Delete), arg0, arg1)
}
