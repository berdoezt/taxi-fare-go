// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/berdoezt/taxi-fare-go/app/service (interfaces: FareService)

// Package mockservice is a generated GoMock package.
package mockservice

import (
	reflect "reflect"

	model "github.com/berdoezt/taxi-fare-go/app/model"
	gomock "github.com/golang/mock/gomock"
)

// MockFareService is a mock of FareService interface.
type MockFareService struct {
	ctrl     *gomock.Controller
	recorder *MockFareServiceMockRecorder
}

// MockFareServiceMockRecorder is the mock recorder for MockFareService.
type MockFareServiceMockRecorder struct {
	mock *MockFareService
}

// NewMockFareService creates a new mock instance.
func NewMockFareService(ctrl *gomock.Controller) *MockFareService {
	mock := &MockFareService{ctrl: ctrl}
	mock.recorder = &MockFareServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFareService) EXPECT() *MockFareServiceMockRecorder {
	return m.recorder
}

// GetFareMeter mocks base method.
func (m *MockFareService) GetFareMeter(arg0 []model.DistanceMeter) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFareMeter", arg0)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFareMeter indicates an expected call of GetFareMeter.
func (mr *MockFareServiceMockRecorder) GetFareMeter(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFareMeter", reflect.TypeOf((*MockFareService)(nil).GetFareMeter), arg0)
}
