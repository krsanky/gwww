package formstuff

import (
	"net/http"

	"oldcode.org/gow/lg"
	"oldcode.org/gow/web"
)


func Index(w http.ResponseWriter, r *http.Request) {
	lg.Log.Printf("formstuff/index.....")
	web.RenderPage(w, "formstuff/index", nil)
}

