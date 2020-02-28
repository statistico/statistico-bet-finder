package bootstrap

import (
	"github.com/statistico/statistico-price-finder/internal/app/bookmaker"
)

func (c Container) BookMaker() bookmaker.BookMaker {
	return bookmaker.NewBookMaker(c.GRPCFixtureClient(), c.BookmakerMarketBuilder(), c.Clock, c.Logger)
}

func (c Container) BookmakerMarketBuilder() bookmaker.MarketBuilder {
	bookmakers := []bookmaker.MarketFactory{
		c.BetFairMarketFactory(),
	}

	return bookmaker.NewMarketBuilder(bookmakers, c.Logger)
}
