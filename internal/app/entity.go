package app

import (
	"github.com/statistico/statistico-bet-finder/internal/app/bookmaker"
	"time"
)

type Book struct {
	Market []*Market `json:"markets"`
	CreatedAt time.Time `json:"created_at"`
}

// Market is a struct for a given fixture and market name with hydrated Statistico and Bookmaker markets
type Market struct {
	FixtureID uint64   `json:"fixture_id"`
	Name      string   `json:"name"`
	Statistico *StatisticoMarket `json:"statistico_market"`
	Bookmaker  []*bookmaker.Market `json:"bookmaker_market"`
}

// StatisticoMarket is a struct containing Statistico calculated odds for a fixture.
type StatisticoMarket struct {
	FixtureID uint64   `json:"fixture_id"`
	Name      string   `json:"name"`
	Runners []Runner   `json:"runners"`
}

// Runner is a struct containing individual runner information.
type Runner struct {
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
}
