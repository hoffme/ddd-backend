package domain

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type CommandDefinition[T any] struct {
	commandType CommandType
}

func NewCommandDefinition[T any](commandType CommandType) CommandDefinition[T] {
	return CommandDefinition[T]{commandType: commandType}
}

func (d CommandDefinition[T]) Type() CommandType {
	return d.commandType
}

func (d CommandDefinition[T]) CreateCommand(params T) Command[any] {
	return Command[any]{
		Type:     d.commandType,
		ID:       uuid.NewString(),
		Datetime: time.Now(),
		Data:     params,
	}
}

func (d CommandDefinition[T]) CreateHandler(handler func(ctx context.Context, command Command[T]) error) CommandHandler {
	return func(ctx context.Context, command Command[any]) error {
		data, ok := command.Data.(T)
		if !ok {
			panic(errors.New("invalid command data"))
		}

		return handler(ctx, Command[T]{
			Type:     command.Type,
			ID:       command.ID,
			Datetime: command.Datetime,
			Data:     data,
		})
	}
}
