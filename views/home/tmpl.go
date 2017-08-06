package home

import 	(
	"github.com/alehano/gobootstrap/sys/tmpl"
	"github.com/alehano/gobootstrap/views/common"
)

var tmplSet = common.BaseTmpl.SetPrefix("views/home/tmpl/")

func init() {

	tmpl.Register("home.index", tmplSet.Add("index.tmpl"))

}

