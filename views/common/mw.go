package common

import (
	"net/http"
	"context"
	"github.com/go-chi/jwtauth"
)

// Gets JWT claims and save it to context by "session" name
// Usage:
// {{.context.Value "session" }}
// {{with .context.Value "session"}}{{ .admin_login }}{{end}}
func Session(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, _ := jwtauth.FromContext(r.Context())
		ctx := context.WithValue(r.Context(), "session", claims)
		// Concrete key
		//ctx := context.WithValue(r.Context(), "admin_login", claims["admin_login"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
