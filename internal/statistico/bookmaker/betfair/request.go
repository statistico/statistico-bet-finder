package betfair

import (
	"fmt"
	"github.com/statistico/statistico-bet-finder/internal/statistico"
	bfClient "github.com/statistico/statistico-betfair-go-client"
)

func buildMarketCatalogueRequest(fix statistico.Fixture, betTypes []string) (*bfClient.ListMarketCatalogueRequest, error) {
	compID, err := competitionMapping(fix.CompetitionID)

	if err != nil {
		return nil, err
	}

	filter := bfClient.MarketFilter{
		CompetitionIDs:  []string{compID},
		TextQuery:       fix.HomeTeam,
		MarketTypeCodes: betTypes,
	}

	request := bfClient.ListMarketCatalogueRequest{
		Filter:           filter,
		MarketProjection: []string{"EVENT", "RUNNER_DESCRIPTION"},
		MaxResults:       len(betTypes),
		Sort:             "FIRST_TO_START",
	}

	return &request, nil
}

func buildRunnerBookRequest(marketID string, selectionID uint64, priceData []string) bfClient.ListRunnerBookRequest {
	projection := bfClient.PriceProjection{
		PriceData: priceData,
	}

	request := bfClient.ListRunnerBookRequest{
		MarketID:        marketID,
		SelectionID:     selectionID,
		PriceProjection: projection,
	}

	return request
}

func competitionMapping(id uint64) (string, error) {
	competitions := map[uint64]string{
		12968: "10932509",
	}

	if val, ok := competitions[id]; ok {
		return val, nil
	}

	return "", fmt.Errorf("competition ID %d is not supported", id)
}
