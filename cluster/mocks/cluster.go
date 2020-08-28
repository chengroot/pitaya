// Code generated by MockGen. DO NOT EDIT.
// Source: cluster.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	cluster "github.com/topfreegames/pitaya/v2/cluster"
	message "github.com/topfreegames/pitaya/v2/conn/message"
	protos "github.com/topfreegames/pitaya/v2/protos"
	route "github.com/topfreegames/pitaya/v2/route"
	session "github.com/topfreegames/pitaya/v2/session"
	reflect "reflect"
)

// MockRPCServer is a mock of RPCServer interface
type MockRPCServer struct {
	ctrl     *gomock.Controller
	recorder *MockRPCServerMockRecorder
}

// MockRPCServerMockRecorder is the mock recorder for MockRPCServer
type MockRPCServerMockRecorder struct {
	mock *MockRPCServer
}

// NewMockRPCServer creates a new mock instance
func NewMockRPCServer(ctrl *gomock.Controller) *MockRPCServer {
	mock := &MockRPCServer{ctrl: ctrl}
	mock.recorder = &MockRPCServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRPCServer) EXPECT() *MockRPCServerMockRecorder {
	return m.recorder
}

// SetPitayaServer mocks base method
func (m *MockRPCServer) SetPitayaServer(arg0 protos.PitayaServer) {
	m.ctrl.Call(m, "SetPitayaServer", arg0)
}

// SetPitayaServer indicates an expected call of SetPitayaServer
func (mr *MockRPCServerMockRecorder) SetPitayaServer(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPitayaServer", reflect.TypeOf((*MockRPCServer)(nil).SetPitayaServer), arg0)
}

// Init mocks base method
func (m *MockRPCServer) Init() error {
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockRPCServerMockRecorder) Init() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockRPCServer)(nil).Init))
}

// AfterInit mocks base method
func (m *MockRPCServer) AfterInit() {
	m.ctrl.Call(m, "AfterInit")
}

// AfterInit indicates an expected call of AfterInit
func (mr *MockRPCServerMockRecorder) AfterInit() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterInit", reflect.TypeOf((*MockRPCServer)(nil).AfterInit))
}

// BeforeShutdown mocks base method
func (m *MockRPCServer) BeforeShutdown() {
	m.ctrl.Call(m, "BeforeShutdown")
}

// BeforeShutdown indicates an expected call of BeforeShutdown
func (mr *MockRPCServerMockRecorder) BeforeShutdown() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeShutdown", reflect.TypeOf((*MockRPCServer)(nil).BeforeShutdown))
}

// Shutdown mocks base method
func (m *MockRPCServer) Shutdown() error {
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown
func (mr *MockRPCServerMockRecorder) Shutdown() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockRPCServer)(nil).Shutdown))
}

// MockRPCClient is a mock of RPCClient interface
type MockRPCClient struct {
	ctrl     *gomock.Controller
	recorder *MockRPCClientMockRecorder
}

// MockRPCClientMockRecorder is the mock recorder for MockRPCClient
type MockRPCClientMockRecorder struct {
	mock *MockRPCClient
}

// NewMockRPCClient creates a new mock instance
func NewMockRPCClient(ctrl *gomock.Controller) *MockRPCClient {
	mock := &MockRPCClient{ctrl: ctrl}
	mock.recorder = &MockRPCClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRPCClient) EXPECT() *MockRPCClientMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockRPCClient) Send(route string, data []byte) error {
	ret := m.ctrl.Call(m, "Send", route, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockRPCClientMockRecorder) Send(route, data interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockRPCClient)(nil).Send), route, data)
}

