package service

import (
	"github.com/ijaybaihaqi/heli-api/model"
	"github.com/ijaybaihaqi/heli-api/repository"
)

type ConsumptionService interface {
	FetchAll(userID uint) ([]model.ConsumptionResponse, error)
	Store(s *model.Consumption) (model.ConsumptionResponse, error)
	Reset(userID uint) error
}

type consumptionService struct {
	consumptionRepository repository.ConsumptionRepository
}

func NewConsumptionService(consumptionRepository repository.ConsumptionRepository) ConsumptionService {
	return &consumptionService{consumptionRepository}
}

func (s *consumptionService) FetchAll(userID uint) ([]model.ConsumptionResponse, error) {
	consumptions, err := s.consumptionRepository.FetchAll(userID)
	if err != nil {
		return nil, err
	}

	return consumptions, nil
}

func (s *consumptionService) Store(consumption *model.Consumption) (model.ConsumptionResponse, error) {
	err := s.consumptionRepository.Store(consumption)
	if err != nil {
		return model.ConsumptionResponse{}, err
	}

	result, err := s.consumptionRepository.FetchByID(consumption.ID)
	if err != nil {
		return model.ConsumptionResponse{}, err
	}

	return *result, nil
}

func (s *consumptionService) Reset(userID uint) error {
	err := s.consumptionRepository.Reset(userID)
	if err != nil {
		return err
	}

	return nil
}
