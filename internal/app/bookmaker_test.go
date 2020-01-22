package app_test

import (
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/statistico/statistico-bet-finder/internal/app"
	"github.com/statistico/statistico-bet-finder/internal/app/mock"
	"github.com/statistico/statistico-bet-finder/internal/app/statistico"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookMaker_CreateBook(t *testing.T) {
	t.Run("returns a Book struct containing statistico and bookmaker markets", func(t *testing.T) {
		t.Helper()

		fixtureClient := new(mock.FixtureClient)
		builder := new(mock.MarketBuilder)
		clock := mock.NewFixedClock()
		logger, hook := test.NewNullLogger()

		bookmaker := app.NewBookMaker(fixtureClient, builder, clock, logger)

		query := app.BookQuery{
			Markets:   []string{"OVER_UNDER_15", "OVER_UNDER_25"},
			FixtureIDs: []uint64{1329},
		}

		fixture := statistico.Fixture{ID:1329}

		fixtureClient.On("FixtureByID", uint64(1329)).Return(&fixture, nil)
		builder.On("FixtureAndMarket", &fixture, "OVER_UNDER_15").Return(&app.Market{}, nil)
		builder.On("FixtureAndMarket", &fixture, "OVER_UNDER_25").Return(&app.Market{}, nil)

		book := bookmaker.CreateBook(&query)

		fixtureClient.AssertExpectations(t)
		builder.AssertExpectations(t)

		assert.Equal(t, 2, len(book.Markets))
		assert.Equal(t, "2019-01-14 11:25:00 +0000 UTC", book.CreatedAt.String())
		assert.Nil(t, hook.LastEntry())
	})
}
