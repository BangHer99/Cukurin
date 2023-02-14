package services

import (
	"errors"
	"hery/cukur/features/login"
	"hery/cukur/middlewares"
)

type loginUsecase struct {
	dataLogin login.DataInterface
}

func New(data login.DataInterface) login.UsecaseInterface {
	return &loginUsecase{
		dataLogin: data,
	}
}
func (usecase *loginUsecase) Login(input login.Core) (login.Core, string, error) {
	if input.Email == "" || input.Password == "" {
		return login.Core{}, "masukan email dan password", errors.New("")
	}
	results, errEmail := usecase.dataLogin.Login(input)
	if errEmail != nil {
		return login.Core{}, "email not found", errEmail
	}
	token, errToken := middlewares.CreateToken(int(results.ID))
	if errToken != nil {
		return login.Core{}, "error to created token", errToken
	}

	return results, token, nil

}
