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

func (service *RoleServiceImpl) UpdateRole(request request.RoleUpdateServiceRequest, pathId int) (map[string]interface{}, error) {
	getRoleById, err := service.repository.GetRole(pathId)

	if err != nil {
		return nil, err
	}

	// validasi jika datanya kosong atau tidak
	if request.RoleName == "" {
		request.RoleName = getRoleById.RoleName
	}

	roleReq := model.Role{
		RoleID:   pathId,
		RoleName: request.RoleName,
	}

	updateRole, errUpdateRole := service.repository.UpdateRole(roleReq)

	if errUpdateRole != nil {
		return nil, errUpdateRole
	}

	return ResponseToJson{"role_name": updateRole.RoleName}, nil
}

func (service *RoleServiceImpl) DeleteData(roleId int) (entity.RoleEntity, error){
	delRole, errDelRole := service.repository.DeleteRole(roleId)

	if errDelRole != nil{
		return entity.RoleEntity{}, errDelRole
	}

	return entity.ToRoleEntity(delRole), nil
}


