package domain

import (
	"context"
	"time"
)

type QueryType string

type Query[T any] struct {
	Type     QueryType
	ID       string
	Datetime time.Time
	Data     T
}

type QueryHandler func(ctx context.Context, query Query[any]) (interface{}, error)

type QueryDispatcher interface {
	Dispatch(ctx context.Context, query Query[any]) (interface{}, error)
}

type QueryRegister interface {
	Register(queryType QueryType, handler QueryHandler)
}

type QueryBus interface {
	QueryDispatcher
	QueryRegister
}
