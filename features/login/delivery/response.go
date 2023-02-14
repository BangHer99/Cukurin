package delivery

import "hery/cukur/features/login"

type LoginResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

func SuccessResponseWithData(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func toResponse(data LoginResponse) login.Core { // tidak pakai ini karna tidak terhubung ke database
	return login.Core{
		ID:    data.ID,
		Name:  data.Name,
		Token: data.Token,
	}
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {

	case "login":
		cnv := core.(login.Core)
		res = LoginResponse{ID: cnv.ID, Name: cnv.Name, Token: cnv.Token}
	}
	return res
}
