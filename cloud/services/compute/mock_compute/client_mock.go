// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock_compute is a generated GoMock package.
package mock_compute

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	core "github.com/oracle/oci-go-sdk/v53/core"
)

// MockComputeClient is a mock of ComputeClient interface.
type MockComputeClient struct {
	ctrl     *gomock.Controller
	recorder *MockComputeClientMockRecorder
}

// MockComputeClientMockRecorder is the mock recorder for MockComputeClient.
type MockComputeClientMockRecorder struct {
	mock *MockComputeClient
}

// NewMockComputeClient creates a new mock instance.
func NewMockComputeClient(ctrl *gomock.Controller) *MockComputeClient {
	mock := &MockComputeClient{ctrl: ctrl}
	mock.recorder = &MockComputeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComputeClient) EXPECT() *MockComputeClientMockRecorder {
	return m.recorder
}

// GetInstance mocks base method.
func (m *MockComputeClient) GetInstance(ctx context.Context, request core.GetInstanceRequest) (core.GetInstanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstance", ctx, request)
	ret0, _ := ret[0].(core.GetInstanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstance indicates an expected call of GetInstance.
func (mr *MockComputeClientMockRecorder) GetInstance(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstance", reflect.TypeOf((*MockComputeClient)(nil).GetInstance), ctx, request)
}

// LaunchInstance mocks base method.
func (m *MockComputeClient) LaunchInstance(ctx context.Context, request core.LaunchInstanceRequest) (core.LaunchInstanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LaunchInstance", ctx, request)
	ret0, _ := ret[0].(core.LaunchInstanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LaunchInstance indicates an expected call of LaunchInstance.
func (mr *MockComputeClientMockRecorder) LaunchInstance(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LaunchInstance", reflect.TypeOf((*MockComputeClient)(nil).LaunchInstance), ctx, request)
}

// ListInstances mocks base method.
func (m *MockComputeClient) ListInstances(ctx context.Context, request core.ListInstancesRequest) (core.ListInstancesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInstances", ctx, request)
	ret0, _ := ret[0].(core.ListInstancesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInstances indicates an expected call of ListInstances.
func (mr *MockComputeClientMockRecorder) ListInstances(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInstances", reflect.TypeOf((*MockComputeClient)(nil).ListInstances), ctx, request)
}

// ListVnicAttachments mocks base method.
func (m *MockComputeClient) ListVnicAttachments(ctx context.Context, request core.ListVnicAttachmentsRequest) (core.ListVnicAttachmentsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVnicAttachments", ctx, request)
	ret0, _ := ret[0].(core.ListVnicAttachmentsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVnicAttachments indicates an expected call of ListVnicAttachments.
func (mr *MockComputeClientMockRecorder) ListVnicAttachments(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVnicAttachments", reflect.TypeOf((*MockComputeClient)(nil).ListVnicAttachments), ctx, request)
}

// TerminateInstance mocks base method.
func (m *MockComputeClient) TerminateInstance(ctx context.Context, request core.TerminateInstanceRequest) (core.TerminateInstanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TerminateInstance", ctx, request)
	ret0, _ := ret[0].(core.TerminateInstanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TerminateInstance indicates an expected call of TerminateInstance.
func (mr *MockComputeClientMockRecorder) TerminateInstance(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TerminateInstance", reflect.TypeOf((*MockComputeClient)(nil).TerminateInstance), ctx, request)
}