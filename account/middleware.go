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

package account

import (
	"net/http"

	"github.com/krsanky/gwww/lg"
	"github.com/krsanky/gwww/session"
)

var UserIdString = "_account__user_id_"

func AddUser(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user_id := session.Session.GetInt(r.Context(), UserIdString)
		//		user_id, err := session.Manager.Load(r).GetInt(UserIdString)
		//		if err != nil {
		//			lg.Log.Printf("AddUser() err:%s", err.Error())
		//		}

		user, err := GetUserById(user_id)
		if false && (err != nil) {
			// might just be no results
			lg.Log.Printf("AddUser() err2:%s", err.Error())
		}

		if user != nil {
			ctx := ContextWithUser(r.Context(), user)
			lg.Log.Printf("AddUser() added user_id:%d to r.Context()", user.Id)
			h.ServeHTTP(w, r.WithContext(ctx))
		} else {
			lg.Log.Printf("AddUser() NO USER")
			h.ServeHTTP(w, r)
		}
	})
}
