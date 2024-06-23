package service

import (
	"event-service/internal/domain/model"
	"event-service/internal/infra/repository"
)

type QuantitiesService struct {
	repo *repository.QuantitiesRepository
}

func NewQuantitiesService(repo *repository.QuantitiesRepository) *QuantitiesService {
	return &QuantitiesService{repo: repo}
}

func (s *QuantitiesService) CreateQuantities(quantity *model.Quantities) error {
	return s.repo.Create(quantity)
}

func (s *QuantitiesService) GetQuantitiesByID(quantityID int) (*model.Quantities, error) {
	return s.repo.GetByID(quantityID)
}

func (s *QuantitiesService) GetAllQuantities() ([]*model.Quantities, error) {
	return s.repo.GetAll()
}

func (s *QuantitiesService) UpdateQuantities(quantity *model.Quantities) error {
	return s.repo.Update(quantity)
}

func (s *QuantitiesService) DeleteQuantities(quantityID int) error {
	return s.repo.Delete(quantityID)
}
