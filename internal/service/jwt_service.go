package service

import (
	"context"
	"go_project/pkg/database"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService struct {
	secretKey []byte
}

func NewJWTService(secretKey string) *JWTService {
	return &JWTService{secretKey: []byte(secretKey)}
}

func (s *JWTService) GenerateToken(userID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 令牌有效期為24小時

	tokenString, err := token.SignedString(s.secretKey)
	if err != nil {
		return "", err
	}

	// 將令牌存儲在 Redis 中
	err = database.RedisClient.Set(context.Background(), tokenString, userID, time.Hour*24).Err()
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *JWTService) ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return s.secretKey, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", jwt.ErrSignatureInvalid
	}

	// 檢查令牌是否存在於 Redis 中
	userID, err := database.RedisClient.Get(context.Background(), tokenString).Result()
	if err != nil {
		return "", err
	}

	return userID, nil
}

func (s *JWTService) RevokeToken(tokenString string) error {
	// 從 Redis 中刪除令牌
	return database.RedisClient.Del(context.Background(), tokenString).Err()
}
