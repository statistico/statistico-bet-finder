package betfair

import (
	"context"
	"fmt"
	"github.com/statistico/statistico-bet-finder/internal/statistico"
	"github.com/statistico/statistico-bet-finder/internal/statistico/bookmaker"
	bfClient "github.com/statistico/statistico-betfair-go-client"
)

const betfair = "Betfair"

// MarketFactory populates BetFair markets meeting specific criteria
type MarketFactory struct {
	client  *bfClient.Client
	runner  bookmaker.RunnerFactory
}

// FixtureAndBetType creates a BetFair bookmaker.Market struct for a specific Fixture and Bet Type
func (b MarketFactory) FixtureAndBetType(fix statistico.Fixture, betType string) (*bookmaker.Market, error) {
	request, err := buildMarketCatalogueRequest(fix, []string{betType})

	if err != nil {
		return nil, err
	}

	market, err := b.parseMarket(request, fix.ID, betType)

	if err != nil {
		return nil, err
	}

	if !fixtureMatchesEvent(fix, market.Event) {
		return nil, fmt.Errorf(
			"event '%s' returned by betfair client does not match fixture '%s'",
			market.Event.Name,
			fmt.Sprintf("%s v %s", fix.HomeTeam, fix.AwayTeam),
		)
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
		run, err := b.runner.CreateRunner(runner.SelectionID, market.MarketID, runner.RunnerName)

		if err != nil {
			return nil, err
		}

		m.Runners = append(m.Runners, *run)
	}

	return &m, nil
}

func (b MarketFactory) parseMarket(req *bfClient.ListMarketCatalogueRequest, fixID uint64, betType string) (*bfClient.MarketCatalogue, error) {
	catalogue, err := b.client.ListMarketCatalogue(context.Background(), *req)

	if err != nil {
		return nil, err
	}

	if len(catalogue) == 0 {
		return nil, fmt.Errorf("no market returned for fixture %d and bet type %s", fixID, betType)
	}

	return &catalogue[0], nil
}
