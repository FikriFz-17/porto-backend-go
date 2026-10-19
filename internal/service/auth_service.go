package service

import (
	"errors"
	"os"
	"porto-backend-go/internal/model/dto"
	"porto-backend-go/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(req dto.LoginRequest) (*dto.LoginResponse, error)
}

type authService struct{
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (s *authService) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.userRepo.FindUsername(req.Username)
	if err != nil {
		return nil, errors.New("username or password not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("username or password not found")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id" : user.UUID,
		"username": user.Username,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	
	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return nil, errors.New("failed to generate token")
	}
	
	return &dto.LoginResponse{
		Token: tokenString,
	}, nil
	
}