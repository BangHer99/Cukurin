package delivery

import (
	"hery/cukur/features/users"
	"hery/cukur/middlewares"
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
	e.GET("/users", handler.GetAllUsers)
	e.PUT("/usersasas", handler.PutUser, middlewares.JWTMiddleware())
	e.PUT("/users", handler.Put, middlewares.JWTMiddleware())

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

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Success Register",
	})
}

func (delivery *UserDelivery) GetAllUsers(c echo.Context) error {
	res, err := delivery.userUsecase.AllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed get All users",
		})

	}
	respon := FromCoreList(res)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   respon,
	})

}

func (delivery *UserDelivery) PutUser(c echo.Context) error {
	var putUser UpdateFormat
	token := middlewares.ExtractToken(c)
	errBind := c.Bind(&putUser)
	if errBind != nil {
		return c.JSON(http.StatusBadGateway, map[string]interface{}{
			"message": "faield update data",
		})
	}

	var update users.Core
	if putUser.Email != "" {
		update.Email = putUser.Email
	}
	if putUser.Password != "" {
		update.Password = putUser.Password
	}
	if putUser.Bio != "" {
		update.Bio = putUser.Bio
	}
	if putUser.Name != "" {
		update.Name = putUser.Name
	}

	update.ID = uint(token)
	_, err := delivery.userUsecase.UpdateUser(update)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed update",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update",
		// "new data": data,
	})
}
func (delivery *UserDelivery) Put(c echo.Context) error {
	var putUser UpdateFormat
	token := middlewares.ExtractToken(c)
	errBind := c.Bind(&putUser)
	if errBind != nil {
		return c.JSON(http.StatusBadGateway, map[string]interface{}{
			"message": "faield update data",
		})
	}

	var update users.Core
	if putUser.Email != "" {
		update.Email = putUser.Email
	}
	if putUser.Password != "" {
		update.Password = putUser.Password
	}
	if putUser.Bio != "" {
		update.Bio = putUser.Bio
	}
	if putUser.Name != "" {
		update.Name = putUser.Name
	}
	if putUser.Phone != 0 {
		update.Phone = putUser.Phone
	}
	if putUser.Image != "" {
		update.Image = putUser.Image
	}

	update.ID = uint(token)
	data, err := delivery.userUsecase.Put(update)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed update",
		})
	}
	return c.JSON(http.StatusUpgradeRequired, map[string]interface{}{
		"message":  "success update",
		"new data": data,
	})
}
