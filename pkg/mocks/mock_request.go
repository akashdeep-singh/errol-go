// Code generated by MockGen. DO NOT EDIT.
// Source: request.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "github.com/akashdeep-singh/errol-go/pkg/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRequestCodec is a mock of RequestCodec interface
type MockRequestCodec struct {
	ctrl     *gomock.Controller
	recorder *MockRequestCodecMockRecorder
}

// MockRequestCodecMockRecorder is the mock recorder for MockRequestCodec
type MockRequestCodecMockRecorder struct {
	mock *MockRequestCodec
}

// NewMockRequestCodec creates a new mock instance
func NewMockRequestCodec(ctrl *gomock.Controller) *MockRequestCodec {
	mock := &MockRequestCodec{ctrl: ctrl}
	mock.recorder = &MockRequestCodecMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRequestCodec) EXPECT() *MockRequestCodecMockRecorder {
	return m.recorder
}

// EncodeRequest mocks base method
func (m *MockRequestCodec) EncodeRequest(request models.Request) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EncodeRequest", request)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// EncodeRequest indicates an expected call of EncodeRequest
func (mr *MockRequestCodecMockRecorder) EncodeRequest(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncodeRequest", reflect.TypeOf((*MockRequestCodec)(nil).EncodeRequest), request)
}

// DecodeRequest mocks base method
func (m *MockRequestCodec) DecodeRequest(encoded []byte) (models.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecodeRequest", encoded)
	ret0, _ := ret[0].(models.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecodeRequest indicates an expected call of DecodeRequest
func (mr *MockRequestCodecMockRecorder) DecodeRequest(encoded interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecodeRequest", reflect.TypeOf((*MockRequestCodec)(nil).DecodeRequest), encoded)
}
