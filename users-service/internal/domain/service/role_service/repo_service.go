package roleservice

import (
	"users-service/internal/domain/model/entity"
	"users-service/internal/interfaces/api/request"
)

type RoleService interface {
	SaveRole(request request.RoleServiceRequest) (map[string]interface{}, error)
	GetRole(roleId int) (entity.RoleEntity, error)
	GetRoles() ([]entity.RoleEntity, error)
	UpdateRole(request request.RoleUpdateServiceRequest, pathId int) (map[string]interface{}, error)
	DeleteData(roleId int) (entity.RoleEntity, error)
}
