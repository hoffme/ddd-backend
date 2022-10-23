package command

import (
	"context"

	"github.com/hoffme/ddd-backend/internal/shared/bus/domain"
)

var _ domain.CommandBus = (*MemoryBus)(nil)

type MemoryBus struct {
	handlers map[domain.CommandType]domain.CommandHandler
}

func New() *MemoryBus {
	return &MemoryBus{
		handlers: make(map[domain.CommandType]domain.CommandHandler),
	}
}

func (b *MemoryBus) Dispatch(ctx context.Context, command domain.Command[any]) error {
	handler, ok := b.handlers[command.Type]
	if !ok {
		return nil
	}

	return handler(ctx, command)
}

func (b *MemoryBus) Register(commandType domain.CommandType, handler domain.CommandHandler) {
	b.handlers[commandType] = handler
}
