package users

import (
	"context"
	"fmt"
	"github.com/avtara/travair-api/businesses/users"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/url"
)

type repoUsers struct {
	DB *gorm.DB
}

func NewRepoMySQL(db *gorm.DB) users.Repository {
	return &repoUsers{
		DB: db,
	}
}

func (ru *repoUsers) StoreNewUsers (ctx context.Context, data *users.Domain) (*users.Domain, error) {
	user := fromDomain(data)
	user.Photo = fmt.Sprintf("https://avatars.dicebear.com/api/miniavs/%s.svg", url.QueryEscape(data.Name))
	if err := ru.DB.Create(&user); err.Error != nil {
		return nil, err.Error
	}
	result := toDomain(user)
	return result, nil
}

func (ru *repoUsers) GetByEmail (ctx context.Context, email string) (*users.Domain, error) {
	var user Users
	if err := ru.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	result := toDomain(&user)
	return result, nil
}

func (ru *repoUsers) UpdateStatus(ctx context.Context, userID uuid.UUID) error {
	var user Users
	if err := ru.DB.Model(&user).Where("user_id = ?", userID).Update("status", 1).Error; err != nil {
		return err
	}

	return nil
}

func (ru *repoUsers) GetByUserID(ctx context.Context, userID uuid.UUID) (*users.Domain, error) {
	var user Users
	if err := ru.DB.Where("user_id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	result := toDomain(&user)
	return result, nil
}

func (ru *repoUsers) GetByEmailAndPassword(ctx context.Context, email string) (*users.Domain, error) {
	var user Users
	if err := ru.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	result := toDomain(&user)
	return result, nil
}
