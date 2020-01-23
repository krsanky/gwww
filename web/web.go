package web

import (
	"html/template"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/Masterminds/sprig"
	"github.com/krsanky/gwww/account"
	"github.com/krsanky/gwww/breadcrumbs"
	"github.com/krsanky/gwww/lg"
)

var GlobalFuncMap template.FuncMap
var _tmpls map[string]*template.Template

func InputChecked(checked bool) string {
	if checked {
		return " checked "
	} else {
		return ""
	}
}

func Test1() string {
	return "Test1..."
}

func init() {
	_tmpls = make(map[string]*template.Template)
	GlobalFuncMap = template.FuncMap{
		"input_checked": InputChecked,
		"test1":         Test1,
	}
	breadcrumbs.AddFuncs(GlobalFuncMap)

	for k, v := range sprig.FuncMap() {
		GlobalFuncMap[k] = v
	}
}

func TmplData(r *http.Request) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	user, _ := account.UserFromContext(r.Context())

	data["page_user"] = user
	return data, nil
}

// the last tmpls is used to name it, so it must be unique
func Render(w http.ResponseWriter, data interface{}, tmpls ...string) {
	if len(tmpls) < 1 {
		lg.Log.Printf("error Render()... tmpls<1")
		panic("at the disco")
	}

	// this name is for my code to cache the template and
	// refer back to it.  In my setup/scheme, it's the last
	// template that is the unique and defining one.
	page := tmpls[len(tmpls)-1]
	//page := filepath.Base(tmpls[0])

	lg.Log.Printf("Render(%s)...", page)
	headers := w.Header()
	headers.Add("Content-Type", "text/html")
	_, tmpl_exists := _tmpls[page]
	if !tmpl_exists {
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		if e := os.Chdir("tmpl/"); e != nil {
			panic(e)
		}

		// this name is like saying {{define "name"}}...
		// which is a feature I don't use in this setup
		name := filepath.Base(tmpls[0])
		t := template.New(name)
		t.Funcs(GlobalFuncMap)

		_, err = t.ParseFiles(tmpls...)
		if err != nil {
			panic(err)
		}
		_tmpls[page] = t

		if e := os.Chdir(dir); e != nil {
			panic(e)
		}
	}
	err := _tmpls[page].Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func LastInPath(u *url.URL) string {
	lg.Log.Printf("LastInPath():%s", u.Path)
	s := strings.SplitN(u.Path, "/", 3)
	lg.Log.Printf("len s:%d %s", len(s), s)

	if len(s) >= 3 {
		val := strings.TrimSuffix(s[2], "/")
		val, err := url.QueryUnescape(val)
		if err != nil {
			return ""
		}
		lg.Log.Printf("return:%s", val)
		return val
	}
	lg.Log.Printf("return:")
	return ""
}
