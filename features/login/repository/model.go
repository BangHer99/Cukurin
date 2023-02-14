package repository

import (
	"hery/cukur/features/login"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

func ToDomain(u User) login.Core {
	return login.Core{
		ID:       u.ID,
		Name:     u.Name,
		Password: u.Password,
	}
}

func FromDomain(du login.Core) User {
	return User{
		Model:    gorm.Model{ID: du.ID},
		Name:     du.Name,
		Email:    du.Email,
		Password: du.Password,
	}
}
