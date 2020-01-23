package bootstrap

import (
	"github.com/statistico/statistico-bet-finder/internal/app"
	"github.com/statistico/statistico-bet-finder/internal/app/bookmaker"
	"github.com/statistico/statistico-bet-finder/internal/app/bookmaker/betfair"
)

func (c Container) BookMaker() app.BookMaker {
	return app.NewBookMaker(c.StatisticoFixtureClient(), c.MarketBuilder(), c.Clock, c.Logger)
}

func (c Container) BetFairRunnerFactory() bookmaker.RunnerFactory {
	return betfair.NewRunnerFactory(c.BetFairClient)
}

func (c Container) BetFairMarketFactory() bookmaker.MarketFactory {
	return betfair.NewMarketFactory(c.BetFairClient, c.BetFairRunnerFactory())
}
