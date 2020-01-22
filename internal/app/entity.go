package app

import (
	"github.com/statistico/statistico-bet-finder/internal/app/bookmaker"
	"github.com/statistico/statistico-bet-finder/internal/app/statistico"
	"time"
)

type Book struct {
	Markets []*Market `json:"markets"`
	CreatedAt time.Time `json:"created_at"`
}

// Market is a struct for a given fixture and market name with hydrated Statistico and Bookmaker markets
type Market struct {
	FixtureID uint64   `json:"fixture_id"`
	Name      string   `json:"name"`
	Statistico *statistico.Market `json:"statistico"`
	Bookmakers  []*bookmaker.Market `json:"bookmakers"`
}
