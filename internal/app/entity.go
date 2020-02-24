package app

import (
	"github.com/statistico/statistico-bet-finder/internal/app/bookmaker"
	"time"
)

type Book struct {
	EventID   uint64    `json:"eventId"`
	Markets   []*Market `json:"markets"`
	CreatedAt time.Time `json:"createdAt"`
}

// Market is a struct for a given market name with hydrated bookmaker markets
type Market struct {
	Name       string              `json:"name"`
	Bookmakers []*bookmaker.Market `json:"bookmakers"`
}
