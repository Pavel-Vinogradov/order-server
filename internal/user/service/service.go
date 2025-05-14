package service

import (
	"database/sql"
	"errors"
	"fmt"
	"order-server/internal/user/entity"
	"order-server/internal/user/repository"
)

var ErrUserNotFound = errors.New("user not found")

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(user entity.User) (entity.User, error) {
	createdUser, err := s.repo.Create(user)
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to create user: %w", err)
	}
	return createdUser, nil
}

func (s *UserService) FindByPhone(phone string) (entity.User, error) {
	user, err := s.repo.FindByPhone(phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, ErrUserNotFound
		}
		return entity.User{}, fmt.Errorf("failed to find user by phone: %w", err)
	}
	return user, nil
}
