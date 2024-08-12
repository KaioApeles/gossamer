// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ChainSafe/gossamer/dot/parachain/runtime (interfaces: RuntimeInstance)
//
// Generated by this command:
//
//	mockgen -destination=mock_runtime_instance_test.go -package availabilitystore github.com/ChainSafe/gossamer/dot/parachain/runtime RuntimeInstance
//

// Package availabilitystore is a generated GoMock package.
package availabilitystore

import (
	reflect "reflect"

	parachaintypes "github.com/ChainSafe/gossamer/dot/parachain/types"
	gomock "go.uber.org/mock/gomock"
)

// MockRuntimeInstance is a mock of RuntimeInstance interface.
type MockRuntimeInstance struct {
	ctrl     *gomock.Controller
	recorder *MockRuntimeInstanceMockRecorder
}

// MockRuntimeInstanceMockRecorder is the mock recorder for MockRuntimeInstance.
type MockRuntimeInstanceMockRecorder struct {
	mock *MockRuntimeInstance
}

// NewMockRuntimeInstance creates a new mock instance.
func NewMockRuntimeInstance(ctrl *gomock.Controller) *MockRuntimeInstance {
	mock := &MockRuntimeInstance{ctrl: ctrl}
	mock.recorder = &MockRuntimeInstanceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRuntimeInstance) EXPECT() *MockRuntimeInstanceMockRecorder {
	return m.recorder
}

// ParachainHostCandidateEvents mocks base method.
func (m *MockRuntimeInstance) ParachainHostCandidateEvents() ([]parachaintypes.CandidateEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParachainHostCandidateEvents")
	ret0, _ := ret[0].([]parachaintypes.CandidateEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParachainHostCandidateEvents indicates an expected call of ParachainHostCandidateEvents.
func (mr *MockRuntimeInstanceMockRecorder) ParachainHostCandidateEvents() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParachainHostCandidateEvents", reflect.TypeOf((*MockRuntimeInstance)(nil).ParachainHostCandidateEvents))
}

// ParachainHostCheckValidationOutputs mocks base method.
func (m *MockRuntimeInstance) ParachainHostCheckValidationOutputs(arg0 parachaintypes.ParaID, arg1 parachaintypes.CandidateCommitments) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParachainHostCheckValidationOutputs", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParachainHostCheckValidationOutputs indicates an expected call of ParachainHostCheckValidationOutputs.
func (mr *MockRuntimeInstanceMockRecorder) ParachainHostCheckValidationOutputs(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParachainHostCheckValidationOutputs", reflect.TypeOf((*MockRuntimeInstance)(nil).ParachainHostCheckValidationOutputs), arg0, arg1)
}

// ParachainHostPersistedValidationData mocks base method.
func (m *MockRuntimeInstance) ParachainHostPersistedValidationData(arg0 uint32, arg1 parachaintypes.OccupiedCoreAssumption) (*parachaintypes.PersistedValidationData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParachainHostPersistedValidationData", arg0, arg1)
	ret0, _ := ret[0].(*parachaintypes.PersistedValidationData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParachainHostPersistedValidationData indicates an expected call of ParachainHostPersistedValidationData.
func (mr *MockRuntimeInstanceMockRecorder) ParachainHostPersistedValidationData(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParachainHostPersistedValidationData", reflect.TypeOf((*MockRuntimeInstance)(nil).ParachainHostPersistedValidationData), arg0, arg1)
}

// ParachainHostValidationCode mocks base method.
func (m *MockRuntimeInstance) ParachainHostValidationCode(arg0 uint32, arg1 parachaintypes.OccupiedCoreAssumption) (*parachaintypes.ValidationCode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParachainHostValidationCode", arg0, arg1)
	ret0, _ := ret[0].(*parachaintypes.ValidationCode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParachainHostValidationCode indicates an expected call of ParachainHostValidationCode.
func (mr *MockRuntimeInstanceMockRecorder) ParachainHostValidationCode(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParachainHostValidationCode", reflect.TypeOf((*MockRuntimeInstance)(nil).ParachainHostValidationCode), arg0, arg1)
}
