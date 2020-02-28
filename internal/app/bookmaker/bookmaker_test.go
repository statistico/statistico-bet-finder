package bookmaker_test

import (
	"errors"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/statistico/statistico-price-finder/internal/app"
	"github.com/statistico/statistico-price-finder/internal/app/bookmaker"
	"github.com/statistico/statistico-price-finder/internal/app/grpc/proto"
	"github.com/statistico/statistico-price-finder/internal/app/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookMaker_CreateBookmakerBook(t *testing.T) {
	t.Run("returns a Book struct containing bookmaker markets", func(t *testing.T) {
		t.Helper()

		fixtureClient := new(mock.FixtureClient)
		builder := new(mock.MarketBuilder)
		clock := mock.NewFixedClock()
		logger, hook := test.NewNullLogger()

		bm := bookmaker.NewBookMaker(fixtureClient, builder, clock, logger)

		query := app.BookQuery{
			Markets: []string{"OVER_UNDER_15", "OVER_UNDER_25"},
			EventID: uint64(1329),
		}

		fixture := proto.Fixture{Id: 1329}

		fixtureClient.On("FixtureByID", uint64(1329)).Return(&fixture, nil)
		builder.On("FixtureAndMarket", &fixture, "OVER_UNDER_15").Return(&bookmaker.Market{}, nil)
		builder.On("FixtureAndMarket", &fixture, "OVER_UNDER_25").Return(&bookmaker.Market{}, nil)

		book, err := bm.CreateBook(&query)

		if err != nil {
			t.Fatalf("Expected nil got %s", err.Error())
		}

		fixtureClient.AssertExpectations(t)
		builder.AssertExpectations(t)

		assert.Equal(t, 2, len(book.Markets))
		assert.Equal(t, "2019-01-14 11:25:00 +0000 UTC", book.CreatedAt.String())
		assert.Nil(t, hook.LastEntry())
	})

	t.Run("returns error if error fetching fixture via fixture client", func(t *testing.T) {
		t.Helper()

		fixtureClient := new(mock.FixtureClient)
		builder := new(mock.MarketBuilder)
		clock := mock.NewFixedClock()
		logger, _ := test.NewNullLogger()

		query := app.BookQuery{
			Markets: []string{"OVER_UNDER_25"},
			EventID: uint64(1329),
		}

		bm := bookmaker.NewBookMaker(fixtureClient, builder, clock, logger)

		fixtureClient.On("FixtureByID", uint64(1329)).Return(&proto.Fixture{}, errors.New("fixture not found"))
		builder.AssertNotCalled(t, "FixtureAndMarket", uint64(1329), "OVER_UNDER_25")

		_, err := bm.CreateBook(&query)

		if err == nil {
			t.Fatal("Expected error got nil")
		}

		fixtureClient.AssertExpectations(t)
		builder.AssertExpectations(t)
	})

	t.Run("logs warning and continues book creation if error building market", func(t *testing.T) {
		t.Helper()

		fixtureClient := new(mock.FixtureClient)
		builder := new(mock.MarketBuilder)
		clock := mock.NewFixedClock()
		logger, hook := test.NewNullLogger()

		bm := bookmaker.NewBookMaker(fixtureClient, builder, clock, logger)

		query := app.BookQuery{
			Markets: []string{"OVER_UNDER_15", "OVER_UNDER_25"},
			EventID: uint64(1329),
		}

		fixture := proto.Fixture{Id: 1329}

		fixtureClient.On("FixtureByID", uint64(1329)).Return(&fixture, nil)
		builder.On("FixtureAndMarket", &fixture, "OVER_UNDER_15").Return(&bookmaker.Market{}, nil)
		builder.On("FixtureAndMarket", &fixture, "OVER_UNDER_25").Return(&bookmaker.Market{}, errors.New("error occurred"))

		book, err := bm.CreateBook(&query)

		if err != nil {
			t.Fatalf("Expected nil got %s", err.Error())
		}

		fixtureClient.AssertExpectations(t)
		builder.AssertExpectations(t)

		assert.Equal(t, 1, len(book.Markets))
		assert.Equal(t, "2019-01-14 11:25:00 +0000 UTC", book.CreatedAt.String())
		assert.Equal(t, "Error building market for event 1329: error occurred", hook.LastEntry().Message)
	})
}
