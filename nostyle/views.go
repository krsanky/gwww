package nostyle

import (
	"net/http"
	"os"
	"text/template"

	"oldcode.org/gow/lg"
)

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
