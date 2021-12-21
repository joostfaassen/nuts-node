// Code generated by MockGen. DO NOT EDIT.
// Source: network/transport/v2/protocol_grpc.pb.go

// Package v2 is a generated GoMock package.
package v2

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockProtocolClient is a mock of ProtocolClient interface.
type MockProtocolClient struct {
	ctrl     *gomock.Controller
	recorder *MockProtocolClientMockRecorder
}

// MockProtocolClientMockRecorder is the mock recorder for MockProtocolClient.
type MockProtocolClientMockRecorder struct {
	mock *MockProtocolClient
}

// NewMockProtocolClient creates a new mock instance.
func NewMockProtocolClient(ctrl *gomock.Controller) *MockProtocolClient {
	mock := &MockProtocolClient{ctrl: ctrl}
	mock.recorder = &MockProtocolClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProtocolClient) EXPECT() *MockProtocolClientMockRecorder {
	return m.recorder
}

// Stream mocks base method.
func (m *MockProtocolClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Protocol_StreamClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Stream", varargs...)
	ret0, _ := ret[0].(Protocol_StreamClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stream indicates an expected call of Stream.
func (mr *MockProtocolClientMockRecorder) Stream(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stream", reflect.TypeOf((*MockProtocolClient)(nil).Stream), varargs...)
}

// MockProtocol_StreamClient is a mock of Protocol_StreamClient interface.
type MockProtocol_StreamClient struct {
	ctrl     *gomock.Controller
	recorder *MockProtocol_StreamClientMockRecorder
}

// MockProtocol_StreamClientMockRecorder is the mock recorder for MockProtocol_StreamClient.
type MockProtocol_StreamClientMockRecorder struct {
	mock *MockProtocol_StreamClient
}

// NewMockProtocol_StreamClient creates a new mock instance.
func NewMockProtocol_StreamClient(ctrl *gomock.Controller) *MockProtocol_StreamClient {
	mock := &MockProtocol_StreamClient{ctrl: ctrl}
	mock.recorder = &MockProtocol_StreamClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProtocol_StreamClient) EXPECT() *MockProtocol_StreamClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockProtocol_StreamClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockProtocol_StreamClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockProtocol_StreamClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockProtocol_StreamClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockProtocol_StreamClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockProtocol_StreamClient)(nil).Context))
}

// Header mocks base method.
func (m *MockProtocol_StreamClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockProtocol_StreamClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockProtocol_StreamClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockProtocol_StreamClient) Recv() (*Envelope, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*Envelope)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockProtocol_StreamClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockProtocol_StreamClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockProtocol_StreamClient) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockProtocol_StreamClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockProtocol_StreamClient)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockProtocol_StreamClient) Send(arg0 *Envelope) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockProtocol_StreamClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockProtocol_StreamClient)(nil).Send), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockProtocol_StreamClient) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockProtocol_StreamClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockProtocol_StreamClient)(nil).SendMsg), m)
}

// Trailer mocks base method.
func (m *MockProtocol_StreamClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockProtocol_StreamClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockProtocol_StreamClient)(nil).Trailer))
}

// MockProtocolServer is a mock of ProtocolServer interface.
type MockProtocolServer struct {
	ctrl     *gomock.Controller
	recorder *MockProtocolServerMockRecorder
}

// MockProtocolServerMockRecorder is the mock recorder for MockProtocolServer.
type MockProtocolServerMockRecorder struct {
	mock *MockProtocolServer
}

// NewMockProtocolServer creates a new mock instance.
func NewMockProtocolServer(ctrl *gomock.Controller) *MockProtocolServer {
	mock := &MockProtocolServer{ctrl: ctrl}
	mock.recorder = &MockProtocolServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProtocolServer) EXPECT() *MockProtocolServerMockRecorder {
	return m.recorder
}

// Stream mocks base method.
func (m *MockProtocolServer) Stream(arg0 Protocol_StreamServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stream", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stream indicates an expected call of Stream.
func (mr *MockProtocolServerMockRecorder) Stream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stream", reflect.TypeOf((*MockProtocolServer)(nil).Stream), arg0)
}

// MockUnsafeProtocolServer is a mock of UnsafeProtocolServer interface.
type MockUnsafeProtocolServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeProtocolServerMockRecorder
}

// MockUnsafeProtocolServerMockRecorder is the mock recorder for MockUnsafeProtocolServer.
type MockUnsafeProtocolServerMockRecorder struct {
	mock *MockUnsafeProtocolServer
}

// NewMockUnsafeProtocolServer creates a new mock instance.
func NewMockUnsafeProtocolServer(ctrl *gomock.Controller) *MockUnsafeProtocolServer {
	mock := &MockUnsafeProtocolServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeProtocolServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeProtocolServer) EXPECT() *MockUnsafeProtocolServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedProtocolServer mocks base method.
func (m *MockUnsafeProtocolServer) mustEmbedUnimplementedProtocolServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedProtocolServer")
}

// mustEmbedUnimplementedProtocolServer indicates an expected call of mustEmbedUnimplementedProtocolServer.
func (mr *MockUnsafeProtocolServerMockRecorder) mustEmbedUnimplementedProtocolServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedProtocolServer", reflect.TypeOf((*MockUnsafeProtocolServer)(nil).mustEmbedUnimplementedProtocolServer))
}

// MockProtocol_StreamServer is a mock of Protocol_StreamServer interface.
type MockProtocol_StreamServer struct {
	ctrl     *gomock.Controller
	recorder *MockProtocol_StreamServerMockRecorder
}

// MockProtocol_StreamServerMockRecorder is the mock recorder for MockProtocol_StreamServer.
type MockProtocol_StreamServerMockRecorder struct {
	mock *MockProtocol_StreamServer
}

// NewMockProtocol_StreamServer creates a new mock instance.
func NewMockProtocol_StreamServer(ctrl *gomock.Controller) *MockProtocol_StreamServer {
	mock := &MockProtocol_StreamServer{ctrl: ctrl}
	mock.recorder = &MockProtocol_StreamServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProtocol_StreamServer) EXPECT() *MockProtocol_StreamServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockProtocol_StreamServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockProtocol_StreamServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockProtocol_StreamServer)(nil).Context))
}

// Recv mocks base method.
func (m *MockProtocol_StreamServer) Recv() (*Envelope, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*Envelope)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockProtocol_StreamServerMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockProtocol_StreamServer)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockProtocol_StreamServer) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockProtocol_StreamServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockProtocol_StreamServer)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockProtocol_StreamServer) Send(arg0 *Envelope) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockProtocol_StreamServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockProtocol_StreamServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockProtocol_StreamServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockProtocol_StreamServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockProtocol_StreamServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockProtocol_StreamServer) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockProtocol_StreamServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockProtocol_StreamServer)(nil).SendMsg), m)
}

// SetHeader mocks base method.
func (m *MockProtocol_StreamServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockProtocol_StreamServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockProtocol_StreamServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockProtocol_StreamServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockProtocol_StreamServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockProtocol_StreamServer)(nil).SetTrailer), arg0)
}