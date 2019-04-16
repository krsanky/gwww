package web

import (
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/template"

	lg "oldcode.org/gow/lg"
)

var tmpls map[string]*template.Template

func GetTmpls() map[string]*template.Template {
	return tmpls
}

func init() {
	tmpls = make(map[string]*template.Template)
}

func RenderPage(w http.ResponseWriter, page string, data interface{}, sub_tmpls ...string) {
	lg.Log.Printf("RenderPage(%s)...", page)
	headers := w.Header()
	headers.Add("Content-Type", "text/html")
	_, tmpl_exists := tmpls[page]
	if !tmpl_exists {
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		if e := os.Chdir("tmpl/"); e != nil {
			panic(e)
		}

		targs := []string{
			"base.html",
			"navbar.tmpl",
			"leftnav.tmpl",
			"js_includes.tmpl",
			"a_z_select.tmpl"}
		if len(sub_tmpls) > 0 {
			lg.Log.Printf("append sub_tmpls...")
			targs = append(targs, sub_tmpls...)
		}
		p := strings.Join([]string{page, "html"}, ".")
		targs = append(targs, p)
		lg.Log.Printf("targs:%s", strings.Join(targs, ","))

		tmpls[page] = template.Must(template.ParseFiles(targs...))
		if e := os.Chdir(dir); e != nil {
			panic(e)
		}
	}
	err := tmpls[page].Execute(w, data)
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
