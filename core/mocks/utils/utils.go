// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	connection "rpm/microservices/core/connection"

	cache "github.com/patrickmn/go-cache"

	mock "github.com/stretchr/testify/mock"

	model "rpm/microservices/core/model"
)

// Utils is an autogenerated mock type for the Utils type
type Utils struct {
	mock.Mock
}

// ConstructApplicationConfig provides a mock function with given fields:
func (_m *Utils) ConstructApplicationConfig() *model.ApplicationConfig {
	ret := _m.Called()

	var r0 *model.ApplicationConfig
	if rf, ok := ret.Get(0).(func() *model.ApplicationConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ApplicationConfig)
		}
	}

	return r0
}

// GetCache provides a mock function with given fields: key, _a1
func (_m *Utils) GetCache(key string, _a1 *cache.Cache) interface{} {
	ret := _m.Called(key, _a1)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, *cache.Cache) interface{}); ok {
		r0 = rf(key, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// GetCacheHandler provides a mock function with given fields:
func (_m *Utils) GetCacheHandler() *cache.Cache {
	ret := _m.Called()

	var r0 *cache.Cache
	if rf, ok := ret.Get(0).(func() *cache.Cache); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cache.Cache)
		}
	}

	return r0
}

// GetCacheKey provides a mock function with given fields: params
func (_m *Utils) GetCacheKey(params ...string) string {
	_va := make([]interface{}, len(params))
	for _i := range params {
		_va[_i] = params[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(...string) string); ok {
		r0 = rf(params...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetClientHandler provides a mock function with given fields:
func (_m *Utils) GetClientHandler() connection.Connection {
	ret := _m.Called()

	var r0 connection.Connection
	if rf, ok := ret.Get(0).(func() connection.Connection); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(connection.Connection)
		}
	}

	return r0
}

// GetDatasourceInfo provides a mock function with given fields:
func (_m *Utils) GetDatasourceInfo() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// SetCache provides a mock function with given fields: key, value, cacheParam
func (_m *Utils) SetCache(key string, value interface{}, cacheParam *cache.Cache) interface{} {
	ret := _m.Called(key, value, cacheParam)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, interface{}, *cache.Cache) interface{}); ok {
		r0 = rf(key, value, cacheParam)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// StringToInt provides a mock function with given fields: input
func (_m *Utils) StringToInt(input string) (int, error) {
	ret := _m.Called(input)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
