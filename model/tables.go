package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);unique;"`
	Password string `json:"password"`
	Name     string `json:"name"`

	Session        Session         `gorm:"foreignKey:Username;references:Username;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserAppliances []UserAppliance `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Session struct {
	gorm.Model
	Token  string    `json:"token"`
	Expiry time.Time `json:"expiry"`

	Username string `json:"username"`
}

type Appliance struct {
	gorm.Model
	Name  string `json:"appliance_name"`
	Image string `json:"image"`

	UserAppliances []UserAppliance `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UserAppliance struct {
	gorm.Model
	Room string `json:"room"`

	UserID      uint `json:"user_id"`
	ApplianceID uint `json:"appliance_id"`

	Consumptions []Consumption `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Consumption struct {
	gorm.Model
	ConsumedAt time.Time `json:"consumed_at"`
	Amount     float64   `json:"amount"`
	Status     string    `json:"status"`

	UserApplianceID uint `json:"user_appliance_id"`
}
