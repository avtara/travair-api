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
	Owner		users.Users
	Name        string
	Category    string
	Price       uint
	AddressID   uint
	Address     Address
	Description string
	Policy      string
	Thumbnail   string
	Photos 		[]Photo
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

type Service interface {
	Add(ctx context.Context, data *Domain, userID uuid.UUID) (*Domain, error)
	ChangeThumbnail(ctx context.Context, unitID string, file *multipart.FileHeader) error
	AddPhoto(ctx context.Context, unitID string, file *multipart.FileHeader) error
	Detail(ctx context.Context, unitID string) (*Domain, error)
}

type Repository interface {
	Store(ctx context.Context, data *Domain) (*Domain, error)
	GetIDByUnitID(ctx context.Context, unitID uuid.UUID) (uint, error)
	UpdatePathByUnitID(ctx context.Context, unitID uuid.UUID, res string) error
	AddPhotoUnit(ctx context.Context, ID uint, path string) error
	GetByUnitID(ctx context.Context, unitID uuid.UUID) (*Domain,error)
	SelectAllPhotosByID(ctx context.Context, ID uint) ([]Photo, error)
	SelectAddressByID(ctx context.Context, ID uint) (Address, error)
}
