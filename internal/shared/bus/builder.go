package bus

import (
	"github.com/hoffme/ddd-backend/internal/shared/bus/domain"

	"github.com/hoffme/ddd-backend/internal/shared/bus/infrastructure/memory/command"
	"github.com/hoffme/ddd-backend/internal/shared/bus/infrastructure/memory/event"
	"github.com/hoffme/ddd-backend/internal/shared/bus/infrastructure/memory/query"
)

type Buses struct {
	Event   domain.EventBus
	Command domain.CommandBus
	Query   domain.QueryBus
}

func Build() *Buses {
	memoryEventBus := event.New()
	memoryCommandBus := command.New()
	memoryQueryBus := query.New()

	return &Buses{
		Event:   memoryEventBus,
		Command: memoryCommandBus,
		Query:   memoryQueryBus,
	}
}
