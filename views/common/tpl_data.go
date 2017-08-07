package common

import (
	"github.com/alehano/gobootstrap/sys/tpl"
	"github.com/alehano/reverse"
	"github.com/alehano/gobootstrap/config"
)

func init()  {
	tpl.RegisterDefaultData(map[string]interface{}{
		"url": reverse.Rev,
		"lang": config.Get().Lang,
	})
}