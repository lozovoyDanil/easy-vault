package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"main.go/internal/model"
	"main.go/internal/repository"
)

const (
	salt       = "euv45675kdfjd458dhg43"
	signingKey = "dfjh47ty34hfd89wofdhf"
	tokenTTL   = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.User) (int, error) {
	user.Password = s.generatePassHash(user.Password)

	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, pass string) (string, error) {
	user, err := s.repo.GetUser(username, s.generatePassHash(pass))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		user.Role,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(token string) (model.UserIdentity, error) {
	var u model.UserIdentity

	userToken, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid jwt signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return u, err
	}

	claims, ok := userToken.Claims.(*tokenClaims)
	if !ok {
		return u, errors.New("wrong claims type")
	}

	u.Id = claims.UserId
	u.Role = claims.Role

	return u, nil
}

func (s *AuthService) generatePassHash(pass string) string {
	hash := sha1.New()
	hash.Write([]byte(pass))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) UserInfo(userId int) (model.User, error) {
	return s.repo.UserInfo(userId)
}

func (s *AuthService) EditUser(userId int, input model.UpdateUserInput) error {
	hashedPass := s.generatePassHash(*input.Password)
	input.Password = &hashedPass

	return s.repo.EditUser(userId, input)
}
