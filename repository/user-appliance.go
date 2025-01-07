package repository

import (
	"github.com/ijaybaihaqi/heli-api/model"
	"gorm.io/gorm"
)

type UserApplianceRepository interface {
	FetchAll(userID uint) ([]model.UserApplianceResponse, error)
	FetchCurrentStatus(userApplianceID uint) (model.CurrentStatusResponse, error)
	FetchByID(id int) (*model.UserApplianceResponse, error)
	FetchUserRooms(id uint) ([]model.UserApplianceRoomResponse, error)
	Store(s *model.UserAppliance) error
	Update(id uint, s *model.UserAppliance) error
	Delete(id string) error
}

type userApplianceRepoImpl struct {
	db *gorm.DB
}

func NewUserApplianceRepo(db *gorm.DB) *userApplianceRepoImpl {
	return &userApplianceRepoImpl{db}
}

func (s *userApplianceRepoImpl) FetchAll(userID uint) ([]model.UserApplianceResponse, error) {
	userAppliances := []model.UserApplianceResponse{}
	res := s.db.Model(&model.UserAppliance{}).Select("user_appliances.id, appliances.name AS appliance, appliances.image, appliances.energy, user_appliances.room").Joins("join appliances on appliances.id = user_appliances.appliance_id").Where("user_appliances.user_id = ?", userID).Scan(&userAppliances)
	if res.Error != nil {
		return nil, res.Error
	}
	return userAppliances, nil
}

func (s *userApplianceRepoImpl) FetchCurrentStatus(userApplianceID uint) (model.CurrentStatusResponse, error) {
	currentStatus := model.CurrentStatusResponse{}
	response := s.db.Model(&model.Consumption{}).Select("status").Where("user_appliance_id = ?", userApplianceID).Order("consumed_at desc").Limit(1).Scan(&currentStatus)
	if err := response.Error; err != nil {
		return model.CurrentStatusResponse{}, err
	}
	return currentStatus, nil
}

func (s *userApplianceRepoImpl) Store(userAppliance *model.UserAppliance) error {
	res := s.db.Create(&userAppliance)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func (s *userApplianceRepoImpl) Update(id uint, userAppliance *model.UserAppliance) error {
	res := s.db.Model(&model.UserAppliance{}).Where("id = ?", id).Updates(userAppliance)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func (s *userApplianceRepoImpl) Delete(id string) error {
	res := s.db.Where("id = ?", id).Delete(&model.UserAppliance{})
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func (s *userApplianceRepoImpl) FetchByID(id int) (*model.UserApplianceResponse, error) {
	userAppliance := model.UserApplianceResponse{}
	response := s.db.Model(&model.UserAppliance{}).Select("user_appliances.id, appliances.name AS appliance, appliances.image, appliances.energy, user_appliances.room").Joins("left join appliances on appliances.id = user_appliances.appliance_id").Find(&model.UserAppliance{}, id).Scan(&userAppliance)
	if err := response.Error; err != nil {
		return nil, err
	}
	return &userAppliance, nil
}

func (s *userApplianceRepoImpl) FetchUserRooms(id uint) ([]model.UserApplianceRoomResponse, error) {
	userApplianceRooms := []model.UserApplianceRoomResponse{}
	response := s.db.Distinct("room").Find(&model.UserAppliance{UserID: id}).Scan(&userApplianceRooms)
	if err := response.Error; err != nil {
		return nil, err
	}
	return userApplianceRooms, nil
}
