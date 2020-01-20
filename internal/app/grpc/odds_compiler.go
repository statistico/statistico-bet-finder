package grpc

import (
	"context"
	"github.com/statistico/statistico-bet-finder/internal/app"
	"github.com/statistico/statistico-bet-finder/internal/app/grpc/proto"
)

type OddsCompilerClient interface {
	GetOverUnderGoalsForFixture(fixtureID uint64, market string) (*app.StatisticoMarket, error)
}

// OddsCompilerClient is a wrapper around the Statistico Odds Compiler service
type oddsCompilerClient struct {
	client proto.OddsCompilerServiceClient
}

// GetOverUnderGoalsForFixture returns a market struct containing data for the requested fixture and market
func (o oddsCompilerClient) GetOverUnderGoalsForFixture(fixtureID uint64, market string) (*app.StatisticoMarket, error) {
	request := proto.OverUnderRequest{
		FixtureId:            fixtureID,
		Market:               market,
	}

	response, err := o.client.GetOverUnderGoalsForFixture(context.Background(), &request)

	if err != nil {
		return nil, err
	}

	return convertResponseToMarket(response), nil
}

func convertResponseToMarket(resp *proto.OverUnderGoalsResponse) *app.StatisticoMarket {
	market := app.StatisticoMarket{
		FixtureID: resp.FixtureId,
		Name:      resp.Market,
	}

	for _, odds := range resp.Odds {
		run := app.Runner{
			Name:  odds.Selection,
			Price: odds.Price,
		}

		market.Runners = append(market.Runners, run)
	}

	return &market
}

func NewOddsCompilerClient(c proto.OddsCompilerServiceClient) OddsCompilerClient {
	return &oddsCompilerClient{client:c}
}
