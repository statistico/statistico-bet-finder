package bookmaker

import "github.com/statistico/statistico-price-finder/internal/app/grpc/proto"

type MarketFactory interface {
	FixtureAndMarket(fix *proto.Fixture, market string) (*Market, error)
}

type RunnerFactory interface {
	CreateRunner(selectionID uint64, marketID, name string) (*Runner, error)
}

type Market struct {
	ID        string   `json:"marketId"`
	Bookmaker string   `json:"bookmaker"`
	Runners   []Runner `json:"runners"`
}

type Runner struct {
	Name        string  `json:"name"`
	SelectionID uint64  `json:"selectionId"`
	Back        []Price `json:"back"`
	Lay         []Price `json:"lay"`
}

type Price struct {
	Price float32 `json:"price"`
	Size  float32 `json:"size"`
}
