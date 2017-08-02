package common

import (
	"github.com/alehano/reverse"
	"html/template"
)

var DefaultTmplFuncMap = template.FuncMap{
	"reverse": reverse.Rev,
}