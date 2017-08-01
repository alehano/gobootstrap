package admin

import (
	"github.com/alehano/gobootstrap/sys/urls"
	"github.com/alehano/reverse"
	"github.com/go-chi/chi"
)

func init() {
	urls.Register(urlGroup)
	urls.RegisterStatic(reverse.Add("static.admin", "/admin/static/"), "/views/admin/static/")
}

func urlGroup(r chi.Router) {
	r.Route(reverse.Add("admin.index", "/admin"), func(r chi.Router) {
		r.Get("/", index)

		r.Route(reverse.AddGr("admin.login", "/admin", "/login"), func(r chi.Router) {
			r.Get("/", login)
			r.Post("/", loginPost)
		})
	})
}
