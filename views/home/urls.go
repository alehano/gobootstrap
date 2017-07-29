package home

import (
	"github.com/go-chi/chi"
)

func URLs(r chi.Router) {
	r.Get("/", Home)
}
