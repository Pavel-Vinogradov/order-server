package service

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"order-server/internal/auth/dto"
	entityAuth "order-server/internal/auth/entity"
	authRepository "order-server/internal/auth/repository"
	entityUser "order-server/internal/user/entity"
	userService "order-server/internal/user/service"
)

type AuthService struct {
	authRepo    *authRepository.AuthRepository
	userService *userService.UserService
}

func (s *AuthService) Auth(body *dto.AuthRequest) (*entityAuth.Auth, error) {
	user, err := s.userService.FindByPhone(body.Phone)
	if err != nil {
		if !errors.Is(err, userService.ErrUserNotFound) {
			return nil, fmt.Errorf("failed to find user: %w", err)
		}
		user, err = s.userService.Create(entityUser.User{Phone: body.Phone})
		if err != nil {
			return nil, fmt.Errorf("failed to create user: %w", err)
		}
	}

	sessionID, err := generateSessionID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate session id: %w", err)
	}

	_, err = s.authRepo.FindByUserId(user.Id)
	if err == nil {
		updated, err := s.authRepo.UpdateSession(user.Id, sessionID)
		if err != nil {
			return nil, fmt.Errorf("failed to update auth session: %w", err)
		}
		return updated, nil
	}

	if !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("failed to find auth session: %w", err)
	}

	newAuth := &entityAuth.Auth{
		UserId:  user.Id,
		Session: sessionID,
	}
	created, err := s.authRepo.Create(newAuth)
	if err != nil {
		return nil, fmt.Errorf("failed to create auth session: %w", err)
	}
	return created, nil
}

func (s *AuthService) GetUserBySession(sessionId string) (*entityAuth.Auth, error) {
	return s.authRepo.FindBySession(sessionId)
}

func NewAuthService(ar *authRepository.AuthRepository, us *userService.UserService) *AuthService {
	return &AuthService{
		authRepo:    ar,
		userService: us,
	}
}
func generateSessionID() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
