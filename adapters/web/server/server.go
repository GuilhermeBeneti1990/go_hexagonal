package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/GuilhermeBeneti1990/go-hexagonal/adapters/web/handlers"
	"github.com/GuilhermeBeneti1990/go-hexagonal/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type Webserver struct {
	Service application.IProductService
}

func MakeNewWebserver() *Webserver {
	return &Webserver{}
}

func (w Webserver) Serve() {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handlers.MakeProductHandlers(r, n, w.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":8080",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log:", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
