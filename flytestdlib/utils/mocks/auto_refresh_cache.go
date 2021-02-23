// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	utils "github.com/flyteorg/flytestdlib/utils"
	mock "github.com/stretchr/testify/mock"
)

// AutoRefreshCache is an autogenerated mock type for the AutoRefreshCache type
type AutoRefreshCache struct {
	mock.Mock
}

type AutoRefreshCache_Get struct {
	*mock.Call
}

func (_m AutoRefreshCache_Get) Return(_a0 utils.CacheItem) *AutoRefreshCache_Get {
	return &AutoRefreshCache_Get{Call: _m.Call.Return(_a0)}
}

func (_m *AutoRefreshCache) OnGet(id string) *AutoRefreshCache_Get {
	c := _m.On("Get", id)
	return &AutoRefreshCache_Get{Call: c}
}

func (_m *AutoRefreshCache) OnGetMatch(matchers ...interface{}) *AutoRefreshCache_Get {
	c := _m.On("Get", matchers...)
	return &AutoRefreshCache_Get{Call: c}
}

// Get provides a mock function with given fields: id
func (_m *AutoRefreshCache) Get(id string) utils.CacheItem {
	ret := _m.Called(id)

	var r0 utils.CacheItem
	if rf, ok := ret.Get(0).(func(string) utils.CacheItem); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(utils.CacheItem)
		}
	}

	return r0
}

type AutoRefreshCache_GetOrCreate struct {
	*mock.Call
}

func (_m AutoRefreshCache_GetOrCreate) Return(_a0 utils.CacheItem, _a1 error) *AutoRefreshCache_GetOrCreate {
	return &AutoRefreshCache_GetOrCreate{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AutoRefreshCache) OnGetOrCreate(item utils.CacheItem) *AutoRefreshCache_GetOrCreate {
	c := _m.On("GetOrCreate", item)
	return &AutoRefreshCache_GetOrCreate{Call: c}
}

func (_m *AutoRefreshCache) OnGetOrCreateMatch(matchers ...interface{}) *AutoRefreshCache_GetOrCreate {
	c := _m.On("GetOrCreate", matchers...)
	return &AutoRefreshCache_GetOrCreate{Call: c}
}

// GetOrCreate provides a mock function with given fields: item
func (_m *AutoRefreshCache) GetOrCreate(item utils.CacheItem) (utils.CacheItem, error) {
	ret := _m.Called(item)

	var r0 utils.CacheItem
	if rf, ok := ret.Get(0).(func(utils.CacheItem) utils.CacheItem); ok {
		r0 = rf(item)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(utils.CacheItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(utils.CacheItem) error); ok {
		r1 = rf(item)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Start provides a mock function with given fields: ctx
func (_m *AutoRefreshCache) Start(ctx context.Context) {
	_m.Called(ctx)
}
