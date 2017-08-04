package tmpl

import (
	"time"
	"html/template"
)

var defaultFilters = template.FuncMap{
	"now":    time.Now,
}

