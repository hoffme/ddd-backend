package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Addr string `json:"addr"`
	User string `json:"user"`
	Pass string `json:"pass"`
	DB   string `json:"db"`
}

type Connection struct {
	db *sql.DB
}

func Connect(config Config) *Connection {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		config.User,
		config.Pass,
		config.Addr,
		config.DB,
	)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	connection := &Connection{
		db: db,
	}

	return connection
}

func (c *Connection) DB() *sql.DB {
	return c.db
}

func (c *Connection) Close() error {
	return c.db.Close()
}
