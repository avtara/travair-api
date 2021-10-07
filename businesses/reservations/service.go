package reservations

import (
	"github.com/avtara/travair-api/businesses"
	"github.com/avtara/travair-api/businesses/units"
	"github.com/avtara/travair-api/businesses/users"
	"golang.org/x/net/context"
	"time"
)

type reservationService struct {
	reservationRepository Repository
	userService           users.Service
	unitService           units.Service
	contextTimeout        time.Duration
}

func NewReservationService(rep Repository, us users.Service, to time.Duration, unitserv units.Service) Service {
	return &reservationService{
		reservationRepository: rep,
		userService:           us,
		contextTimeout:        to,
		unitService:           unitserv,
	}
}

func (rs *reservationService) Reservation(ctx context.Context, data *Domain) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, rs.contextTimeout)
	defer cancel()
	var err error

	data.CustomerID, err = rs.userService.GetID(ctx, data.CustomerUUID)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	data.UnitID, err = rs.unitService.GetID(ctx, data.UnitUUID)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}
	err = rs.reservationRepository.GetByDate(ctx, data)
	if err == nil {
		return nil, businesses.ErrUnitReserved
	}
	res, err := rs.reservationRepository.Store(ctx, data)

	return res, nil
}
