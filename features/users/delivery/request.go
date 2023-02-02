package delivery

import "hery/cukur/features/users"

type UserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func toCore(data UserRequest) users.Core {
	return users.Core{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
}
