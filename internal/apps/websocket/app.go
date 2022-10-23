package websocket

import (
	"github.com/hoffme/ddd-backend/internal/apps/app"
	"github.com/hoffme/ddd-backend/internal/apps/websocket/api"
	"net/http"

	"github.com/hoffme/ddd-backend/internal/shared/bus"
)

var _ app.App = (*App)(nil)

type App struct {
	config  Config
	handler http.Handler
}

func New(config Config, buses *bus.Buses) *App {
	router := api.New(buses).Router()

	return &App{
		config:  config,
		handler: router,
	}
}

func (h *App) Name() string {
	return "websocket"
}

func (h *App) Entrypoint() string {
	return "ws://" + h.config.Addr
}

func (h *App) Run() error {
	return http.ListenAndServe(h.config.Addr, h.handler)
}
