package home

import (
	"net/http"
	"github.com/alehano/gobootstrap/sys/tmpl"
	"context"
)

func index(w http.ResponseWriter, r *http.Request) {
	r = r.WithContext(context.WithValue(r.Context(), "ctxValue", "ctxValueOK"))
	tmpl.Render(w, r, "home.index", tmpl.D{"testValue": "testValueOK"})
}
