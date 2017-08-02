package admin

import (
	"net/http"
	"github.com/go-chi/jwtauth"
	"github.com/alehano/reverse"
)

func AdminAuthenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, claims, err := jwtauth.FromContext(r.Context())
		redirect := func() {
			http.Redirect(w, r, reverse.Rev("admin.login"), http.StatusFound)
		}
		if err != nil {
			redirect()
			return
		}
		if token == nil || !token.Valid {
			redirect()
			return
		}
		// Check admin rights
		if isAdmin, exists := claims.Get("is_admin"); !exists || !isAdmin.(bool) {
			redirect()
			return
		}
		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}
