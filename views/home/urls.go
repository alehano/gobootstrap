package home

import (
	"github.com/go-chi/chi"
	"github.com/alehano/gobootstrap/sys/urls"
)

func init() {
	urls.Register(urlGroup)
}

func urlGroup(r chi.Router) {
	r.Get("/", Home)
}