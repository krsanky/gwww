package db

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"oldcode.org/gow/lg"
	"oldcode.org/gow/settings"
)

var BeetsDB *sqlx.DB
var DB *sql.DB
var DBX *sqlx.DB

func InitDB() {
	lg.Log.Printf("init pg db start ...")

	password := settings.GetString("db.password")
	user := settings.GetString("db.user")
	db_name := settings.GetString("db.name")

	var err error
	connect_string := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		user, password, db_name)

	DB, err = sql.Open("postgres", connect_string)
	if err != nil {
		panic(err)
	}
	if err = DB.Ping(); err != nil {
		panic(err)
	}

	DBX, err = sqlx.Open("postgres", connect_string)
	if err != nil {
		panic(err)
	}
	if err = DBX.Ping(); err != nil {
		panic(err)
	}

	//TestDB()
}

func Open() {
	beets_db_file := settings.GetString("db.beets_file")

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

func TestDB() {
	var err error
	if err = DB.Ping(); err != nil {
		panic(err)
	}

	rows, err := DB.Query("SELECT name FROM test1")
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			panic(err)
		}
		lg.Log.Printf("name:%s", name)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
}
