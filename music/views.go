package music

import (
	"net/http"
	"strconv"

	"oldcode.org/gow/breadcrumbs"
	"oldcode.org/gow/lg"
	"oldcode.org/gow/model"
	"oldcode.org/gow/views"
	"oldcode.org/gow/web"
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
	data["view"] = "filter";
	tmpls := []string{
		"base.html",
		"nav.tmpl",
		"breadcrumbs.tmpl",
		"music/filter_results.html",
		"music/filter_filter.html",
		"music/filter.html"}
	web.Render(w, data, tmpls...)
}

