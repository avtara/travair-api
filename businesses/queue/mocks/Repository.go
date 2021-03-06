// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// EmailUsers provides a mock function with given fields: userID, name, email, payloadType
func (_m *Repository) EmailUsers(userID uuid.UUID, name string, email string, payloadType string) {
	_m.Called(userID, name, email, payloadType)
}
