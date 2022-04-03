// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetShoppingCartByUserIdQuery provides a mock function with given fields: userID
func (_m *Repository) GetShoppingCartByUserIdQuery(userID string) ([][]string, error) {
	ret := _m.Called(userID)

	var r0 [][]string
	if rf, ok := ret.Get(0).(func(string) [][]string); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}