package response

import (
	"github.com/avtara/travair-api/businesses/units"
	"github.com/google/uuid"
	"time"
)

type AddUnit struct {
	UserID    uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func FromDomain(domain *units.Domain) *AddUnit {
	return &AddUnit{
		UserID:    domain.UnitID,
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt,
	}
}

type Detail struct {
	OwnerName    string        `json:"owner_name"`
	OwnerPicture string        `json:"owner_picture"`
	Name         string        `json:"name"`
	Category     string        `json:"category"`
	Price        uint          `json:"price"`
	Address      units.Address `json:"address"`
	Description  string        `json:"description"`
	Policy       string        `json:"policy"`
	Thumbnail    string        `json:"thumbnail"`
	Photos       []units.Photo `json:"photos"`
}

type Address struct {
	Street     string  `json:"street"`
	City       string  `json:"city"`
	State      string  `json:"state"`
	Country    string  `json:"country"`
	PostalCode string  `json:"postal_code"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

type Photo struct {
	Path string `json:"path"`
}

func DetailFromDomain(domain *units.Domain) *Detail {
	return &Detail{
		OwnerName: domain.Owner.Name,
		OwnerPicture: domain.Owner.Photo,
		Name: domain.Name,
		Category: domain.Category,
		Price: domain.Price,
		Address: domain.Address,
		Description: domain.Description,
		Policy: domain.Policy,
		Thumbnail: domain.Thumbnail,
		Photos: domain.Photos,
	}
}
