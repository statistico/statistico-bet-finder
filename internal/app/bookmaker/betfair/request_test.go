package betfair

import (
	"github.com/statistico/statistico-bet-finder/internal/app/statistico"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_buildMarketCatalogueRequest(t *testing.T) {
	t.Run("returns new betfair market catalogue request", func(t *testing.T) {
		fix := statistico.Fixture{
			ID:            99,
			CompetitionID: 8,
			HomeTeam:      "West Ham United",
			AwayTeam:      "Chelsea",
			Date:          time.Time{},
		}

		types := []string{"OVER_UNDER_25"}

		request, err := buildMarketCatalogueRequest(fix, types)

		if err != nil {
			t.Fatalf("Error building request expected nil got %s", err)
		}

		filter := request.Filter

		assert.Equal(t, []string{"10932509"}, filter.CompetitionIDs)
		assert.Equal(t, "West Ham", filter.TextQuery)
		assert.Equal(t, types, filter.MarketTypeCodes)
		assert.Equal(t, []string{"EVENT", "RUNNER_DESCRIPTION"}, request.MarketProjection)
		assert.Equal(t, 10, request.MaxResults)
		assert.Equal(t, "FIRST_TO_START", request.Sort)
	})
}

func Test_buildRunnerBookRequest(t *testing.T) {
	t.Run("returns new betfair runner book request", func(t *testing.T) {
		request := buildRunnerBookRequest("1.16879", 47892, []string{"EX_BEST_OFFERS"})

		assert.Equal(t, "1.16879", request.MarketID)
		assert.Equal(t, uint64(47892), request.SelectionID)
		assert.Equal(t, []string{"EX_BEST_OFFERS"}, request.PriceProjection.PriceData)
	})
}
