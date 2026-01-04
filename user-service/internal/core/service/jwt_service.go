package service

import (
	"time"
	"user-service/config"

	"github.com/golang-jwt/jwt/v5"
)

type JwtServiceInterface interface {
	GenerateToken(userID int64) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type JWTService struct {
	secretKey string
	issuer    string
}

// GenerateToken implements [JWTServiceInterface].
func (j *JWTService) GenerateToken(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"iss":     j.issuer,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

// ValidateToken implements [JWTServiceInterface].
func (j *JWTService) ValidateToken(encodeToken string) (*jwt.Token, error) {
	return jwt.Parse(encodeToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(j.secretKey), nil
	})
}

func NewJwtService(cfg *config.Config) JwtServiceInterface {
	return &JWTService{
		secretKey: cfg.App.JwtSecretKey,
		issuer:    cfg.App.JwtIssuer,
	}
}
