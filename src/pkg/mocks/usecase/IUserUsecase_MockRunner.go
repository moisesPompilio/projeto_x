// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/moisesPompilio/projeto_x/src/internal/UseCase/UserUseCase (interfaces: IUserUsecase)

// Package _mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"

	input "github.com/moisesPompilio/projeto_x/src/internal/interfaces/input"
	output "github.com/moisesPompilio/projeto_x/src/internal/interfaces/output"
)

// MockIUserUsecase is a mock of IUserUsecase interface.
type MockIUserUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockIUserUsecaseMockRecorder
}

// MockIUserUsecaseMockRecorder is the mock recorder for MockIUserUsecase.
type MockIUserUsecaseMockRecorder struct {
	mock *MockIUserUsecase
}

// NewMockIUserUsecase creates a new mock instance.
func NewMockIUserUsecase(ctrl *gomock.Controller) *MockIUserUsecase {
	mock := &MockIUserUsecase{ctrl: ctrl}
	mock.recorder = &MockIUserUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserUsecase) EXPECT() *MockIUserUsecaseMockRecorder {
	return m.recorder
}

// CreateUserUseCase mocks base method.
func (m *MockIUserUsecase) CreateUserUseCase(arg0 input.CreateUserDTO, arg1 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserUseCase", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUserUseCase indicates an expected call of CreateUserUseCase.
func (mr *MockIUserUsecaseMockRecorder) CreateUserUseCase(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserUseCase", reflect.TypeOf((*MockIUserUsecase)(nil).CreateUserUseCase), arg0, arg1)
}

// DeleteUsersById mocks base method.
func (m *MockIUserUsecase) DeleteUsersById(arg0 string, arg1 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUsersById", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUsersById indicates an expected call of DeleteUsersById.
func (mr *MockIUserUsecaseMockRecorder) DeleteUsersById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUsersById", reflect.TypeOf((*MockIUserUsecase)(nil).DeleteUsersById), arg0, arg1)
}

// GetByID mocks base method.
func (m *MockIUserUsecase) GetByID(arg0 context.Context, arg1 uuid.UUID) (*output.UserOutpDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0, arg1)
	ret0, _ := ret[0].(*output.UserOutpDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockIUserUsecaseMockRecorder) GetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIUserUsecase)(nil).GetByID), arg0, arg1)
}

// LoginUseCase mocks base method.
func (m *MockIUserUsecase) LoginUseCase(arg0 input.Login, arg1 context.Context) (output.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginUseCase", arg0, arg1)
	ret0, _ := ret[0].(output.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginUseCase indicates an expected call of LoginUseCase.
func (mr *MockIUserUsecaseMockRecorder) LoginUseCase(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUseCase", reflect.TypeOf((*MockIUserUsecase)(nil).LoginUseCase), arg0, arg1)
}
