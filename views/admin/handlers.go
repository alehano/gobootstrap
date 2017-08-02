package admin

import (
	"net/http"
	"github.com/alehano/reverse"
	"github.com/alehano/gobootstrap/sys/tmpl"
	"context"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Admin Index"))
}

func login(w http.ResponseWriter, r *http.Request) {
	r = r.WithContext(context.WithValue(r.Context(), "ctx", "ctxOK"))
	tmpl.Render(w, r, "admin.login", tmpl.D{"test": "OK"})
}

func loginPost(w http.ResponseWriter, r *http.Request) {
	// Check credentials ...

	// If not auth, redirect
	http.Redirect(w, r, reverse.Rev("admin.index"), http.StatusFound)
}
