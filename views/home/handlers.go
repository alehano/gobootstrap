package home

import (
	"context"
	"net/http"
	"github.com/alehano/gobootstrap/sys/tpl"
)

func index(w http.ResponseWriter, r *http.Request) {
	r = r.WithContext(context.WithValue(r.Context(), "ctxValue", "ctxValueOK"))
	tpl.Render(w, r, "home.index", tpl.D{
		"page_title": "Homepage",
		"testValue": "testValueOK", 
	})
}
