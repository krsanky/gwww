package db

import (
	"database/sql"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jmoiron/sqlx"
)

var beets_db_file = "/home/wise/go/src/oldcode.org/gow/beets.db"
var BeetsDB *sqlx.DB
var DB *sql.DB

func Open() {
	var err error
	BeetsDB, err = sqlx.Open("sqlite3", beets_db_file)
	if err != nil {
		panic(err)
	}
}

func Drivers() {
	for _, d := range sql.Drivers() {
		fmt.Printf("driver:%s\n", d)
	}
}

func GetRawArtists() ([]string, error) {
	if BeetsDB == nil {
		Open()
	}

	rows, err := BeetsDB.Query(`
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
