package common

import 	"github.com/alehano/gobootstrap/sys/tmpl"

var BaseTmpl = tmpl.NewSet().SetPrefix("views/common/tmpl/").Add("base.tmpl")

func init()  {

	tmpl.Register("common.robots_txt", BaseTmpl.Add("robots_txt.tmpl"))

}

