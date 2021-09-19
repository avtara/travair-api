package controller

import (
	"fmt"
	"github.com/avtara/travair-api/models"
	"github.com/google/uuid"
	"time"
)

type RegistrationUserRequest struct {
	Name     string      `json:"name" validate:"required"`
	Username string      `json:"username" validate:"required,min=5,max=50"`
	Email    string      `json:"email" validate:"required,email"`
	Password string      `json:"password" validate:"required,password"`
	Phone    string      `json:"phone"`
	Role     models.Role `json:"role" validate:"required,role"`
}

type RegistrationUserResponse struct {
	UserID    uuid.UUID   `json:"user_id"`
	Name      string      `json:"name"`
	Role      models.Role `json:"role"`
	CreatedAt time.Time   `json:"created_at"`
}

func RegistrationUserRequestToModelUser(data *RegistrationUserRequest) *models.User {
	return &models.User{
		Name:     data.Name,
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
		Role:     data.Role,
		Phone:    data.Phone,
		Photo:    fmt.Sprintf("https://avatars.dicebear.com/api/micah/%s.svg", data.Username),
	}
}

func ModelUserToRegistrationUserResponse(data *models.User) *RegistrationUserResponse {
	return &RegistrationUserResponse{
		UserID:    data.UserID,
		Name:      data.Name,
		Role:      data.Role,
		CreatedAt: data.CreatedAt,
	}
}
