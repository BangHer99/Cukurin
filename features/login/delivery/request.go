package delivery

import "hery/cukur/features/login"

type LoginRequest struct {
	ID       uint   `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LoginFormat struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ToDomain(i interface{}) login.Core { // sama dengan cara di baris 27
	switch i.(type) {

	case LoginFormat:
		cnv := i.(LoginFormat)
		return login.Core{Email: cnv.Email, Password: cnv.Password}
	}
	return login.Core{}
}

func toCore(data LoginFormat) login.Core { // sama dengan cara di baris 17
	return login.Core{
		Email:    data.Email,
		Password: data.Password,
	}
}
