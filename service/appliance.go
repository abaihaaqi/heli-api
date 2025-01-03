package service

import (
	"github.com/ijaybaihaqi/heli-api/model"
	"github.com/ijaybaihaqi/heli-api/repository"
)

type ApplianceService interface {
	FetchAll() ([]model.ApplianceResponse, error)
}

type applianceService struct {
	applianceRepository repository.ApplianceRepository
}

func NewApplianceService(applianceRepository repository.ApplianceRepository) ApplianceService {
	return &applianceService{applianceRepository}
}

func (s *applianceService) FetchAll() ([]model.ApplianceResponse, error) {
	appliances, err := s.applianceRepository.FetchAll()
	if err != nil {
		return nil, err
	}

	return appliances, nil
}
