package repository

import (
	"fmt"

	"github.com/ijaybaihaqi/heli-api/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Add(user model.User) error
	FetchByUsername(username string) (*model.User, error)
	CheckAvail(user model.User) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) Add(user model.User) error {
	res := u.db.Create(&user)
	err := res.Error
	if err != nil {
		return err
	}
	return nil
}

func (s *userRepository) FetchByUsername(username string) (*model.User, error) {
	user := model.User{}
	response := s.db.Model(&model.User{}).Where("username = ?", username).Scan(&user)
	if err := response.Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) CheckAvail(user model.User) (model.User, error) {
	result := model.User{}
	u.db.Model(&model.User{}).Where("username = ?", user.Username).Find(&result)
	if result.Username == "" {
		return model.User{}, fmt.Errorf("username not found")
	}
	return result, nil
}
