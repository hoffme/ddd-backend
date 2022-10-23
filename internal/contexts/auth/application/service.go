package application

import (
	"time"

	busDomain "github.com/hoffme/ddd-backend/internal/shared/bus/domain"

	"github.com/hoffme/ddd-backend/internal/contexts/auth/domain"
)

type Ports struct {
	EventBusEmitter busDomain.EventEmitter
	UserRepository  domain.UserRepository
	Password        domain.PasswordCrypt
	JWT             domain.JWTCrypt
}

type Configs struct {
	RefreshTokenExpiration time.Duration `json:"refresh_token_expiration"`
	AccessTokenExpiration  time.Duration `json:"access_token_expiration"`
}

type Service struct {
	configs Configs
	ports   Ports
}

func New(configs Configs, ports Ports) *Service {
	return &Service{
		configs: configs,
		ports:   ports,
	}
}
