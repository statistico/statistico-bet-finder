package app_test

import (
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/statistico/statistico-bet-finder/internal/app"
	"github.com/statistico/statistico-bet-finder/internal/app/bookmaker"
	"github.com/statistico/statistico-bet-finder/internal/app/mock"
	"github.com/statistico/statistico-bet-finder/internal/app/statistico"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarketBuilder_FixtureAndBetType(t *testing.T) {
	t.Run("returns a hydrated market struct", func(t *testing.T) {
		t.Helper()

		oddsClient := new(mock.OddsCompilerClient)
		factory := new(mock.MarketFactory)
		books := []bookmaker.MarketFactory{factory}
		logger, hook := test.NewNullLogger()

		builder := app.NewMarketBuilder(oddsClient, books, logger)

		fixture := statistico.Fixture{
			ID:            45381,
			CompetitionID: 42,
		}

		odds := statistico.Market{
			FixtureID:45381,
			Runners: []statistico.Runner{
				{
					Name:  "OVER",
					Price: 1.59,
				},
			},
		}

		m := bookmaker.Market{
			ID:        "1.14567",
			FixtureID: 45381,
			Bookmaker: "Betfair",
			Name:      "OVER_UNDER_25",
			Runners:   []bookmaker.Runner{
				{
					Name:        "OVER",
					SelectionID: 13459,
					Back:        []bookmaker.Price{
						{
							Size:  145.41,
							Price: 1.79,
						},
					},
				},
			},
		}

		oddsClient.On("GetOverUnderGoalsForFixture", uint64(45381), "OVER_UNDER_25").Return(&odds, nil)
		factory.On("FixtureAndMarket", fixture, "OVER_UNDER_25").Return(&m, nil)

		market := builder.FixtureAndMarket(&fixture, "OVER_UNDER_25")

		if market == nil {
			t.Fatalf("Error building market expected struct got nil")
		}

		oddsClient.AssertExpectations(t)
		factory.AssertExpectations(t)
		assert.Equal(t, uint64(45381), market.FixtureID)
		assert.Equal(t, uint64(45381), market.FixtureID)
		assert.Nil(t, hook.LastEntry())
	})

	t.Run("error is logged and nil returned if error returned from odds client", func(t *testing.T) {
		t.Helper()

		oddsClient := new(mock.OddsCompilerClient)
		factory := new(mock.MarketFactory)
		books := []bookmaker.MarketFactory{factory}
		logger, hook := test.NewNullLogger()

		builder := app.NewMarketBuilder(oddsClient, books, logger)

		fixture := statistico.Fixture{
			ID:            45381,
			CompetitionID: 42,
		}

		oddsClient.On("GetOverUnderGoalsForFixture", uint64(45381), "OVER_UNDER_25").Return(&statistico.Market{},errors.New("error occurred"))
		factory.AssertNotCalled(t, "FixtureAndMarket", fixture, "OVER_UNDER_25")

		market := builder.FixtureAndMarket(&fixture, "OVER_UNDER_25")

		if market != nil {
			t.Fatalf("Error building market expected nil got %+v", market)
		}

		oddsClient.AssertExpectations(t)
		factory.AssertExpectations(t)
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.WarnLevel, hook.LastEntry().Level)
		assert.Equal(
			t,
			"Error 'error occurred' building statistico odds for fixture '45381' and market 'OVER_UNDER_25'",
			hook.LastEntry().Message,
		)
	})

	t.Run("error is logged and nil returned if error returned from bookmaker factory", func(t *testing.T) {
		t.Helper()

		oddsClient := new(mock.OddsCompilerClient)
		factory := new(mock.MarketFactory)
		books := []bookmaker.MarketFactory{factory}
		logger, hook := test.NewNullLogger()

		builder := app.NewMarketBuilder(oddsClient, books, logger)

		fixture := statistico.Fixture{
			ID:            45381,
			CompetitionID: 42,
		}

		odds := statistico.Market{
			FixtureID:45381,
			Runners: []statistico.Runner{
				{
					Name:  "OVER",
					Price: 1.59,
				},
			},
		}

		oddsClient.On("GetOverUnderGoalsForFixture", uint64(45381), "OVER_UNDER_25").Return(&odds, nil)
		factory.On("FixtureAndMarket", fixture, "OVER_UNDER_25").Return(&bookmaker.Market{}, errors.New("error occurred"))

		market := builder.FixtureAndMarket(&fixture, "OVER_UNDER_25")

		if market != nil {
			t.Fatalf("Error building market expected nil got %+v", market)
		}

		oddsClient.AssertExpectations(t)
		factory.AssertExpectations(t)
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.WarnLevel, hook.LastEntry().Level)
		assert.Equal(
			t,
			"Error 'error occurred' building bookmaker odds for fixture '45381' and market 'OVER_UNDER_25'",
			hook.LastEntry().Message,
		)
	})
}
