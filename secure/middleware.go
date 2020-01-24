// Copyright (c) 2020 Paul Wisehart paul@oldcode.org
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

package secure

import (
	"net/http"

	"github.com/krsanky/gwww/account"
	"github.com/krsanky/gwww/lg"
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
