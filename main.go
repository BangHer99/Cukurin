package main

import (
	"fmt"
	"hery/cukur/config"
	"hery/cukur/factory"
	"hery/cukur/migration"

	"hery/cukur/utils/database/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	e := echo.New()
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	migration.InitMigrate(db)
	factory.InitFactory(e, db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))
}
