package common

import (
	"github.com/alehano/gobootstrap/sys/urls"
	"github.com/alehano/reverse"
	"github.com/go-chi/chi"
)

func init() {
	urls.Register(urlGroup)
	urls.RegisterStatic(reverse.Add("static", "/static/"), "/views/common/static/")
}


func urlGroup(r chi.Router) {

	r.NotFound(notFound)

	urls.GetAndHead(r, "/robots.txt", RenderTmpl("common.robots_txt"))
}
