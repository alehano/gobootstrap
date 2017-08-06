package common

import (
	"github.com/alehano/reverse"
	"html/template"
)

var DefaultTmplFuncMap = template.FuncMap{
	"url": reverse.Rev,
}

// You can add more... See examples: https://github.com/flosch/pongo2/blob/master/filters_builtin.go
