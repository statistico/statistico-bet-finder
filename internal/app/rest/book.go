package rest

import (
	"encoding/json"
	"github.com/statistico/statistico-bet-finder/internal/app"
	"net/http"
)

type BookHandler struct {
	bookmaker app.BookMaker
}

func (b BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	query, err := parseBookQuery(r)

	if err != nil {
		failResponse(w, http.StatusBadRequest, err)
		return
	}

	book := b.bookmaker.CreateBook(query)

	response := bookResponse{Book:book}

	successResponse(w, http.StatusOK, response)
}

func parseBookQuery(r *http.Request) (*app.BookQuery, error) {
	body := struct {
		FixtureIDs []uint64 `json:"fixtureIds"`
		Markets []string `json:"markets"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		return nil, errBadRequest
	}

	query := app.BookQuery{
		Markets:    body.Markets,
		FixtureIDs: body.FixtureIDs,
	}

	return &query, nil
}
