package repository

import (
	"github.com/ijaybaihaqi/heli-api/model"
	"gorm.io/gorm"
)

type UserApplianceRepository interface {
	FetchAll(userID uint) ([]model.UserApplianceResponse, error)
	FetchByID(id int) (*model.UserApplianceResponse, error)
	Store(s *model.UserAppliance) error
	Update(id string, s *model.UserAppliance) error
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
	res := s.db.Model(&model.UserAppliance{}).Select("user_appliances.id, appliances.name AS appliance, user_appliances.room").Joins("join appliances on appliances.id = user_appliances.appliance_id").Where("user_appliances.user_id = ?", userID).Scan(&userAppliances)
	if res.Error != nil {
		return nil, res.Error
	}
	return userAppliances, nil
}

func (s *userApplianceRepoImpl) Store(userAppliance *model.UserAppliance) error {
	res := s.db.Create(&userAppliance)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func (s *userApplianceRepoImpl) Update(id string, userAppliance *model.UserAppliance) error {
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
	response := s.db.Model(&model.UserAppliance{}).Select("user_appliances.id, appliances.name AS appliance, user_appliances.room").Joins("left join appliances on appliances.id = user_appliances.appliance_id").Find(&model.UserAppliance{}, id).Scan(&userAppliance)
	if err := response.Error; err != nil {
		return nil, err
	}
	return &userAppliance, nil
}
