package response

import (
	"github.com/avtara/travair-api/businesses/users"
	"github.com/google/uuid"
	"time"
)

type Users struct {
	UserID    uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain *users.Domain) *Users {
	return &Users{
		UserID:    domain.UserID,
		Name:      domain.Name,
		Email:     domain.Email,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
