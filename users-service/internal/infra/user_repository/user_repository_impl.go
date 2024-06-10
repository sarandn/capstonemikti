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

func (repo *UserRepositoryImpl) UpdateUsers(user model.User) (model.User, error) {
	err := repo.db.Model(model.User{}).Where("user_id = ?", user.UserID).Updates(user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo *UserRepositoryImpl) DeleteUser(Id int) (model.User, error) {
	var userData model.User

	err := repo.db.Where("user_id = ?", Id).Delete(&userData).Error

	if err != nil {
		return model.User{}, err
	}

	return userData, nil
}

func (repo *UserRepositoryImpl) FindUserByEmail(email string) (*model.User, error) {
	user := new(model.User)

	if err := repo.db.Where("email = ?", email).Take(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
