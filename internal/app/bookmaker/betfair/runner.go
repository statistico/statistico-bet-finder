package betfair

import (
	"context"
	"fmt"
	bfClient "github.com/statistico/statistico-betfair-go-client"
	"github.com/statistico/statistico-price-finder/internal/app/bookmaker"
)

// RunnerFactory is a wrapper around the BetFair API Client that is responsible for creating new
// bookmaker.Runner struct
type RunnerFactory struct {
	client *bfClient.Client
}

// CreateRunner uses the arguments provided to call the BetFair API and parse the response into a
// bookmaker.Runner struct
func (r RunnerFactory) CreateRunner(selectionID uint64, marketID, name string) (*bookmaker.Runner, error) {
	request := buildRunnerBookRequest(marketID, selectionID, []string{"EX_BEST_OFFERS"})

	run, err := r.parseRunner(request)

	if err != nil {
		return nil, err
	}

	runner := bookmaker.Runner{
		Name:        name,
		Back:        buildPrices(run.EX.AvailableToBack),
		Lay:         buildPrices(run.EX.AvailableToLay),
		SelectionID: selectionID,
	}

	return &runner, nil
}

func (r RunnerFactory) parseRunner(req bfClient.ListRunnerBookRequest) (*bfClient.Runner, error) {
	book, err := r.client.ListRunnerBook(context.Background(), req)

	if err != nil {
		return nil, err
	}

	if len(book) == 0 || len(book[0].Runners) == 0 {
		return nil, fmt.Errorf(
			"runner book does not exist for Market '%s' and Selection '%d'",
			req.MarketID,
			req.SelectionID,
		)
	}

	return &book[0].Runners[0], nil
}

func NewRunnerFactory(c *bfClient.Client) *RunnerFactory {
	return &RunnerFactory{client: c}
}
