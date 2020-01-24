package bookmaker

import (
	"github.com/statistico/statistico-bet-finder/internal/app/statistico"
)

type MarketFactory interface {
	FixtureAndMarket(fix statistico.Fixture, market string) (*Market, error)
}

type RunnerFactory interface {
	CreateRunner(selectionID uint64, marketID, name string) (*Runner, error)
}

type Market struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Bookmaker string   `json:"bookmaker"`
	Runners   []Runner `json:"book"`
}

type Runner struct {
	Name        string  `json:"name"`
	SelectionID uint64  `json:"selection_id"`
	Back        []Price `json:"back"`
	Lay         []Price `json:"lay"`
}

type Price struct {
	Price float32 `json:"price"`
	Size  float32 `json:"size"`
}
