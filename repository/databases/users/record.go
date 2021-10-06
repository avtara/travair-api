package users

import (
	"github.com/avtara/travair-api/businesses/users"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID       uint    `gorm:"primary_key:auto_increment"`
	UserID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name     string    `gorm:"type:varchar(255)"`
	Email    string    `gorm:"uniqueIndex;type:varchar(255)"`
	Password string    `gorm:"->;<-;not null" `
	Photo    string    `gorm:"type:varchar(255)"`
	Role     string    `gorm:"type:varchar(255)"`
	Status   int       `gorm:"default:0;size:10"`
}

func toDomain(rec *Users) *users.Domain {
	return &users.Domain{
		ID: rec.ID,
		UserID:    rec.UserID,
		Name:      rec.Name,
		Email:     rec.Email,
		Password:  rec.Password,
		Photo:     rec.Photo,
		Role:      rec.Role,
		Status:    rec.Status,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(domain *users.Domain) *Users {
	return &Users{
		Name:     domain.Name,
		Email:    domain.Email,
		Password: domain.Password,
		Photo:    domain.Photo,
		Role:     domain.Role,
		Status:   domain.Status,
	}
}
