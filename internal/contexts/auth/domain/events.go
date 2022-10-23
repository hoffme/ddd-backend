package domain

import (
	"github.com/hoffme/ddd-backend/internal/shared/bus/domain"
)

// Sign In

const EventSignInType domain.EventType = "event.auth.signIn"

type EventSignInData struct {
	CommandId    string `json:"commandId"`
	RefreshToken string `json:"refreshToken"`
	ExpireAt     int64  `json:"expireAt"`
}

var EventSignInDefinition = domain.NewEventDefinition[EventSignInData](EventSignInType)

// Sign Out

const EventSignOutType domain.EventType = "event.auth.signOut"

type EventSignOutData struct {
	RefreshToken string `json:"refreshToken"`
	CommandId    string `json:"commandId"`
}

var EventSignOutDefinition = domain.NewEventDefinition[EventSignOutData](EventSignOutType)

// Access

const EventAccessType domain.EventType = "event.auth.access"

type EventAccessData struct {
	CommandId    string `json:"commandId"`
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
}

var EventAccessDefinition = domain.NewEventDefinition[EventAccessData](EventAccessType)
