package home

import (
	"net/http"
	"github.com/alehano/gobootstrap/sys/tmpl"
	"context"
	"github.com/alehano/gobootstrap/sys/tpl"
)

func index(w http.ResponseWriter, r *http.Request) {
	r = r.WithContext(context.WithValue(r.Context(), "ctxValue", "ctxValueOK"))
	tmpl.Render(w, r, "home.index", tmpl.D{"testValue": "testValueOK"})
}

////

func testPongo(w http.ResponseWriter, r *http.Request) {
	r = r.WithContext(context.WithValue(r.Context(), "ctxValue", map[string]interface{}{
		"nested": "nestedValue",
	}))
	tpl.Render(w, r, "home.test", tpl.D{"query": "OK"})
}
