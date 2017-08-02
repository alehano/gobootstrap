package admin

import (
	"net/http"
	"github.com/alehano/reverse"
	"github.com/alehano/gobootstrap/sys/tmpl"
)

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "admin.index")
}

func login(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "admin.login")
}

func loginPost(w http.ResponseWriter, r *http.Request) {
	// Check credentials ...

	// If not auth, redirect
	http.Redirect(w, r, reverse.Rev("admin.index"), http.StatusFound)
}
