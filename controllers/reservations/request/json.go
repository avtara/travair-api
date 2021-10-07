package request

import (
	"github.com/avtara/travair-api/businesses/reservations"
	"github.com/google/uuid"
	"time"
)

type Reservation struct {
	UnitUUID     uuid.UUID `json:"unit_id" validate:"required"`
	CheckInDate  string    `json:"check_in_date" validate:"required"`
	CheckOutDate string    `json:"check_out_date" validate:"required"`
	Price        int       `json:"price" validate:"required"`
}

func ToDomain(data *Reservation, claim uuid.UUID) *reservations.Domain {
	startDate, _ := time.Parse("2006-01-02", data.CheckInDate)
	endDate, _ := time.Parse("2006-01-02", data.CheckOutDate)
	return &reservations.Domain{
		CustomerUUID: claim,
		UnitUUID:     data.UnitUUID,
		CheckInDate:  startDate,
		CheckOutDate: endDate,
		Price:        data.Price,
	}
}
