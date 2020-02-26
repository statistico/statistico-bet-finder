package bookmaker

import (
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-price-finder/internal/app/grpc/proto"
)

type MarketBuilder interface {
	FixtureAndMarket(f *proto.Fixture, bet string) (*Market, error)
}

// MarketBuilder builds markets for bookmakers.
type marketBuilder struct {
	bookmakers []MarketFactory
	logger     *logrus.Logger
}

// FixtureAndBetType creates a Market struct for a given Fixture and Market i.e. OVER_UNDER_25.
func (m marketBuilder) FixtureAndMarket(f *proto.Fixture, market string) (*Market, error) {
	mark := Market{
		Name: market,
	}

	for _, bookie := range m.bookmakers {
		mk, err := bookie.FixtureAndMarket(f, market)

		if err != nil {
			m.logger.Warnf(
				"Error '%s' building bookmaker odds for fixture '%d' and market '%s'",
				err.Error(),
				f.Id,
				market,
			)
			continue
		}

		mark.Bookmakers = append(mark.Bookmakers, mk)
	}

	return &mark, nil
}

func NewMarketBuilder(book []MarketFactory, log *logrus.Logger) MarketBuilder {
	return &marketBuilder{bookmakers: book, logger: log}
}
