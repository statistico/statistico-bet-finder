package app_test

import (
	"errors"
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
			Markets:    []string{"OVER_UNDER_15", "OVER_UNDER_25"},
			FixtureIDs: []uint64{1329},
		}

		fixture := statistico.Fixture{ID: 1329}

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

	t.Run("logs error and continues book creation if error fetching fixture via fixture client", func(t *testing.T) {
		t.Helper()

		fixtureClient := new(mock.FixtureClient)
		builder := new(mock.MarketBuilder)
		clock := mock.NewFixedClock()
		logger, hook := test.NewNullLogger()

		query := app.BookQuery{
			Markets:    []string{"OVER_UNDER_25"},
			FixtureIDs: []uint64{1329, 45901},
		}

		bookmaker := app.NewBookMaker(fixtureClient, builder, clock, logger)

		fixtureClient.On("FixtureByID", uint64(1329)).Return(&statistico.Fixture{}, errors.New("error occurred"))
		builder.AssertNotCalled(t, "FixtureAndMarket", uint64(1329), "OVER_UNDER_25")

		fixture := statistico.Fixture{ID: 1329}

		fixtureClient.On("FixtureByID", uint64(45901)).Return(&fixture, nil)
		builder.On("FixtureAndMarket", &fixture, "OVER_UNDER_25").Return(&app.Market{}, nil)

		book := bookmaker.CreateBook(&query)

		fixtureClient.AssertExpectations(t)
		builder.AssertExpectations(t)

		assert.Equal(t, 1, len(book.Markets))
		assert.Equal(t, "2019-01-14 11:25:00 +0000 UTC", book.CreatedAt.String())
		assert.Equal(t, "Error 'error occurred' fetching fixture '1329' when creating a book", hook.LastEntry().Message)
	})

	t.Run("logs error and continues book creation if error building market", func(t *testing.T) {
		t.Helper()

		fixtureClient := new(mock.FixtureClient)
		builder := new(mock.MarketBuilder)
		clock := mock.NewFixedClock()
		logger, hook := test.NewNullLogger()

		bookmaker := app.NewBookMaker(fixtureClient, builder, clock, logger)

		query := app.BookQuery{
			Markets:    []string{"OVER_UNDER_15", "OVER_UNDER_25"},
			FixtureIDs: []uint64{1329},
		}

		fixture := statistico.Fixture{ID: 1329}

		fixtureClient.On("FixtureByID", uint64(1329)).Return(&fixture, nil)
		builder.On("FixtureAndMarket", &fixture, "OVER_UNDER_15").Return(&app.Market{}, nil)
		builder.On("FixtureAndMarket", &fixture, "OVER_UNDER_25").Return(&app.Market{}, errors.New("error occurred"))

		book := bookmaker.CreateBook(&query)

		fixtureClient.AssertExpectations(t)
		builder.AssertExpectations(t)

		assert.Equal(t, 1, len(book.Markets))
		assert.Equal(t, "2019-01-14 11:25:00 +0000 UTC", book.CreatedAt.String())
		assert.Equal(t, "error occurred", hook.LastEntry().Message)
	})
}
