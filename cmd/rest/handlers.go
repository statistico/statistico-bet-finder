package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-bet-finder/internal/app"
	"net/http"
	"strconv"
)

type bookHandler struct {
	bookmaker app.BookMaker
}

func (b bookHandler) PostBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	query, err := parseBookQuery(r, ps)

	if err != nil {
		failResponse(w, http.StatusBadRequest, err)
		return
	}

	book, err := b.bookmaker.CreateBook(query)

	if err != nil {
		failResponse(w, http.StatusNotFound, err)
		return
	}

	response := bookResponse{Book: book}

	successResponse(w, http.StatusOK, response)
}

func parseBookQuery(r *http.Request, ps httprouter.Params) (*app.BookQuery, error) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		return nil, errBadRequestPath
	}

	body := struct {
		Markets    []string `json:"markets"`
	}{}

	err = json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		logrus.Error(err.Error())
		return nil, errBadRequestBody
	}

	query := app.BookQuery{
		EventID:    uint64(id),
		Markets:    body.Markets,
	}

	return &query, nil
}

func newBookHandler(b app.BookMaker) *bookHandler {
	return &bookHandler{bookmaker: b}
}
