package query

import (
	"context"

	"github.com/hoffme/ddd-backend/internal/shared/bus/domain"
)

var _ domain.QueryBus = (*MemoryBus)(nil)

type MemoryBus struct {
	handlers map[domain.QueryType]domain.QueryHandler
}

func New() *MemoryBus {
	return &MemoryBus{
		handlers: make(map[domain.QueryType]domain.QueryHandler),
	}
}

func (b *MemoryBus) Dispatch(ctx context.Context, query domain.Query[any]) (interface{}, error) {
	handler, ok := b.handlers[query.Type]
	if !ok {
		return nil, nil
	}

	return handler(ctx, query)
}

func (b *MemoryBus) Register(queryType domain.QueryType, handler domain.QueryHandler) {
	b.handlers[queryType] = handler
}
