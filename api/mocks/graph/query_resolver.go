// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "rpm/microservices/api/graph/model"
)

// QueryResolver is an autogenerated mock type for the QueryResolver type
type QueryResolver struct {
	mock.Mock
}

// ProductCatalog provides a mock function with given fields: ctx
func (_m *QueryResolver) ProductCatalog(ctx context.Context) (*model.AbstractModel, error) {
	ret := _m.Called(ctx)

	var r0 *model.AbstractModel
	if rf, ok := ret.Get(0).(func(context.Context) *model.AbstractModel); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AbstractModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShoppingCart provides a mock function with given fields: ctx
func (_m *QueryResolver) ShoppingCart(ctx context.Context) (*model.AbstractModel, error) {
	ret := _m.Called(ctx)

	var r0 *model.AbstractModel
	if rf, ok := ret.Get(0).(func(context.Context) *model.AbstractModel); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AbstractModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// User provides a mock function with given fields: ctx
func (_m *QueryResolver) User(ctx context.Context) (*model.AbstractModel, error) {
	ret := _m.Called(ctx)

	var r0 *model.AbstractModel
	if rf, ok := ret.Get(0).(func(context.Context) *model.AbstractModel); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AbstractModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}