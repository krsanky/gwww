package account

import (
	"net/http"

	"oldcode.org/repo/go/gow/lg"
	"oldcode.org/repo/go/gow/session"
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

