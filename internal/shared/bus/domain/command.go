package domain

import (
	"context"
	"time"
)

type CommandType string

type Command[T interface{}] struct {
	Type     CommandType
	ID       string
	Datetime time.Time
	Data     T
}

type CommandHandler func(ctx context.Context, command Command[any]) error

type CommandDispatcher interface {
	Dispatch(ctx context.Context, command Command[any]) error
}

type CommandRegister interface {
	Register(commandType CommandType, handler CommandHandler)
}

type CommandBus interface {
	CommandDispatcher
	CommandRegister
}
