package wsTiingo

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/market-tracker/market-tracker/pkg/errorHandler"
	"github.com/market-tracker/market-tracker/pkg/wsWrapper"
	"nhooyr.io/websocket"
)

// TODO: dependency injection strategy
type WsTiingo struct {
	conn      *websocket.Conn
	wsWrapper *wsWrapper.WsWrapper
	opts      *TiingoOptions
}

type EventDataTiingo struct {
	ThresholdLevel int `json:"thresholdLevel"`
}

type TiingoOptions struct {
	EventName     string           `json:"eventName"`
	Authorization string           `json:"authorization"`
	EventData     *EventDataTiingo `json:"eventData"`
}

// TODO: create an options struct, to defined the variables to setup this websocket
func NewWsTiingo(ctx context.Context, tiingoWsUrl string, opts *TiingoOptions) *WsTiingo {
	dialOps := &websocket.DialOptions{}
	c, _, err := websocket.Dial(ctx, tiingoWsUrl, dialOps)
	errorHandler.PanicError(err)
	wsWrapper := wsWrapper.NewWsWrapper(16)
	return &WsTiingo{
		conn:      c,
		wsWrapper: wsWrapper,
		opts:      opts,
	}
}

func (w *WsTiingo) Close() error {
	w.wsWrapper.Close(w.conn)
	return nil
}

func (w *WsTiingo) Subscribe(ctx context.Context) {
	interrupt := make(chan os.Signal, 1)

	msg, err := json.Marshal(w.opts)
	errorHandler.LogError(err)
	if err = w.conn.Write(ctx, websocket.MessageText, msg); err != nil {
		// TODO: It is necesary to panic in this part if some websocket failed
		errorHandler.PanicError(err)
	}
	// subscribeOptions := &wsWrapper.SubscribeOptions{
	// 	Event:      "",
	// 	NotifierFn: w.notifierFunc,
	// }
	// w.wsWrapper.Subscribe(ctx, w.conn, subscribeOptions)

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			msgType, message, err := w.conn.Read(ctx)
			fmt.Println()
			fmt.Println("message")
			log.Printf("recv: %s", message)
			fmt.Println(msgType)
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	select {
	case <-interrupt:
		log.Println("interrupt")
	case <-done:
		return
	}
	fmt.Println("message2")
	err = w.conn.Close(websocket.StatusNormalClosure, "")
	if err != nil {
		log.Println("write close:", err)
		return
	}
	<-done
}

func (w *WsTiingo) notifierFunc(parameters interface{}) {}
