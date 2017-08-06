/*
Templates store and rendering.
Root template must have name "base".

Context value available by "context" name.
To get value use: {{.context.Value "key"}}

Example:

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

func NewSet() Set {
	s := Set{}
	s.AddFuncMap(defaultFilters)
	return s
}

// A set of template files with a list of FuncMap
type Set struct {
	prefix    string
	filenames []string
	funcMaps  []template.FuncMap
}

// AddFuncMap adds template FuncMap for the Set
func (s Set) AddFuncMap(fm template.FuncMap) Set {
	ns := s.copy()
	ns.funcMaps = append(ns.funcMaps, fm)
	return ns
}

func (s Set) AddFunc(name string, fn interface{}) Set {
	fm := template.FuncMap{}
	fm[name]=fn
	return s.AddFuncMap(fm)
}

// Set template filenames prefix
func (s Set) SetPrefix(pathPrefix string) Set {
	ns := s.copy()
	ns.prefix = pathPrefix
	return ns
}

// Adds files with set prefix
func (s Set) Add(filename ...string) Set {
	ns := s.copy()
	for _, fn := range filename {
		ns.filenames = append(ns.filenames, ns.prefix+fn)
	}
	return ns
}

// Adds files with given prefix, ignoring set by default
func (s Set) AddWithPrefix(prefix string, filename ...string) Set {
	ns := s.copy()
	for _, fn := range filename {
		ns.filenames = append(ns.filenames, prefix+fn)
	}
	return ns
}

// Adds files without any prefix
func (s Set) AddNoPrefix(filename ...string) Set {
	return s.AddWithPrefix("", filename...)
}

// Copy creates new Set
func (s Set) copy() Set {
	newSet := Set{prefix:s.prefix}
	newSet.filenames = append(newSet.filenames, s.filenames...)
	newSet.funcMaps = append(newSet.funcMaps, s.funcMaps...)
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
		return 
	}
	if config.Get().Debug {
		storeSets[name] = set
	}
	compileTmpl(name, set)
}

func compileTmpl(name string, set Set) {
	var absFileNames []string
	for _, fn := range set.filenames {
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
