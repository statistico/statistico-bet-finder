package betfair

import (
	"context"
	"fmt"
	"github.com/statistico/statistico-bet-finder/internal/statistico/bookmaker"
	bfClient "github.com/statistico/statistico-betfair-go-client"
)

type RunnerFactory struct {
	client  *bfClient.Client
}

func (r RunnerFactory) CreateRunner(selectionID uint64, name, marketID string) (*bookmaker.Runner, error) {
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
	return &RunnerFactory{client:c}
}
