package home

import (
	"net/http"
	"github.com/alehano/gobootstrap/sys/tpl"
	"github.com/alehano/gobootstrap/helpers"
	"github.com/go-chi/render"
)

func index(w http.ResponseWriter, r *http.Request) {
	helpers.Context.AddValueToRequest(r, "ctxValue", "ctxValueOK")
	tpl.Render(w, r, "home.index", tpl.D{
		"page_title": "Homepage",
		"testValue":  "testValueOK",
	})
}

// Rest example
func json(w http.ResponseWriter, r *http.Request) {
	type data struct {
		Title string
		Count int
	}
	d := data{
		Title: "Test",
		Count: 42,
	}
	render.JSON(w, r, d)
}
