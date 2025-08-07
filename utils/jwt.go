package utils

import (
	"errors"
	"os"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(getJWTSecret())

func getJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "defaultsecret" // fallback saat dev
	}
	return secret
}

type JWTClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

var (
	validTokens   = make(map[string]time.Time)
	validTokensMu sync.RWMutex
)

func GenerateJWT(userID uint) (string, error) {
	exp := time.Now().Add(24 * time.Hour)

	claims := JWTClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	// Masukkan ke whitelist
	AddValidToken(signedToken, exp)
	return signedToken, nil
}

func VerifyJWT(tokenStr string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if !IsTokenValid(tokenStr) {
		return nil, errors.New("token has been invalidated")
	}

	return claims, nil
}

func AddValidToken(token string, exp time.Time) {
	validTokensMu.Lock()
	defer validTokensMu.Unlock()
	validTokens[token] = exp
}

func IsTokenValid(token string) bool {
	validTokensMu.RLock()
	defer validTokensMu.RUnlock()

	exp, exists := validTokens[token]
	if !exists {
		return false
	}

	if time.Now().After(exp) {
		// otomatis hapus token expired
		go RemoveToken(token)
		return false
	}

	return true
}

func RemoveToken(token string) {
	validTokensMu.Lock()
	defer validTokensMu.Unlock()
	delete(validTokens, token)
}
