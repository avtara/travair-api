package units

import (
	"context"
	"fmt"
	"github.com/avtara/travair-api/businesses/units"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repoUnit struct {
	DB *gorm.DB
}

func NewRepoMySQL(db *gorm.DB) units.Repository {
	return &repoUnit{
		DB: db,
	}
}

func (ur *repoUnit) Store(ctx context.Context, data *units.Domain) (*units.Domain, error) {
	unit := fromDomain(data)
	if err := ur.DB.Create(&unit).Error; err != nil {
		return nil, err
	}

	return addUnitToDomain(*unit), nil
}

func (ur *repoUnit) GetIDByUnitID(ctx context.Context, unitID uuid.UUID) (uint, error) {
	var unit Units
	if err := ur.DB.Where("unit_id = ?", unitID).First(&unit).Error; err != nil {
		return 0, err
	}
	return unit.ID, nil
}

func (ur *repoUnit) UpdatePathByUnitID(ctx context.Context, unitID uuid.UUID, res string) error {
	var unit Units
	if err := ur.DB.Model(&unit).Where("unit_id = ?", unitID).Update("thumbnail", res).Error; err != nil {
		return err
	}

	return nil
}

func (ur *repoUnit) AddPhotoUnit(ctx context.Context, ID uint, path string) error {
	if err := ur.DB.Create(&Photos{UnitID: ID, Path: path}).Error; err != nil {
		return err
	}
	return nil
}

func (ur *repoUnit) SelectAllPhotosByID(ctx context.Context, ID uint) ([]units.Photo, error) {
	var photos []Photos
	if err := ur.DB.Find(&photos).Where("id = ?", ID).Error; err != nil {
		return nil, err
	}
	return photosToDomain(photos), nil
}

func (ur *repoUnit) GetByUnitID(ctx context.Context, unitID uuid.UUID) (*units.Domain,error) {
	var res Units
	if err := ur.DB.Preload("Users").Find(&res).Where("unit_id = ?", unitID).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return detailToDomain(res), nil
}

func (ur *repoUnit) SelectAddressByID(ctx context.Context, ID uint) (units.Address, error) {
	var address Address
	if err := ur.DB.Find(&address).Where("id = ?", ID).Error; err != nil {
		fmt.Println(err)
		return units.Address{}, err
	}

	return addressToDomain(address), nil
}