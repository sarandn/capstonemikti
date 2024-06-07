package userservice

import (
	"users-service/internal/domain/model"
	"users-service/internal/domain/model/entity"
	userrepository "users-service/internal/infra/user_repository"
	"users-service/internal/interfaces/api/request"

	"golang.org/x/crypto/bcrypt"
)

type ResponseToJson map[string]interface{}

type UserServiceImpl struct {
	repo userrepository.UserRepository
}

func NewUserService(repo userrepository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		repo: repo,
	}
}

func (service *UserServiceImpl) SaveUser(request request.UserServiceRequest) (map[string]interface{}, error) {

	// hash untuk password
	passHash, errHash := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)

	if errHash != nil {
		return nil, errHash
	}

	userReq := model.User{
		RoleIDFK: request.RoleIDFK,
		FullName: request.FullName,
		Email:    request.Email,
		Password: string(passHash),
		PhoneNum: request.PhoneNum,
		Address:  request.Address,
	}

	saveUser, errSaveUser := service.repo.SaveUser(userReq)

	if errSaveUser != nil {
		return nil, errSaveUser
	}

	return ResponseToJson{
		"role_id_fk": saveUser.RoleIDFK,
		"full_name":  saveUser.FullName,
		"email":      saveUser.Email,
		"password":   saveUser.Password,
		"phone_num":  saveUser.PhoneNum,
		"address":    saveUser.Address,
	}, nil
}

func (service *UserServiceImpl) GetUser(userId int) (entity.UserEntity, error) {
	getUser, errGetUser := service.repo.GetUser(userId)

	if errGetUser != nil {
		return entity.UserEntity{}, nil
	}

	return entity.ToUserEntity(getUser), nil
}

func (service *UserServiceImpl) GetUsers() ([]entity.UserEntity, error) {
	getUserList, errGetUserList := service.repo.GetUsers()

	if errGetUserList != nil {
		return []entity.UserEntity{}, errGetUserList
	}

	return entity.ToUserListEntity(getUserList), nil
}
