package factory

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	userDelivery "hery/cukur/features/users/delivery"
	userData "hery/cukur/features/users/repository"
	userUsecase "hery/cukur/features/users/services"

	loginDelivery "hery/cukur/features/login/delivery"
	loginData "hery/cukur/features/login/repository"
	loginUsecase "hery/cukur/features/login/services"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userDataF := userData.New(db)
	userUsecaseF := userUsecase.New(userDataF)
	userDelivery.New(e, userUsecaseF)

	loginDataF := loginData.New(db)
	loginUsecase := loginUsecase.New(loginDataF)
	loginDelivery.New(e, loginUsecase)
}
