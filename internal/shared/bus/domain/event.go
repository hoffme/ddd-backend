package domain

import (
	"context"
	"time"
)

type EventType string

type Event[T any] struct {
	Type     EventType
	ID       string
	Datetime time.Time
	Data     T
}

type EventHandler func(ctx context.Context, event Event[any])

type EventSubscription interface {
	Unsubscribe()
}

type EventEmitter interface {
	Emit(ctx context.Context, event Event[any])
}

type EventSubscriber interface {
	Subscribe(eventType EventType, handler EventHandler) EventSubscription
}

type EventBus interface {
	EventEmitter
	EventSubscriber
}
