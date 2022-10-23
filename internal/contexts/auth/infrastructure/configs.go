package infrastructure

import (
	"github.com/hoffme/ddd-backend/internal/contexts/auth/infrastructure/jwt"
	"github.com/hoffme/ddd-backend/internal/contexts/auth/infrastructure/mysql"
	"github.com/hoffme/ddd-backend/internal/contexts/auth/infrastructure/password"
)

type Config struct {
	JWT            jwt.Config       `json:"jwt"`
	Password       password.Config  `json:"password"`
	UserRepository mysql.UserConfig `json:"user_repository"`
}
