package request

import "github.com/avtara/travair-api/businesses/users"

type UserRegistration struct {
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,password"`
	Role      string `json:"role" validate:"required,role"`
}

func (rec *UserRegistration) ToDomain() *users.Domain{
	return &users.Domain{
		Name      :rec.Name,
		Email     :rec.Email,
		Password  :rec.Password,
		Role      :rec.Role,
	}
}