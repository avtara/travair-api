package reservations

import (
	"context"
	"github.com/avtara/travair-api/businesses/units"
	"github.com/avtara/travair-api/businesses/users"
	"github.com/google/uuid"
	"time"
)

type Domain struct {
	ID           uint
	CustomerID       uint
	CustomerUUID uuid.UUID
	Customer     users.Domain
	UnitID       uint
	UnitUUID     uuid.UUID
	Unit         units.Domain
	CheckInDate  time.Time
	CheckOutDate time.Time
	Status       string
	Price        int
}

type Service interface {
	Reservation(ctx context.Context, data *Domain) (*Domain, error)
}

type Repository interface {
	Store(ctx context.Context, domain *Domain) (*Domain, error)
	GetByDate (ctx context.Context, domain *Domain) error
}
