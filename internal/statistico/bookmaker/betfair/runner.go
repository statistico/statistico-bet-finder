package betfair

import (
	"context"
	"fmt"
	"github.com/statistico/statistico-bet-finder/internal/statistico/bookmaker"
	bfClient "github.com/statistico/statistico-betfair-go-client"
)


func createRunners(client *bfClient.Client, runners []bfClient.RunnerCatalogue, marketID string) ([]bookmaker.Runner, error) {
	var run []bookmaker.Runner

	for _, runner := range runners {
		request := buildRunnerBookRequest(marketID, runner.SelectionID, []string{"EX_BEST_OFFERS"})

		y, err := parseRunner(client, request)

		if err != nil {
			return nil, err
		}

		r := bookmaker.Runner{
			Name:        runner.RunnerName,
			Back:        buildPrices(y.EX.AvailableToBack),
			Lay:         buildPrices(y.EX.AvailableToLay),
			SelectionID: runner.SelectionID,
		}

		run = append(run, r)
	}

	return run, nil
}

func parseRunner(client *bfClient.Client, req bfClient.ListRunnerBookRequest) (*bfClient.Runner, error) {
	book, err := client.ListRunnerBook(context.Background(), req)

	if err != nil {
		return nil, err
	}

	if len(book) == 0 || len(book[0].Runners) == 0 {
		return nil, fmt.Errorf(
			"runner book does not exist for Market %s and Selection %d",
			req.MarketID,
			req.SelectionID,
		)
	}

	return &book[0].Runners[0], nil
}
