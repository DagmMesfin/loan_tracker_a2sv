// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "loan_tracker_api/domain"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// DeleteUser provides a mock function with given fields: uid
func (_m *UserRepository) DeleteUser(uid string) error {
	ret := _m.Called(uid)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(uid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ForgotPassword provides a mock function with given fields: email
func (_m *UserRepository) ForgotPassword(email string) error {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for ForgotPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoginUser provides a mock function with given fields: user
func (_m *UserRepository) LoginUser(user domain.User) (string, string, error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for LoginUser")
	}

	var r0 string
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(domain.User) (string, string, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(domain.User) string); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(domain.User) string); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(domain.User) error); ok {
		r2 = rf(user)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// LogoutUser provides a mock function with given fields: uid
func (_m *UserRepository) LogoutUser(uid string) error {
	ret := _m.Called(uid)

	if len(ret) == 0 {
		panic("no return value specified for LogoutUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(uid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegisterUser provides a mock function with given fields: user
func (_m *UserRepository) RegisterUser(user *domain.User) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for RegisterUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResetPassword provides a mock function with given fields: token, newPassword
func (_m *UserRepository) ResetPassword(token string, newPassword string) error {
	ret := _m.Called(token, newPassword)

	if len(ret) == 0 {
		panic("no return value specified for ResetPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(token, newPassword)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TokenRefresh provides a mock function with given fields: uid
func (_m *UserRepository) TokenRefresh(uid string) (string, error) {
	ret := _m.Called(uid)

	if len(ret) == 0 {
		panic("no return value specified for TokenRefresh")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(uid)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(uid)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUserDetails provides a mock function with given fields: user
func (_m *UserRepository) UpdateUserDetails(user *domain.User) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUserDetails")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserProfile provides a mock function with given fields: uid
func (_m *UserRepository) UserProfile(uid string) (domain.User, error) {
	ret := _m.Called(uid)

	if len(ret) == 0 {
		panic("no return value specified for UserProfile")
	}

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (domain.User, error)); ok {
		return rf(uid)
	}
	if rf, ok := ret.Get(0).(func(string) domain.User); ok {
		r0 = rf(uid)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyUserEmail provides a mock function with given fields: token
func (_m *UserRepository) VerifyUserEmail(token string) error {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for VerifyUserEmail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ViewAllUsers provides a mock function with given fields:
func (_m *UserRepository) ViewAllUsers() ([]domain.User, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ViewAllUsers")
	}

	var r0 []domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]domain.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []domain.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
