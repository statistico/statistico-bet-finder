package bootstrap

import (
	"github.com/statistico/statistico-price-finder/internal/app/bookmaker"
)

func (c Container) MarketBuilder() bookmaker.MarketBuilder {
	bookmakers := []bookmaker.MarketFactory{
		c.BetFairMarketFactory(),
	}

	return bookmaker.NewMarketBuilder(bookmakers, c.Logger)
}
