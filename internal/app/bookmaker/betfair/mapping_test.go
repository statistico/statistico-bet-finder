package betfair

import (
	bfClient "github.com/statistico/statistico-betfair-go-client"
	"github.com/statistico/statistico-price-finder/internal/app/grpc/proto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fixtureMatchesEvent(t *testing.T) {
	t.Run("returns true if fixture matches event", func(t *testing.T) {
		t.Helper()

		home := proto.Team{Name: "West Ham United"}
		away := proto.Team{Name: "AFC Bournemouth"}

		fixture := proto.Fixture{
			HomeTeam: &home,
			AwayTeam: &away,
		}

		event := bfClient.Event{
			Name: "West Ham v Bournemouth",
		}

		assert.True(t, fixtureMatchesEvent(&fixture, event))
	})

	t.Run("returns false if fixture does not match event", func(t *testing.T) {
		t.Helper()

		home := proto.Team{Name: "West Ham United"}
		away := proto.Team{Name: "Chelsea"}

		fixture := proto.Fixture{
			HomeTeam: &home,
			AwayTeam: &away,
		}

		event := bfClient.Event{
			Name: "West Ham v Bournemouth",
		}

		assert.False(t, fixtureMatchesEvent(&fixture, event))
	})
}
