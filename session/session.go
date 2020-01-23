package session

import (
	"net/http"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
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
