package repository

import (
	"hery/cukur/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Image    string
	Phone    int
	Email    string `gorm:"unique"`
	Password string
	Bio      string
}

func FromDomain(data users.Core) User {
	return User{
		Model:    gorm.Model{ID: data.ID},
		Name:     data.Name,
		Phone:    data.Phone,
		Email:    data.Email,
		Password: data.Password,
		Bio:      data.Bio,
		Image:    data.Image,
	}
}

func FromCore(dataCore users.Core) User {
	dataModel := User{
		Name:     dataCore.Name,
		Image:    dataCore.Image,
		Email:    dataCore.Email,
		Password: dataCore.Password,
		Phone:    dataCore.Phone,
		Bio:      dataCore.Bio,
	}
	return dataModel
}

func (data *User) toCore() users.Core {
	return users.Core{
		ID:       data.ID,
		Name:     data.Name,
		Image:    data.Image,
		Phone:    data.Phone,
		Bio:      data.Bio,
		Email:    data.Email,
		Password: data.Password,
	}
}

func ToDomain(u User) users.Core {
	return users.Core{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Phone:    u.Phone,
		Bio:      u.Bio,
		Image:    u.Image,
	}
}

func toUserList(user []User) []users.Core {
	var userAll []users.Core
	for _, v := range user {
		userAll = append(userAll, users.Core{
			ID:    v.ID,
			Name:  v.Name,
			Image: v.Image,
			Bio:   v.Bio,
		})
	}
	return userAll
}
