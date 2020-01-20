package app

import (
	"github.com/statistico/statistico-bet-finder/internal/app/bookmaker"
	"github.com/statistico/statistico-bet-finder/internal/app/grpc"
)

type MarketBuilder interface {
	FixtureAndBetType(f *Fixture, bet string) *Market
}

// MarketBuilder builds markets from Statistico and associated bookmakers.
type marketBuilder struct {
	oddsClient grpc.OddsCompilerClient
	bookmakers []bookmaker.MarketFactory
}

// FixtureAndBetType creates a Market struct for a given Fixture and bet type.
func (m marketBuilder) FixtureAndBetType(f *Fixture, bet string) *Market {
	market := Market{
		FixtureID:  f.ID,
		Name:       bet,
	}

	odds, err := m.oddsClient.GetOverUnderGoalsForFixture(f.ID, bet)

	if err != nil {
		// Log error here
		return nil
	}

	market.Statistico = odds

	for _, bookie := range m.bookmakers {
		m, err := bookie.FixtureAndBetType(*f, bet)

		if err != nil {
			// Log error here
			continue
		}

		market.Bookmaker = append(market.Bookmaker, m)
	}

	return &market
}

func NewMarketBuilder(odds grpc.OddsCompilerClient, book []bookmaker.MarketFactory) MarketBuilder {
	return &marketBuilder{oddsClient: odds, bookmakers: book}
}
