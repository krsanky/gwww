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

package account

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/krsanky/gwww/db"
	"github.com/krsanky/gwww/lg"
	"github.com/krsanky/gwww/session"
)

type User struct {
	Id         int
	Email      string
	Password   string
	Username   sql.NullString
	First_name string
	Last_name  string
	Is_super   bool
	Is_staff   bool
	Is_active  bool
}

// write whatever we have to new record
func (u *User) SaveNew() error {
	lg.Log.Printf("account.SaveNew() email:%s", u.Email)

	sql := `INSERT INTO account
(password, is_superuser, username, first_name, last_name, 
email, is_staff, is_active)
VALUES ($1, $2, $3, $4, $5, 
$6, $7, $8)`
	_, err := db.DBX.Exec(sql,
		u.Password, u.Is_super, u.Username, u.First_name, u.Last_name,
		u.Email, u.Is_staff, u.Is_active)
	return err
}

func (u *User) String() string {
	return fmt.Sprintf("[id:%d username:%v email:%s pw:%s act:%t su:%t]",
		u.Id, u.Username, u.Email, u.Password, u.Is_active, u.Is_super)
}

func (u *User) Url() template.HTML {
	return template.HTML(fmt.Sprintf("<a href='/xyz/user?u=%d'>%s</a>", u.Id, u.Email))
}

func GetUsers() ([]User, error) {
	var us []User
	db := db.DBX.Unsafe()
	rows, err := db.Queryx("SELECT * FROM account")
	for rows.Next() {
		var u User
		err = rows.StructScan(&u)
		if err != nil {
			return nil, err
		}
		us = append(us, u)
	}
	return us, nil
}

func GetUserById(id int) (*User, error) {
	row := db.DB.QueryRow(`
SELECT id, password, is_superuser, username, first_name, last_name, 
email, is_staff, is_active
FROM account WHERE id=$1`, id)

	u := &User{}
	err := row.Scan(
		&u.Id,
		&u.Password,
		&u.Is_super,
		&u.Username,
		&u.First_name,
		&u.Last_name,
		&u.Email,
		&u.Is_staff,
		&u.Is_active)

	// add check for ErrNoRows ....
	if err != nil {
		return nil, err
	}
	return u, nil
}

func GetUserByName(username string) (*User, error) {
	row := db.DB.QueryRow(`
SELECT id, password, is_superuser, username, first_name, last_name, 
email, is_staff, is_active
FROM account WHERE username=$1`, username)

	u := &User{}
	err := row.Scan(
		&u.Id,
		&u.Password,
		&u.Is_super,
		&u.Username,
		&u.First_name,
		&u.Last_name,
		&u.Email,
		&u.Is_staff,
		&u.Is_active)

	// add check for ErrNoRows ....
	if err != nil {
		return nil, err
	} else {
		return u, nil
	}
}

func GetUserByEmail(email string) (*User, error) {
	db := db.DBX.Unsafe()
	row := db.QueryRowx(`SELECT * FROM account WHERE email=$1`, email)
	lg.Log.Printf("row:%v", row)
	u := &User{}
	err := row.StructScan(u)

	//var ErrNoRows = errors.New("sql: no rows in result set")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	} else {
		return u, nil
	}
}

func PasswordMatch(u *User, formpassword string) bool {
	lg.Log.Printf("user:%s", u)
	lg.Log.Printf("u-pw:%s f-pw:%s", u.Password, formpassword)
	return u.Password == formpassword
}

func (u *User) Login(w http.ResponseWriter, r *http.Request) {
	LoginUser(w, r, u)
}

func LoginUser(w http.ResponseWriter, r *http.Request, u *User) {
	session.Session.Put(r.Context(), UserIdString, u.Id)
}

// user is implied as current logged in user
func Logout(w http.ResponseWriter, r *http.Request) {
	err := session.Session.Destroy(r.Context())
	if err != nil {
		lg.Log.Printf("sess.Destroy() error")
	}
}

func (u *User) Auth(w http.ResponseWriter, r *http.Request, password string) bool {
	if u.Password == password {
		u.Login(w, r)
		return true
	}
	return false
}

func AuthUser(w http.ResponseWriter, r *http.Request, username string, password string) bool {
	u, err := GetUserByName(username)
	if err == nil {
		lg.Log.Printf("user email:%s", u.Email)
		if PasswordMatch(u, password) {
			lg.Log.Printf("AuthUser MATCH")
			session.Session.Put(r.Context(), UserIdString, u.Id)
			return true
		} else {
			lg.Log.Printf("AuthUser NO MATCH")
		}
	}

	lg.Log.Printf("AuthUser err:%s", err)
	session.Session.PopInt(r.Context(), UserIdString)
	session.Session.PopString(r.Context(), "user_username")
	return false
}

// This implies a user can just be an email address.
func (u *User) Register() {

}
