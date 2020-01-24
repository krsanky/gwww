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

package session

import (
	"net/http"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
	"github.com/krsanky/gwww/db"
	"github.com/krsanky/gwww/lg"
)

var Session *scs.Session

func Init() {
	Session = scs.NewSession()
	Session.Store = postgresstore.New(db.DB)

	Session.Cookie.Name = "les_super_sesh"
	Session.Cookie.Persist = true
	lg.Log.Printf("session.Init()...")
}

func GetWithDefault(r *http.Request, key string, def string) (value string, err error) {
	lg.Log.Printf("get sess value for k:%s", key)

	value = Session.GetString(r.Context(), key)
	//value, err = Manager.Load(r).GetString(key)
	//	if err != nil {
	//		lg.Log.Printf("err: %s", err.Error())
	//	}
	lg.Log.Printf("value: %s", value)
	if value == "" {
		value = def
	}
	return
}
