package service

import (
	"fmt"
	"time"

	"github.com/ijaybaihaqi/heli-api/model"
	"github.com/ijaybaihaqi/heli-api/repository"
)

type SessionService interface {
	AddSession(session model.Session) error
	UpdateSession(session model.Session) error
	DeleteSession(sessionToken string) error
	SessionAvailName(username string) error
	TokenExpired(session model.Session) bool
	TokenValidity(token string) (model.Session, error)
}

type sessionService struct {
	sessionRepository repository.SessionsRepository
}

func NewSessionService(sessionRepository repository.SessionsRepository) SessionService {
	return &sessionService{sessionRepository}
}

func (s *sessionService) SessionAvailName(username string) error {
	return s.sessionRepository.SessionAvailName(username)
}

func (s *sessionService) AddSession(session model.Session) error {
	return s.sessionRepository.AddSession(session)
}

func (s *sessionService) UpdateSession(session model.Session) error {
	return s.sessionRepository.UpdateSession(session)
}

func (s *sessionService) DeleteSession(sessionToken string) error {
	return s.sessionRepository.DeleteSession(sessionToken)
}

func (s *sessionService) TokenValidity(token string) (model.Session, error) {
	session, err := s.sessionRepository.SessionAvailToken(token)
	if err != nil {
		return model.Session{}, err
	}

	if s.TokenExpired(session) {
		err := s.sessionRepository.DeleteSession(token)
		if err != nil {
			return model.Session{}, err
		}
		return model.Session{}, fmt.Errorf("token is expired")
	}

	return session, nil
}

func (s *sessionService) TokenExpired(session model.Session) bool {
	return session.Expiry.Before(time.Now())
}
