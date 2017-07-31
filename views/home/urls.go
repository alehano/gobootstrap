package home

import (
	"github.com/go-chi/chi"
	"github.com/alehano/gobootstrap/sys/urls"
	"github.com/alehano/reverse"
)

func init() {
	urls.Register(urlGroup)
}

func urlGroup(r chi.Router) {
	r.Get("/", home)

	//r.Route("/articles", func(r chi.Router) {
	//	r.Get("/", listArticles)
	//	r.Route("/{articleID}", func(r chi.Router) {
	//		r.Get("/", getArticle)
	//	})
	//})

	r.Route("/articles", func(r chi.Router) {
		r.Get(reverse.AddGr("list_articles", "/articles", "/"), listArticles)
		r.Route("/{articleID}", func(r chi.Router) {
			r.Get(reverse.AddGr("get_article", "/articles{articleID}", "/", "{articleID}"), getArticle)
		})
	})

}
