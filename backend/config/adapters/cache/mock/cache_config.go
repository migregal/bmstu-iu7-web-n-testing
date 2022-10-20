// Code generated by mockery v2.11.0. DO NOT EDIT.

package mock

import (
	cache "neural_storage/cache/core/services/interactors/cache"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// CacheConfig is an autogenerated mock type for the CacheConfig type
type CacheConfig struct {
	mock.Mock
}

// Adapter provides a mock function with given fields:
func (_m *CacheConfig) Adapter() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ConnParams provides a mock function with given fields:
func (_m *CacheConfig) ConnParams() cache.Params {
	ret := _m.Called()

	var r0 cache.Params
	if rf, ok := ret.Get(0).(func() cache.Params); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(cache.Params)
	}

	return r0
}

// IsMocked provides a mock function with given fields:
func (_m *CacheConfig) IsMocked() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewCacheConfig creates a new instance of CacheConfig. It also registers a cleanup function to assert the mocks expectations.
func NewCacheConfig(t testing.TB) *CacheConfig {
	mock := &CacheConfig{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}