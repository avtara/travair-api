package units_test

import (
	"errors"
	"fmt"
	"github.com/avtara/travair-api/businesses"
	"github.com/avtara/travair-api/businesses/iplocator"
	_ipMock "github.com/avtara/travair-api/businesses/iplocator/mocks"
	"github.com/avtara/travair-api/businesses/units"
	_unitsMock "github.com/avtara/travair-api/businesses/units/mocks"
	_uploadMock "github.com/avtara/travair-api/businesses/uploads/mocks"
	_usersMock "github.com/avtara/travair-api/businesses/users/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	"mime/multipart"
	"testing"
	"time"
)

var (
	mockUnitRepository _unitsMock.Repository
	mockUserService    _usersMock.Service
	mockUploadRepo     _uploadMock.Repository
	mockIpapiRepo      _ipMock.Repository
	unitsService       units.Service
	domainTest         units.Domain
)

func TestMain(m *testing.M) {
	unitsService = units.NewUnitService(&mockUnitRepository, &mockUserService, 1*time.Second, &mockUploadRepo, &mockIpapiRepo)
	domainTest = units.Domain{
		ID: 2,
	}
	m.Run()
}

func TestUnitService_Add(t *testing.T) {
	t.Run("error store", func(t *testing.T) {
		mockUserService.On("GetID", mock.Anything, mock.Anything).Return(uint(0), nil).Once()
		mockUnitRepository.On("Store", mock.Anything, mock.Anything).Return(nil, errors.New("Error: bla bla")).Once()
		uuidID, _ := uuid.Parse("c5c838ba-24f8-11ec-9621-0242ac130002")
		req := &units.Domain{
			ID: 2,
		}

		res, err := unitsService.Add(context.Background(), req, uuidID)
		assert.Nil(t, res)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("success store", func(t *testing.T) {
		mockUserService.On("GetID", mock.Anything, mock.Anything).Return(uint(0), nil).Once()
		mockUnitRepository.On("Store", mock.Anything, mock.Anything).Return(&domainTest, nil).Once()
		uuidID, _ := uuid.Parse("c5c838ba-24f8-11ec-9621-0242ac130002")
		req := &units.Domain{
			ID: 2,
		}

		res, err := unitsService.Add(context.Background(), req, uuidID)
		fmt.Println(err)
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
}

func TestUnitService_ChangeThumbnail(t *testing.T) {
	t.Run("error parse", func(t *testing.T) {
		resss := multipart.FileHeader{}

		err := unitsService.ChangeThumbnail(context.Background(), "c5c838ba-24f8-11ec-9621-0242130002", &resss)

		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("error not found", func(t *testing.T) {
		mockUnitRepository.On("GetIDByUnitID", mock.Anything, mock.Anything).Return(uint(0), errors.New("not found")).Once()
		resss := multipart.FileHeader{}

		err := unitsService.ChangeThumbnail(context.Background(), "c5c838ba-24f8-11ec-9621-0242ac130002", &resss)

		assert.Equal(t, businesses.ErrUnitNotFound, err)
	})

	t.Run("error sql", func(t *testing.T) {
		mockUnitRepository.On("GetIDByUnitID", mock.Anything, mock.Anything).Return(uint(0), errors.New("Error:")).Once()
		resss := multipart.FileHeader{}

		err := unitsService.ChangeThumbnail(context.Background(), "c5c838ba-24f8-11ec-9621-0242ac130002", &resss)

		assert.NotNil(t, err)
	})

	t.Run("error upload", func(t *testing.T) {
		mockUnitRepository.On("GetIDByUnitID", mock.Anything, mock.Anything).Return(uint(1), nil).Once()
		mockUploadRepo.On("UploadLocal", mock.Anything, mock.Anything).Return("localhost", errors.New("Error:")).Once()
		resss := multipart.FileHeader{}

		err := unitsService.ChangeThumbnail(context.Background(), "c5c838ba-24f8-11ec-9621-0242ac130002", &resss)

		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("error update path", func(t *testing.T) {
		mockUnitRepository.On("GetIDByUnitID", mock.Anything, mock.Anything).Return(uint(1), nil).Once()
		mockUploadRepo.On("UploadLocal", mock.Anything, mock.Anything).Return("localhost", nil).Once()
		mockUnitRepository.On("UpdatePathByUnitID", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("Error:")).Once()
		resss := multipart.FileHeader{}

		err := unitsService.ChangeThumbnail(context.Background(), "c5c838ba-24f8-11ec-9621-0242ac130002", &resss)

		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("success", func(t *testing.T) {
		mockUnitRepository.On("GetIDByUnitID", mock.Anything, mock.Anything).Return(uint(1), nil).Once()
		mockUploadRepo.On("UploadLocal", mock.Anything, mock.Anything).Return("localhost", nil).Once()
		mockUnitRepository.On("UpdatePathByUnitID", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		resss := multipart.FileHeader{}

		err := unitsService.ChangeThumbnail(context.Background(), "c5c838ba-24f8-11ec-9621-0242ac130002", &resss)

		assert.Nil(t, err)
	})
}

func TestUnitService_AddPhoto(t *testing.T) {
	t.Run("error parse", func(t *testing.T) {
		resss := multipart.FileHeader{}

		err := unitsService.AddPhoto(context.Background(), "c5c838ba-24f8-11ec-9621-0242130002", &resss)

		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("error not found", func(t *testing.T) {
		mockUnitRepository.On("GetIDByUnitID", mock.Anything, mock.Anything).Return(uint(0), errors.New("not found")).Once()
		resss := multipart.FileHeader{}

		err := unitsService.AddPhoto(context.Background(), "c5c838ba-24f8-11ec-9621-0242ac130002", &resss)

		assert.Equal(t, businesses.ErrUnitNotFound, err)
	})

	t.Run("error sql", func(t *testing.T) {
		mockUnitRepository.On("GetIDByUnitID", mock.Anything, mock.Anything).Return(uint(0), errors.New("Error:")).Once()
		resss := multipart.FileHeader{}

		err := unitsService.AddPhoto(context.Background(), "c5c838ba-24f8-11ec-9621-0242ac130002", &resss)

		assert.NotNil(t, err)
	})

	t.Run("error upload", func(t *testing.T) {
		mockUnitRepository.On("GetIDByUnitID", mock.Anything, mock.Anything).Return(uint(1), nil).Once()
		mockUploadRepo.On("UploadLocal", mock.Anything, mock.Anything).Return("localhost", errors.New("Error:")).Once()
		resss := multipart.FileHeader{}

		err := unitsService.AddPhoto(context.Background(), "c5c838ba-24f8-11ec-9621-0242ac130002", &resss)

		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("error update path", func(t *testing.T) {
		mockUnitRepository.On("GetIDByUnitID", mock.Anything, mock.Anything).Return(uint(1), nil).Once()
		mockUploadRepo.On("UploadLocal", mock.Anything, mock.Anything).Return("localhost", nil).Once()
		mockUnitRepository.On("AddPhotoUnit", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("Error:")).Once()
		resss := multipart.FileHeader{}

		err := unitsService.AddPhoto(context.Background(), "c5c838ba-24f8-11ec-9621-0242ac130002", &resss)

		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("success", func(t *testing.T) {
		mockUnitRepository.On("GetIDByUnitID", mock.Anything, mock.Anything).Return(uint(1), nil).Once()
		mockUploadRepo.On("UploadLocal", mock.Anything, mock.Anything).Return("localhost", nil).Once()
		mockUnitRepository.On("AddPhotoUnit", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		resss := multipart.FileHeader{}

		err := unitsService.AddPhoto(context.Background(), "c5c838ba-24f8-11ec-9621-0242ac130002", &resss)

		assert.Nil(t, err)
	})
}

func TestUnitService_GetID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUnitRepository.On("GetIDByUnitID", mock.Anything, mock.Anything).Return(uint(1), nil).Once()
		uuidID, _ := uuid.Parse("c5c838ba-24f8-11ec-9621-0242ac130002")
		res, err := unitsService.GetID(context.Background(), uuidID)

		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
}

func TestUnitService_Detail(t *testing.T) {
	t.Run("error parse", func(t *testing.T) {
		res, err := unitsService.Detail(context.Background(), "c5c838ba-24f8-11ec-9621-0242130002")

		assert.Equal(t, businesses.ErrInternalServer, err)
		assert.Nil(t, res)
	})

	t.Run("error not found", func(t *testing.T) {
		mockUnitRepository.On("GetIDByUnitID", mock.Anything, mock.Anything).Return(uint(0), errors.New("not found")).Once()
		res, err := unitsService.Detail(context.Background(), "c5c838ba-24f8-11ec-9621-0242ac130002")

		assert.Equal(t, businesses.ErrUnitNotFound, err)
		assert.Nil(t, res)
	})

	t.Run("error sql", func(t *testing.T) {
		mockUnitRepository.On("GetIDByUnitID", mock.Anything, mock.Anything).Return(uint(0), errors.New("Error:")).Once()
		res, err := unitsService.Detail(context.Background(), "c5c838ba-24f8-11ec-9621-0242ac130002")

		assert.Equal(t, errors.New("Error:"), err)
		assert.Nil(t, res)
	})

	t.Run("error get unit", func(t *testing.T) {
		mockUnitRepository.On("GetIDByUnitID", mock.Anything, mock.Anything).Return(uint(0), nil).Once()
		mockUnitRepository.On("GetByUnitID", mock.Anything, mock.Anything).Return(nil, errors.New("Error:")).Once()
		res, err := unitsService.Detail(context.Background(), "c5c838ba-24f8-11ec-9621-0242ac130002")

		assert.Nil(t, res)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("error get address", func(t *testing.T) {
		mockUnitRepository.On("GetIDByUnitID", mock.Anything, mock.Anything).Return(uint(0), nil).Once()
		mockUnitRepository.On("GetByUnitID", mock.Anything, mock.Anything).Return(&domainTest, nil).Once()
		mockUnitRepository.On("SelectAddressByID", mock.Anything, mock.Anything).Return(units.Address{}, errors.New("Error:")).Once()
		res, err := unitsService.Detail(context.Background(), "c5c838ba-24f8-11ec-9621-0242ac130002")

		assert.Nil(t, res)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("error get photo", func(t *testing.T) {
		mockUnitRepository.On("GetIDByUnitID", mock.Anything, mock.Anything).Return(uint(0), nil).Once()
		mockUnitRepository.On("GetByUnitID", mock.Anything, mock.Anything).Return(&domainTest, nil).Once()
		mockUnitRepository.On("SelectAddressByID", mock.Anything, mock.Anything).Return(units.Address{}, nil).Once()
		mockUnitRepository.On("SelectAllPhotosByID", mock.Anything, mock.Anything).Return([]units.Photo{}, errors.New("Error:")).Once()
		res, err := unitsService.Detail(context.Background(), "c5c838ba-24f8-11ec-9621-0242ac130002")

		assert.Nil(t, res)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("success", func(t *testing.T) {
		mockUnitRepository.On("GetIDByUnitID", mock.Anything, mock.Anything).Return(uint(0), nil).Once()
		mockUnitRepository.On("GetByUnitID", mock.Anything, mock.Anything).Return(&domainTest, nil).Once()
		mockUnitRepository.On("SelectAddressByID", mock.Anything, mock.Anything).Return(units.Address{}, nil).Once()
		mockUnitRepository.On("SelectAllPhotosByID", mock.Anything, mock.Anything).Return([]units.Photo{}, nil).Once()
		res, err := unitsService.Detail(context.Background(), "c5c838ba-24f8-11ec-9621-0242ac130002")

		assert.NotNil(t, res)
		assert.Nil(t, err)
	})
}

func TestUnitService_UnitsByGeo(t *testing.T) {
	t.Run("error get ip", func(t *testing.T) {
		mockIpapiRepo.On("GetLocationByIP", mock.Anything, mock.Anything).Return(&iplocator.Domain{}, errors.New("")).Once()
		res, err := unitsService.UnitsByGeo(context.Background(),"192.109.2.1","","")

		assert.Equal(t, businesses.ErrInternalServer, err)
		assert.Nil(t, res)
	})

	t.Run("error get unit by geo", func(t *testing.T) {
		mockUnitRepository.On("GetUnitsByGeo", mock.Anything, mock.Anything, mock.Anything).Return([]units.Result{}, errors.New("")).Once()
		res, err := unitsService.UnitsByGeo(context.Background(),"","1","2")

		assert.Equal(t, businesses.ErrInternalServer, err)
		assert.Nil(t, res)
	})

	t.Run("error not found by ip", func(t *testing.T) {
		mockIpapiRepo.On("GetLocationByIP", mock.Anything, mock.Anything).Return(&iplocator.Domain{}, nil).Once()
		mockUnitRepository.On("GetUnitsByGeo", mock.Anything, mock.Anything, mock.Anything).Return([]units.Result{}, errors.New("")).Once()
		res, err := unitsService.UnitsByGeo(context.Background(),"192.109.2.1","","")

		assert.Equal(t, businesses.ErrInternalServer, err)
		assert.Nil(t, res)
	})

	t.Run("error not found", func(t *testing.T) {
		mockIpapiRepo.On("GetLocationByIP", mock.Anything, mock.Anything).Return(&iplocator.Domain{}, nil).Once()
		mockUnitRepository.On("GetUnitsByGeo", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil).Once()
		res, err := unitsService.UnitsByGeo(context.Background(),"192.109.2.1","","")

		assert.Equal(t, businesses.ErrUnitNotFound, err)
		assert.Nil(t, res)
	})



	t.Run("success", func(t *testing.T) {
		mockIpapiRepo.On("GetLocationByIP", mock.Anything, mock.Anything).Return(&iplocator.Domain{}, nil).Once()
		mockUnitRepository.On("GetUnitsByGeo", mock.Anything, mock.Anything, mock.Anything).Return([]units.Result{}, nil).Once()
		res, err := unitsService.UnitsByGeo(context.Background(),"192.109.2.1","","")

		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
}
