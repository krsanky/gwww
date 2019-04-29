package music

import "net/http"

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/music", Index)
	mux.HandleFunc("/music/artists", Artists)
	mux.HandleFunc("/music/artist", Artist)
	mux.HandleFunc("/music/album", Album)
	mux.HandleFunc("/music/items", Items)
}
