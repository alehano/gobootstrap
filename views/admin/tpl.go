package admin

import (
	"github.com/alehano/gobootstrap/sys/tpl"
)

func init() {

	tpl.RegisterMulti("views/admin/tpl/", map[string]string{
		"admin.index": "index.tpl",
		"admin.login":  "login.tpl",
	})
}
