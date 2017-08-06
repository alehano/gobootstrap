package common

import (
	"github.com/alehano/gobootstrap/sys/tpl"
	"github.com/alehano/reverse"
)

func init()  {
	tpl.RegisterDefaultData(map[string]interface{}{
		"url": reverse.Rev,
	})
}