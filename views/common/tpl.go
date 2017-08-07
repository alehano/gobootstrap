package common

import (
	"github.com/alehano/gobootstrap/sys/tpl"
)

func init() {

	tpl.RegisterMulti("views/common/tpl/", map[string]string{
		"common.robots_txt": "robots_txt.tpl",
		"common.not_found":  "404.tpl",
	})
}
