package radio

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"oldcode.org/repo/go/gow/breadcrumbs"
	"oldcode.org/repo/go/gow/lg"
	"oldcode.org/repo/go/gow/web"
)

/*
{"icestats":
	{"admin":"radio@oldcode.org","host":"radio.oldcode.org",
	"location":"Vermont-or-New-Jersey","server_id":"Icecast 2.4.4",
	"server_start":"Sun, 01 Dec 2019 23:14:27 +0000",
	"server_start_iso8601":"2019-12-01T23:14:27+0000",
	"source":
		{"audio_info":"bitrate=128;channels=2;samplerate=44100","bitrate":128,"channels":2,
		"genre":"RockNRoll","listener_peak":2,"listeners":0,"listenurl":"http://radio.oldcode.org:8000/stream",
		"samplerate":44100,"server_description":"This is a stream description",
		"server_name":"My Stream","server_type":"audio/mpeg",
		"server_url":"http://www.oddsock.org","stream_start":"Thu, 05 Dec 2019 17:29:47 +0000",
		"stream_start_iso8601":"2019-12-05T17:29:47+0000","title":"King Diamond - Tea","dummy":null}}}

*/

type Icestats struct {
	Location string `json:"location"`
}

type JsonTop struct {
	Stats Icestats `json:"icestats"`
}

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/radio/stat", Status)
}

func Status(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	bcs := breadcrumbs.New().Append("Home", "/").Append("Radio", "/radio")
	bcs.AppendActive("Status")
	data["breadcrumbs"] = bcs

	url := "http://oldcode.org:8000/status-json.xsl"
	var res *http.Response
	var err error
	var body []byte
	stats := JsonTop{}

	client := http.Client{Timeout: time.Second * 2}
	res, err = client.Get(url)
	if err != nil {
		lg.Log.Printf("ERR:%v", err)
		goto Render
	}
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		lg.Log.Printf("ERR:%v", err)
		goto Render
	}
	data["res_body"] = string(body)
	json.Unmarshal(body, &stats)
	//lg.Log.Printf("stats.Location:%s", stats.Location)
	data["stats"] = stats

Render:
	tmpls := []string{
		"base.html",
		"breadcrumbs.tmpl",
		"radio/status.html"}
	web.Render(w, data, tmpls...)
}
