package music

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/justinas/nosurf"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/breadcrumbs"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/lg"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/model"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/views"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/web"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/music", Index)
	mux.HandleFunc("/music/artists", Artists)
	mux.HandleFunc("/music/artist", Artist)
	mux.HandleFunc("/music/album", Album)
	mux.HandleFunc("/music/items", Items)
	mux.HandleFunc("/music/filter", Filter)
	//mux.HandleFunc("/music/playsong", Playsong)
}

func Index(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	data["breadcrumbs"] = breadcrumbs.New().Append("Home", "/").AppendActive("Music")
	tmpls := []string{
		"base.html",
		"breadcrumbs.tmpl",
		"music/index.html"}
	web.Render(w, data, tmpls...)
}

func Items(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)

	data["A_Z"] = views.A_Z

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
}

func Artist(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)

	q := r.URL.Query()
	artist := q.Get("a")
	lg.Log.Printf("Artist() artist:%s", artist)
	data["artist"] = artist

	bcs := breadcrumbs.New().Append("Home", "/")
	bcs.Append("Music", "/music")
	bcs.Append("Artists", "/music/artists")
	bcs.AppendActive(artist)
	data["breadcrumbs"] = bcs

	albums, err := model.Albums(artist)
	if err != nil {
		panic(err)
	}
	data["albums"] = albums

	tmpls := []string{
		"base.html",
		"breadcrumbs.tmpl",
		"music/artist.html"}
	web.Render(w, data, tmpls...)
}

func Album(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	data, _ := web.TmplData(r)

	id, err := strconv.Atoi(q.Get("ai"))
	if err != nil {
		lg.Log.Printf("err:%s", err.Error())
	}
	album, err := model.AlbumById(id)
	if err != nil {
		lg.Log.Printf("AlbumById() err:%s", err.Error())
	}
	data["album"] = album

	bcs := breadcrumbs.New().Append("Home", "/")
	bcs.Append("Music", "/music")
	bcs.Append("Artists", "/music/artists")
	bcs.Append(album.AlbumArtist, fmt.Sprintf("/music/artist?a=%s", album.AlbumArtist))
	bcs.AppendActive(album.Title)
	data["breadcrumbs"] = bcs

	items, err := album.Items()
	if err != nil {
		lg.Log.Printf("err:%s", err.Error())
	}
	data["items"] = items

	tmpls := []string{
		"base.html",
		"breadcrumbs.tmpl",
		"music/album.html"}

	web.Render(w, data, tmpls...)
}

func Filter(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	bcs := breadcrumbs.New().Append("Home", "/").Append("Music", "/music")
	bcs.AppendActive("Filter")
	data["breadcrumbs"] = bcs
	data["A_Z"] = views.A_Z
	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"breadcrumbs.tmpl",
		"music/filter_results.html",
		"music/filter_filter.html",
		"a_z_select.tmpl",
		"music/filter.html"}
	web.Render(w, data, tmpls...)
}

func Artists(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	data["A_Z"] = views.A_Z
	data["token"] = nosurf.Token(r)

	bcs := breadcrumbs.New().Append("Home", "/")
	bcs.Append("Music", "/music")
	bcs.AppendActive("Artists")
	data["breadcrumbs"] = bcs

	tmpls := []string{
		"base.html",
		"breadcrumbs.tmpl",
		"a_z_select.tmpl",
		"music/artists.html"}

	if "POST" == r.Method {
		//views.LogFormData(r)
		artist_startswith := r.PostFormValue("artist_startswith")
		artists, err := model.GetArtists(artist_startswith)
		data["artists"] = artists
		data["artist_startswith"] = artist_startswith
		data["a_z_select_value"] = artist_startswith
		lg.Log.Printf("len artists:%v", len(artists))
		if err != nil {
			lg.Log.Printf("err:%v", err)
		}
	}

	web.Render(w, data, tmpls...)
}
