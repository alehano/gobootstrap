package common

import 	"github.com/alehano/gobootstrap/sys/tmpl"


func init()  {

	tmpl.Register("common.robots_txt", tmpl.NewSet("views/common/tmpl/", "robots_txt.tmpl"))

}

