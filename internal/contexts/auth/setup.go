package auth

import (
	"github.com/hoffme/ddd-backend/internal/shared/bus"

	"github.com/hoffme/ddd-backend/internal/contexts/auth/domain"

	mysqlShared "github.com/hoffme/ddd-backend/internal/contexts/shared/infrastructure/mysql"

	"github.com/hoffme/ddd-backend/internal/contexts/auth/application"
	"github.com/hoffme/ddd-backend/internal/contexts/auth/infrastructure"
	"github.com/hoffme/ddd-backend/internal/contexts/auth/infrastructure/jwt"
	"github.com/hoffme/ddd-backend/internal/contexts/auth/infrastructure/mysql"
	"github.com/hoffme/ddd-backend/internal/contexts/auth/infrastructure/password"
)

type Config struct {
	Service        application.Configs   `json:"service"`
	Infrastructure infrastructure.Config `json:"infrastructure"`
}

type Adapters struct {
	MysqlConnection *mysqlShared.Connection
}

func Setup(config Config, buses *bus.Buses, dependencies Adapters) {
	ports := application.Ports{}

	ports.UserRepository = mysql.NewUserRepository(
		dependencies.MysqlConnection,
		config.Infrastructure.UserRepository,
	)
	ports.JWT = jwt.New(config.Infrastructure.JWT)
	ports.Password = password.New(config.Infrastructure.Password)
	ports.EventBusEmitter = buses.Event

	service := application.New(config.Service, ports)

	buses.Command.Register(
		domain.CommandSignInDefinition.Type(),
		domain.CommandSignInDefinition.CreateHandler(service.CommandSignIn),
	)
	buses.Command.Register(
		domain.CommandSignOutDefinition.Type(),
		domain.CommandSignOutDefinition.CreateHandler(service.CommandSignOut),
	)
	buses.Command.Register(
		domain.CommandAccessDefinition.Type(),
		domain.CommandAccessDefinition.CreateHandler(service.CommandAccess),
	)
}
