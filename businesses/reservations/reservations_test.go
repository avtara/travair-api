package reservations_test

import (
	"errors"
	"github.com/avtara/travair-api/businesses"
	"github.com/avtara/travair-api/businesses/reservations"
	_resMock "github.com/avtara/travair-api/businesses/reservations/mocks"
	_unitsMock "github.com/avtara/travair-api/businesses/units/mocks"
	_usersMock "github.com/avtara/travair-api/businesses/users/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	"testing"
	"time"
)

var (
	mockReservationRepo _resMock.Repository
	mockUsersService    _usersMock.Service
	mockUnitService     _unitsMock.Service
	reservationsService reservations.Service
	domainTest reservations.Domain
)

func TestMain(m *testing.M) {
	reservationsService = reservations.NewReservationService(&mockReservationRepo,&mockUsersService,time.Second*1,&mockUnitService)
	domainTest = reservations.Domain{
		ID: 2,
	}
	m.Run()
}

func TestReservationService_Reservation(t *testing.T) {
	t.Run("error get user id", func(t *testing.T) {
		mockUsersService.On("GetID", mock.Anything, mock.Anything).Return(uint(1), errors.New("Error: ")).Once()
		uuidUserID, _ := uuid.Parse("c5c838ba-24f8-11ec-9621-0242ac130002")
		req := &reservations.Domain{
			CustomerUUID: uuidUserID,
			CustomerID: 2,
		}

		res, err := reservationsService.Reservation(context.Background(), req)
		assert.Nil(t, res)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("error get unit id", func(t *testing.T) {
		mockUsersService.On("GetID", mock.Anything, mock.Anything).Return(uint(1), nil).Once()
		mockUnitService.On("GetID", mock.Anything, mock.Anything).Return(uint(1), errors.New("Error: ")).Once()
		uuidID, _ := uuid.Parse("c5c838ba-24f8-11ec-9621-0242ac130002")
		req := &reservations.Domain{
			CustomerUUID: uuidID,
			CustomerID: 2,
		}

		res, err := reservationsService.Reservation(context.Background(), req)
		assert.Nil(t, res)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("fail get id", func(t *testing.T) {
		mockUsersService.On("GetID", mock.Anything, mock.Anything).Return(uint(1), nil).Once()
		mockUnitService.On("GetID", mock.Anything, mock.Anything).Return(uint(1), nil).Once()
		mockReservationRepo.On("GetByDate", mock.Anything, mock.Anything).Return(nil).Once()
		uuidID, _ := uuid.Parse("c5c838ba-24f8-11ec-9621-0242ac130002")
		req := &reservations.Domain{
			CustomerUUID: uuidID,
			CustomerID: 2,
		}

		res, err := reservationsService.Reservation(context.Background(), req)
		assert.Nil(t, res)
		assert.Equal(t, businesses.ErrUnitReserved, err)
	})

	t.Run("fail get id", func(t *testing.T) {
		mockUsersService.On("GetID", mock.Anything, mock.Anything).Return(uint(1), nil).Once()
		mockUnitService.On("GetID", mock.Anything, mock.Anything).Return(uint(1), nil).Once()
		mockReservationRepo.On("GetByDate", mock.Anything, mock.Anything).Return(errors.New("Error: not found")).Once()
		mockReservationRepo.On("Store", mock.Anything, mock.Anything).Return(&domainTest, nil).Once()
		uuidID, _ := uuid.Parse("c5c838ba-24f8-11ec-9621-0242ac130002")
		req := &reservations.Domain{
			CustomerUUID: uuidID,
			CustomerID: 2,
		}

		res, err := reservationsService.Reservation(context.Background(), req)
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
}


