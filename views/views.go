package views

import (
	"net/http"
	"strconv"

	"oldcode.org/gow/lg"
	"oldcode.org/gow/model"
	"oldcode.org/gow/tmplutil"
	"oldcode.org/gow/web"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	web.RenderPage(w, "index", nil)
}

func Items(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	data["A_Z"] = tmplutil.A_Z

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

	var items []model.Item
	data["items"] = items

	web.RenderPage(w, "items", data)
}

func Artists(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	artists, err := model.GetAllArtists()
	if err != nil {
		panic(err)
	}
	data["artists"] = artists

	web.RenderPage(w, "artists", data)
}

func Artist(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	q := r.URL.Query()
	artist := model.Artist{q.Get("a")}
	lg.Log.Printf("Artist() artist:%s", artist.Name)
	data["artist"] = artist

	albums, err := artist.Albums()
	if err != nil {
		panic(err)
	}
	data["albums"] = albums

	web.RenderPage(w, "artist", data)
}

func Album(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	data := make(map[string]interface{})
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

	web.RenderPage(w, "album", data)
}

func Track(w http.ResponseWriter, r *http.Request) {
	web.RenderPage(w, "track", nil)
}
