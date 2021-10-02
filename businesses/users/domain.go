package users

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Domain struct {
	UserID    uuid.UUID
	Name      string
	Email     string
	Password  string
	Photo     string
	Role      string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	Registration(ctx context.Context, data *Domain) (*Domain, error)
}

type Repository interface {
	StoreNewUsers(ctx context.Context, data *Domain) (*Domain, error)
	GetByEmail(ctx context.Context, email string) (*Domain, error)
}
