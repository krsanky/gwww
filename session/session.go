package session

import (
	"net/http"

	"github.com/alexedwards/scs"
	"github.com/alexedwards/scs/stores/pgstore"
	"github.com/krsanky/lg"
)

var Manager *scs.Manager

func Init() {
	store := pgstore.New(db.DB, 0)
	//engine := boltstore.New(sess_db.DB, 0)

	Manager = scs.NewManager(store)
	Manager.Name("les_super_sesh")
	Manager.Persist(true)
}

func GetWithDefault(r *http.Request, key string, def string) (value string, err error) {
	lg.Log.Printf("get sess value for k:%s", key)
	value, err = Manager.Load(r).GetString(key)
	if err != nil {
		lg.Log.Printf("err: %s", err.Error())
	}
	lg.Log.Printf("value: %s", value)
	if value == "" {
		value = def
	}
	return
}
