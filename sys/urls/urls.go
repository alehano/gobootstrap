// URL store
package urls

import 	"github.com/go-chi/chi"

var store = []func(r chi.Router){}

func RegisterRouter(fn func(r chi.Router)) {
	store = append(store, fn)
}

func BuildAllURLs(r *chi.Mux) {
	for _, fn := range store {
		r.Group(fn)
	}
}
