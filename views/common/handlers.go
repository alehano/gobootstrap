package common

import (
	"net/http"
	"github.com/alehano/gobootstrap/sys/tpl"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	tpl.Render(w, r, "common.not_found", tpl.D{"page_title": "Page not found"})
}
