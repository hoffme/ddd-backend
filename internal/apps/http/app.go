package http

import (
	"net/http"

	"github.com/hoffme/ddd-backend/internal/apps/http/api"

	"github.com/hoffme/ddd-backend/internal/shared/bus"
)

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
	return "http"
}

func (h *App) Entrypoint() string {
	return "http://" + h.config.Addr
}

func (h *App) Run() error {
	return http.ListenAndServe(h.config.Addr, h.handler)
}
