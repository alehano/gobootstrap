package helpers

import (
	"net/http"
	"context"
)

var Context = ctxHelpers{}

type ctxHelpers struct {
}

// Adds value to Request's Context
func (h ctxHelpers) AddValueToRequest(r *http.Request, key, value interface{}) {
	*r = *r.WithContext(context.WithValue(r.Context(), key, value))
}