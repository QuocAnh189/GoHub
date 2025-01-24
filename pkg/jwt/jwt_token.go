package jwt

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"gohub/internal/libs/logger"

	"gohub/configs"
	"gohub/pkg/utils"
)

const (
	AccessTokenExpiredTime  = 2 * 24 * 3600 // 5 hours
	RefreshTokenExpiredTime = 30 * 24 * 3600
	AccessTokenType         = "x-access"  // 5 minutes
	RefreshTokenType        = "x-refresh" // 30 days
)

func GenerateAccessToken(payload map[string]interface{}) string {
	cfg := configs.GetConfig()
	payload["type"] = AccessTokenType
	tokenContent := jwt.MapClaims{
		"payload": payload,
		"exp":     time.Now().Add(time.Second * AccessTokenExpiredTime).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte(cfg.AuthSecret))
	if err != nil {
		logger.Error("Failed to generate access token: ", err)
		return ""
	}

	return token
}

func GenerateRefreshToken(payload map[string]interface{}) string {
	cfg := configs.GetConfig()
	payload["type"] = RefreshTokenType
	tokenContent := jwt.MapClaims{
		"payload": payload,
		"exp":     time.Now().Add(time.Second * RefreshTokenExpiredTime).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte(cfg.AuthSecret))
	if err != nil {
		logger.Error("Failed to generate refresh token: ", err)
		return ""
	}

	return token
}

func ValidateToken(jwtToken string) (map[string]interface{}, error) {
	cfg := configs.GetConfig()
	cleanJWT := strings.Replace(jwtToken, "Bearer ", "", -1)

	if IsTokenBlacklisted(cleanJWT) {
		return nil, jwt.ErrInvalidKey
	}

	tokenData := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(cleanJWT, tokenData, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.AuthSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrInvalidKey
	}

	var data map[string]interface{}
	utils.MapStruct(&data, tokenData["payload"])

	return data, nil
}

var tokenBlacklist = make(map[string]time.Time)

func BlacklistToken(token string) {
	tokenBlacklist[token] = time.Now()
}

func IsTokenBlacklisted(token string) bool {
	_, exists := tokenBlacklist[token]
	if !exists {
		return false
	}

	return true
}
