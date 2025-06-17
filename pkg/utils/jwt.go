package utils

import (
	"encoding/base64"
	"go_ecommerce/pkg/setting"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(username string, role string, id string) (string, error) {
	expireTime, err := time.ParseDuration("8h")

	if err != nil {
		panic("Invalid time duration. Should be time.ParseDuration string")
	}

	claims := &jwt.MapClaims{
		"username": base64.StdEncoding.EncodeToString([]byte(username)),
		"role":     base64.StdEncoding.EncodeToString([]byte(role)),
		"exp":      time.Now().Add(expireTime).Unix(),
		"id":       base64.StdEncoding.EncodeToString([]byte(id)),
	}

	// Generate encoded token and send it as response.
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString([]byte(setting.AppSetting.JwtSecret))

	return token, err
}

// Verify verifies the jwt token against the secret
func VerifyToken(token string) (*jwt.MapClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(setting.AppSetting.JwtSecret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*jwt.MapClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
