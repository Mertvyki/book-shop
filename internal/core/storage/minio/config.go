package core_minio

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Endpoint      string `envconfig:"ENDPOINT" required:"true"`
	AccessKey     string `envconfig:"ACCESS_KEY" required:"true"`
	SecretKey     string `envconfig:"SECRET_KEY" required:"true"`
	Bucket        string `envconfig:"BUCKET" required:"true"`
	UseSSL        bool   `envconfig:"USE_SSL" default:"false"`
	PublicBaseURL string `envconfig:"PUBLIC_BASE_URL"`
}

func NewConfig() (Config, error) {
	var cfg Config
	if err := envconfig.Process("MINIO", &cfg); err != nil {
		return Config{}, fmt.Errorf("process minio envconfig: %w", err)
	}
	return cfg, nil
}

func NewConfigMust() Config {
	cfg, err := NewConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}
