package services

import (
	"hery/cukur/features/login"
	"hery/cukur/middlewares"

	"golang.org/x/crypto/bcrypt"
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
	res, err := usecase.dataLogin.Login(input)
	if err != nil {
		return login.Core{}, "", err
	}
	pass := login.Core{Password: res.Password}
	check := bcrypt.CompareHashAndPassword([]byte(pass.Password), []byte(input.Password))
	if check != nil {
		return login.Core{}, "", check
	}
	// if input.Email == "" || input.Password == "" {
	// 	return login.Core{}, "masukan email dan password", errors.New("")
	// }
	// results, errEmail := usecase.dataLogin.Login(input)
	// if errEmail != nil {
	// 	return login.Core{}, "email not found", errEmail
	// }
	token, errToken := middlewares.CreateToken(int(res.ID))
	if errToken != nil {
		return login.Core{}, "error to created token", errToken
	}

	return res, token, nil

}
