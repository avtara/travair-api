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
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	Registration(ctx context.Context, data *Domain) (*Domain, error)
	Activation(ctx context.Context, userID string) (*Domain, error)
	Login(ctx context.Context, email, password string) (*Domain, error)
}

type Repository interface {
	StoreNewUsers(ctx context.Context, data *Domain) (*Domain, error)
	GetByEmail(ctx context.Context, email string) (*Domain, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) (*Domain, error)
	UpdateStatus(ctx context.Context, userID uuid.UUID) error
	GetByEmailAndPassword(ctx context.Context, email string) (*Domain, error)
}
