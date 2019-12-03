package nostyle

import (
	"net/http"
	"os"
	"strings"
	"text/template"

	"oldcode.org/repo/go/gow/lg"
)

var tmpls map[string]*template.Template

func init() {
	tmpls = make(map[string]*template.Template)
}

func RenderPage(w http.ResponseWriter, page string, data interface{}, sub_tmpls ...string) {
	lg.Log.Printf("nostyle.RenderPage(%s)...", page)
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
			"nostyle/base.html",
			//"navbar.tmpl",
			"js_includes.tmpl"}
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
	if data == nil {
		lg.Log.Printf("data is nil")
	}
	tmpls[page].Execute(w, data)
}

func V2(w http.ResponseWriter, r *http.Request) {
	RenderPage(w, "nostyle/index", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	lg.Log.Printf("nostyle/index.....")
	data := make(map[string]interface{})
	data["asdasd"] = "asdasd..."

	headers := w.Header()
	headers.Add("Content-Type", "text/html")

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if e := os.Chdir("tmpl/"); e != nil {
		panic(e)
	}

	page := template.Must(template.ParseFiles("nostyle/index.html"))

	if e := os.Chdir(dir); e != nil {
		panic(e)
	}

	lg.Log.Printf("Execute...")
	page.Execute(w, data)
}
