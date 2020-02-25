package main

import (
	"bytes"
	"github.com/julienschmidt/httprouter"
	"github.com/statistico/statistico-bet-finder/internal/app"
	"github.com/statistico/statistico-bet-finder/internal/app/mock"
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
		bookHandler := newBookHandler(bookmaker)

		var body = `{"markets": ["OVER_UNDER_15", "OVER_UNDER_25"]}`

		request, err := http.NewRequest("POST", "/api/v1/event/18279/book", ioutil.NopCloser(bytes.NewBufferString(body)))

		if err != nil {
			t.Fatal(err)
		}

		query := app.BookQuery{
			EventID:    18279,
			Markets:    []string{"OVER_UNDER_15", "OVER_UNDER_25"},
		}

		book := app.Book{
			EventID:    18279,
			Markets: []*app.Market{
				{
					Name:       "OVER_UNDER_25",
					Bookmakers: nil,
				},
			},
			CreatedAt: time.Date(2019, 01, 14, 11, 25, 00, 00, time.UTC),
		}

		bookmaker.On("CreateBook", &query).Return(&book, nil)

		response := httptest.NewRecorder()
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bookHandler.PostBook(w, r, httprouter.Params{{Key: "id", Value: "18279"}})
		})
		handler.ServeHTTP(response, request)

		expected := `{"status":"success","data":{"book":{"eventId":18279,"markets":[{"name":"OVER_UNDER_25","bookmakers":null}],"createdAt":"2019-01-14T11:25:00Z"}}}`

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, expected, response.Body.String())
	})

	t.Run("returns 400 response if unable to parse request body", func(t *testing.T) {
		t.Helper()

		bookmaker := new(mock.Bookmaker)
		bookHandler := newBookHandler(bookmaker)

		request, err := http.NewRequest("POST", "/api/v1/event/18279book", ioutil.NopCloser(bytes.NewBufferString(`[]`)))

		if err != nil {
			t.Fatal(err)
		}

		bookmaker.AssertNotCalled(t, "CreateBook")

		response := httptest.NewRecorder()
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bookHandler.PostBook(w, r, httprouter.Params{{Key: "id", Value: "18279"}})
		})

		handler.ServeHTTP(response, request)

		expected := `{"status":"fail","data":[{"message":"unable to parse request body","code":1}]}`

		assert.Equal(t, 400, response.Code)
		assert.Equal(t, expected, response.Body.String())
	})

	t.Run("returns 400 response if request body does not match expected schema", func(t *testing.T) {
		t.Helper()

		bookmaker := new(mock.Bookmaker)
		bookHandler := newBookHandler(bookmaker)

		var body = `{"markets": [13986]}`

		request, err := http.NewRequest("POST", "/api/v1/event/18279/book", ioutil.NopCloser(bytes.NewBufferString(body)))

		if err != nil {
			t.Fatal(err)
		}

		bookmaker.AssertNotCalled(t, "CreateBook")

		response := httptest.NewRecorder()
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bookHandler.PostBook(w, r, httprouter.Params{{Key: "id", Value: "18279"}})
		})

		handler.ServeHTTP(response, request)

		expected := `{"status":"fail","data":[{"message":"unable to parse request body","code":1}]}`

		assert.Equal(t, 400, response.Code)
		assert.Equal(t, expected, response.Body.String())
	})
}
