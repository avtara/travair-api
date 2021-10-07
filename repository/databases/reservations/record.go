package reservations

import (
	"github.com/avtara/travair-api/businesses/reservations"
	"github.com/avtara/travair-api/repository/databases/units"
	"github.com/avtara/travair-api/repository/databases/users"
	"gorm.io/gorm"
	"time"
)

type Reservation struct {
	gorm.Model
	CustomerID   uint
	Customer     users.Users
	UnitID       uint
	Unit         units.Units
	CheckInDate  time.Time
	CheckOutDate time.Time
	Status       string
	Price        int
}

func fromDomain(domain *reservations.Domain) *Reservation {
	return &Reservation{
		CustomerID: domain.CustomerID,
		UnitID: domain.UnitID,
		CheckInDate: domain.CheckInDate,
		CheckOutDate: domain.CheckOutDate,
		Status: domain.Status,
		Price: domain.Price,
	}
}

func toDomain(reservation *Reservation) *reservations.Domain {
	return &reservations.Domain {
		CustomerID: reservation.CustomerID,
		UnitID: reservation.UnitID,
		CheckInDate: reservation.CheckInDate,
		CheckOutDate: reservation.CheckOutDate,
		Status: "unconfirmed",
		Price: reservation.Price,
	}
}