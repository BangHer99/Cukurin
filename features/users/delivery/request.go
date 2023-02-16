package delivery

import "hery/cukur/features/users"

type UserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UpdateFormat struct {
	Name     string `json:"name"`
	Phone    int    `json:"phone"`
	Email    string `json:"email"`
	Image    string `json:"image"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}

func toCore(data UserRequest) users.Core {
	return users.Core{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
}
func toUpdate(data UpdateFormat) users.Core {
	return users.Core{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
}
