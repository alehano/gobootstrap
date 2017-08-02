package admin

import (
	"net/http"
	"github.com/alehano/reverse"
	"github.com/alehano/gobootstrap/sys/tmpl"
	"github.com/alehano/gobootstrap/views/common"
	"github.com/alehano/gobootstrap/config"
	"golang.org/x/crypto/bcrypt"
	"github.com/go-chi/jwtauth"
)

func index(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tmpl.Render(w, r, "admin.index", tmpl.D{
		"is_admin": claims["is_admin"],
		"admin_login": claims["admin_login"],
	})
}

func login(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	if isAdmin, exists := claims.Get("is_admin"); exists && isAdmin.(bool) {
		http.Redirect(w, r, reverse.Rev("admin.index"), http.StatusFound)
		return 
	}
	tmpl.Render(w, r, "admin.login")
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
