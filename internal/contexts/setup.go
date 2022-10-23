package contexts

import (
	"github.com/hoffme/ddd-backend/internal/shared/bus"

	"github.com/hoffme/ddd-backend/internal/contexts/shared/infrastructure/mysql"

	"github.com/hoffme/ddd-backend/internal/contexts/auth"
)

type Config struct {
	Auth           auth.Config          `json:"auth"`
	Infrastructure ConfigInfrastructure `json:"infrastructure"`
}

type ConfigInfrastructure struct {
	MysqlConnection mysql.Config `json:"mysql_connection"`
}

func Setup(config Config, buses *bus.Buses) {
	mysqlConnection := mysql.Connect(config.Infrastructure.MysqlConnection)

	auth.Setup(config.Auth, buses, auth.Adapters{
		MysqlConnection: mysqlConnection,
	})

}
