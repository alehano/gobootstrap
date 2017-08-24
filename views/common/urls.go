package common

import (
	"github.com/alehano/gobootstrap/sys/urls"
	"github.com/alehano/reverse"
	"github.com/go-chi/chi"
	"github.com/alehano/gobootstrap/sys/tpl"
	"github.com/alehano/gobootstrap/config"
)

func init() {
	urls.Register(urlGroup)
	urls.RegisterStatic(reverse.Add("static", "/static/"), "/views/common/static/")
}

func urlGroup(r chi.Router) {

	r.NotFound(notFound)

	r.Get("/robots.txt", tpl.RenderHandler("common.robots_txt",
		tpl.D{"host": config.WebsiteURL()}))
}
