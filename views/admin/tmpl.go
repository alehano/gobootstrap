package admin

import 	"github.com/alehano/gobootstrap/sys/tmpl"

var BaseTmpl = tmpl.NewSet("views/admin/tmpl/", "base.tmpl")

func init()  {

	tmpl.Register("admin.login", BaseTmpl.With("views/admin/tmpl/", "login.tmpl"))

}

