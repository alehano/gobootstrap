package common

import (
	"net/http"
	"github.com/alehano/gobootstrap/sys/tmpl"
)

// RenderTmpl returns URL handler with rendered template
func RenderTmpl(name string, data ...map[string]interface{}) func (http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Render(w, r, name, data...)
	}
}
