package web

import (
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/template"

	"oldcode.org/gow/account"
	lg "oldcode.org/gow/lg"
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

func init() {
	_tmpls = make(map[string]*template.Template)
	GlobalFuncMap = template.FuncMap{
		"input_checked": InputChecked,
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

	page := tmpls[len(tmpls)-1]

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

		//_tmpls[page] = template.Must(template.ParseFiles(tmpls...))
		
		t, _ := template.ParseFiles(tmpls...)
		_tmpls[page] = t



		//_tmpls[page] = template.Must(template.New(page).Funcs(GlobalFuncMap).ParseFiles(tmpls...))

		//		_tmpls[page] = template.New(page).Funcs(GlobalFuncMap)
		//		_, err = _tmpls[page].ParseFiles(tmpls...)
		//		if err != nil {
		//			panic(err)
		//		}

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

func CookieTest() {

}
