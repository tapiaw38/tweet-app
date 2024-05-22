// Code generated by MockGen. DO NOT EDIT.
// Source: internal/usecases/user/create.go
//
// Generated by this command:
//
//	mockgen -source=internal/usecases/user/create.go -destination=internal/usecases/user/mocks/create.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	domain "github.com/tapiaw38/tweet-app/internal/domain"
	user "github.com/tapiaw38/tweet-app/internal/usecases/user"
	gomock "go.uber.org/mock/gomock"
)

// MockCreateUsecase is a mock of CreateUsecase interface.
type MockCreateUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockCreateUsecaseMockRecorder
}

// MockCreateUsecaseMockRecorder is the mock recorder for MockCreateUsecase.
type MockCreateUsecaseMockRecorder struct {
	mock *MockCreateUsecase
}

// NewMockCreateUsecase creates a new mock instance.
func NewMockCreateUsecase(ctrl *gomock.Controller) *MockCreateUsecase {
	mock := &MockCreateUsecase{ctrl: ctrl}
	mock.recorder = &MockCreateUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreateUsecase) EXPECT() *MockCreateUsecaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockCreateUsecase) Execute(arg0 context.Context, arg1 domain.User) (user.CreateOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", arg0, arg1)
	ret0, _ := ret[0].(user.CreateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockCreateUsecaseMockRecorder) Execute(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockCreateUsecase)(nil).Execute), arg0, arg1)
}
