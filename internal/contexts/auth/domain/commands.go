package domain

import "github.com/hoffme/ddd-backend/internal/shared/bus/domain"

// Sign In

const CommandSignInType domain.CommandType = "command.auth.signIn"

type CommandSignInData struct {
	Nick     string `json:"nick"`
	Password string `json:"password"`
}

var CommandSignInDefinition = domain.NewCommandDefinition[CommandSignInData](CommandSignInType)

// Sign Out

const CommandSignOutType domain.CommandType = "command.auth.signOut"

type CommandSignOutData struct {
	RefreshToken string `json:"refreshToken"`
}

var CommandSignOutDefinition = domain.NewCommandDefinition[CommandSignOutData](CommandSignOutType)

// Access

const CommandAccessType domain.CommandType = "command.auth.access"

type CommandAccessData struct {
	RefreshToken string `json:"refreshToken"`
}

var CommandAccessDefinition = domain.NewCommandDefinition[CommandAccessData](CommandAccessType)
