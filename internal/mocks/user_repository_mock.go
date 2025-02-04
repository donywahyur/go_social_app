// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/repositories/user_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	model "go_social_app/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// ActivationUser mocks base method.
func (m *MockUserRepository) ActivationUser(token string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActivationUser", token)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActivationUser indicates an expected call of ActivationUser.
func (mr *MockUserRepositoryMockRecorder) ActivationUser(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActivationUser", reflect.TypeOf((*MockUserRepository)(nil).ActivationUser), token)
}

// CompareHash mocks base method.
func (m *MockUserRepository) CompareHash(password, passwordHash string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompareHash", password, passwordHash)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CompareHash indicates an expected call of CompareHash.
func (mr *MockUserRepositoryMockRecorder) CompareHash(password, passwordHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompareHash", reflect.TypeOf((*MockUserRepository)(nil).CompareHash), password, passwordHash)
}

// DeleteUser mocks base method.
func (m *MockUserRepository) DeleteUser(userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserRepositoryMockRecorder) DeleteUser(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserRepository)(nil).DeleteUser), userID)
}

// GenereateJWTToken mocks base method.
func (m *MockUserRepository) GenereateJWTToken(userID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenereateJWTToken", userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenereateJWTToken indicates an expected call of GenereateJWTToken.
func (mr *MockUserRepositoryMockRecorder) GenereateJWTToken(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenereateJWTToken", reflect.TypeOf((*MockUserRepository)(nil).GenereateJWTToken), userID)
}

// GetUserByEmail mocks base method.
func (m *MockUserRepository) GetUserByEmail(email string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", email)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockUserRepositoryMockRecorder) GetUserByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockUserRepository)(nil).GetUserByEmail), email)
}

// GetUserByID mocks base method.
func (m *MockUserRepository) GetUserByID(userID string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", userID)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUserRepositoryMockRecorder) GetUserByID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserRepository)(nil).GetUserByID), userID)
}

// HashPassword mocks base method.
func (m *MockUserRepository) HashPassword(password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashPassword", password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HashPassword indicates an expected call of HashPassword.
func (mr *MockUserRepositoryMockRecorder) HashPassword(password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashPassword", reflect.TypeOf((*MockUserRepository)(nil).HashPassword), password)
}

// RegisterAndInviteUser mocks base method.
func (m *MockUserRepository) RegisterAndInviteUser(user model.User, userInvitation model.UserInvitation) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterAndInviteUser", user, userInvitation)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterAndInviteUser indicates an expected call of RegisterAndInviteUser.
func (mr *MockUserRepositoryMockRecorder) RegisterAndInviteUser(user, userInvitation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAndInviteUser", reflect.TypeOf((*MockUserRepository)(nil).RegisterAndInviteUser), user, userInvitation)
}
