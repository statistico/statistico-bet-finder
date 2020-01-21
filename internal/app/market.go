package app

import (
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-bet-finder/internal/app/bookmaker"
	"github.com/statistico/statistico-bet-finder/internal/app/statistico"
)

type MarketBuilder interface {
	FixtureAndMarket(f *statistico.Fixture, bet string) *Market
}

// MarketBuilder builds markets for Statistico and associated bookmakers.
type marketBuilder struct {
	oddsClient statistico.OddsCompilerClient
	bookmakers []bookmaker.MarketFactory
	logger     *logrus.Logger
}

// FixtureAndBetType creates a Market struct for a given Fixture and market.
func (m marketBuilder) FixtureAndMarket(f *statistico.Fixture, market string) *Market {
	mark := Market{
		FixtureID:  f.ID,
		Name:       market,
	}

	odds, err := m.oddsClient.GetOverUnderGoalsForFixture(f.ID, market)

	if err != nil {
		m.logger.Warnf("Error '%s' building statistico odds for fixture '%d' and market '%s'", err.Error(), f.ID, market)
		return nil
	}

	mark.Statistico = odds

	for _, bookie := range m.bookmakers {
		mk, err := bookie.FixtureAndMarket(*f, market)

		if err != nil {
			m.logger.Warnf("Error '%s' building bookmaker odds for fixture '%d' and market '%s'", err.Error(), f.ID, market)
			return nil
		}

		mark.Bookmaker = append(mark.Bookmaker, mk)
	}

	return &mark
}

func NewMarketBuilder(odds statistico.OddsCompilerClient, book []bookmaker.MarketFactory, log *logrus.Logger) MarketBuilder {
	return &marketBuilder{oddsClient: odds, bookmakers: book, logger: log}
}
