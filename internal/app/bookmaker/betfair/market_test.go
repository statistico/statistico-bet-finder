package betfair

import (
	"errors"
	"github.com/statistico/statistico-bet-finder/internal/app/bookmaker"
	"github.com/statistico/statistico-bet-finder/internal/app/mock"
	"github.com/statistico/statistico-bet-finder/internal/app/statistico"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarketFactory_FixtureAndBetType(t *testing.T) {
	t.Run("returns a market struct for a given fixture and bet type", func(t *testing.T) {
		t.Helper()

		server := mock.ResponseServer(t, marketCatalogueResponse, 200, "https://api.betfair.com/test/listMarketCatalogue/")

		client := mock.BetfairClient(server)
		runners := new(mock.RunnerFactory)

		factory := MarketFactory{
			client: client,
			runner: runners,
		}

		fixture := newStatisticoFixture("Liverpool", "Manchester United", 148270, 8)

		runnerOne := newBookmakerRunner("Under 2.5 Goals", 47972)
		runnerTwo := newBookmakerRunner("Under 2.5 Goals", 47973)

		runners.On("CreateRunner", uint64(47972), "1.167019590", "Under 2.5 Goals").Return(runnerOne, nil)
		runners.On("CreateRunner", uint64(47973), "1.167019590", "Over 2.5 Goals").Return(runnerTwo, nil)

		market, err := factory.FixtureAndMarket(fixture, "OVER_UNDER_25")

		if err != nil {
			t.Fatalf("Error creating market expected nil got %s", err)
		}

		runners.AssertExpectations(t)

		assert.Equal(t, "1.167019590", market.ID)
		assert.Equal(t, uint64(148270), market.FixtureID)
		assert.Equal(t, "Betfair", market.Bookmaker)
		assert.Equal(t, "OVER_UNDER_25", market.Name)
		assert.Equal(t, 2, len(market.Runners))
	})

	t.Run("returns an error if fixture competition is not supported", func(t *testing.T) {
		t.Helper()

		server := mock.ResponseServer(t, marketCatalogueResponse, 200, "https://api.betfair.com/test/listMarketCatalogue/")

		client := mock.BetfairClient(server)
		runners := new(mock.RunnerFactory)

		factory := MarketFactory{
			client: client,
			runner: runners,
		}

		fixture := newStatisticoFixture("Liverpool", "Manchester United", 148270, 44)

		market, err := factory.FixtureAndMarket(fixture, "OVER_UNDER_25")

		if market != nil {
			t.Fatalf("Error expected nil got %+v", market)
		}

		if err == nil {
			t.Fatal("Error expected error got nil")
		}

		assert.Equal(t, "competition ID 44 is not supported", err.Error())
	})

	t.Run("returns an error if unable to parse market from response", func(t *testing.T) {
		t.Helper()

		server := mock.ResponseServer(t, `[]`, 200, "https://api.betfair.com/test/listMarketCatalogue/")

		client := mock.BetfairClient(server)
		runners := new(mock.RunnerFactory)

		factory := MarketFactory{
			client: client,
			runner: runners,
		}

		fixture := newStatisticoFixture("Liverpool", "Manchester United", 148270, 8)

		market, err := factory.FixtureAndMarket(fixture, "OVER_UNDER_25")

		if market != nil {
			t.Fatalf("Error expected nil got %+v", market)
		}

		if err == nil {
			t.Fatal("Error expected error got nil")
		}

		assert.Equal(t, "no market returned for fixture 148270 and bet type OVER_UNDER_25", err.Error())
	})

	t.Run("returns error if event returned by betfair does not match fixture", func(t *testing.T) {
		t.Helper()

		server := mock.ResponseServer(t, marketCatalogueResponse, 200, "https://api.betfair.com/test/listMarketCatalogue/")

		client := mock.BetfairClient(server)
		runners := new(mock.RunnerFactory)

		factory := MarketFactory{
			client: client,
			runner: runners,
		}

		fixture := newStatisticoFixture("Liverpool", "Manchester City", 148270, 8)

		market, err := factory.FixtureAndMarket(fixture, "OVER_UNDER_25")

		if market != nil {
			t.Fatalf("Error expected nil got %+v", market)
		}

		if err == nil {
			t.Fatal("Error expected error got nil")
		}

		assert.Equal(
			t,
			"event 'Liverpool v Man Utd' returned by betfair client does not match fixture 'Liverpool v Manchester City'",
			err.Error(),
		)
	})

	t.Run("returns an error if error occurs creating runners", func(t *testing.T) {
		t.Helper()

		server := mock.ResponseServer(t, marketCatalogueResponse, 200, "https://api.betfair.com/test/listMarketCatalogue/")

		client := mock.BetfairClient(server)
		runners := new(mock.RunnerFactory)

		factory := MarketFactory{
			client: client,
			runner: runners,
		}

		fixture := newStatisticoFixture("Liverpool", "Manchester United", 148270, 8)

		runnerOne := newBookmakerRunner("Under 2.5 Goals", 47972)

		runners.On("CreateRunner", uint64(47972), "1.167019590", "Under 2.5 Goals").Return(runnerOne, nil)
		runners.On("CreateRunner", uint64(47973), "1.167019590", "Over 2.5 Goals").Return(&bookmaker.Runner{}, errors.New("oh no"))

		market, err := factory.FixtureAndMarket(fixture, "OVER_UNDER_25")

		if market != nil {
			t.Fatalf("Error expected nil got %+v", market)
		}

		if err == nil {
			t.Fatal("Error expected error got nil")
		}

		assert.Equal(t, "oh no", err.Error())

		runners.AssertExpectations(t)
	})
}

var marketCatalogueResponse = `[
  {
    "marketId": "1.167019590",
    "marketName": "Over/Under 2.5 Goals",
    "totalMatched": 233.98,
    "runners": [
      {
        "selectionId": 47972,
        "runnerName": "Under 2.5 Goals",
        "handicap": 0.0,
        "sortPriority": 1
      },
      {
        "selectionId": 47973,
        "runnerName": "Over 2.5 Goals",
        "handicap": 0.0,
        "sortPriority": 2
      }
    ],
    "event": {
      "id": "29637723",
      "name": "Liverpool v Man Utd",
      "countryCode": "GB",
      "timezone": "GMT",
      "openDate": "2020-01-19T16:30:00.000Z"
    }
  }
]`

func newStatisticoFixture(home, away string, fixtureID, competitionID uint64) statistico.Fixture {
	return statistico.Fixture{
		ID:            fixtureID,
		CompetitionID: competitionID,
		HomeTeam:      home,
		AwayTeam:      away,
	}
}

func newBookmakerRunner(name string, selectionID uint64) *bookmaker.Runner {
	return &bookmaker.Runner{
		Name:        name,
		SelectionID: selectionID,
	}
}
