package wsWrapper

import (
	"context"

	"nhooyr.io/websocket"
)

type IWsWrapper interface {
	Subscribe(ctx context.Context, conn *websocket.Conn) error
	Unsubscribe(conn *websocket.Conn, event string) error
	Emit(message interface{}) error
	Close() error
}

type Subscriber struct {
	msgs      chan []byte
	closeSlow func()
}

// It is simmilar to the callback in javascript
// TODO: define parameters
type NotifierFunc func(interface{})

type SubscribeOptions struct {
	Event      string
	NotifierFn NotifierFunc
}
