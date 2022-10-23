package event

import (
	"context"

	"github.com/hoffme/ddd-backend/internal/shared/bus/domain"
)

var _ domain.EventBus = (*MemoryBus)(nil)

type MemoryBus struct {
	eventsSubscriptions map[domain.EventType]map[string]*memorySubscription
}

func New() *MemoryBus {
	return &MemoryBus{
		eventsSubscriptions: make(map[domain.EventType]map[string]*memorySubscription),
	}
}

func (b *MemoryBus) Emit(ctx context.Context, event domain.Event[any]) {
	subscriptions, ok := b.eventsSubscriptions[event.Type]
	if !ok {
		return
	}

	for _, subscription := range subscriptions {
		subscription.handler(ctx, event)
	}
}

func (b *MemoryBus) Subscribe(eventType domain.EventType, handler domain.EventHandler) domain.EventSubscription {
	subscriptions, ok := b.eventsSubscriptions[eventType]
	if !ok {
		subscriptions = make(map[string]*memorySubscription, 1)
	}

	subscription := newSubscription(b, eventType, handler)
	subscriptions[subscription.id] = subscription

	b.eventsSubscriptions[eventType] = subscriptions

	return subscription
}

func (b *MemoryBus) unsubscribe(subscription *memorySubscription) {
	subscriptions, ok := b.eventsSubscriptions[subscription.eventType]
	if !ok {
		return
	}

	delete(subscriptions, subscription.id)

	if len(subscriptions) > 0 {
		b.eventsSubscriptions[subscription.eventType] = subscriptions
	} else {
		delete(b.eventsSubscriptions, subscription.eventType)
	}
}
