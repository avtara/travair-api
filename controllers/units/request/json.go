package request

import (
	"github.com/avtara/travair-api/businesses/units"
)

type Unit struct {
	Name        string  `json:"name" validate:"required"`
	Category    string  `json:"category" validate:"required,unit_category"`
	Price       uint   `json:"price" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Policy      string  `json:"policy" validate:"required"`
	Thumbnail   string  `json:"thumbnail" validate:"required"`
	Address     Address `json:"address" validate:"required"`
}

type Address struct {
	Street     string  `json:"street" validate:"required"`
	City       string  `json:"city" validate:"required"`
	State      string  `json:"state" validate:"required"`
	Country    string  `json:"country" validate:"required"`
	PostalCode string  `json:"postal_code" validate:"required"`
	Latitude   float64 `json:"latitude" validate:"required"`
	Longitude  float64 `json:"longitude" validate:"required"`
}


func (rec *Unit) AddUnitToDomain() *units.Domain{
	return &units.Domain{
		Name        : rec.Name,
		Category    : rec.Category,
		Price       : rec.Price,
		Description : rec.Description,
		Policy      : rec.Policy,
		Thumbnail   : rec.Thumbnail,
		Address     : units.Address(rec.Address),
	}
}