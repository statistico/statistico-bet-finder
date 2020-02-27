package bootstrap

import (
	"github.com/statistico/statistico-price-finder/internal/app/bookmaker"
	"github.com/statistico/statistico-price-finder/internal/app/bookmaker/betfair"
)

func (c Container) BookMaker() bookmaker.BookMaker {
	return bookmaker.NewBookMaker(c.GRPCFixtureClient(), c.MarketBuilder(), c.Clock, c.Logger)
}

func (c Container) BetFairRunnerFactory() bookmaker.RunnerFactory {
	return betfair.NewRunnerFactory(c.BetFairClient)
}

func (c Container) BetFairMarketFactory() bookmaker.MarketFactory {
	return betfair.NewMarketFactory(c.BetFairClient, c.BetFairRunnerFactory())
}
