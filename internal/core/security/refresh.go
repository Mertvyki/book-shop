package core_security

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type RefreshTokenService struct{}

func NewRefreshTokenService() *RefreshTokenService {
	return &RefreshTokenService{}
}

func (s *RefreshTokenService) Generate() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func (s *RefreshTokenService) Hash(token string) (string, error) {
	sum := sha256.Sum256([]byte(token))
	return fmt.Sprintf("%x", sum), nil
}

func (s *RefreshTokenService) Compare(token, hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(token)) == nil
}
