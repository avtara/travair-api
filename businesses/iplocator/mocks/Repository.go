// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	iplocator "github.com/avtara/travair-api/businesses/iplocator"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetLocationByIP provides a mock function with given fields: ctx, ip
func (_m *Repository) GetLocationByIP(ctx context.Context, ip string) (*iplocator.Domain, error) {
	ret := _m.Called(ctx, ip)

	var r0 *iplocator.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) *iplocator.Domain); ok {
		r0 = rf(ctx, ip)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iplocator.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, ip)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
