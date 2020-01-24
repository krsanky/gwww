// Copyright (c) 2020 Paul Wisehart paul@oldcode.org
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

package db

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/krsanky/gwww/lg"
	"github.com/krsanky/gwww/settings"
	_ "github.com/lib/pq"
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
