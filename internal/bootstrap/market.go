package bootstrap

import (
	"github.com/statistico/statistico-bet-finder/internal/app"
	"github.com/statistico/statistico-bet-finder/internal/app/bookmaker"
)

func (c Container) MarketBuilder() app.MarketBuilder {
	bookmakers := []bookmaker.MarketFactory{
		c.BetFairMarketFactory(),
	}

	return app.NewMarketBuilder(c.StatisticoOddsCompilerClient(), bookmakers, c.Logger)
}

