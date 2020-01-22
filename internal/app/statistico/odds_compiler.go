package statistico

import (
	"context"
	"github.com/statistico/statistico-bet-finder/internal/app/grpc/proto"
)

// OddsCompilerClient is a wrapper around the Statistico Odds Compiler service
type OddsCompilerClient interface {
	GetOverUnderGoalsForFixture(fixtureID uint64, market string) (*Market, error)
}

type gRPCOddsCompilerClient struct {
	client proto.OddsCompilerServiceClient
}

// GetOverUnderGoalsForFixture returns a market struct containing data for the requested fixture and market
func (o gRPCOddsCompilerClient) GetOverUnderGoalsForFixture(fixtureID uint64, market string) (*Market, error) {
	request := proto.OverUnderRequest{
		FixtureId: fixtureID,
		Market:    market,
	}

	response, err := o.client.GetOverUnderGoalsForFixture(context.Background(), &request)

	if err != nil {
		return nil, err
	}

	return convertResponseToMarket(response), nil
}

func convertResponseToMarket(resp *proto.OverUnderGoalsResponse) *Market {
	market := Market{
		FixtureID: resp.FixtureId,
		Name:      resp.Market,
	}

	for _, odds := range resp.Odds {
		run := Runner{
			Name:  odds.Selection,
			Price: odds.Price,
		}

		market.Runners = append(market.Runners, run)
	}

	return &market
}

func NewGRPCOddsCompilerClient(c proto.OddsCompilerServiceClient) OddsCompilerClient {
	return &gRPCOddsCompilerClient{client: c}
}
