package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/statistico/statistico-price-finder/internal/bootstrap"
	"log"
	"net/http"
)

func main() {
	container := bootstrap.BuildContainer(bootstrap.BuildConfig())

	handler := newBookHandler(container.BookMaker(), container.StatisticoBookmaker())

	router := httprouter.New()
	router.POST("/api/v1/event/:id/bookmaker-book", handler.PostBookmakerBook)
	router.POST("/api/v1/event/:id/statistico-book", handler.PostStatisticoBook)

	log.Fatal(http.ListenAndServe(":80", router))
}
