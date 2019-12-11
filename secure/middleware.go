package secure

import (
	"net/http"

	"oldcode.org/repo/go/gow/account"
	"oldcode.org/repo/go/gow/lg"
)
/*
three options to secure a page:
1. path based.  This is easy but less secure, because if
  the view somehow had a different url/path to it, this won't secure it.
2.View based internally.  The view itself checks for valid user.
This is most secure, but have to repaeat for each view and in the view.
3. In the view code, but at the routing mux setup.
This is in between other two.  It's pretty secure, because you
can see all the routes assigned and secure there.
... It seems like if you are gonna do this you might as well do it in the view.
... I guess this decouples the view from the security, like maybe a backend site
with custom auth would not want the view secured here, but the

UNKNOWN: a mux and sub-mux based approach.  Block a whole sub mux for ease, and full
security.
*/

func SuperOnlyMiddleware(h http.Handler) http.Handler {
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

func SuperOnly(handler func(http.ResponseWriter, *http.Request)) http.Handler {
	return SuperOnlyMiddleware(http.HandlerFunc(handler))
}

