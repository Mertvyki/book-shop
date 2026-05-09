package core_security

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type JWTConfig struct {
	Secret             string        `envconfig:"SECRET" required:"true"`
	AccessTokenExpiry  time.Duration `envconfig:"ACCESS_TOKEN_EXPIRY" default:"15m"`
	RefreshTokenExpiry time.Duration `envconfig:"REFRESH_TOKEN_EXPIRY" default:"168h"`
	Issuer             string        `envconfig:"ISSUER" default:"bookshop"`
}

func NewJWTConfig() (JWTConfig, error) {
	var cfg JWTConfig
	if err := envconfig.Process("JWT", &cfg); err != nil {
		return JWTConfig{}, fmt.Errorf("process envconfig for JWT: %w", err)
	}

	return cfg, nil
}

func NewJWTConfigMust() JWTConfig {
	cfg, err := NewJWTConfig()
	if err != nil {
		panic(fmt.Errorf("get JWT config: %w", err))
	}
	return cfg
}
