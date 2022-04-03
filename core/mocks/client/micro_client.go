// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	go_micro_srv_core "rpm/microservices/core/proto"

	mock "github.com/stretchr/testify/mock"
)

// MicroClient is an autogenerated mock type for the MicroClient type
type MicroClient struct {
	mock.Mock
}

// GetMicroServiceClientCart provides a mock function with given fields:
func (_m *MicroClient) GetMicroServiceClientCart() go_micro_srv_core.ShoppingCartService {
	ret := _m.Called()

	var r0 go_micro_srv_core.ShoppingCartService
	if rf, ok := ret.Get(0).(func() go_micro_srv_core.ShoppingCartService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(go_micro_srv_core.ShoppingCartService)
		}
	}

	return r0
}

// GetMicroServiceClientCatalog provides a mock function with given fields:
func (_m *MicroClient) GetMicroServiceClientCatalog() go_micro_srv_core.ProductCatalogService {
	ret := _m.Called()

	var r0 go_micro_srv_core.ProductCatalogService
	if rf, ok := ret.Get(0).(func() go_micro_srv_core.ProductCatalogService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(go_micro_srv_core.ProductCatalogService)
		}
	}

	return r0
}

// GetMicroServiceClientUser provides a mock function with given fields:
func (_m *MicroClient) GetMicroServiceClientUser() go_micro_srv_core.UserService {
	ret := _m.Called()

	var r0 go_micro_srv_core.UserService
	if rf, ok := ret.Get(0).(func() go_micro_srv_core.UserService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(go_micro_srv_core.UserService)
		}
	}

	return r0
}