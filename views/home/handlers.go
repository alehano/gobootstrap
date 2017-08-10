package home

import (
	"net/http"
	"github.com/alehano/gobootstrap/sys/tpl"
	"github.com/alehano/gobootstrap/helpers"
)

func index(w http.ResponseWriter, r *http.Request) {
	helpers.Context.AddValueToRequest(r, "ctxValue", "ctxValueOK")
	tpl.Render(w, r, "home.index", tpl.D{
		"page_title": "Homepage",
		"testValue": "testValueOK", 
	})
}
