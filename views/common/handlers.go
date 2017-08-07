package common

import (
	"net/http"
	"github.com/alehano/gobootstrap/sys/tpl"
	"github.com/alehano/gobootstrap/config"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	tpl.Render(w, r, "common.not_found", tpl.D{"page_title": "Page not found"})
}

func robotsTxt(w http.ResponseWriter, r *http.Request) {
	tpl.Render(w, r, "common.robots_txt", tpl.D{"host": config.WebsiteURL()})
}
