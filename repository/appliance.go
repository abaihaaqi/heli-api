package repository

import (
	"github.com/ijaybaihaqi/heli-api/model"
	"gorm.io/gorm"
)

type ApplianceRepository interface {
	FetchAll() ([]model.ApplianceResponse, error)
}

type applianceRepoImpl struct {
	db *gorm.DB
}

func NewApplianceRepo(db *gorm.DB) *applianceRepoImpl {
	return &applianceRepoImpl{db}
}

func (s *applianceRepoImpl) FetchAll() ([]model.ApplianceResponse, error) {
	appliances := []model.ApplianceResponse{}
	res := s.db.Find(&[]model.Appliance{}).Scan(&appliances)
	if res.Error != nil {
		return nil, res.Error
	}
	return appliances, nil
}
