package delivery

import "hery/cukur/features/users"

type UserResponse struct {
	ID       uint
	Name     string
	Image    string
	Password string
	Phone    int
	Bio      string
	Email    string
}

func FromCore(data users.Core) UserResponse {
	return UserResponse{
		ID:    data.ID,
		Name:  data.Name,
		Image: data.Image,
		Phone: data.Phone,
		Email: data.Email,
		Bio:   data.Bio,
	}
}
