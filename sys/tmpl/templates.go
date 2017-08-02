/*
Templates store and rendering.
Root template must have name "base".

Context value available by "context" name.
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
	if len(s.funcMaps) == 1 {
		return s.funcMaps[0]
	}
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
	t := template.Must(template.New(name).Funcs(set.getFuncMapCombined()).ParseFiles(absFileNames...))
	store[name] = *t
}

func Render(w http.ResponseWriter, r *http.Request, name string, data ...map[string]interface{}) {
	renderTmpl(w, r, name, data...)
}

func renderTmpl(w http.ResponseWriter, r *http.Request, name string, data ...map[string]interface{}) {
	// Recompile in debug mode
	if config.Get().Debug {
		compileTmpl(name, storeSets[name])
	}
	t, ok := store[name]
	if !ok {
		http.Error(w, fmt.Sprintf("Template %q does not exists", name), http.StatusInternalServerError)
	}
	newData := map[string]interface{}{}
	newData["context"] = r.Context()
	if len(data) > 0 && data[0] != nil {
		for k, v := range data[0] {
			newData[k] = v
		}
	}
	err := t.ExecuteTemplate(w, "base", newData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
