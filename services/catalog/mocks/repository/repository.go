// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetProductCatalogQuery provides a mock function with given fields:
func (_m *Repository) GetProductCatalogQuery() ([][]string, error) {
	ret := _m.Called()

	var r0 [][]string
	if rf, ok := ret.Get(0).(func() [][]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
