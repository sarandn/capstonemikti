package rolerepository

import (
	"errors"
	"fmt"
	"users-service/internal/domain/model"

	"gorm.io/gorm"
)

type RoleRepositoryImpl struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepositoryImpl {
	return &RoleRepositoryImpl{db: db}
}

func (repo *RoleRepositoryImpl) SaveRole(role model.Role) (model.Role, error) {
	err := repo.db.Create(&role).Error

	if err != nil {
		return model.Role{}, err
	}

	return role, nil
}

func (repo *RoleRepositoryImpl) GetRole(Id int) (model.Role, error) {
	var roleData model.Role

	err := repo.db.Preload("User").First(&roleData, "role_id = ?", Id).Error

	if err != nil {
		return model.Role{}, errors.New("role tidak ditemukan")
	}

	return roleData, nil
}

func (repo *RoleRepositoryImpl) GetAllRole() ([]model.Role, error) {
	var roles []model.Role

	err := repo.db.Preload("User").Find(&roles).Error

	if err != nil {
		return []model.Role{}, err
	}

	fmt.Println(roles[1].User == nil)
	return roles, nil
}

func (repo *RoleRepositoryImpl) UpdateRole(role model.Role) (model.Role, error) {
	err := repo.db.Model(model.Role{}).Where("role_id = ?", role.RoleID).Updates(role).Error

	if err != nil {
		return role, err
	}

	return role, nil
}

func (repo *RoleRepositoryImpl) DeleteRole(Id int) (model.Role, error) {
	var roleData model.Role

	err := repo.db.Where("role_id = ?", Id).Delete(&roleData).Error

	if err != nil {
		return model.Role{}, err
	}

	return roleData, nil
}
