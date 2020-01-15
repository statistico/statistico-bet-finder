package betfair

import (
	"github.com/statistico/statistico-bet-finder/internal/statistico"
	bfClient "github.com/statistico/statistico-betfair-go-client"
)

func buildMarketCatalogueRequest(fix statistico.Fixture, betTypes []string) (bfClient.ListMarketCatalogueRequest, error) {
	// ToDo map Fixture Competition ID to associated Competition ID mapping
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

	return request, nil
}

func buildRunnerBookRequest(marketID string, selectionID uint64, priceData []string) bfClient.ListRunnerBookRequest {
	projection := bfClient.PriceProjection{
		PriceData: priceData,
	}

	request := bfClient.ListRunnerBookRequest{
		MarketID:    marketID,
		SelectionID: selectionID,
		PriceProjection: projection,
	}

	return request
}