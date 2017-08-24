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
	r.Get(reverse.Add("home.index", "/"), index)
	r.Get(reverse.Add("home.json", "/json"), json)
}
