package response

import (
	"github.com/avtara/travair-api/businesses/users"
	"github.com/google/uuid"
	"time"
)

type Users struct {
	UserID    uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginUsers struct {
	UserID    uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Token     string    `json:"token"`
}

func FromDomain(domain *users.Domain) *Users {
	return &Users{
		UserID:    domain.UserID,
		Name:      domain.Name,
		Email:     domain.Email,
		CreatedAt: domain.CreatedAt,
	}
}

func LoginFromDomain(domain *users.Domain) *LoginUsers {
	return &LoginUsers{
		UserID: domain.UserID,
		Name:   domain.Name,
		Token:  domain.Token,
	}
}
