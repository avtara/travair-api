package reservations

import (
	"github.com/avtara/travair-api/businesses/reservations"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type repoUnit struct {
	DB *gorm.DB
}

func NewRepoMySQL(db *gorm.DB) reservations.Repository {
	return &repoUnit{
		DB: db,
	}
}

func (ur *repoUnit) Store(ctx context.Context, domain *reservations.Domain) (*reservations.Domain, error) {
	res := fromDomain(domain)

	ur.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&res).Error; err != nil {
			return err
		}
		return nil
	})

	return toDomain(res), nil
}

func (ur *repoUnit) GetByDate(ctx context.Context, domain *reservations.Domain) error {
	var rev Reservation
	if err := ur.DB.Where("check_in_date >= ? AND check_out_date <= ? AND unit_id = ?",
		domain.CheckInDate, domain.CheckOutDate, domain.UnitID).First(&rev).Error; err != nil {
		return err
	}
	return nil
}

