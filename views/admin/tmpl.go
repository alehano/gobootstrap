package admin

import (
	"github.com/alehano/gobootstrap/sys/tmpl"
	"github.com/alehano/gobootstrap/views/common"
)

var (
	BaseTmpl = tmpl.NewSet().
		SetPrefix("views/admin/tmpl/").
		Add("base.tmpl").
		AddFuncMap(common.DefaultTmplFuncMap)
)


func init()  {

	tmpl.Register("admin.index", BaseTmpl.Add("index.tmpl"))
	tmpl.Register("admin.login", BaseTmpl.Add("login.tmpl"))

}

