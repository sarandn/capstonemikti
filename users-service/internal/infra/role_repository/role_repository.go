package rolerepository

import "users-service/internal/domain/model"

type RoleRepository interface {
	SaveRole(role model.Role) (model.Role, error)
	GetRole(Id int) (model.Role, error)
	GetAllRole() ([]model.Role, error)
}
