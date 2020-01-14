package betfair

import (
	"context"
	"fmt"
	"github.com/statistico/statistico-bet-finder/internal/statistico/bookmaker"
	bfClient "github.com/statistico/statistico-betfair-go-client"
)

type RunnerCreator struct {
	Client *bfClient.Client
}

func (r RunnerCreator) Create(marketID string, runner bfClient.RunnerCatalogue) (*bookmaker.Runner, error) {
	request := bfClient.ListRunnerBookRequest{
		MarketID:                      marketID,
		SelectionID:                   runner.SelectionID,
		PriceProjection: bfClient.PriceProjection{
			PriceData:             []string{"EX_BEST_OFFERS"},
		},
	}

	book, err := r.Client.ListRunnerBook(context.Background(), request)

	if err != nil {
		return nil, err
	}

	if len(book) == 0 || len(book[0].Runners) == 0 {
		return nil, fmt.Errorf("runner book does not exist for Market %s and Selection %d", marketID, runner.SelectionID)
	}

	x := book[0].Runners[0]

	run := bookmaker.Runner{
		Name:        runner.RunnerName,
		Price:       x.LastPriceTraded,
		SelectionID: runner.SelectionID,
	}

	return &run, nil
}
