package admin

import (
	"github.com/alehano/gobootstrap/sys/urls"
	"github.com/alehano/gobootstrap/views/common"
	"github.com/alehano/reverse"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
)


func init() {
	urls.Register(urlGroup)
	urls.RegisterStatic(reverse.Add("static.admin", "/admin/static/"), "/views/admin/static/")
}

func urlGroup(r chi.Router) {
	r.Route(reverse.Add("admin.login", "/admin/login"), func(r chi.Router) {
		r.Use(jwtauth.Verifier(common.JwtTokenAuth))
		r.Get("/", login)
		r.With(middleware.Throttle(1)).Post("/", loginPost)
	})
	r.Get(reverse.Add("admin.logout", "/admin/logout"), logout)
	r.Route(reverse.Add("admin.index", "/admin"), func(r chi.Router) {
		r.Use(jwtauth.Verifier(common.JwtTokenAuth))
		r.Use(AdminAuthenticator)
		r.Get("/", index)
	})
}
