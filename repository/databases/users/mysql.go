package users

import (
	"context"
	"fmt"
	"github.com/avtara/travair-api/businesses/users"
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
	fmt.Println("BEFORE: ", user)
	user.Photo = fmt.Sprintf("https://avatars.dicebear.com/api/miniavs/%s.svg", url.QueryEscape(data.Name))
	if err := ru.DB.Create(&user); err.Error != nil {
		return nil, err.Error
	}
	fmt.Println("AFTER: ", user)
	result := toDomain(user)
	return result, nil
}
func (ru *repoUsers) GetByEmail (ctx context.Context, data *users.Domain) (*users.Domain, error) {
	var user Users
	fmt.Println("BEFORE: ", data.Email)

	if err := ru.DB.Where("email = ?", data.Email).First(&user).Error; err != nil {
		return nil, err
	}
	fmt.Println("AFTER: ", user)
	result := toDomain(&user)
	return result, nil
}