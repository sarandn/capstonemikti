package service

import (
	"event-service/internal/domain/model"
	"event-service/internal/infra/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(category *model.Category) error {
	return s.repo.Create(category)
}

func (s *CategoryService) GetCategoryByID(categoryID int) (*model.Category, error) {
	return s.repo.GetByID(categoryID)
}

func (s *CategoryService) GetAllCategories() ([]*model.Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) UpdateCategory(category *model.Category) error {
	return s.repo.Update(category)
}

func (s *CategoryService) DeleteCategory(categoryID int) error {
	return s.repo.Delete(categoryID)
}
