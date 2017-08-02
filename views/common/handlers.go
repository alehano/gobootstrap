package common

import (
	"net/http"
	"github.com/alehano/gobootstrap/sys/tmpl"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	tmpl.Render(w, r, "common.not_found")
}
