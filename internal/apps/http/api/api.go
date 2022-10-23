package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/hoffme/ddd-backend/internal/shared/bus"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type API struct {
	buses *bus.Buses
}

func New(buses *bus.Buses) API {
	return API{
		buses: buses,
	}
}

func (a API) Router() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	r.Use(middleware.Timeout(time.Minute))

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-aType", "X-CSRF-RefreshToken"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler)

	r.Route("/status", a.RouterStatus)
	r.Route("/auth", a.RouterAuth)

	return r
}

func (a API) sendJSON(w http.ResponseWriter, status int, result interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if result == nil {
		return
	}

	data := struct {
		Result interface{} `json:"result,omitempty"`
		Error  interface{} `json:"error,omitempty"`
	}{}

	if status >= 200 && status < 300 {
		data.Result = result
	} else {
		data.Error = result
	}

	err := json.NewEncoder(w).Encode(&data)
	if err != nil {
		w.WriteHeader(500)
	}
}
