package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/statistico/statistico-bet-finder/internal/bootstrap"
	"log"
	"net/http"
)

func main() {
	container := bootstrap.BuildContainer(bootstrap.BuildConfig())

	handler := newBookHandler(container.BookMaker())

	router := httprouter.New()
	router.POST("/api/v1/event/:id/book", handler.PostBook)

	log.Fatal(http.ListenAndServe(":80", router))
}