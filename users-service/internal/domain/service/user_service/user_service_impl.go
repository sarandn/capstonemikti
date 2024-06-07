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

func (service *UserServiceImpl) UpdateUser(request request.UserUpdateServiceRequest, pathId int) (map[string]interface{}, error) {
	getUserById, err := service.repo.GetUser(pathId)

	if err != nil {
		return nil, err
	}

	// validasi jika datanya kosong atau tidak
	if request.FullName == "" {
		request.FullName = getUserById.FullName
	}

	if request.Email == "" {
		request.Email = getUserById.Email
	}

	if request.PhoneNum == "" {
		request.PhoneNum = getUserById.PhoneNum
	}

	if request.Address == "" {
		request.Address = getUserById.Address
	}

	userReq := model.User{
		UserID:   pathId,
		FullName: request.FullName,
		Email:    request.Email,
		Password: getUserById.Password,
		PhoneNum: request.PhoneNum,
		Address:  request.Address,
		RoleIDFK: getUserById.RoleIDFK,
	}

	// melakukan update dari userReq
	updateUser, errUpdateUser := service.repo.UpdateUsers(userReq)

	if errUpdateUser != nil {
		return nil, errUpdateUser
	}

	return ResponseToJson{"full_name": updateUser.FullName, "email": updateUser.Email}, nil
}

func (service *UserServiceImpl) DeleteData(userId int) (entity.UserEntity, error) {
	delUser, errDelUser := service.repo.DeleteUser(userId)

	if errDelUser != nil {
		return entity.UserEntity{}, errDelUser
	}

	return entity.ToUserEntity(delUser), nil
}
