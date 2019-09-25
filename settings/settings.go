package settings

import (
	"github.com/pelletier/go-toml"
)

var settings *toml.Tree

func Init(sfile string) {
	var err error
	settings, err = toml.LoadFile(sfile)
	if err != nil {
		panic(err)
	}
}

func GetInt(key string) int64 {
	return settings.Get(key).(int64)
}

func GetString(key string) string {
	return settings.Get(key).(string)
}

func ReadEmailHost() string {
	return GetString("email.host")
}
