package music

import (
	"net/http"
	"strconv"

	"github.com/justinas/nosurf"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/breadcrumbs"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/lg"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/model"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/views"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/web"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	data["breadcrumbs"] = breadcrumbs.New().Append("Home", "/").AppendActive("Music")
	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"breadcrumbs.tmpl",
		"music/index.html"}
	web.Render(w, data, tmpls...)
}

// old view render style doesnt work ...
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

func Artists(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	artists, err := model.GetAllArtists()
	if err != nil {
		panic(err)
	}
	data["artists"] = artists

	bcs := breadcrumbs.New().Append("Home", "/")
	bcs.Append("Music", "/music")
	bcs.AppendActive("Artists")
	data["breadcrumbs"] = bcs
	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"breadcrumbs.tmpl",
		"music/artists_pagination.tmpl",
		"music/artists.html"}
	web.Render(w, data, tmpls...)
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

	data["breadcrumbs"] = breadcrumbs.New().Append("Home", "/").AppendActive("Music")
	tmpls := []string{
		"base.html",
		"nav.tmpl",
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

func Music(w http.ResponseWriter, r *http.Request) {
	data, _ := web.TmplData(r)
	data["A_Z"] = views.A_Z
	data["token"] = nosurf.Token(r)
	tmpls := []string{
		"ttown/base.html",
		"a_z_select.tmpl",
		"ttown/music.html"}

	if "POST" == r.Method {
		//views.LogFormData(r)
		artist_startswith := r.PostFormValue("artist_startswith")
		artists, err := model.GetArtists(artist_startswith)
		data["artists"] = artists
		lg.Log.Printf("len artists:%v", len(artists))
		if err != nil {
			lg.Log.Printf("err:%v", err)
		}
	}

	web.Render(w, data, tmpls...)
}
