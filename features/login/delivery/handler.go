package delivery

import (
	"hery/cukur/features/login"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoginDelivery struct {
	loginUsecase login.UsecaseInterface
}

func New(e *echo.Echo, usecase login.UsecaseInterface) {
	handler := LoginDelivery{
		loginUsecase: usecase,
	}
	e.POST("/login", handler.Login)
}

func (delivery *LoginDelivery) Login(c echo.Context) error {
	var req LoginFormat
	errBind := c.Bind(&req)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed login server error",
		})
	}
	cnv := toCore(req)
	res, token, err := delivery.loginUsecase.Login(cnv)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "user not found",
		})
	}
	if token == "please input email and password" || token == "email not found" || token == "wrong password" {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed login",
		})
	}

	res.Token = token

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"data":    ToResponse(res, "login"),
	})

}
