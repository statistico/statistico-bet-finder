package bookmaker

import "github.com/statistico/statistico-price-finder/internal/app/grpc/proto"

type MarketFactory interface {
	FixtureAndMarket(fix *proto.Fixture, market string) (*SubMarket, error)
}

type RunnerFactory interface {
	CreateRunner(selectionID uint64, marketID, name string) (*Runner, error)
}