// SendPush mocks base method
func (m *MockRPCClient) SendPush(userID string, frontendSv *cluster.Server, push *protos.Push) error {
	ret := m.ctrl.Call(m, "SendPush", userID, frontendSv, push)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendPush indicates an expected call of SendPush
func (mr *MockRPCClientMockRecorder) SendPush(userID, frontendSv, push interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendPush", reflect.TypeOf((*MockRPCClient)(nil).SendPush), userID, frontendSv, push)
}

// SendKick mocks base method
func (m *MockRPCClient) SendKick(userID, serverType string, kick *protos.KickMsg) error {
	ret := m.ctrl.Call(m, "SendKick", userID, serverType, kick)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendKick indicates an expected call of SendKick
func (mr *MockRPCClientMockRecorder) SendKick(userID, serverType, kick interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendKick", reflect.TypeOf((*MockRPCClient)(nil).SendKick), userID, serverType, kick)
}

// BroadcastSessionBind mocks base method
func (m *MockRPCClient) BroadcastSessionBind(uid string) error {
	ret := m.ctrl.Call(m, "BroadcastSessionBind", uid)
	ret0, _ := ret[0].(error)
	return ret0
}

// BroadcastSessionBind indicates an expected call of BroadcastSessionBind
func (mr *MockRPCClientMockRecorder) BroadcastSessionBind(uid interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastSessionBind", reflect.TypeOf((*MockRPCClient)(nil).BroadcastSessionBind), uid)
}

// Call mocks base method
func (m *MockRPCClient) Call(ctx context.Context, rpcType protos.RPCType, route *route.Route, session session.Session, msg *message.Message, server *cluster.Server) (*protos.Response, error) {
	ret := m.ctrl.Call(m, "Call", ctx, rpcType, route, session, msg, server)
	ret0, _ := ret[0].(*protos.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call indicates an expected call of Call
func (mr *MockRPCClientMockRecorder) Call(ctx, rpcType, route, session, msg, server interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockRPCClient)(nil).Call), ctx, rpcType, route, session, msg, server)
}

// Init mocks base method
func (m *MockRPCClient) Init() error {
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockRPCClientMockRecorder) Init() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockRPCClient)(nil).Init))
}

// AfterInit mocks base method
func (m *MockRPCClient) AfterInit() {
	m.ctrl.Call(m, "AfterInit")
}

// AfterInit indicates an expected call of AfterInit
func (mr *MockRPCClientMockRecorder) AfterInit() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterInit", reflect.TypeOf((*MockRPCClient)(nil).AfterInit))
}

// BeforeShutdown mocks base method
func (m *MockRPCClient) BeforeShutdown() {
	m.ctrl.Call(m, "BeforeShutdown")
}

// BeforeShutdown indicates an expected call of BeforeShutdown
func (mr *MockRPCClientMockRecorder) BeforeShutdown() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeShutdown", reflect.TypeOf((*MockRPCClient)(nil).BeforeShutdown))
}

// Shutdown mocks base method
func (m *MockRPCClient) Shutdown() error {
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown
func (mr *MockRPCClientMockRecorder) Shutdown() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockRPCClient)(nil).Shutdown))
}

// MockSDListener is a mock of SDListener interface
type MockSDListener struct {
	ctrl     *gomock.Controller
	recorder *MockSDListenerMockRecorder
}

// MockSDListenerMockRecorder is the mock recorder for MockSDListener
type MockSDListenerMockRecorder struct {
	mock *MockSDListener
}

// NewMockSDListener creates a new mock instance
func NewMockSDListener(ctrl *gomock.Controller) *MockSDListener {
	mock := &MockSDListener{ctrl: ctrl}
	mock.recorder = &MockSDListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSDListener) EXPECT() *MockSDListenerMockRecorder {
	return m.recorder
}

// AddServer mocks base method
func (m *MockSDListener) AddServer(arg0 *cluster.Server) {
	m.ctrl.Call(m, "AddServer", arg0)
}

// AddServer indicates an expected call of AddServer
func (mr *MockSDListenerMockRecorder) AddServer(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddServer", reflect.TypeOf((*MockSDListener)(nil).AddServer), arg0)
}

// RemoveServer mocks base method
func (m *MockSDListener) RemoveServer(arg0 *cluster.Server) {
	m.ctrl.Call(m, "RemoveServer", arg0)
}

// RemoveServer indicates an expected call of RemoveServer
func (mr *MockSDListenerMockRecorder) RemoveServer(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveServer", reflect.TypeOf((*MockSDListener)(nil).RemoveServer), arg0)
}

// MockRemoteBindingListener is a mock of RemoteBindingListener interface
type MockRemoteBindingListener struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteBindingListenerMockRecorder
}

// MockRemoteBindingListenerMockRecorder is the mock recorder for MockRemoteBindingListener
type MockRemoteBindingListenerMockRecorder struct {
	mock *MockRemoteBindingListener
}

// NewMockRemoteBindingListener creates a new mock instance
func NewMockRemoteBindingListener(ctrl *gomock.Controller) *MockRemoteBindingListener {
	mock := &MockRemoteBindingListener{ctrl: ctrl}
	mock.recorder = &MockRemoteBindingListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRemoteBindingListener) EXPECT() *MockRemoteBindingListenerMockRecorder {
	return m.recorder
}

// OnUserBind mocks base method
func (m *MockRemoteBindingListener) OnUserBind(uid, fid string) {
	m.ctrl.Call(m, "OnUserBind", uid, fid)
}

// OnUserBind indicates an expected call of OnUserBind
func (mr *MockRemoteBindingListenerMockRecorder) OnUserBind(uid, fid interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnUserBind", reflect.TypeOf((*MockRemoteBindingListener)(nil).OnUserBind), uid, fid)
}
