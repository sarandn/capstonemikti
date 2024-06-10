package entity

import "users-service/internal/domain/model"

type UserEntity struct {
	UserID   int    `json:"user_id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	PhoneNum string `json:"phone_num"`
	Address  string `json:"address"`
}

func ToUserEntity(user model.User) UserEntity {
	return UserEntity{
		UserID:   user.UserID,
		Email:    user.Email,
		FullName: user.FullName,
		PhoneNum: user.PhoneNum,
		Address:  user.Address,
	}
}

func ToUserListEntity(users []model.User) []UserEntity {
	UserData := []UserEntity{}

	for _, user := range users {
		UserData = append(UserData, ToUserEntity(user))
	}

	return UserData
}
