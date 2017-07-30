package urls

import "github.com/go-chi/chi"

var store []func(r chi.Router)

func Register(fn func(r chi.Router))  {
	store = append(store, fn)
}

// Add all registered URLs into Mux
func AddAll(r *chi.Mux) {
	for _, fn := range store {
		r.Group(fn)
	}
}