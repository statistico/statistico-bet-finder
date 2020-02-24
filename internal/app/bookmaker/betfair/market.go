package betfair

import (
	"context"
	"fmt"
	"github.com/statistico/statistico-bet-finder/internal/app/bookmaker"
	"github.com/statistico/statistico-bet-finder/internal/app/statistico"
	bfClient "github.com/statistico/statistico-betfair-go-client"
)

const betfair = "Betfair"

// MarketFactory populates BetFair markets meeting specific criteria
type MarketFactory struct {
	client *bfClient.Client
	runner bookmaker.RunnerFactory
}

// FixtureAndBetType creates a BetFair bookmaker.Market struct for a specific Fixture and Market
func (b MarketFactory) FixtureAndMarket(fix statistico.Fixture, market string) (*bookmaker.Market, error) {
	request, err := buildMarketCatalogueRequest(fix, []string{market})

	if err != nil {
		return nil, err
	}

	catalogues, err := b.parseMarketCatalogue(request, fix.ID, market)

	if err != nil {
		return nil, err
	}

	catalogue, err := parseCatalogue(catalogues, &fix)

	if err != nil {
		return nil, err
	}

	m := bookmaker.Market{
		ID:        catalogue.MarketID,
		Bookmaker: betfair,
		Runners:   nil,
	}

	for _, runner := range catalogue.Runners {
		run, err := b.runner.CreateRunner(runner.SelectionID, catalogue.MarketID, runner.RunnerName)

		if err != nil {
			return nil, err
		}

		m.Runners = append(m.Runners, *run)
	}

	return &m, nil
}

func (b MarketFactory) parseMarketCatalogue(req *bfClient.ListMarketCatalogueRequest, fixID uint64, betType string) ([]bfClient.MarketCatalogue, error) {
	catalogue, err := b.client.ListMarketCatalogue(context.Background(), *req)

	if err != nil {
		return nil, err
	}

	if len(catalogue) == 0 {
		return nil, fmt.Errorf("no market returned for fixture %d and bet type %s", fixID, betType)
	}

	return catalogue, nil
}

func NewMarketFactory(c *bfClient.Client, r bookmaker.RunnerFactory) *MarketFactory {
	return &MarketFactory{client: c, runner: r}
}
