package bookmaker

import "github.com/statistico/statistico-bet-finder/internal/statistico"

type Service interface {
	MarketBooks(query ServiceQuery) ([]Book, error)
}

type ServiceQuery struct {
	BetType  []string
	Fixtures []statistico.Fixture
}

type Book struct {
	Bookmaker string   `json:"bookmaker"`
	FixtureID uint64   `json:"fixture_id"`
	BetType   string   `json:"bet_type"`
	Market    []Market `json:"market"`
}

type Market struct {
	Selection   string  `json:"selection"`
	Price       float32 `json:"price"`
	MarketID    string  `json:"market_id"`
	SelectionID string  `json:"selection_id"`
}
