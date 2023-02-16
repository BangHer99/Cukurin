package services

import (
	"errors"
	"hery/cukur/features/users"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userData users.DataInterface
}

func New(data users.DataInterface) users.UsecaseInterface {
	return &userUsecase{
		userData: data,
	}
}

func (usecase *userUsecase) Create(data users.Core) (int, error) {
	if data.Name == "" {
		return -1, errors.New("masukan nama")
	}
	if data.Email == "" {
		return -1, errors.New("masukan email")
	}
	if data.Password == "" {
		return -1, errors.New("masukan password")
	}
	hashPass, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}
	data.Password = string(hashPass)
	row, err := usecase.userData.Register(data)
	if err != nil {
		return -1, err
	}
	return row, nil
}

func (usecase *userUsecase) AllUsers() ([]users.Core, error) {
	user, err := usecase.userData.GetUser()
	if err != nil {
		return nil, errors.New("failed get data")
	} else if len(user) == 0 {
		return nil, errors.New("users not available")
	} else {
		return user, nil
	}

}

func (usecase *userUsecase) UpdateUser(data users.Core) (int, error) {
	row, err := usecase.userData.UpdateId(data)

	return row, err
}
func (usecase *userUsecase) Put(data users.Core) (users.Core, error) {
	input, err := usecase.userData.Upgrade(data)

	return input, err
}
