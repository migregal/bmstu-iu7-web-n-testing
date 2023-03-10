// Code generated by mockery v2.15.0. DO NOT EDIT.

package mock

import (
	repositories "neural_storage/cube/core/ports/repositories"

	mock "github.com/stretchr/testify/mock"

	time "time"

	user "neural_storage/cube/core/entities/user"

	userstat "neural_storage/cube/core/entities/user/userstat"
)

// UserInfoRepository is an autogenerated mock type for the UserInfoRepository type
type UserInfoRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: _a0
func (_m *UserInfoRepository) Add(_a0 user.Info) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(user.Info) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(user.Info) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: _a0
func (_m *UserInfoRepository) Delete(_a0 user.Info) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(user.Info) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: filter
func (_m *UserInfoRepository) Find(filter repositories.UserInfoFilter) ([]user.Info, int64, error) {
	ret := _m.Called(filter)

	var r0 []user.Info
	if rf, ok := ret.Get(0).(func(repositories.UserInfoFilter) []user.Info); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.Info)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(repositories.UserInfoFilter) int64); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(repositories.UserInfoFilter) error); ok {
		r2 = rf(filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Get provides a mock function with given fields: id
func (_m *UserInfoRepository) Get(id string) (user.Info, error) {
	ret := _m.Called(id)

	var r0 user.Info
	if rf, ok := ret.Get(0).(func(string) user.Info); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(user.Info)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddStat provides a mock function with given fields: from, to
func (_m *UserInfoRepository) GetAddStat(from time.Time, to time.Time) ([]*userstat.Info, error) {
	ret := _m.Called(from, to)

	var r0 []*userstat.Info
	if rf, ok := ret.Get(0).(func(time.Time, time.Time) []*userstat.Info); ok {
		r0 = rf(from, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*userstat.Info)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(time.Time, time.Time) error); ok {
		r1 = rf(from, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUpdateStat provides a mock function with given fields: from, to
func (_m *UserInfoRepository) GetUpdateStat(from time.Time, to time.Time) ([]*userstat.Info, error) {
	ret := _m.Called(from, to)

	var r0 []*userstat.Info
	if rf, ok := ret.Get(0).(func(time.Time, time.Time) []*userstat.Info); ok {
		r0 = rf(from, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*userstat.Info)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(time.Time, time.Time) error); ok {
		r1 = rf(from, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0
func (_m *UserInfoRepository) Update(_a0 user.Info) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(user.Info) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUserInfoRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserInfoRepository creates a new instance of UserInfoRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserInfoRepository(t mockConstructorTestingTNewUserInfoRepository) *UserInfoRepository {
	mock := &UserInfoRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
