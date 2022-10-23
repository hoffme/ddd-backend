package domain

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type QueryDefinition[T any, V any] struct {
	queryType QueryType
}

func NewQueryDefinition[T any, V any](queryType QueryType) QueryDefinition[T, V] {
	return QueryDefinition[T, V]{queryType: queryType}
}

func (d QueryDefinition[T, V]) Type() QueryType {
	return d.queryType
}

func (d QueryDefinition[T, V]) CreateQuery(params T) Query[any] {
	return Query[any]{
		Type:     d.queryType,
		ID:       uuid.NewString(),
		Datetime: time.Now(),
		Data:     params,
	}
}

func (d QueryDefinition[T, V]) CreateHandler(handler func(ctx context.Context, query Query[T]) (V, error)) QueryHandler {
	return func(ctx context.Context, query Query[any]) (interface{}, error) {
		data, ok := query.Data.(T)
		if !ok {
			panic(errors.New("invalid query data"))
		}

		return handler(ctx, Query[T]{
			Type:     query.Type,
			ID:       query.ID,
			Datetime: query.Datetime,
			Data:     data,
		})
	}
}
