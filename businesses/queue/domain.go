package queue

import "github.com/google/uuid"

type Repository interface {
	EmailUsers(userID uuid.UUID, name, email, payloadType string) error
}