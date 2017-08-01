package common

import (
	"net/http"
	"github.com/alehano/gobootstrap/sys/tmpl"
)

func robotsTxt(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, "common.robots_txt", nil)
}
