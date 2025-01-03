package repository

import (
	"fmt"

	"github.com/ijaybaihaqi/heli-api/model"
	"gorm.io/gorm"
)

type SessionsRepository interface {
	AddSession(session model.Session) error
	DeleteSession(token string) error
	UpdateSession(session model.Session) error
	SessionAvailName(name string) error
	SessionAvailToken(token string) (model.Session, error)
}

type sessionsRepoImpl struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) *sessionsRepoImpl {
	return &sessionsRepoImpl{db}
}

func (s *sessionsRepoImpl) AddSession(session model.Session) error {
	res := s.db.Create(&session)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func (s *sessionsRepoImpl) DeleteSession(token string) error {
	res := s.db.Where("token = ?", token).Delete(&model.Session{})
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func (s *sessionsRepoImpl) UpdateSession(session model.Session) error {
	res := s.db.Model(&model.Session{}).Where("username = ?", session.Username).Updates(session)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func (s *sessionsRepoImpl) SessionAvailName(name string) error {
	result := model.Session{}
	response := s.db.Model(&model.Session{}).Where("username = ?", name).Find(&result)
	if err := response.Error; err != nil {
		return err
	}
	if result.Username == "" {
		return fmt.Errorf("user not found")
	}
	return nil
}

func (s *sessionsRepoImpl) SessionAvailToken(token string) (model.Session, error) {
	result := model.Session{}
	response := s.db.Model(&model.Session{}).Where("token = ?", token).Find(&result)
	if err := response.Error; err != nil {
		return model.Session{}, err
	}
	if result.Token == "" {
		return model.Session{}, fmt.Errorf("token not found")
	}
	return result, nil
}
