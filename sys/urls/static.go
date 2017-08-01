package urls

import (
	"strings"
	"github.com/go-chi/chi"
	"net/http"
)

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func fileServer(r chi.Router, url string, path http.FileSystem) {
	if strings.ContainsAny(url, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}
	fs := http.StripPrefix(url, http.FileServer(path))
	if url != "/" && url[len(url)-1] != '/' {
		r.Get(url, http.RedirectHandler(url+"/", 301).ServeHTTP)
		url += "/"
	}
	url += "*"
	r.Get(url, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}