// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Layr-Labs/incredible-squaring-avs/core/chainio (interfaces: AvsReaderer)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/avs_reader.go -package=mocks github.com/Layr-Labs/incredible-squaring-avs/core/chainio AvsReaderer
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	contractIncredibleSquaringTaskManager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	contractERC20Mock "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/MockERC20"
	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	common "github.com/ethereum/go-ethereum/common"
	gomock "go.uber.org/mock/gomock"
)

// MockAvsReaderer is a mock of AvsReaderer interface.
type MockAvsReaderer struct {
	ctrl     *gomock.Controller
	recorder *MockAvsReadererMockRecorder
}

// MockAvsReadererMockRecorder is the mock recorder for MockAvsReaderer.
type MockAvsReadererMockRecorder struct {
	mock *MockAvsReaderer
}

// NewMockAvsReaderer creates a new mock instance.
func NewMockAvsReaderer(ctrl *gomock.Controller) *MockAvsReaderer {
	mock := &MockAvsReaderer{ctrl: ctrl}
	mock.recorder = &MockAvsReadererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAvsReaderer) EXPECT() *MockAvsReadererMockRecorder {
	return m.recorder
}

// CheckSignatures mocks base method.
func (m *MockAvsReaderer) CheckSignatures(arg0 context.Context, arg1 [32]byte, arg2 []byte, arg3 uint32, arg4 contractIncredibleSquaringTaskManager.IBLSSignatureCheckerTypesNonSignerStakesAndSignature) (contractIncredibleSquaringTaskManager.IBLSSignatureCheckerTypesQuorumStakeTotals, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckSignatures", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(contractIncredibleSquaringTaskManager.IBLSSignatureCheckerTypesQuorumStakeTotals)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckSignatures indicates an expected call of CheckSignatures.
func (mr *MockAvsReadererMockRecorder) CheckSignatures(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSignatures", reflect.TypeOf((*MockAvsReaderer)(nil).CheckSignatures), arg0, arg1, arg2, arg3, arg4)
}

// GetErc20Mock mocks base method.
func (m *MockAvsReaderer) GetErc20Mock(arg0 context.Context, arg1 common.Address) (*contractERC20Mock.ContractMockERC20, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetErc20Mock", arg0, arg1)
	ret0, _ := ret[0].(*contractERC20Mock.ContractMockERC20)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetErc20Mock indicates an expected call of GetErc20Mock.
func (mr *MockAvsReadererMockRecorder) GetErc20Mock(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetErc20Mock", reflect.TypeOf((*MockAvsReaderer)(nil).GetErc20Mock), arg0, arg1)
}

// GetOperatorId mocks base method.
func (m *MockAvsReaderer) GetOperatorId(arg0 *bind.CallOpts, arg1 common.Address) ([32]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorId", arg0, arg1)
	ret0, _ := ret[0].([32]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperatorId indicates an expected call of GetOperatorId.
func (mr *MockAvsReadererMockRecorder) GetOperatorId(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorId", reflect.TypeOf((*MockAvsReaderer)(nil).GetOperatorId), arg0, arg1)
}

// IsOperatorRegistered mocks base method.
func (m *MockAvsReaderer) IsOperatorRegistered(arg0 *bind.CallOpts, arg1 common.Address) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsOperatorRegistered", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsOperatorRegistered indicates an expected call of IsOperatorRegistered.
func (mr *MockAvsReadererMockRecorder) IsOperatorRegistered(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOperatorRegistered", reflect.TypeOf((*MockAvsReaderer)(nil).IsOperatorRegistered), arg0, arg1)
}
