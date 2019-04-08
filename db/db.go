package db

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var beets_db_file = "/home/wise/go/src/oldcode.org/gow/beets.db"

func Drivers() {
	for _, d := range sql.Drivers() {
		fmt.Printf("driver:%s\n", d)
	}
}

func GetOpenDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", beets_db_file)
	if err != nil {
		panic(err)
	}
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}
	return db
}

func TestSql() {
	db, err := sql.Open("sqlite3", beets_db_file)
	if err != nil {
		panic(err)
	}
	stmt, err := db.Prepare("SELECT DISTINCT albumartist FROM albums")
	if err != nil {
		panic(err)
	}
	stmt.Exec()

	rows, err := db.Query("SELECT DISTINCT albumartist FROM albums ORDER by albumartist")
	if err != nil {
		panic(err)
	}
	var s string
	for rows.Next() {
		rows.Scan(&s)
		fmt.Printf("row albumartist:%s\n", s)
	}
}

func TestGormSql() {
	fmt.Printf("test (sql) beets: %s\n", beets_db_file)
	db, err := gorm.Open("sqlite3", beets_db_file)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	//var ODB *sql.DB
	ODB := db.DB()
	if err = ODB.Ping(); err != nil {
		panic("failed to ping database")
	}

	type Result struct {
		Id    int
		Title string
	}
	var result Result
	db.Raw("SELECT id, title from items where id = 10").Scan(&result)
	fmt.Printf("id:%d title:%s\n", result.Id, result.Title)
}
