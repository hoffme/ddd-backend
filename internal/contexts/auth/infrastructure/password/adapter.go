package password

import (
	"github.com/hoffme/ddd-backend/internal/contexts/auth/domain"
)

var _ domain.PasswordCrypt = (*Adapter)(nil)

type Config struct {
	Cost int `json:"cost"`
}

type Adapter struct {
	config Config
}

func New(config Config) *Adapter {
	return &Adapter{config: config}
}
