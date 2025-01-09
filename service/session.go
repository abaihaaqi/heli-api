package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
	GenerateJWT(user model.User, expiresAt time.Time) (string, error)
	ValidateJWT(tokenString string) (jwt.MapClaims, error)
}

type sessionService struct {
	sessionRepository repository.SessionsRepository
	jwtSecret         []byte
}

func NewSessionService(sessionRepository repository.SessionsRepository, jwtSecret []byte) SessionService {
	return &sessionService{sessionRepository, jwtSecret}
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

func (s *sessionService) GenerateJWT(user model.User, expiresAt time.Time) (string, error) {
	claims := jwt.MapClaims{
		"username": user.Username,
		"name":     user.Name,
		"exp":      expiresAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecret)
}

func (s *sessionService) ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token's signing method is correct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
