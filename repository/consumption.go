package repository

import (
	"github.com/ijaybaihaqi/heli-api/model"
	"gorm.io/gorm"
)

type ConsumptionRepository interface {
	FetchAll(userID uint) ([]model.ConsumptionResponse, error)
	Store(s *model.Consumption) error
	Reset(userID uint) error
}

type consumptionRepoImpl struct {
	db *gorm.DB
}

func NewConsumptionRepo(db *gorm.DB) *consumptionRepoImpl {
	return &consumptionRepoImpl{db}
}

func (s *consumptionRepoImpl) FetchAll(userID uint) ([]model.ConsumptionResponse, error) {
	consumptions := []model.ConsumptionResponse{}
	res := s.db.Model(&model.Consumption{}).Select("consumptions.id AS id, consumptions.consumed_at, appliances.name AS appliance, user_appliances.room, consumptions.amount, consumptions.status").Joins("left join user_appliances on user_appliances.id = consumptions.user_appliance_id").Joins("left join appliances on appliances.id = user_appliances.appliance_id").Where("user_appliances.user_id = ?", userID).Order("consumptions.consumed_at desc").Scan(&consumptions)
	if res.Error != nil {
		return nil, res.Error
	}
	return consumptions, nil
}

func (s *consumptionRepoImpl) Store(consumption *model.Consumption) error {
	res := s.db.Create(&consumption)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func (s *consumptionRepoImpl) Reset(userID uint) error {
	consumptions := []model.Consumption{}
	res := s.db.Model(&model.Consumption{}).Select("consumptions.id").Joins("left join user_appliances on user_appliances.id = consumptions.user_appliance_id").Where("user_appliances.user_id = ?", userID).Scan(&consumptions)
	if res.Error != nil {
		return res.Error
	}

	res = s.db.Delete(&consumptions)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}
