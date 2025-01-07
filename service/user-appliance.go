package service

import (
	"github.com/ijaybaihaqi/heli-api/model"
	"github.com/ijaybaihaqi/heli-api/repository"
)

type UserApplianceService interface {
	FetchAll(userID uint) (map[string][]model.UserApplianceResponse, error)
	FetchUserRooms(userID uint) ([]string, error)
	FetchByID(id int) (*model.UserApplianceResponse, error)
	Store(s *model.UserAppliance) error
	Update(id uint, s *model.UserAppliance) error
	Delete(id string) error
}

type userApplianceService struct {
	userApplianceRepository repository.UserApplianceRepository
}

func NewUserApplianceService(userApplianceRepository repository.UserApplianceRepository) UserApplianceService {
	return &userApplianceService{userApplianceRepository}
}

func (s *userApplianceService) FetchAll(userID uint) (map[string][]model.UserApplianceResponse, error) {
	userAppliances, err := s.userApplianceRepository.FetchAll(userID)
	if err != nil {
		return nil, err
	}

	userAppliancesByRoom := map[string][]model.UserApplianceResponse{}

	for _, userAppliance := range userAppliances {
		status, err := s.userApplianceRepository.FetchCurrentStatus(userAppliance.ID)
		if err != nil {
			return nil, err
		}
		if status.Status != "" {
			userAppliance.CurrentStatus = status.Status
		} else {
			userAppliance.CurrentStatus = "off"
		}
		if len(userAppliancesByRoom[userAppliance.Room]) == 0 {
			userAppliancesByRoom[userAppliance.Room] = []model.UserApplianceResponse{}
		}
		userAppliancesByRoom[userAppliance.Room] = append(userAppliancesByRoom[userAppliance.Room], userAppliance)
	}

	return userAppliancesByRoom, nil
}

func (s *userApplianceService) FetchUserRooms(userID uint) ([]string, error) {
	userApplianceRooms, err := s.userApplianceRepository.FetchUserRooms(userID)
	if err != nil {
		return nil, err
	}

	rooms := []string{}
	for _, userApplianceRoom := range userApplianceRooms {
		rooms = append(rooms, userApplianceRoom.Room)
	}

	return rooms, nil
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

func (s *userApplianceService) Update(id uint, userAppliance *model.UserAppliance) error {
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
