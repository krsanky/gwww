package views

import (
	"net/http"
	"strconv"
	"strings"

	"oldcode.org/gow/lg"
	"oldcode.org/gow/model"
	"oldcode.org/gow/web"
)

var A_Z = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I",
	"J", "K", "L", "M", "N", "O", "P", "Q", "R",
	"S", "T", "U", "V", "W", "X", "Y", "Z"}

func ShowFormData(r *http.Request) {
	r.ParseForm()
	for k, v := range r.Form {
		lg.Log.Printf("k: %s v: %s", k, strings.Join(v, ""))
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	if r.URL.Path != "/" {
		lg.Log.Printf("views.Index(): NOT FOUND %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}
	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"index.html"}
	web.Render(w, data, tmpls...)
}

func Msg(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"msg.html"}
	web.Render(w, data, tmpls...)
}

func Items(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)

	data["A_Z"] = A_Z

	artist := r.FormValue("artist")
	artist_startswith := r.FormValue("artist_startswith")
	lg.Log.Printf("artist[%s] artist_startswith[%s]", artist, artist_startswith)
	data["artist"] = artist
	data["artist_startswith"] = artist_startswith

	artists, err := model.GetArtists(artist_startswith)
	if err != nil {
		panic(err)
	}
	data["artists"] = artists

	//	albums, err := artist.GetAlbums()
	//	if err != nil {
	//		panic(err)
	//	}
	//	data["albums"] = albums

	var items []model.Item
	data["items"] = items

	//web.RenderPage(w, "items", data)
}

func Artists(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	artists, err := model.GetAllArtists()
	if err != nil {
		panic(err)
	}
	data["artists"] = artists

	//web.RenderPage(w, "artists", data)
}

func Artist(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)

	q := r.URL.Query()
	artist := model.Artist{q.Get("a")}
	lg.Log.Printf("Artist() artist:%s", artist.Name)
	data["artist"] = artist

	albums, err := artist.Albums()
	if err != nil {
		panic(err)
	}
	data["albums"] = albums

	//web.RenderPage(w, "artist", data)
}

func Album(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	data, _ := web.TmplData(r)

	id, err := strconv.Atoi(q.Get("ai"))
	if err != nil {
		lg.Log.Printf("err:%s", err.Error())
	}
	album, err := model.AlbumByID(id)
	if err != nil {
		lg.Log.Printf("AlbumByID() err:%s", err.Error())
	}
	data["album"] = album

	items, err := album.Items()
	if err != nil {
		lg.Log.Printf("AlbumByID() err:%s", err.Error())
	}
	data["items"] = items

	//web.RenderPage(w, "album", data)
}
