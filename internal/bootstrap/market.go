package bootstrap

import (
	"github.com/statistico/statistico-price-finder/internal/app"
	"github.com/statistico/statistico-price-finder/internal/app/bookmaker"
)

func (c Container) MarketBuilder() app.MarketBuilder {
	bookmakers := []bookmaker.MarketFactory{
		c.BetFairMarketFactory(),
	}

	return app.NewMarketBuilder(bookmakers, c.Logger)
}
