package userrepository

import (
	"errors"
	"users-service/internal/domain/model"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (repo *UserRepositoryImpl) SaveUser(user model.User) (model.User, error) {
	err := repo.db.Create(&user).Error

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (repo *UserRepositoryImpl) GetUser(Id int) (model.User, error) {
	var userData model.User

	err := repo.db.First(&userData, "user_id = ?", Id).Error

	if err != nil {
		return model.User{}, errors.New("User tidak ditemukan")
	}

	return userData, nil
}

func (repo *UserRepositoryImpl) GetUsers() ([]model.User, error) {
	var users []model.User

	err := repo.db.Find(&users).Error

	if err != nil {
		return []model.User{}, err
	}

	return users, nil
}
