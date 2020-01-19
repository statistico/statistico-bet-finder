package grpc

import (
	"context"
	"github.com/statistico/statistico-bet-finder/internal/app"
	"github.com/statistico/statistico-bet-finder/internal/app/grpc/proto"
)

type OddsCompilerClient struct {
	client proto.OddsCompilerServiceClient
}

func (o OddsCompilerClient) GetOverUnderGoalsForFixture(fixtureID uint64, market string) (*app.Market, error) {
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

func convertResponseToMarket(resp *proto.OverUnderGoalsResponse) *app.Market {
	market := app.Market{
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
