package urls

import (
	"github.com/go-chi/chi"
	"net/http"
	"os"
	"path/filepath"
)

var store []func(r chi.Router)
var storeStatic = [][]string{}

// Register URLs as a router Group
func Register(fn func(r chi.Router)) {
	store = append(store, fn)
}

// Register static files from path, accessible by url
func RegisterStatic(url, path string) {
	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, path)
	storeStatic = append(storeStatic, []string{url, filesDir})
}

// Add all registered URLs into Mux
func AddAll(r *chi.Mux) {
	for _, fn := range store {
		r.Group(fn)
	}
	for _, item := range storeStatic {
		fileServer(r, item[0], http.Dir(item[1]))
	}
}
