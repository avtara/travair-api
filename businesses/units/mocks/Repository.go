// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	units "github.com/avtara/travair-api/businesses/units"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// AddPhotoUnit provides a mock function with given fields: ctx, ID, path
func (_m *Repository) AddPhotoUnit(ctx context.Context, ID uint, path string) error {
	ret := _m.Called(ctx, ID, path)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint, string) error); ok {
		r0 = rf(ctx, ID, path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByUnitID provides a mock function with given fields: ctx, unitID
func (_m *Repository) GetByUnitID(ctx context.Context, unitID uuid.UUID) (*units.Domain, error) {
	ret := _m.Called(ctx, unitID)

	var r0 *units.Domain
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *units.Domain); ok {
		r0 = rf(ctx, unitID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*units.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, unitID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIDByUnitID provides a mock function with given fields: ctx, unitID
func (_m *Repository) GetIDByUnitID(ctx context.Context, unitID uuid.UUID) (uint, error) {
	ret := _m.Called(ctx, unitID)

	var r0 uint
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) uint); ok {
		r0 = rf(ctx, unitID)
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, unitID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUnitsByGeo provides a mock function with given fields: ctx, lat, long
func (_m *Repository) GetUnitsByGeo(ctx context.Context, lat float64, long float64) ([]units.Result, error) {
	ret := _m.Called(ctx, lat, long)

	var r0 []units.Result
	if rf, ok := ret.Get(0).(func(context.Context, float64, float64) []units.Result); ok {
		r0 = rf(ctx, lat, long)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]units.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, float64, float64) error); ok {
		r1 = rf(ctx, lat, long)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectAddressByID provides a mock function with given fields: ctx, ID
func (_m *Repository) SelectAddressByID(ctx context.Context, ID uint) (units.Address, error) {
	ret := _m.Called(ctx, ID)

	var r0 units.Address
	if rf, ok := ret.Get(0).(func(context.Context, uint) units.Address); ok {
		r0 = rf(ctx, ID)
	} else {
		r0 = ret.Get(0).(units.Address)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectAllPhotosByID provides a mock function with given fields: ctx, ID
func (_m *Repository) SelectAllPhotosByID(ctx context.Context, ID uint) ([]units.Photo, error) {
	ret := _m.Called(ctx, ID)

	var r0 []units.Photo
	if rf, ok := ret.Get(0).(func(context.Context, uint) []units.Photo); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]units.Photo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, data
func (_m *Repository) Store(ctx context.Context, data *units.Domain) (*units.Domain, error) {
	ret := _m.Called(ctx, data)

	var r0 *units.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *units.Domain) *units.Domain); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*units.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *units.Domain) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePathByUnitID provides a mock function with given fields: ctx, unitID, res
func (_m *Repository) UpdatePathByUnitID(ctx context.Context, unitID uuid.UUID, res string) error {
	ret := _m.Called(ctx, unitID, res)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string) error); ok {
		r0 = rf(ctx, unitID, res)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}