package betfair

import (
	"context"
	"fmt"
	bfClient "github.com/statistico/statistico-betfair-go-client"
	"github.com/statistico/statistico-price-finder/internal/app/bookmaker"
	"github.com/statistico/statistico-price-finder/internal/app/grpc/proto"
)

const betfair = "Betfair"

// MarketFactory populates BetFair markets meeting specific criteria
type MarketFactory struct {
	client *bfClient.Client
	runner bookmaker.RunnerFactory
}

// FixtureAndBetType creates a BetFair bookmaker.Market struct for a specific Fixture and Market
func (b MarketFactory) FixtureAndMarket(fix *proto.Fixture, market string) (*bookmaker.SubMarket, error) {
	request, err := buildMarketCatalogueRequest(fix, []string{market})

	if err != nil {
		return nil, err
	}

	catalogues, err := b.parseMarketCatalogue(request, uint64(fix.Id), market)

	if err != nil {
		return nil, err
	}

	catalogue, err := parseCatalogue(catalogues, fix)

	if err != nil {
		return nil, err
	}

	m := bookmaker.SubMarket{
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
