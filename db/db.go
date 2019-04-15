package db

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jmoiron/sqlx"
)

var beets_db_file = "/home/wise/go/src/oldcode.org/gow/beets.db"
var DB *sqlx.DB

func Open() {
	var err error
	DB, err = sqlx.Open("sqlite3", beets_db_file)
	if err != nil {
		panic(err)
	}
}

func Drivers() {
	for _, d := range sql.Drivers() {
		fmt.Printf("driver:%s\n", d)
	}
}

func GormOpenDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", beets_db_file)
	if err != nil {
		panic(err)
	}
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}
	return db
}

func GetRawArtists() ([]string, error) {
	if DB == nil {
		Open()
	}

	rows, err := DB.Query(`
SELECT DISTINCT albumartist 
FROM albums 
WHERE albumartist <> ''
ORDER by albumartist
`)
	if err != nil {
		return nil, err
	}

	var s string
	artists := make([]string, 0)
	for rows.Next() {
		rows.Scan(&s)
		artists = append(artists, s)
	}
	return artists, nil
}

func TestSql() {
	artists, err := GetRawArtists()
	if err != nil {
		panic(err)
	}
	for a := range artists {
		fmt.Printf("row albumartist:%s\n", a)
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
