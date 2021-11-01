package replicators

import (
	"log"

	"github.com/market-tracker/market-tracker/wsMsg"
)

// Dummy follows the Replicator interface of
type Dummy struct{}

func (d *Dummy) Publish(msg *wsMsg.MarketTrackerMsg) {
	log.Printf("Market: %v", msg)
}
