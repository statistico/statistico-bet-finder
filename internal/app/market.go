package app

import (
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-bet-finder/internal/app/bookmaker"
	"github.com/statistico/statistico-bet-finder/internal/app/statistico"
)

type MarketBuilder interface {
	FixtureAndMarket(f *statistico.Fixture, bet string) (*Market, error)
}

// MarketBuilder builds markets for Statistico and associated bookmakers.
type marketBuilder struct {
	bookmakers []bookmaker.MarketFactory
	logger     *logrus.Logger
}

// FixtureAndBetType creates a Market struct for a given Fixture and Market i.e. OVER_UNDER_25.
func (m marketBuilder) FixtureAndMarket(f *statistico.Fixture, market string) (*Market, error) {
	mark := Market{
		FixtureID: f.ID,
		Name:      market,
	}

	for _, bookie := range m.bookmakers {
		mk, err := bookie.FixtureAndMarket(*f, market)

		if err != nil {
			m.logger.Warnf(
				"Error '%s' building bookmaker odds for fixture '%d' and market '%s'",
				err.Error(),
				f.ID,
				market,
			)
			continue
		}

		mark.Bookmakers = append(mark.Bookmakers, mk)
	}

	return &mark, nil
}

func NewMarketBuilder(book []bookmaker.MarketFactory, log *logrus.Logger) MarketBuilder {
	return &marketBuilder{bookmakers: book, logger: log}
}
