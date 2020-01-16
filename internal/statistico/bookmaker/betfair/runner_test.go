package betfair

import (
	"bytes"
	"github.com/statistico/statistico-bet-finder/internal/statistico/mock"
	bfClient "github.com/statistico/statistico-betfair-go-client"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_createRunners(t *testing.T) {
	t.Run("returns hydrated runners based on runner catalogue and market", func(t *testing.T) {
		t.Helper()

		server := mock.HttpClient(func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewBufferString(runnersResponse)),
			}, nil
		})

		client := mock.BetfairClient(server)

		runners := []bfClient.RunnerCatalogue{
			{
				SelectionID:  47973,
				RunnerName:   "Under 2.5 Goals",
			},
			{
				SelectionID:  45766,
				RunnerName:   "Over 2.5 Goals",
			},
		}

		fetched, err := createRunners(&client, runners, "1.567278")

		if err != nil {
			t.Fatalf("Error creating runners expected nil got %s", err)
		}

		assert.Equal(t, 2, len(fetched))
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