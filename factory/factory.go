package factory

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	userDelivery "hery/cukur/features/users/delivery"
	userData "hery/cukur/features/users/repository"
	userUsecase "hery/cukur/features/users/services"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userDataF := userData.New(db)
	userUsecaseF := userUsecase.New(userDataF)
	userDelivery.New(e, userUsecaseF)
}
