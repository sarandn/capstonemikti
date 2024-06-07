package userrepository

import "users-service/internal/domain/model"

type UserRepository interface {
	SaveUser(user model.User) (model.User, error)
	GetUser(Id int) (model.User, error)
	GetUsers() ([]model.User, error)
}
