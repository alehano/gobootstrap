package common

import (
	"net/http"
	"context"
	"github.com/go-chi/jwtauth"
	"time"
	"github.com/alehano/gobootstrap/sys/memcache"
	"github.com/alehano/gobootstrap/config"
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

// Progressive pause on loading auth page if password was wrong.
// You have to set login value to a Context by a key "throttle_login" when auth failed.
// Or just send credentials as "login", "password" form values.
func LoginThrottle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			cooldown          = 5 * time.Minute
			basePause         = time.Second
			factor    float64 = 1.5
			login             = ""
		)
		if r.Context().Value("throttle_login") != nil {
			login = r.Context().Value("throttle_login").(string)
		} else if r.FormValue("login") != "" && r.FormValue("password") != "" {
			login = r.FormValue("login")
		}
		if login != "" {
			key := config.CacheKeys.AuthThrottle(login)
			pause, _, _ := memcached.GetInt(key)
			newPause := int(float64(pause) * factor)
			if newPause == 0 {
				newPause = int(basePause)
			}
			memcached.SetInt(key, newPause, memcached.SetExpiration(cooldown))
			if pause > 0 {
				time.Sleep(time.Duration(pause))
			}
		}
		next.ServeHTTP(w, r)
	})
}
