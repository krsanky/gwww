package views

import (
	"net/http"

	"oldcode.org/gow/db"
	"oldcode.org/gow/lg"
	"oldcode.org/gow/model"
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
	lg.Log.Printf("views.Items.....")

	odb := db.GormOpenDB()
	defer odb.Close()

	data := make(map[string]interface{})

	item := model.Item{}
	odb.Limit(15).First(&item, 10)
	data["item"] = item

	var items []model.Item
	odb.Limit(15).Find(&items)
	data["items"] = items

	web.RenderPage(w, "items", data)
}

func Artists(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	artists, err := model.GetArtists()
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




