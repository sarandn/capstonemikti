package entity

import "users-service/internal/domain/model"

type RoleEntity struct {
	RoleID   int         `json:"role_id"`
	RoleName string      `json:"role_name"`
	User     interface{} `json:"user"`
}

func ToRoleEntity(role model.Role) RoleEntity {
	if role.User != nil {
		user := ToUserListEntity(role.User)

		return RoleEntity{
			RoleID:   role.RoleID,
			RoleName: role.RoleName,
			User:     user,
		}
	}

	return RoleEntity{
		RoleID:   role.RoleID,
		RoleName: role.RoleName,
		User:     "User belum terisi",
	}
}

func ToRoleListEntity(roles []model.Role) []RoleEntity {
	roleData := []RoleEntity{}

	for _, role := range roles {
		roleData = append(roleData, ToRoleEntity(role))
	}

	return roleData
}
