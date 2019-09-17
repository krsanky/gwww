package settings

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml"
)

var settings *toml.Tree

func Init() {
	var err error
	d, _ := os.Getwd()
	fmt.Printf("settings dir:%s\n", d)
	settings, err = toml.LoadFile("settings.toml")
	if err != nil {
		panic(err)
	}
}

func GetInt(key string) int64 {
	if settings == nil {
		Init()
	}
	return settings.Get(key).(int64)
}

func GetString(key string) string {
	if settings == nil {
		Init()
	}
	return settings.Get(key).(string)
}

func ReadEmailHost() string {
	return GetString("email.host")
}

//config, _ := toml.Load(`
//[postgres]
//user = "pelletier"
//password = "mypassword"`)
//// retrieve data directly
//user := config.Get("postgres.user").(string)
//// or using an intermediate object
//postgresConfig := config.Get("postgres").(*toml.Tree)
//password := postgresConfig.Get("password").(string)
