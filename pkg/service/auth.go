package service

import (
	"Resrik"
	"Resrik/pkg/repository"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const salt = "wudgqugdgu121ddhishd"
const sighingKey = "ugsuigdasbwtdqwvsjh"
const TokenTTL = 12 * time.Hour

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreateUser(user Resrik.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{jwt.StandardClaims{
		ExpiresAt: time.Now().Add(TokenTTL).Unix(),
		IssuedAt:  time.Now().Unix(),
	},
		user.Id,
	})
	return token.SignedString([]byte(sighingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid method")
		}
		return []byte(sighingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("tokenClaims are not of type tokenClaims")
	}

	return claims.UserId, nil
}
