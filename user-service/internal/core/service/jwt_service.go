package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/terenjit/ecommerce-veggie/user-service/config"
)

type JwtServiceInterface interface {
	GenerateToken(userId int64) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type JwtService struct {
	secretKey string
	issuer    string
}

func NewJwtService(cfg *config.Config) JwtServiceInterface {
	return &JwtService{
		secretKey: cfg.App.JWTSecretKey,
		issuer:    cfg.App.JWTIssuer,
	}
}

func (s *JwtService) GenerateToken(userId int64) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"iss":     s.issuer,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	return claims.SignedString([]byte(s.secretKey))
}

func (s *JwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(s.secretKey), nil
	})
}
