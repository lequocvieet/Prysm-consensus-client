// Code generated by MockGen. DO NOT EDIT.
// Source: validator/client/beacon-api/state_validators.go
//
// Generated by this command:
//
//	mockgen -package=mock -source=validator/client/beacon-api/state_validators.go -destination=validator/client/beacon-api/mock/state_validators_mock.go
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	structs "github.com/prysmaticlabs/prysm/v5/api/server/structs"
	primitives "github.com/prysmaticlabs/prysm/v5/consensus-types/primitives"
	gomock "go.uber.org/mock/gomock"
)

// MockStateValidatorsProvider is a mock of StateValidatorsProvider interface.
type MockStateValidatorsProvider struct {
	ctrl     *gomock.Controller
	recorder *MockStateValidatorsProviderMockRecorder
}

// MockStateValidatorsProviderMockRecorder is the mock recorder for MockStateValidatorsProvider.
type MockStateValidatorsProviderMockRecorder struct {
	mock *MockStateValidatorsProvider
}

// NewMockStateValidatorsProvider creates a new mock instance.
func NewMockStateValidatorsProvider(ctrl *gomock.Controller) *MockStateValidatorsProvider {
	mock := &MockStateValidatorsProvider{ctrl: ctrl}
	mock.recorder = &MockStateValidatorsProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateValidatorsProvider) EXPECT() *MockStateValidatorsProviderMockRecorder {
	return m.recorder
}

// StateValidators mocks base method.
func (m *MockStateValidatorsProvider) StateValidators(arg0 context.Context, arg1 []string, arg2 []primitives.ValidatorIndex, arg3 []string) (*structs.GetValidatorsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateValidators", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*structs.GetValidatorsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateValidators indicates an expected call of StateValidators.
func (mr *MockStateValidatorsProviderMockRecorder) StateValidators(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateValidators", reflect.TypeOf((*MockStateValidatorsProvider)(nil).StateValidators), arg0, arg1, arg2, arg3)
}

// StateValidatorsForHead mocks base method.
func (m *MockStateValidatorsProvider) StateValidatorsForHead(arg0 context.Context, arg1 []string, arg2 []primitives.ValidatorIndex, arg3 []string) (*structs.GetValidatorsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateValidatorsForHead", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*structs.GetValidatorsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateValidatorsForHead indicates an expected call of StateValidatorsForHead.
func (mr *MockStateValidatorsProviderMockRecorder) StateValidatorsForHead(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateValidatorsForHead", reflect.TypeOf((*MockStateValidatorsProvider)(nil).StateValidatorsForHead), arg0, arg1, arg2, arg3)
}

// StateValidatorsForSlot mocks base method.
func (m *MockStateValidatorsProvider) StateValidatorsForSlot(arg0 context.Context, arg1 primitives.Slot, arg2 []string, arg3 []primitives.ValidatorIndex, arg4 []string) (*structs.GetValidatorsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateValidatorsForSlot", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*structs.GetValidatorsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateValidatorsForSlot indicates an expected call of StateValidatorsForSlot.
func (mr *MockStateValidatorsProviderMockRecorder) StateValidatorsForSlot(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateValidatorsForSlot", reflect.TypeOf((*MockStateValidatorsProvider)(nil).StateValidatorsForSlot), arg0, arg1, arg2, arg3, arg4)
}