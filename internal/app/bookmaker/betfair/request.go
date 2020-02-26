package betfair

import (
	bfClient "github.com/statistico/statistico-betfair-go-client"
	"github.com/statistico/statistico-price-finder/internal/app/grpc/proto"
)

func buildMarketCatalogueRequest(fix *proto.Fixture, betTypes []string) (*bfClient.ListMarketCatalogueRequest, error) {
	compID, err := parseCompetitionMapping(uint64(fix.Competition.Id))

	if err != nil {
		return nil, err
	}

	filter := bfClient.MarketFilter{
		CompetitionIDs:  []string{compID},
		TextQuery:       parseTeamMapping(fix.HomeTeam.Name),
		MarketTypeCodes: betTypes,
	}

	request := bfClient.ListMarketCatalogueRequest{
		Filter:           filter,
		MarketProjection: []string{"EVENT", "RUNNER_DESCRIPTION"},
		MaxResults:       10,
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
