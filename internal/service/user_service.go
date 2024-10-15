// internal/service/user_service.go
package service

import (
	"go_project/internal/model"
	"go_project/internal/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Register(user *model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return s.userRepo.Create(user)
}

func (s *UserService) Login(username, password string) (string, error) {
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return "查無此用戶", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "密碼錯誤", err
	}

	// 這裡應該生成一個 JWT token
	// 為了簡化，這裡返回一個假的 token
	return "fake-jwt-token", nil
}
