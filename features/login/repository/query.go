package repository

import (
	"hery/cukur/features/login"
	"log"

	"gorm.io/gorm"
)

type loginData struct {
	db *gorm.DB
}

func New(db *gorm.DB) login.DataInterface {
	return &loginData{
		db: db,
	}
}

func (repo *loginData) Login(input login.Core) (login.Core, error) {

	cnv := FromDomain(input)
	if err := repo.db.Where("email = ?", cnv.Email).Find(&cnv).Error; err != nil {
		log.Fatal("error get data")
		return login.Core{}, err
	}
	input = ToDomain(cnv)
	return input, nil
}
