package mysql

import (
	"fmt"
	"hery/cukur/config"
	"log"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func InitDB(config *config.AppConfig) *gorm.DB {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
