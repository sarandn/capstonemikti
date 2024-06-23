package userservice

import (
	"users-service/internal/domain/model/entity"
	"users-service/internal/interfaces/api/request"
)

type UserService interface {
	SaveUser(request request.UserServiceRequest) (map[string]interface{}, error)
	GetUser(userId int) (entity.UserEntity, error)
	GetUsers() ([]entity.UserEntity, error)
	UpdateUser(request request.UserUpdateServiceRequest, pathId int) (map[string]interface{}, error)
	DeleteData(userId int) (entity.UserEntity, error)
	LoginUser(email string, password string) (map[string]interface{}, error)
}
