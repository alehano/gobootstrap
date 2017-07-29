package home

import (
	"github.com/alehano/gobootstrap/sys/urls"
	"github.com/go-chi/chi"
	"net/http"
)

func init() {
	urls.RegisterRouter(urlGroup)
}

func urlGroup(r chi.Router)  {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
}