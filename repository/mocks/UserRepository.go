// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	model "hacktiv8-chapter-three-session-two/models"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: user
func (_m *UserRepository) CreateUser(user model.User) (*model.User, error) {
	ret := _m.Called(user)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(model.User) (*model.User, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(model.User) *model.User); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(model.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserCheck provides a mock function with given fields: userId
func (_m *UserRepository) UserCheck(userId string) (*model.User, error) {
	ret := _m.Called(userId)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.User, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserCheckByEmail provides a mock function with given fields: email
func (_m *UserRepository) UserCheckByEmail(email string) (*model.User, error) {
	ret := _m.Called(email)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.User, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
