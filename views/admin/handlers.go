package admin

import (
	"net/http"
	"github.com/alehano/reverse"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Admin Index"))
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Admin Login"))
}

func loginPost(w http.ResponseWriter, r *http.Request) {
	// Check credentials ...

	// If not auth, redirect
	http.Redirect(w, r, reverse.Rev("admin.index"), http.StatusFound)
}
