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
func FromCoreList(data []users.Core) []UserResponse {
	var userAll []UserResponse
	for _, v := range data {
		userAll = append(userAll, UserResponse{
			ID:    v.ID,
			Name:  v.Name,
			Image: v.Image,
			Bio:   v.Bio,
		})
	}
	return userAll
}
