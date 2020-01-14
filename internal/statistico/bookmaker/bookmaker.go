package bookmaker

import "github.com/statistico/statistico-bet-finder/internal/statistico"

type Service interface {
	Book(query ServiceQuery) <-chan *Book
}

type ServiceQuery struct {
	BetTypes  []string
	Fixtures []statistico.Fixture
}

type Book struct {
	FixtureID uint64   `json:"fixture_id"`
	Bookmaker string   `json:"bookmaker"`
	Markets   []Market `json:"markets"`
}

type Market struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Runners []Runner `json:"book"`
}

type Runner struct {
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Available   float32 `json:"available"`
	SelectionID uint64  `json:"selection_id"`
}
