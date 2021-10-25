package main

import (
	"context"
	"log"

	"github.com/market-tracker/market-tracker/config"
	"github.com/market-tracker/market-tracker/server"
	"github.com/market-tracker/market-tracker/wsTiingo"
)

func init() {
	// Containers of global dependencies
	log.SetFlags(0)
	ctx := context.Background()
	c := config.GetConfiguration()
	tiingoOpts := &wsTiingo.TiingoOptions{
		EventName:     "subscribe",
		Authorization: c.TiingoApiToken,
		EventData: &wsTiingo.EventDataTiingo{
			ThresholdLevel: 5,
		},
	}
	ws := wsTiingo.GetWsTiingo(ctx, c.TiingoApiUrl, tiingoOpts)

	// run in a go rutine because in the subscription, the subscriber is waiting
	// for msgs
	go ws.Subscribe(ctx)
}

func main() {
	c := config.GetConfiguration()
	s := server.InitServer(c.Port)
	s.Start()
}
