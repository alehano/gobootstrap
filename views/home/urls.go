package home

import (
	"github.com/go-chi/chi"
	"github.com/alehano/gobootstrap/sys/urls"
	"github.com/alehano/reverse"
)

func init() {
	urls.Register(urlGroup)
}

func urlGroup(r chi.Router) {
	urls.GetAndHead(r, reverse.Add("home.index", "/"), index)
	urls.GetAndHead(r, reverse.Add("home.rest", "/rest"), rest)
}
