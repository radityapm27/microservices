// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	context "context"
	go_micro_srv_core "rpm/microservices/core/proto"

	mock "github.com/stretchr/testify/mock"
)

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

// GetProductCatalog provides a mock function with given fields: ctx, req, rsp
func (_m *Handler) GetProductCatalog(ctx context.Context, req *go_micro_srv_core.CoreRequest, rsp *go_micro_srv_core.ProductCatalogResponse) error {
	ret := _m.Called(ctx, req, rsp)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *go_micro_srv_core.CoreRequest, *go_micro_srv_core.ProductCatalogResponse) error); ok {
		r0 = rf(ctx, req, rsp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}