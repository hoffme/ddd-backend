package jwt

import (
	"github.com/hoffme/ddd-backend/internal/contexts/auth/domain"
)

var _ domain.JWTCrypt = (*Adapter)(nil)

type Config struct {
	RefreshTokenSecret string `json:"refresh_token_secret"`
	AccessTokenSecret  string `json:"access_token_secret"`
}

type Adapter struct {
	config Config
}

func New(config Config) *Adapter {
	return &Adapter{config: config}
}
