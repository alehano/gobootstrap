package common

import (
	"github.com/alehano/gobootstrap/sys/urls"
	"github.com/alehano/reverse"
)

func init() {
	urls.RegisterStatic(reverse.Add("static", "/static/"), "/views/common/static/")
}
