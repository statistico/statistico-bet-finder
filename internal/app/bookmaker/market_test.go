package bookmaker_test

import (
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/statistico/statistico-price-finder/internal/app/bookmaker"
	"github.com/statistico/statistico-price-finder/internal/app/grpc/proto"
	"github.com/statistico/statistico-price-finder/internal/app/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarketBuilder_FixtureAndBetType(t *testing.T) {
	t.Run("returns a hydrated market struct", func(t *testing.T) {
		t.Helper()

		factory := new(mock.MarketFactory)
		books := []bookmaker.MarketFactory{factory}
		logger, hook := test.NewNullLogger()

		builder := bookmaker.NewMarketBuilder(books, logger)

		fixture := proto.Fixture{
			Id:          45381,
			Competition: &proto.Competition{Id: 42},
		}

		m := bookmaker.SubMarket{
			ID:        "1.14567",
			Bookmaker: "Betfair",
			Runners: []bookmaker.Runner{
				{
					Name:        "OVER",
					SelectionID: 13459,
					Back: []bookmaker.Price{
						{
							Size:  145.41,
							Price: 1.79,
						},
					},
				},
			},
		}

		factory.On("FixtureAndMarket", &fixture, "OVER_UNDER_25").Return(&m, nil)

		market, err := builder.FixtureAndMarket(&fixture, "OVER_UNDER_25")

		if err != nil {
			t.Fatalf("Error building market expected nil got %s", err)
		}

		if market == nil {
			t.Fatalf("Error building market expected struct got nil")
		}

		factory.AssertExpectations(t)
		assert.Equal(t, "OVER_UNDER_25", market.Name)
		assert.Equal(t, 1, len(market.Bookmakers))
		assert.Nil(t, hook.LastEntry())
	})

	t.Run("error is logged but execution continues if error creating bookmaker market", func(t *testing.T) {
		t.Helper()

		factory := new(mock.MarketFactory)
		books := []bookmaker.MarketFactory{factory}
		logger, hook := test.NewNullLogger()

		builder := bookmaker.NewMarketBuilder(books, logger)

		fixture := proto.Fixture{
			Id:          45381,
			Competition: &proto.Competition{Id: 42},
		}

		factory.On("FixtureAndMarket", &fixture, "OVER_UNDER_25").Return(&bookmaker.SubMarket{}, errors.New("error occurred"))

		market, err := builder.FixtureAndMarket(&fixture, "OVER_UNDER_25")

		if err != nil {
			t.Fatalf("Error building market expected nil got %s", err)
		}

		factory.AssertExpectations(t)
		assert.Equal(t, 0, len(market.Bookmakers))
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.WarnLevel, hook.LastEntry().Level)
		assert.Equal(
			t,
			"Error 'error occurred' building bookmaker odds for fixture '45381' and market 'OVER_UNDER_25'",
			hook.LastEntry().Message,
		)
	})
}
