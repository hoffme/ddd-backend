package apps

import (
	"github.com/hoffme/ddd-backend/internal/apps/http"
	"github.com/hoffme/ddd-backend/internal/apps/websocket"
)

type Config struct {
	Http      http.Config      `json:"http"`
	Websocket websocket.Config `json:"websocket"`
}
