package api

import (
	"log"
	"net/http"

	"github.com/hoffme/ddd-backend/internal/shared/bus"

	"nhooyr.io/websocket"
)

type API struct {
	pool *Pool
}

func New(buses *bus.Buses) API {
	pool := newPool(buses)

	return API{pool: pool}
}

func (a API) Router() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Accept(w, r, nil)
		if err != nil {
			log.Printf("error to accept Socket: %s\n", err.Error())
			return
		}

		a.pool.AddSocket(newSocket(r.Context(), conn))
	})
}
