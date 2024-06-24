package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(apiKey string, expSeconds int64) (string, error) {
	// 分割 API key
	parts := strings.Split(apiKey, ".")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid api key format")
	}
	id, secret := parts[0], parts[1]

	// 构建 payload
	payload := jwt.MapClaims{
		"api_key":   id,
		"exp":       time.Now().Add(time.Duration(expSeconds) * time.Second).Unix(),
		"timestamp": time.Now().Unix(),
	}

	// 签名算法和头部
	signingMethod := jwt.SigningMethodHS256
	header := map[string]interface{}{
		"alg":       signingMethod.Alg(),
		"sign_type": "SIGN",
	}

	// 生成 JWT token
	token := jwt.NewWithClaims(signingMethod, payload)
	token.Header = header
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("令牌生成失败: %v", err)
	}

	return signedToken, nil
}
