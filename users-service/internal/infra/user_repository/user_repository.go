package userrepository

import "users-service/internal/domain/model"

type UserRepository interface {
	SaveUser(user model.User) (model.User, error)
	GetUser(Id int) (model.User, error)
	GetUsers() ([]model.User, error)
	UpdateUsers(user model.User) (model.User, error)
	DeleteUser(Id int) (model.User, error)
	FindUserByEmail(email string) (*model.User, error)
}
