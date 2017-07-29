package main

import (
	"fmt"
	"net/http"
	"github.com/alehano/gobootstrap/sys/cli"
	"github.com/alehano/gobootstrap/sys/urls"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/alehano/gobootstrap/config"
	// Include views
	_ "github.com/alehano/gobootstrap/views/home"
)

func main() {
	cli.CheckAndRun()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	urls.BuildAllURLs(r)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Get().Port), r)
}
