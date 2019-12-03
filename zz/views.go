package zz

import (
	"net/http"

	"oldcode.org/repo/go/gow/web"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/zz", Index)
}

func Index(w http.ResponseWriter, r *http.Request) {
	tmpls := []string{"zzbase.html"}
	web.Render(w, nil, tmpls...)
}
