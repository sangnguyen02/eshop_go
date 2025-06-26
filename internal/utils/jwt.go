package utils

import (
	"encoding/base64"
	"errors"
	"go_ecommerce/pkg/setting"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserClaims struct {
	Username string
	Role     string
	ID       string
}

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

func ExtractUserFromClaims(claims *jwt.MapClaims) (*UserClaims, error) {
	getStr := func(key string) (string, error) {
		v, ok := (*claims)[key].(string)
		if !ok {
			return "", errors.New("missing or invalid " + key)
		}
		decoded, err := base64.StdEncoding.DecodeString(v)
		if err != nil {
			return "", errors.New("failed to decode " + key)
		}
		return string(decoded), nil
	}

	username, err := getStr("username")
	if err != nil {
		return nil, err
	}
	role, err := getStr("role")
	if err != nil {
		return nil, err
	}
	id, err := getStr("id")
	if err != nil {
		return nil, err
	}

	return &UserClaims{
		Username: username,
		Role:     role,
		ID:       id,
	}, nil
}
