package web

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)
import "github.com/justinas/alice"

type App struct {
	Router    *mux.Router
	Middeware *Middeware
}

func (a *App) Initalize() {
	a.Router = mux.NewRouter()
	a.Middeware = &Middeware{}
	a.registerHandler()
}
func (a *App) registerHandler() {
	chain := alice.New(a.Middeware.LoggingHandler, a.Middeware.RecoverHandler)
	a.Router.Handle("/api/shorten", chain.ThenFunc(createShortUrl)).Methods("POST")
	a.Router.Handle("/api/info", chain.ThenFunc(getShortlinkInfo)).Methods("GET")
	a.Router.Handle("/{shortUrl:[a-zA-Z0-9]{1,11}}", chain.ThenFunc(redirect)).Methods("GET")
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
