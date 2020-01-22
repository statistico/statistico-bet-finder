package rest_test

import (
	"bytes"
	"github.com/statistico/statistico-bet-finder/internal/app"
	"github.com/statistico/statistico-bet-finder/internal/app/mock"
	"github.com/statistico/statistico-bet-finder/internal/app/rest"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestBookHandler_CreateBook(t *testing.T) {
	t.Run("returns 200 response containing book data", func(t *testing.T) {
		t.Helper()

		bookmaker := new(mock.Bookmaker)
		bookHandler := rest.NewBookHandler(bookmaker)

		var body = `{"fixture_ids": [18279, 17289], "markets": ["OVER_UNDER_15", "OVER_UNDER_25"]}`

		request, err := http.NewRequest("POST", "/book", ioutil.NopCloser(bytes.NewBufferString(body)))

		if err != nil {
			t.Fatal(err)
		}
		
		query := app.BookQuery{
			FixtureIDs:    []uint64{18279, 17289},
			Markets: []string{"OVER_UNDER_15", "OVER_UNDER_25"},
		}
		
		book := app.Book{
			Markets: []*app.Market{
				&app.Market{
					FixtureID:  18279,
					Name:       "OVER_UNDER_25",
					Statistico: nil,
					Bookmakers:  nil,
				},
			},
			CreatedAt: time.Date(2019, 01, 14, 11, 25, 00, 00, time.UTC),
		}
		
		bookmaker.On("CreateBook", &query).Return(&book)

		response := httptest.NewRecorder()
		handler := http.HandlerFunc(bookHandler.CreateBook)
		
		handler.ServeHTTP(response, request)

		expected := `{"message":"success","data":{"book":{"markets":[{"fixture_id":18279,"name":"OVER_UNDER_25","statistico":null,"bookmakers":null}],"created_at":"2019-01-14T11:25:00Z"}}}`

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, expected, response.Body.String())
	})
}
