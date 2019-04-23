package view

import (
	"net/http"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/account/login", Login)

	//func Routes(r *router.Router) {
	//	r.Get("/account/login", http.HandlerFunc(LoginPage))
	//	r.Post("/account/login", http.HandlerFunc(LoginPage))
	//	r.Get("/account/logout", http.HandlerFunc(LogoutPage))
	//}
}
