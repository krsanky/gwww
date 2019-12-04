package secure

import (
	"net/http"

	"oldcode.org/repo/go/gow/account"
	"oldcode.org/repo/go/gow/lg"
)

// https://golang.org/pkg/net/http/#HandlerFunc
// https://golang.org/pkg/net/http/#Redirect

func SuperOnly(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lg.Log.Println("SuperOnly Before")
		defer lg.Log.Println("SuperOnly After")

		u, ok := account.UserFromContext(r.Context())
		if (u == nil) || (!ok) {
			http.Redirect(w, r, "/account/login2", 302)
		} else {
			if u.Is_super {
				h.ServeHTTP(w, r)
			} else {
				http.Redirect(w, r, "/account/login2", 302)
			}
		}
	})
}

func SuperOnlyFunc(handler func(http.ResponseWriter, *http.Request)) http.Handler {
	return SuperOnly(http.HandlerFunc(handler))
}
