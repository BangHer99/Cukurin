package repository

import (
	"errors"
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
func (repo *userData) GetUser() ([]users.Core, error) {
	dataUsers := []User{}
	tx := repo.db.Find(&dataUsers)
	if tx.Error != nil {
		return nil, tx.Error
	}
	allUser := toUserList(dataUsers)
	return allUser, nil
}

func (repo *userData) UpdateId(data users.Core) (row int, err error) {
	tx := repo.db.Model(&User{}).Where("id = ?", data.ID).Updates(FromCore(data))
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed to update data")
	}

	return int(tx.RowsAffected), nil
	// tx := repo.db.Model(&User{}).Where("id = ?", data.ID).Updates(FromCore(data))
	// if tx.Error != nil {
	// 	return -1, errors.New("failed update data")
	// }

	// return int(tx.RowsAffected), nil
}
func (repo *userData) Upgrade(input users.Core) (row users.Core, err error) {
	tx := repo.db.Model(&User{}).Where("id = ?", input.ID).Updates(FromCore(input))
	if tx.Error != nil {
		return users.Core{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return users.Core{}, errors.New("failed to update data")
	}

	return input, nil

}
