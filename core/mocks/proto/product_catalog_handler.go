// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	context "context"
	go_micro_srv_core "rpm/microservices/core/proto"

	mock "github.com/stretchr/testify/mock"
)

// ProductCatalogHandler is an autogenerated mock type for the ProductCatalogHandler type
type ProductCatalogHandler struct {
	mock.Mock
}

// GetProductCatalog provides a mock function with given fields: _a0, _a1, _a2
func (_m *ProductCatalogHandler) GetProductCatalog(_a0 context.Context, _a1 *go_micro_srv_core.CoreRequest, _a2 *go_micro_srv_core.ProductCatalogResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *go_micro_srv_core.CoreRequest, *go_micro_srv_core.ProductCatalogResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}