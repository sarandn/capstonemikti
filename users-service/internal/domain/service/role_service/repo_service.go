package roleservice

import (
	"users-service/internal/domain/model/entity"
	"users-service/internal/interfaces/api/request"
)

type RoleService interface {
	SaveRole(request request.RoleServiceRequest) (map[string]interface{}, error)
	GetRole(roleId int) (entity.RoleEntity, error)
	GetRoles() ([]entity.RoleEntity, error)
}
