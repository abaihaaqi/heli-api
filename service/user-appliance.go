package service

import (
	"github.com/ijaybaihaqi/heli-api/model"
	"github.com/ijaybaihaqi/heli-api/repository"
)

type UserApplianceService interface {
	FetchAll(userID uint) ([]model.UserApplianceResponse, error)
	FetchByID(id int) (*model.UserApplianceResponse, error)
	Store(s *model.UserAppliance) error
	Update(id string, s *model.UserAppliance) error
	Delete(id string) error
}

type userApplianceService struct {
	userApplianceRepository repository.UserApplianceRepository
}

func NewUserApplianceService(userApplianceRepository repository.UserApplianceRepository) UserApplianceService {
	return &userApplianceService{userApplianceRepository}
}

func (s *userApplianceService) FetchAll(userID uint) ([]model.UserApplianceResponse, error) {
	userAppliances, err := s.userApplianceRepository.FetchAll(userID)
	if err != nil {
		return nil, err
	}

	return userAppliances, nil
}

func (s *userApplianceService) FetchByID(id int) (*model.UserApplianceResponse, error) {
	userAppliance, err := s.userApplianceRepository.FetchByID(id)
	if err != nil {
		return nil, err
	}

	return userAppliance, nil
}

func (s *userApplianceService) Store(userAppliance *model.UserAppliance) error {
	err := s.userApplianceRepository.Store(userAppliance)
	if err != nil {
		return err
	}

	return nil
}

func (s *userApplianceService) Update(id string, userAppliance *model.UserAppliance) error {
	err := s.userApplianceRepository.Update(id, userAppliance)
	if err != nil {
		return err
	}

	return nil
}

func (s *userApplianceService) Delete(id string) error {
	err := s.userApplianceRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
