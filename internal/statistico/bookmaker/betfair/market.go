package betfair

import (
	"context"
	"fmt"
	"github.com/statistico/statistico-bet-finder/internal/statistico"
	"github.com/statistico/statistico-bet-finder/internal/statistico/bookmaker"
	bfClient "github.com/statistico/statistico-betfair-go-client"
)

const betfair = "Betfair"

type MarketFactory struct {
	Client  *bfClient.Client
	RunnerFactory
}

func (b MarketFactory) FixtureAndBetType(fix statistico.Fixture, betType string) (*bookmaker.Market, error) {
	request, err := buildMarketCatalogueRequest(fix, []string{betType})

	if err != nil {
		return nil, err
	}

	market, err := b.parseMarket(request)

	if err != nil {
		return nil, err
	}

	if !fixtureMatchesEvent(fix, market.Event) {
		return nil, fmt.Errorf("rvent %+v returned by betfair client does not match fixture %+v", market.Event, fix)
	}

	m := bookmaker.Market{
		ID:        market.MarketID,
		FixtureID: fix.ID,
		Bookmaker: betfair,
		Name:      market.MarketName,
		BetType:   betType,
		Runners:   nil,
	}

	for _, runner := range market.Runners {
		run, err := b.CreateRunner(runner.SelectionID, market.MarketID, runner.RunnerName)

		if err != nil {
			return nil, err
		}

		m.Runners = append(m.Runners, *run)
	}

	return &m, nil
}

func (b MarketFactory) parseMarket(req *bfClient.ListMarketCatalogueRequest) (*bfClient.MarketCatalogue, error) {
	catalogue, err := b.Client.ListMarketCatalogue(context.Background(), *req)

	if err != nil {
		return nil, err
	}

	if len(catalogue) == 0 {
		return nil, nil
	}

	return &catalogue[0], nil
}
