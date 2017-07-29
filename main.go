package main

import (
	"fmt"
	"net/http"
	"github.com/alehano/gobootstrap/sys/cli"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/alehano/gobootstrap/config"
	"github.com/alehano/gobootstrap/views/home"
)

func main() {
	cli.CheckAndRun()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Add URLs
	r.Group(home.URLs)

	http.ListenAndServe(fmt.Sprintf(":%d", config.Get().Port), r)
}
