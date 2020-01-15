package betfair

import (
	"context"
	"github.com/statistico/statistico-bet-finder/internal/statistico"
	"github.com/statistico/statistico-bet-finder/internal/statistico/bookmaker"
	bfClient "github.com/statistico/statistico-betfair-go-client"
)

type BookCreator struct {
	Client *bfClient.Client
	RunnerCreator
}

func (b BookCreator) FixtureAndBetType(fix statistico.Fixture, betType string) (*bookmaker.Market, error) {
	request := buildMarketCatalogueRequest(fix, []string{betType})

	catalogue, err := b.Client.ListMarketCatalogue(context.Background(), request)

	if err != nil {
		return nil, err
	}

	if len(catalogue) == 0 {
		return nil, nil
	}

	market := catalogue[0]

	m := bookmaker.Market{
		ID:        market.MarketID,
		FixtureID: fix.ID,
		Bookmaker: "Betfair",
		Name:      market.MarketName,
		BetType:   betType,
		Runners:   nil,
	}

	runners, err := b.RunnerCreator.Create(market.Runners, market.MarketID, []string{"EX_BEST_OFFERS"})

	if err != nil {
		return nil, err
	}

	m.Runners = runners

	return &m, nil
}

func buildMarketCatalogueRequest(fix statistico.Fixture, betTypes []string) bfClient.ListMarketCatalogueRequest {
	filter := bfClient.MarketFilter{
		CompetitionIDs:  []string{"10932509"},
		TextQuery:       fix.HomeTeam,
		MarketTypeCodes: betTypes,
	}

	request := bfClient.ListMarketCatalogueRequest{
		Filter:           filter,
		MarketProjection: []string{"EVENT", "RUNNER_DESCRIPTION"},
		MaxResults:       1,
		Sort:             "FIRST_TO_START",
	}

	return request
}
