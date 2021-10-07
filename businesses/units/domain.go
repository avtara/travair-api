package units

import (
	"context"
	"github.com/avtara/travair-api/repository/databases/users"
	"github.com/google/uuid"
	"mime/multipart"
	"time"
)

type Domain struct {
	ID          uint
	UnitID      uuid.UUID
	OwnerID     uint
	Owner       users.Users
	Name        string
	Category    string
	Price       uint
	AddressID   uint
	Address     Address
	Description string
	Policy      string
	Thumbnail   string
	Photos      []Photo
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Address struct {
	Street     string
	City       string
	State      string
	Country    string
	PostalCode string
	Latitude   float64
	Longitude  float64
}

type Photo struct {
	Path string
}

type Result struct {
	Distance   float64 `json:"distance"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Name       string  `json:"name"`
	Street     string  `json:"street"`
	City       string  `json:"city"`
	State      string  `json:"state"`
	Country    string  `json:"country"`
	PostalCode string  `json:"postal_code"`
}

type Service interface {
	Add(ctx context.Context, data *Domain, userID uuid.UUID) (*Domain, error)
	ChangeThumbnail(ctx context.Context, unitID string, file *multipart.FileHeader) error
	AddPhoto(ctx context.Context, unitID string, file *multipart.FileHeader) error
	Detail(ctx context.Context, unitID string) (*Domain, error)
	GetID(ctx context.Context, unitID uuid.UUID) (uint, error)
	UnitsByGeo(ctx context.Context, ip string, long, lat string) ([]Result, error)
}

type Repository interface {
	Store(ctx context.Context, data *Domain) (*Domain, error)
	GetIDByUnitID(ctx context.Context, unitID uuid.UUID) (uint, error)
	UpdatePathByUnitID(ctx context.Context, unitID uuid.UUID, res string) error
	AddPhotoUnit(ctx context.Context, ID uint, path string) error
	GetByUnitID(ctx context.Context, unitID uuid.UUID) (*Domain, error)
	SelectAllPhotosByID(ctx context.Context, ID uint) ([]Photo, error)
	SelectAddressByID(ctx context.Context, ID uint) (Address, error)
	GetUnitsByGeo(ctx context.Context, lat, long float64) ([]Result, error)
}
