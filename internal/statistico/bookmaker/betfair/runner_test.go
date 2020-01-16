package betfair

import (
	"github.com/statistico/statistico-bet-finder/internal/statistico/bookmaker"
	"github.com/statistico/statistico-bet-finder/internal/statistico/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunnerFactory_CreateRunner(t *testing.T) {
	t.Run("returns hydrated runners based on runner catalogue and market", func(t *testing.T) {
		t.Helper()

		server := mock.ResponseServer(t, runnersResponse, 200, "https://api.betfair.com/test/listRunnerBook/")

		client := mock.BetfairClient(server)

		factory := RunnerFactory{client:client}

		fetched, err := factory.CreateRunner(47973, "1.567278", "Under 2.5 Goals")

		if err != nil {
			t.Fatalf("Error creating runners expected nil got %s", err)
		}

		back := []bookmaker.Price{
			{
				Price: 1.86,
				Size: 19.16,
			},
			{
				Price: 1.85,
				Size: 87.82,
			},
			{
				Price: 1.81,
				Size: 247.25,
			},
		}

		lay := []bookmaker.Price{
			{
				Price: 1.98,
				Size: 21.6,
			},
			{
				Price: 2.0,
				Size: 10.6,
			},
			{
				Price: 2.04,
				Size: 33.97,
			},
		}

		assert.Equal(t, "Under 2.5 Goals", fetched.Name)
		assert.Equal(t, uint64(47973), fetched.SelectionID)
		assert.Equal(t, back, fetched.Back)
		assert.Equal(t, lay, fetched.Lay)
	})

	t.Run("returns error if runner does not exist for market and selection", func(t *testing.T) {
		t.Helper()

		server := mock.ResponseServer(t, `[]`, 200, "https://api.betfair.com/test/listRunnerBook/")


		client := mock.BetfairClient(server)

		factory := RunnerFactory{client:client}

		fetched, err := factory.CreateRunner(47973, "1.567278", "Under 2.5 Goals")

		if err == nil {
			t.Fatal("Error expected got nil")
		}

		if fetched != nil {
			t.Fatalf("Expected nil got %+v", fetched)
		}

		assert.Equal(t, "runner book does not exist for Market '1.567278' and Selection '47973'", err.Error())
	})

	t.Run("returns error if error returned from betfair client", func(t *testing.T) {
		t.Helper()

		server := mock.ResponseServer(t, `Error occurred`, 400, "https://api.betfair.com/test/listRunnerBook/")

		client := mock.BetfairClient(server)

		factory := RunnerFactory{client:client}

		fetched, err := factory.CreateRunner(47973, "1.567278", "Under 2.5 Goals")

		if err == nil {
			t.Fatal("Error expected got nil")
		}

		if fetched != nil {
			t.Fatalf("Expected nil got %+v", fetched)
		}
	})

	t.Run("returns error if response does not return any runnes", func(t *testing.T) {
		t.Helper()

		server := mock.ResponseServer(t, emptyRunnerResponse, 200, "https://api.betfair.com/test/listRunnerBook/")


		client := mock.BetfairClient(server)

		factory := RunnerFactory{client:client}

		fetched, err := factory.CreateRunner(47973, "1.567278", "Under 2.5 Goals")

		if err == nil {
			t.Fatal("Error expected got nil")
		}

		if fetched != nil {
			t.Fatalf("Expected nil got %+v", fetched)
		}

		assert.Equal(t, "runner book does not exist for Market '1.567278' and Selection '47973'", err.Error())
	})
}

var runnersResponse = `[
  {
    "marketId": "1.167019437",
    "isMarketDataDelayed": true,
    "status": "OPEN",
    "betDelay": 0,
    "bspReconciled": false,
    "complete": true,
    "inplay": false,
    "numberOfWinners": 1,
    "numberOfRunners": 2,
    "numberOfActiveRunners": 2,
    "lastMatchTime": "2020-01-16T11:30:52.965Z",
    "totalMatched": 863.62,
    "totalAvailable": 63875.59,
    "crossMatching": true,
    "runnersVoidable": false,
    "version": 3116445978,
    "runners": [
      {
        "selectionId": 47973,
        "handicap": 0.0,
        "status": "ACTIVE",
        "lastPriceTraded": 1.94,
        "totalMatched": 0.0,
        "ex": {
          "availableToBack": [
            {
              "price": 1.86,
              "size": 19.16
            },
            {
              "price": 1.85,
              "size": 87.82
            },
            {
              "price": 1.81,
              "size": 247.25
            }
          ],
          "availableToLay": [
            {
              "price": 1.98,
              "size": 21.6
            },
            {
              "price": 2.0,
              "size": 10.6
            },
            {
              "price": 2.04,
              "size": 33.97
            }
          ],
          "tradedVolume": []
        }
      }
    ]
  }
]`

var emptyRunnerResponse = `[
  {
    "marketId": "1.167019437",
    "isMarketDataDelayed": true,
    "status": "OPEN",
    "betDelay": 0,
    "bspReconciled": false,
    "complete": true,
    "inplay": false,
    "numberOfWinners": 1,
    "numberOfRunners": 2,
    "numberOfActiveRunners": 2,
    "lastMatchTime": "2020-01-16T11:30:52.965Z",
    "totalMatched": 863.62,
    "totalAvailable": 63875.59,
    "crossMatching": true,
    "runnersVoidable": false,
    "version": 3116445978,
    "runners": []
  }
]`