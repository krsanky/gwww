package views

import (
	"net/http"

	svg "github.com/ajstarks/svgo"
	"oldcode.org/gow/db"
	"oldcode.org/gow/lg"
	"oldcode.org/gow/model"
	"oldcode.org/gow/web"
)

func Index(w http.ResponseWriter, r *http.Request) {
	web.RenderPage(w, "index", nil)
}

func Items(w http.ResponseWriter, r *http.Request) {
	lg.Log.Printf("views.Items.....")

	odb := db.GetOpenDB()
	defer odb.Close()

	data := make(map[string]interface{})

	item := model.Item{}
	odb.Limit(15).First(&item, 10)
	data["item"] = item

	var items []model.Item
	odb.Limit(15).Find(&items)
	data["items"] = items

	web.RenderPage(w, "items", data, "item_search.tmpl")
}

func Page3(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title1 string
		Items  []string
	}{
		Title1: "My page",
		Items: []string{
			"My photos",
			"My blog",
		},
	}

	web.RenderPage(w, "page3", data)
}

func Circle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(500, 500)
	s.Circle(250, 250, 125, "fill:none;stroke:black")
	s.End()
}
