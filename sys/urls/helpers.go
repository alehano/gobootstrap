package urls

import (
	"net/http"
	"github.com/go-chi/chi"
)

// Register Get() and Head() methods at once
// Due to issue: https://github.com/go-chi/chi/issues/238#event-1189509880
func GetHead(r chi.Router, pattern string, h http.HandlerFunc) {
	r.Get(pattern, h)
	r.Head(pattern, h)
}
