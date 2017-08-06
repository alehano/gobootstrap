package home

import (
	"github.com/alehano/gobootstrap/sys/tpl"
)

func init() {

	tpl.RegisterMulti("views/home/tpl/", map[string]string{
		"home.test": "example.tpl",
	})

}
