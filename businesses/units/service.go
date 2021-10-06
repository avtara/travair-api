package units

import (
	"context"
	"github.com/avtara/travair-api/businesses"
	"github.com/avtara/travair-api/businesses/uploads"
	"github.com/avtara/travair-api/businesses/users"
	"github.com/google/uuid"
	"mime/multipart"
	"strings"
	"time"
)

type unitService struct {
	unitRepository Repository
	userService    users.Service
	contextTimeout time.Duration
	uploadRepo     uploads.Repository
}

func NewUnitService(rep Repository, us users.Service, to time.Duration, ur uploads.Repository) Service {
	return &unitService{
		unitRepository: rep,
		userService:    us,
		contextTimeout: to,
		uploadRepo: ur,
	}
}

func (us *unitService) Add(ctx context.Context, data *Domain, userID uuid.UUID) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, us.contextTimeout)
	defer cancel()

	id, _ := us.userService.GetID(ctx, userID)

	data.OwnerID = id
	res, err := us.unitRepository.Store(ctx, data)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (us *unitService) ChangeThumbnail(ctx context.Context, unitID string,file *multipart.FileHeader) error {
	ctx, cancel := context.WithTimeout(ctx, us.contextTimeout)
	defer cancel()

	uuidUnitID, err := uuid.Parse(unitID)
	if err != nil {
		return businesses.ErrInternalServer
	}

	if _, err := us.unitRepository.GetIDByUnitID(ctx, uuidUnitID); err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
		return businesses.ErrUnitNotFound
	}

	res, err := us.uploadRepo.UploadLocal(file, "unit/"+unitID)
	if err != nil {
		return businesses.ErrInternalServer
	}

	if err = us.unitRepository.UpdatePathByUnitID(ctx, uuidUnitID, res); err != nil {
		return businesses.ErrInternalServer
	}
	return nil
}

func (us *unitService) AddPhoto(ctx context.Context, unitID string, file *multipart.FileHeader) error {
	ctx, cancel := context.WithTimeout(ctx, us.contextTimeout)
	defer cancel()

	uuidUnitID, err := uuid.Parse(unitID)
	if err != nil {
		return businesses.ErrInternalServer
	}

	id, err := us.unitRepository.GetIDByUnitID(ctx, uuidUnitID)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
		return businesses.ErrUnitNotFound
	}
	res, err := us.uploadRepo.UploadLocal(file, "unit/"+unitID)
	if err != nil {
		return businesses.ErrInternalServer
	}

	if err := us.unitRepository.AddPhotoUnit(ctx, id, res); err != nil {
		return businesses.ErrInternalServer
	}

	return nil
}

func (us *unitService) Detail(ctx context.Context, unitID string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, us.contextTimeout)
	defer cancel()

	uuidUnitID, err := uuid.Parse(unitID)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	id, err := us.unitRepository.GetIDByUnitID(ctx, uuidUnitID)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return nil, err
		}
		return nil, businesses.ErrUnitNotFound
	}

	res, err := us.unitRepository.GetByUnitID(ctx, uuidUnitID)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}
	res.Address, err = us.unitRepository.SelectAddressByID(ctx, id)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}
	res.Photos, err = us.unitRepository.SelectAllPhotosByID(ctx,id)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}
