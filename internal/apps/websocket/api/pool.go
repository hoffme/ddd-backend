package api

import (
	"context"
	"log"

	"github.com/hoffme/ddd-backend/internal/shared/bus"
	"github.com/hoffme/ddd-backend/internal/shared/bus/domain"

	"nhooyr.io/websocket"
)

type Message struct {
	Action    string           `json:"action"`
	EventType domain.EventType `json:"event"`
	Data      interface{}      `json:"data,omitempty"`
}

type Subscription struct {
	subscription domain.EventSubscription
	sockets      map[string]*Socket
}

type Pool struct {
	buses   *bus.Buses
	sockets map[string]*Socket
	events  map[domain.EventType]*Subscription
	actions map[string]func(socket *Socket, message Message)
}

func newPool(buses *bus.Buses) *Pool {
	pool := &Pool{
		buses:   buses,
		sockets: make(map[string]*Socket),
		events:  make(map[domain.EventType]*Subscription),
		actions: make(map[string]func(socket *Socket, message Message)),
	}

	pool.actions["subscribe"] = pool.actionSubscribe
	pool.actions["unsubscribe"] = pool.actionUnsubscribe

	return pool
}

func (p *Pool) AddSocket(socket *Socket) {
	socket.onClose = p.DeleteSocket

	defer socket.Close(websocket.StatusNormalClosure, "close")

	p.sockets[socket.id] = socket

	for {
		p.listenMessage(socket)
	}
}

func (p *Pool) DeleteSocket(socket *Socket) {
	delete(p.sockets, socket.id)
}

func (p *Pool) listenMessage(socket *Socket) {
	message := Message{}
	err := socket.ReadJSON(&message)
	if err != nil {
		return
	}

	actionFunc, ok := p.actions[message.Action]
	if !ok {
		return
	}

	actionFunc(socket, message)
}

func (p *Pool) actionSubscribe(socket *Socket, message Message) {
	eventSubscribers, ok := p.events[message.EventType]
	if !ok {
		eventSubscribers = &Subscription{
			subscription: p.buses.Event.Subscribe(message.EventType, p.emitEvent),
			sockets:      map[string]*Socket{},
		}
	}

	eventSubscribers.sockets[socket.id] = socket

	p.events[message.EventType] = eventSubscribers
}

func (p *Pool) actionUnsubscribe(socket *Socket, message Message) {
	eventSubscribers, ok := p.events[message.EventType]
	if !ok {
		return
	}

	delete(eventSubscribers.sockets, socket.id)

	if len(eventSubscribers.sockets) == 0 {
		eventSubscribers.subscription.Unsubscribe()
		delete(p.events, message.EventType)
	}
}

func (p *Pool) emitEvent(ctx context.Context, event domain.Event[any]) {
	eventSubscribers, ok := p.events[event.Type]
	if !ok {
		return
	}

	for _, socket := range eventSubscribers.sockets {
		err := socket.SendJSONCTX(ctx, Message{
			Action:    "event",
			EventType: event.Type,
			Data:      event.Data,
		})
		if err != nil {
			log.Printf("error to send event: %s\n", err.Error())
		}
	}
}
