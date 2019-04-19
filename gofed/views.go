package gofed

import (
	"net/http"

	"oldcode.org/gow/web"
)

func Index(w http.ResponseWriter, r *http.Request) {
	web.Render(w, nil, "gofed/base.html", "gofed/index.html")
}
