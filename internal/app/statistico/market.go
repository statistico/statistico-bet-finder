package statistico

import (
	"github.com/statistico/statistico-price-finder/internal/app/grpc"
	"github.com/statistico/statistico-price-finder/internal/app/grpc/proto"
)

type MarketBuilder interface {
	FixtureAndMarket(f *proto.Fixture, name string) (*Market, error)
}

type marketBuilder struct {
	oddsClient grpc.OddsCompilerClient
}

func (m marketBuilder) FixtureAndMarket(f *proto.Fixture, name string) (*Market, error) {
	market := Market{Name: name}

	response, err := m.oddsClient.EventMarket(uint64(f.Id), name)

	if err != nil {
		return &market, err
	}

	for _, odds := range response.Odds {
		o := Runner{
			Name:  odds.Selection,
			Price: odds.Price,
		}

		market.Runners = append(market.Runners, &o)
	}

	return &market, nil
}

func NewMarketBuilder(o grpc.OddsCompilerClient) MarketBuilder {
	return &marketBuilder{oddsClient: o}
}
