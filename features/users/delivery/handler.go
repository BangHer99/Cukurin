package delivery

import (
	"hery/cukur/features/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userUsecase users.UsecaseInterface
}

func New(e *echo.Echo, usecase users.UsecaseInterface) {
	handler := &UserDelivery{
		userUsecase: usecase,
	}
	e.POST("/registers", handler.RegisterUser)
}

func (delivery *UserDelivery) RegisterUser(c echo.Context) error {
	var dataRequest UserRequest
	errBind := c.Bind(&dataRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to bind",
		})
	}
	_, err := delivery.userUsecase.Create(toCore(dataRequest))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Register",
		})
	}
	// if row != 1 {
	// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 		"message": "failed register new user",
	// 	})
	// }
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Success Register",
	})
}
