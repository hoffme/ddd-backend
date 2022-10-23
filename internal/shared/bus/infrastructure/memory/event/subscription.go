package event

import (
	"github.com/hoffme/ddd-backend/internal/shared/bus/domain"

	"github.com/google/uuid"
)

var _ domain.EventSubscription = (*memorySubscription)(nil)

type memorySubscription struct {
	eventType domain.EventType
	id        string
	bus       *MemoryBus
	handler   domain.EventHandler
}

func newSubscription(bus *MemoryBus, eventType domain.EventType, handler domain.EventHandler) *memorySubscription {
	return &memorySubscription{
		eventType: eventType,
		id:        uuid.NewString(),
		bus:       bus,
		handler:   handler,
	}
}

func (s *memorySubscription) Unsubscribe() {
	s.bus.unsubscribe(s)
}
