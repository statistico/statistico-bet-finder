package app

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-bet-finder/internal/app/bookmaker"
	"github.com/statistico/statistico-bet-finder/internal/app/statistico"
)

type MarketBuilder interface {
	FixtureAndMarket(f *statistico.Fixture, bet string) (*Market, error)
}

// MarketBuilder builds markets for Statistico and associated bookmakers.
type marketBuilder struct {
	oddsClient statistico.OddsCompilerClient
	bookmakers []bookmaker.MarketFactory
	logger     *logrus.Logger
}

// FixtureAndBetType creates a Market struct for a given Fixture and Market i.e. OVER_UNDER_25.
func (m marketBuilder) FixtureAndMarket(f *statistico.Fixture, market string) (*Market, error) {
	mark := Market{
		FixtureID: f.ID,
		Name:      market,
	}

	odds, err := m.oddsClient.GetOverUnderGoalsForFixture(f.ID, market)

	if err != nil {
		return nil, fmt.Errorf(
			"error '%s' building statistico odds for fixture '%d' and market '%s'",
			err.Error(),
			f.ID,
			market,
		)
	}

	mark.Statistico = odds

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

func NewMarketBuilder(odds statistico.OddsCompilerClient, book []bookmaker.MarketFactory, log *logrus.Logger) MarketBuilder {
	return &marketBuilder{oddsClient: odds, bookmakers: book, logger: log}
}
