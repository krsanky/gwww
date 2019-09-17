package zz

import (
	"net/http"

	"oldcode.org/home/wise/repo/go/oldcode.org/gow/web"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/zz", Index)
}

func Index(w http.ResponseWriter, r *http.Request) {
	tmpls := []string{"zzbase.html"}
	web.Render(w, nil, tmpls...)
}
