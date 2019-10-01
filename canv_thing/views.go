package music

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/canv-thing", Index)
}

func Index(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	tmpls := []string{
		"ttown/base.html",
		"canv_thing/index.html"}
	web.Render(w, data, tmpls...)
}

