package views

import (
	"net/http"

	"oldcode.org/gow/db"
	"oldcode.org/gow/lg"
	"oldcode.org/gow/model"
	"oldcode.org/gow/web"
)

func Items(w http.ResponseWriter, r *http.Request) {
	lg.Log.Printf("views.Index.....")

	odb := db.GetOpenDB()
	defer odb.Close()

	data := make(map[string]interface{})

	item := model.Item{}
	odb.First(&item, 10)
	data["item"] = item

	var items []model.Item
	odb.Find(&items)
	data["items"] = items

	web.RenderPage(w, "items", data)
}

