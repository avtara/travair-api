package units

import (
	"github.com/avtara/travair-api/businesses/units"
	"github.com/avtara/travair-api/repository/databases/users"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Units struct {
	gorm.Model
	UnitID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	UsersID     uint
	Users       users.Users
	Name        string `gorm:"type:varchar(255)"`
	Category    string `gorm:"type:varchar(20)"`
	Price       uint
	AddressID   uint
	Address     Address
	Description string
	Policy      string
	Thumbnail   string `gorm:"type:varchar(255)"`
}

type Address struct {
	gorm.Model
	Street     string `gorm:"type:varchar(255)"`
	City       string `gorm:"type:varchar(255)"`
	State      string `gorm:"type:varchar(255)"`
	Country    string `gorm:"type:varchar(255)"`
	PostalCode string `gorm:"type:varchar(255)"`
	Latitude   float64
	Longitude  float64
}

type Photos struct {
	UnitID uint
	Unit   Units  `gorm:"foreignKey:UnitID"`
	Path   string `gorm:"type:varchar(255)"`
}

type Result struct {
	Distance   float64
	Latitude   float64
	Longitude  float64
	Name       string
	Street     string
	City       string
	State      string
	Country    string
	PostalCode string
}

func resultToDomain(data Result) units.Result {
	return units.Result{
		Name:       data.Name,
		Latitude:   data.Latitude,
		Longitude:  data.Longitude,
		Distance:   data.Distance,
		Street:     data.Street,
		City:       data.City,
		State:      data.State,
		Country:    data.Country,
		PostalCode: data.PostalCode,
	}
}

func resultsToDomain(data []Result) []units.Result {
	var res []units.Result
	for _, s := range data {
		if s.Distance < 20 {
			res = append(res, resultToDomain(s))
		}
	}
	return res
}

func addUnitToDomain(data Units) *units.Domain {
	return &units.Domain{
		UnitID:      data.UnitID,
		OwnerID:     data.UsersID,
		Name:        data.Name,
		Category:    data.Category,
		Price:       data.Price,
		Description: data.Description,
		Policy:      data.Policy,
		Thumbnail:   data.Thumbnail,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func fromDomain(data *units.Domain) *Units {
	return &Units{
		Name:        data.Name,
		UsersID:     data.OwnerID,
		Category:    data.Category,
		Price:       data.Price,
		AddressID:   data.AddressID,
		Description: data.Description,
		Policy:      data.Policy,
		Thumbnail:   data.Thumbnail,
		Address:     *addressFromDomain(data)}
}

func addressFromDomain(domain *units.Domain) *Address {
	return &Address{
		Street:     domain.Address.Street,
		City:       domain.Address.City,
		State:      domain.Address.State,
		Country:    domain.Address.Country,
		PostalCode: domain.Address.PostalCode,
		Latitude:   domain.Address.Latitude,
		Longitude:  domain.Address.Longitude,
	}
}

func photoToDomain(data Photos) units.Photo {
	return units.Photo{
		Path: data.Path,
	}
}

func photosToDomain(data []Photos) []units.Photo {
	var res []units.Photo
	for _, s := range data {
		res = append(res, photoToDomain(s))
	}
	return res
}

func detailToDomain(data Units) *units.Domain {
	return &units.Domain{
		UnitID:      data.UnitID,
		Name:        data.Name,
		Category:    data.Category,
		Price:       data.Price,
		Description: data.Description,
		Policy:      data.Policy,
		Thumbnail:   data.Thumbnail,
		Owner:       data.Users,
		OwnerID:     data.UsersID,
	}
}

func addressToDomain(data Address) units.Address {
	return units.Address{
		Street:     data.Street,
		City:       data.City,
		State:      data.State,
		Country:    data.Country,
		PostalCode: data.PostalCode,
		Latitude:   data.Latitude,
		Longitude:  data.Longitude,
	}
}
