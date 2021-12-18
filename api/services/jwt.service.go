package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type JwtService struct {
	secretKey string
	issuer    string
}
type jwtCustomClaims struct {
	InstituteID string `json:"institute_id"`
	jwt.StandardClaims
}

func NewJWtService() JwtService {
	return JwtService{
		issuer:    "vnoticeboard-service",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "ydmfgfgh"
	}
	return secretKey
}
func (jwtService JwtService) GenerateToken(InstituteID string) string {
	claims := &jwtCustomClaims{
		InstituteID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    jwtService.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(jwtService.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (jwtService JwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(jwtService.secretKey), nil
	})

}
