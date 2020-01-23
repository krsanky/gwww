package lg

import (
	"fmt"
	"log"
	"os"
)

var Log *log.Logger

func init() {
	log_file := "log.txt"
	f, err := os.OpenFile(log_file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("error opening file:%s err:%s", log_file, err.Error()))
	}

	Log = log.New(f, ":", log.Lshortfile)
	Log.Println("opening logfile ...")
}
