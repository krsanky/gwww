package view

import (
	"net/http"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/account/login", Login)
	mux.HandleFunc("/account/logout", Logout)
}
