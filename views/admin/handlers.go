package admin

import (
	"net/http"
	"github.com/alehano/reverse"
	"github.com/alehano/gobootstrap/sys/tpl"
	"github.com/alehano/gobootstrap/views/common"
	"github.com/alehano/gobootstrap/config"
	"golang.org/x/crypto/bcrypt"
	"github.com/go-chi/jwtauth"
)

func index(w http.ResponseWriter, r *http.Request) {
	tpl.Render(w, r, "admin.index")
}

func login(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	if isAdmin, exists := claims.Get("is_admin"); exists && isAdmin.(bool) {
		http.Redirect(w, r, reverse.Rev("admin.index"), http.StatusFound)
		return
	}
	tpl.Render(w, r, "admin.login", tpl.D{"key": "value"})
}

func loginPost(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("login") == config.Get().AdminLogin &&
		bcrypt.CompareHashAndPassword([]byte(config.Get().AdminPasswordHash),
			[]byte(r.FormValue("password"))) == nil {

		// Auth
		_, tokenString, err := common.JwtTokenAuth.Encode(jwtauth.Claims{
			"is_admin":    true,
			"admin_login": config.Get().AdminLogin,
		})
		if err == nil {
			http.SetCookie(w, &http.Cookie{Name: "jwt", Value: tokenString})
			http.Redirect(w, r, reverse.Rev("admin.index"), http.StatusFound)
		}
	}
	http.Redirect(w, r, reverse.Rev("admin.login"), http.StatusFound)
}

func logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "jwt", Value: "", MaxAge: -1})
	http.Redirect(w, r, reverse.Rev("admin.login"), http.StatusFound)
}
