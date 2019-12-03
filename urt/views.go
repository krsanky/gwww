package urt

import (
	"net/http"

	"github.com/krsanky/go-urt-server-query/urt"
	"oldcode.org/repo/go/gow/breadcrumbs"
	"oldcode.org/repo/go/gow/lg"
	"oldcode.org/repo/go/gow/views"
	"oldcode.org/repo/go/gow/web"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/urt", Index)
	mux.HandleFunc("/urt/rdio", RadioFrame)
	mux.HandleFunc("/urt/rdo", RadioFrame)
	mux.HandleFunc("/urt/radio", Radio)
	mux.HandleFunc("/urt/radio/key", RadioKey)
	mux.HandleFunc("/urt/servers", Servers)
	mux.HandleFunc("/urt/urtctf", Urtctf)
}

func Index(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	data["breadcrumbs"] = breadcrumbs.New().Append("Home", "/").AppendActive("URT")
	tmpls := []string{
		"urt/base.html",
		"nav.tmpl",
		"breadcrumbs.tmpl",
		"urt/index.html"}
	web.Render(w, data, tmpls...)
}

func RadioFrame(w http.ResponseWriter, r *http.Request) {
	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"urt/radio_frame.html"}
	web.Render(w, nil, tmpls...)
}
func Radio(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	bcs := breadcrumbs.New().Append("Home", "/")
	bcs.Append("URT", "/urt")
	bcs.AppendActive("Radio")
	data["breadcrumbs"] = bcs

	if "POST" == r.Method {
		views.LogFormData(r)
	}

	tmpls := []string{
		"urt/base.html",
		"nav.tmpl",
		"breadcrumbs.tmpl",
		"urt/numpad.tmpl",
		"urt/radio.html"}
	web.Render(w, data, tmpls...)
}

func RadioKey(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	bcs := breadcrumbs.New().Append("Home", "/")
	bcs.Append("URT", "/urt")
	bcs.AppendActive("Radio Key")
	data["breadcrumbs"] = bcs
	tmpls := []string{
		"urt/base.html",
		"nav.tmpl",
		"breadcrumbs.tmpl",
		"urt/radio_key.html"}
	web.Render(w, data, tmpls...)
}

func Servers(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)

	UrtCtf(data)

	bcs := breadcrumbs.New().Append("Home", "/")
	bcs.Append("URT", "/urt")
	bcs.AppendActive("URT Servers")
	data["breadcrumbs"] = bcs
	tmpls := []string{
		"urt/base.html",
		"nav.tmpl",
		"breadcrumbs.tmpl",
		"urt/servers.html"}
	web.Render(w, data, tmpls...)
}

func UrtCtf(page_data map[string]interface{}) {
	data, err := urt.GetRawStatus("216.52.148.134:27961") // urtctf
	if err != nil {
		lg.Log.Printf("ERR:%s", err)
		return
	}
	//fmt.Println(string(data))

	_, err = urt.ServerVars(data)
	if err != nil {
		lg.Log.Printf("ERR:%s", err)
		return
	}
	//fmt.Printf("len vars:%d\n", len(vars))

	players, err := urt.Players(data)
	if err != nil {
		lg.Log.Printf("ERR:%s", err)
		return
	}
	//fmt.Printf("len players:%d\n", len(players))
	//for _, p := range players {
	//	fmt.Println(p)
	//}

	page_data["players"] = players
}

func Urtctf(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	data["breadcrumbs"] = breadcrumbs.New().Append("Home", "/").AppendActive("Urtctf")
	UrtCtf(data)
	tmpls := []string{
		"base.html",
		"breadcrumbs.tmpl",
		"urt/urtctf.html"}
	web.Render(w, data, tmpls...)
}
