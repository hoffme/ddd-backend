package domain

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type EventDefinition[T any] struct {
	eventType EventType
}

func NewEventDefinition[T any](eventType EventType) EventDefinition[T] {
	return EventDefinition[T]{eventType: eventType}
}

func (d EventDefinition[T]) Type() EventType {
	return d.eventType
}

func (d EventDefinition[T]) CreateEvent(params T) Event[any] {
	return Event[any]{
		Type:     d.eventType,
		ID:       uuid.NewString(),
		Datetime: time.Now(),
		Data:     params,
	}
}

func (d EventDefinition[T]) CreateHandler(handler func(ctx context.Context, event Event[T])) EventHandler {
	return func(ctx context.Context, event Event[any]) {
		data, ok := event.Data.(T)
		if !ok {
			panic(errors.New("invalid event data"))
		}

		handler(ctx, Event[T]{
			Type:     event.Type,
			ID:       event.ID,
			Datetime: event.Datetime,
			Data:     data,
		})
	}
}
