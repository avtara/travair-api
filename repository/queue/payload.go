package queue

import "github.com/google/uuid"

type UsersPayload struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Token       string `json:"token,omitempty"`
	UserID      uuid.UUID `json:"user_id"`
	PayloadType string `json:"payload_type"`
}

func FromDomainUsers(userID uuid.UUID, name, email, payloadType string) UsersPayload {
	return UsersPayload{
		Name:        name,
		Email:       email,
		UserID:      userID,
		PayloadType: payloadType,
	}
}
