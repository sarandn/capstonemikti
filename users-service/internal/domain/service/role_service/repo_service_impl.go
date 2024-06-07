package roleservice

import (
	"users-service/internal/domain/model"
	"users-service/internal/domain/model/entity"
	rolerepository "users-service/internal/infra/role_repository"
	"users-service/internal/interfaces/api/request"
)

type ResponseToJson map[string]interface{}

type RoleServiceImpl struct {
	repository rolerepository.RoleRepository
}

func NewRoleService(repository rolerepository.RoleRepository) *RoleServiceImpl {
	return &RoleServiceImpl{
		repository: repository,
	}
}

func (service *RoleServiceImpl) SaveRole(request request.RoleServiceRequest) (map[string]interface{}, error) {
	RoleRer := model.Role{
		RoleID:   request.RoleID,
		RoleName: request.RoleName,
	}

	saverole, errSaveRole := service.repository.SaveRole(RoleRer)

	if errSaveRole != nil {
		return nil, errSaveRole
	}

	return ResponseToJson{"role_id": saverole.RoleID, "role_name": saverole.RoleName}, nil
}

func (service *RoleServiceImpl) GetRole(roleId int) (entity.RoleEntity, error) {
	getRoleById, errGetRole := service.repository.GetRole(roleId)

	if errGetRole != nil {
		return entity.RoleEntity{}, errGetRole
	}

	return entity.ToRoleEntity(getRoleById), nil
}

func (service *RoleServiceImpl) GetRoles() ([]entity.RoleEntity, error) {
	getRole, err := service.repository.GetAllRole()

	if err != nil {
		return nil, err
	}

	return entity.ToRoleListEntity(getRole), nil
}
