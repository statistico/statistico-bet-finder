package bootstrap

import "github.com/statistico/statistico-price-finder/internal/app/statistico"

func (c Container) StatisticoBookmaker() statistico.BookMaker {
	return statistico.NewBookMaker(c.GRPCFixtureClient(), c.StatisticoMarketBuilder(), c.Clock, c.Logger)
}

func (c Container) StatisticoMarketBuilder() statistico.MarketBuilder {
	return statistico.NewMarketBuilder(c.GRPCOddsCompilerClient())
}
