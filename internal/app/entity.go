package app

import (
	"github.com/statistico/statistico-bet-finder/internal/app/bookmaker"
	"time"
)

type Book struct {
	Markets   []*Market `json:"markets"`
	CreatedAt time.Time `json:"created_at"`
}

// Market is a struct for a given fixture and market name with hydrated Statistico and Bookmaker markets
type Market struct {
	FixtureID  uint64              `json:"fixture_id"`
	Name       string              `json:"name"`
	Bookmakers []*bookmaker.Market `json:"bookmakers"`
}
