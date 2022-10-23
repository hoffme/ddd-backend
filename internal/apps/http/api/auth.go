package api

import (
	"context"
	"encoding/json"
	"net/http"

	busDomain "github.com/hoffme/ddd-backend/internal/shared/bus/domain"

	"github.com/hoffme/ddd-backend/internal/contexts/auth/domain"

	"github.com/go-chi/chi/v5"
)

func (a API) RouterAuth(r chi.Router) {
	r.Post("/signIn", a.AuthSignIn)
	r.Post("/signOut", a.AuthSignOut)
	r.Post("/access", a.AuthAccess)
}

func (a API) AuthSignIn(w http.ResponseWriter, r *http.Request) {
	params := domain.CommandSignInData{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		a.sendJSON(w, http.StatusBadRequest, "invalid params")
		return
	}

	command := domain.CommandSignInDefinition.CreateCommand(params)

	var subscription busDomain.EventSubscription
	eventHandler := domain.EventSignInDefinition.CreateHandler(
		func(ctx context.Context, event busDomain.Event[domain.EventSignInData]) {
			if event.Data.CommandId != command.ID {
				return
			}

			a.sendJSON(w, http.StatusOK, event.Data)
			subscription.Unsubscribe()
		},
	)
	subscription = a.buses.Event.Subscribe(domain.EventSignInDefinition.Type(), eventHandler)

	err = a.buses.Command.Dispatch(r.Context(), command)
	if err != nil {
		a.sendJSON(w, http.StatusBadRequest, err.Error())
		subscription.Unsubscribe()
	}
}

func (a API) AuthSignOut(w http.ResponseWriter, r *http.Request) {
	params := domain.CommandSignOutData{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		a.sendJSON(w, http.StatusBadRequest, "invalid params")
		return
	}

	command := domain.CommandSignOutDefinition.CreateCommand(params)
	err = a.buses.Command.Dispatch(r.Context(), command)
	if err != nil {
		a.sendJSON(w, http.StatusBadRequest, err.Error())
	}

	a.sendJSON(w, http.StatusOK, true)
}

func (a API) AuthAccess(w http.ResponseWriter, r *http.Request) {
	params := domain.CommandAccessData{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		a.sendJSON(w, http.StatusBadRequest, "invalid params")
		return
	}

	command := domain.CommandAccessDefinition.CreateCommand(params)

	var subscription busDomain.EventSubscription
	eventHandler := domain.EventAccessDefinition.CreateHandler(
		func(ctx context.Context, event busDomain.Event[domain.EventAccessData]) {
			if event.Data.CommandId != command.ID {
				return
			}

			a.sendJSON(w, http.StatusOK, event.Data)
			subscription.Unsubscribe()
		},
	)
	subscription = a.buses.Event.Subscribe(domain.EventAccessDefinition.Type(), eventHandler)

	err = a.buses.Command.Dispatch(r.Context(), command)
	if err != nil {
		a.sendJSON(w, http.StatusBadRequest, err.Error())
		subscription.Unsubscribe()
	}
}
