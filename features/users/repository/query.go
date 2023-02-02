package repository

import (
	"hery/cukur/features/users"

	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.DataInterface {
	return &userData{
		db: db,
	}
}

func (repo *userData) Register(data users.Core) (int, error) {
	dataModel := FromCore(data)
	tx := repo.db.Create(&dataModel)
	if tx != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}
