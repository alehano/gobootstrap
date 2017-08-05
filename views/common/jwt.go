package common

import (
	"github.com/go-chi/jwtauth"
	"github.com/alehano/gobootstrap/config"
)

var JwtTokenAuth = jwtauth.New("HS256", []byte(config.Get().JWTSecret), nil)
