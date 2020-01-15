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

func (r RunnerCreator) Create(runners []bfClient.RunnerCatalogue, marketID string, data []string) ([]bookmaker.Runner, error) {
	var x []bookmaker.Runner

	for _, runner := range runners {
		request := bfClient.ListRunnerBookRequest{
			MarketID:    marketID,
			SelectionID: runner.SelectionID,
			PriceProjection: bfClient.PriceProjection{
				PriceData: data,
			},
		}

		book, err := r.Client.ListRunnerBook(context.Background(), request)

		if err != nil {
			return nil, err
		}

		if len(book) == 0 || len(book[0].Runners) == 0 {
			return nil, fmt.Errorf("runner book does not exist for Market %s and Selection %d", marketID, runner.SelectionID)
		}

		y := book[0].Runners[0]

		run := bookmaker.Runner{
			Name:        runner.RunnerName,
			Back:        buildPrices(y.EX.AvailableToBack),
			Lay:         buildPrices(y.EX.AvailableToLay),
			SelectionID: runner.SelectionID,
		}

		x = append(x, run)
	}


	return x, nil
}
