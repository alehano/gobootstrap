package home

import "net/http"

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}

func listArticles(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list articles"))
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get article"))
}