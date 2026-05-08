package core_security

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTManager struct {
	secret []byte
	expiry time.Duration
	issuer string
}

func NewJWTManager(secret string, expiry time.Duration, issuer string) *JWTManager {
	return &JWTManager{
		secret: []byte(secret),
		expiry: expiry,
		issuer: issuer,
	}
}

func (m *JWTManager) GenerateAccessToken(userID int, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"iss":     m.issuer,
		"exp":     time.Now().Add(m.expiry).Unix(),
		"iat":     time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(m.secret)
}

func (m *JWTManager) Parse(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return m.secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if iss, ok := claims["iss"].(string); !ok || iss != m.issuer {
			return nil, fmt.Errorf("invalid issuer")
		}
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
