package api

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"nhooyr.io/websocket"
)

type Socket struct {
	closed   bool
	ctx      context.Context
	conn     *websocket.Conn
	id       string
	datetime time.Time
	onClose  func(socket *Socket)
}

func newSocket(ctx context.Context, conn *websocket.Conn) *Socket {
	return &Socket{
		ctx:      ctx,
		conn:     conn,
		closed:   false,
		id:       uuid.NewString(),
		datetime: time.Now(),
	}
}

func (c *Socket) ReadJSON(v any) error {
	messageType, reader, err := c.conn.Reader(c.ctx)
	if err != nil {
		return err
	}
	if messageType != websocket.MessageText {
		return errors.New("invalid message type")
	}

	return json.NewDecoder(reader).Decode(v)
}

func (c *Socket) ReadJSONCTX(ctx context.Context, v any) error {
	messageType, reader, err := c.conn.Reader(ctx)
	if err != nil {
		return err
	}
	if messageType != websocket.MessageText {
		return errors.New("invalid message type")
	}

	return json.NewDecoder(reader).Decode(v)
}

func (c *Socket) SendJSON(data interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return c.conn.Write(c.ctx, websocket.MessageText, bytes)
}

func (c *Socket) SendJSONCTX(ctx context.Context, data interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return c.conn.Write(ctx, websocket.MessageText, bytes)
}

func (c *Socket) Close(status websocket.StatusCode, message string) {
	c.closed = true

	err := c.conn.Close(status, message)
	if err != nil {
		log.Printf("Socket closed with error")
	}

	if c.onClose != nil {
		c.onClose(c)
	}
}
