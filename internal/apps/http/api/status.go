package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a API) RouterStatus(r chi.Router) {
	r.Get("/ping", a.ControllerStatusPing)
}

func (a API) ControllerStatusPing(w http.ResponseWriter, _ *http.Request) {
	a.sendJSON(w, http.StatusOK, "pong")
}
