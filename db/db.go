package db

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"oldcode.org/repo/go/gow/lg"
	"oldcode.org/repo/go/gow/settings"
)

var DB *sql.DB
var DBX *sqlx.DB

func handle_err(e error) {
	if e != nil {
		panic(e)
	}
}

func Init() {
	lg.Log.Printf("init pg db start ...")

	password := settings.GetString("db.password")
	user := settings.GetString("db.user")
	db_name := settings.GetString("db.name")

	var err error
	connect_string := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		user, password, db_name)

	DB, err = sql.Open("postgres", connect_string)
	handle_err(err)
	handle_err(DB.Ping())

	DBX, err = sqlx.Open("postgres", connect_string)
	handle_err(err)
	handle_err(DBX.Ping())
}

func Drivers() {
	for _, d := range sql.Drivers() {
		fmt.Printf("driver:%s\n", d)
	}
}

func TestDB() {
	fmt.Printf("db.TestDB()\n")
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
		lg.Log.Printf("name:%s\n", name)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
}
