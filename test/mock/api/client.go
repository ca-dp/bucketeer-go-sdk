// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package api is a generated GoMock package.
package api

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"

	gateway "github.com/ca-dp/bucketeer-go-server-sdk/proto/gateway"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Ping mocks base method
func (m *MockClient) Ping(ctx context.Context, in *gateway.PingRequest, opts ...grpc.CallOption) (*gateway.PingResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Ping", varargs...)
	ret0, _ := ret[0].(*gateway.PingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ping indicates an expected call of Ping
func (mr *MockClientMockRecorder) Ping(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockClient)(nil).Ping), varargs...)
}

// GetEvaluations mocks base method
func (m *MockClient) GetEvaluations(ctx context.Context, in *gateway.GetEvaluationsRequest, opts ...grpc.CallOption) (*gateway.GetEvaluationsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEvaluations", varargs...)
	ret0, _ := ret[0].(*gateway.GetEvaluationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEvaluations indicates an expected call of GetEvaluations
func (mr *MockClientMockRecorder) GetEvaluations(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvaluations", reflect.TypeOf((*MockClient)(nil).GetEvaluations), varargs...)
}

// GetEvaluation mocks base method
func (m *MockClient) GetEvaluation(ctx context.Context, in *gateway.GetEvaluationRequest, opts ...grpc.CallOption) (*gateway.GetEvaluationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEvaluation", varargs...)
	ret0, _ := ret[0].(*gateway.GetEvaluationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEvaluation indicates an expected call of GetEvaluation
func (mr *MockClientMockRecorder) GetEvaluation(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvaluation", reflect.TypeOf((*MockClient)(nil).GetEvaluation), varargs...)
}

// RegisterEvents mocks base method
func (m *MockClient) RegisterEvents(ctx context.Context, in *gateway.RegisterEventsRequest, opts ...grpc.CallOption) (*gateway.RegisterEventsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RegisterEvents", varargs...)
	ret0, _ := ret[0].(*gateway.RegisterEventsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterEvents indicates an expected call of RegisterEvents
func (mr *MockClientMockRecorder) RegisterEvents(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterEvents", reflect.TypeOf((*MockClient)(nil).RegisterEvents), varargs...)
}

// Close mocks base method
func (m *MockClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
}
