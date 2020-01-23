package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/statistico/statistico-bet-finder/internal/bootstrap"
	"log"
	"net/http"
)

func main() {
	container := bootstrap.BuildContainer(bootstrap.BuildConfig())

	router := httprouter.New()
	router.POST("/book", container.BookHandler().PostBook)

	log.Fatal(http.ListenAndServe(":80", router))
}
