package betfair

import (
	"github.com/statistico/statistico-bet-finder/internal/app/statistico"
	bfClient "github.com/statistico/statistico-betfair-go-client"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fixtureMatchesEvent(t *testing.T) {
	t.Run("returns true if fixture matches event", func(t *testing.T) {
		t.Helper()

		fixture := statistico.Fixture{
			HomeTeam: "West Ham United",
			AwayTeam: "AFC Bournemouth",
		}

		event := bfClient.Event{
			Name: "West Ham v Bournemouth",
		}

		assert.True(t, fixtureMatchesEvent(&fixture, event))
	})

	t.Run("returns false if fixture does not match event", func(t *testing.T) {
		t.Helper()

		fixture := statistico.Fixture{
			HomeTeam: "West Ham United",
			AwayTeam: "Chelsea",
		}

		event := bfClient.Event{
			Name: "West Ham v Bournemouth",
		}

		assert.False(t, fixtureMatchesEvent(&fixture, event))
	})
}
