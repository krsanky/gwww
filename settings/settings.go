package settings

import (
	"fmt"

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

func check() {
	if settings == nil {
		panic(fmt.Errorf("settings: need settings.Init(<settings-file>)\n"))
	}
}

func GetInt(key string) int64 {
	check()
	return settings.Get(key).(int64)
}

func GetString(key string) string {
	check()
	return settings.Get(key).(string)
}

func ReadEmailHost() string {
	check()
	return GetString("email.host")
}
