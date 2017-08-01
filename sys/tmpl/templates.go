/*
Templates store and rendering.
Root template must have name "base".

Context passes by "context" name.
To get value use: {{.context.Value "key"}}

Example:

BaseTmpl := tmpl.NewSet("views/common/tmpl/",
	"base.tmpl", "inc/header.tmpl", inc/footer.tmpl")

BaseTmpl.AddFuncMap(...)

tmpl.Register("home", BaseTmpl.With("views/home/tmpl/", "home.tmpl"))

 */
package tmpl

import (
	"html/template"
	"sync"
	"fmt"
	"github.com/alehano/gobootstrap/config"
	"context"
	"net/http"
)

// Helper type for template data
type D map[string]interface{}

var (
	mu        = sync.RWMutex{}
	store     = map[string]template.Template{}
	storeSets = map[string]Set{}
)

func NewSet(pathPrefix string, names ...string) Set {
	s := Set{}
	s.AddFuncMap(defaultFilters)
	for _, name := range names {
		s.fileNames = append(s.fileNames, pathPrefix+name)
	}
	return s
}

// A set of template files with a list of FuncMap
type Set struct {
	fileNames []string
	funcMaps  []template.FuncMap
}

// AddFuncMap register template FuncMap for the Set
func (s *Set) AddFuncMap(fm template.FuncMap) {
	s.funcMaps = append(s.funcMaps, fm)
}

// With creates new Set based on parent and adds new template files
func (s *Set) With(pathPrefix string, names ...string) Set {
	newSet := Set{}
	newSet.fileNames = append(newSet.fileNames, s.fileNames...)
	newSet.funcMaps = append(newSet.funcMaps, s.funcMaps...)
	secondSet := NewSet(pathPrefix, names...)
	newSet.fileNames = append(newSet.fileNames, secondSet.fileNames...)
	newSet.funcMaps = append(newSet.funcMaps, secondSet.funcMaps...)
	return newSet
}

func (s Set) getFuncMapCombined() template.FuncMap {
	combFm := template.FuncMap{}
	for _, fm := range s.funcMaps {
		for key, fn := range fm {
			combFm[key] = fn
		}
	}
	return combFm
}

// Register adds template to a storage by a name
func Register(name string, set Set) {
	if _, ok := store[name]; ok {
		panic(fmt.Sprintf("Template %q already exists", name))
	}
	if config.Get().Debug {
		storeSets[name] = set
	}
	compileTmpl(name, set)
}

func compileTmpl(name string, set Set) {
	var absFileNames []string
	for _, fn := range set.fileNames {
		absFileNames = append(absFileNames, fmt.Sprintf("%s/%s", config.Get().ProjectPath, fn))
	}
	mu.Lock()
	defer mu.Unlock()
	t := template.Must(template.New("").Funcs(set.getFuncMapCombined()).ParseFiles(absFileNames...))
	store[name] = *t
}

func Render(w http.ResponseWriter, name string, data map[string]interface{}) {
	renderTmpl(w, nil, name, data)
}

func RenderCtx(w http.ResponseWriter, c context.Context, name string, data map[string]interface{}) {
	renderTmpl(w, c, name, data)
}

func renderTmpl(w http.ResponseWriter, c context.Context, name string, data map[string]interface{}) {
	// Recompile in debug mode
	if config.Get().Debug {
		compileTmpl(name, storeSets[name])
	}
	t, ok := store[name]
	if !ok {
		http.Error(w, fmt.Sprintf("Template %q does not exists", name), http.StatusInternalServerError)
	}
	tmplData := map[string]interface{}{}
	if c != nil {
		tmplData["context"] = c
	}
	if data != nil && len(data) > 0 {
		for k := range data {
			tmplData[k] = data[k]
		}
	}
	err := t.ExecuteTemplate(w, "base", tmplData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
