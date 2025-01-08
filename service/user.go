package service

import (
	"fmt"

	"github.com/ijaybaihaqi/heli-api/model"
	"github.com/ijaybaihaqi/heli-api/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetCurrentUsername() string
	SetCurrentUsername(username string)
	FetchByUsername(username string) (*model.User, error)
	Login(user model.User) (model.User, error)
	Register(user model.User) error
	CheckPassLength(pass string) bool
	CheckPassAlphabet(pass string) bool
	HashPassword(password string) (string, error)
}

type userService struct {
	currentUsername string
	userRepository  repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{"", userRepository}
}

func (s *userService) GetCurrentUsername() string {
	return s.currentUsername
}

func (s *userService) SetCurrentUsername(username string) {
	s.currentUsername = username
}

func (s *userService) FetchByUsername(username string) (*model.User, error) {
	user, err := s.userRepository.FetchByUsername(username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Login(reqUser model.User) (model.User, error) {
	user, err := s.userRepository.CheckAvail(reqUser)
	if err != nil {
		return model.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqUser.Password))
	if err != nil {
		return model.User{}, fmt.Errorf("password not match")
	}

	return user, nil
}

func (s *userService) Register(user model.User) error {
	err := s.userRepository.Add(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) CheckPassLength(pass string) bool {
	if len(pass) <= 5 {
		return true
	} else {
		return false
	}

}

func (s *userService) CheckPassAlphabet(pass string) bool {
	for _, charVariable := range pass {
		if (charVariable < 'a' || charVariable > 'z') && (charVariable < 'A' || charVariable > 'Z') {
			return false
		}
	}
	return true
}

func (s *userService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
