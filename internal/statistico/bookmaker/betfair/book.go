package betfair

import (
	"context"
	"github.com/statistico/statistico-bet-finder/internal/statistico"
	"github.com/statistico/statistico-bet-finder/internal/statistico/bookmaker"
	bfClient "github.com/statistico/statistico-betfair-go-client"
)

type BookCreator struct {
	Client *bfClient.Client
	MarketBuilder
}

func (b BookCreator) CreateForFixture(fix statistico.Fixture, types []string) (*bookmaker.Book, error) {
	book := bookmaker.Book{
		FixtureID: fix.ID,
		Bookmaker: "Betfair",
	}

	request := buildMarketCatalogueRequest(fix, types)

	catalogue, err := b.Client.ListMarketCatalogue(context.Background(), request)

	if err != nil {
		return nil, err
	}

	for _, market := range catalogue {
		// Sanity check catalogue Event matches fixture here - ensure home and away team names match
		m, err := b.MarketBuilder.Build(&market)

		if err != nil {
			return nil, err
		}

		book.Markets = append(book.Markets, *m)
	}

	return &book, nil
}

func buildMarketCatalogueRequest(fix statistico.Fixture, types []string) bfClient.ListMarketCatalogueRequest {
	filter := bfClient.MarketFilter{
		CompetitionIDs:  []string{"10932509"},
		TextQuery:       fix.HomeTeam,
		MarketTypeCodes: types,
	}

	request := bfClient.ListMarketCatalogueRequest{
		Filter:           filter,
		MarketProjection: []string{"EVENT", "RUNNER_DESCRIPTION"},
		MaxResults:       1,
		Sort:             "FIRST_TO_START",
	}

	return request
}
