package account

import (
	"fmt"
	"net/http"

	"oldcode.org/gow/db"
	"oldcode.org/gow/lg"
	"oldcode.org/gow/session"
)

type User struct {
	Id           int
	Password     string
	Is_superuser bool
	Username     string
	First_name   string
	Last_name    string
	Email        string
	Is_staff     bool
	Is_active    bool
	//Timezone carchar(128)
}

func (u *User) String() string {
	return fmt.Sprintf("[id:%d username:%s email:%s pw:%s]", u.Id, u.Username, u.Email, u.Password)
}

func (u *User) Url() string {
	return fmt.Sprintf("<a href='/xyz/user/%d'>%s</a>", u.Id, u.Email)
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
		&u.Is_superuser,
		&u.Username,
		&u.First_name,
		&u.Last_name,
		&u.Email,
		&u.Is_staff,
		&u.Is_active)

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
		&u.Is_superuser,
		&u.Username,
		&u.First_name,
		&u.Last_name,
		&u.Email,
		&u.Is_staff,
		&u.Is_active)

	if err != nil {
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

func AuthUser(w http.ResponseWriter, r *http.Request, username string, password string) bool {
	u, err := GetUserByName(username)
	sess := session.Manager.Load(r)
	if err == nil {
		lg.Log.Printf("user email:%s", u.Email)
		if PasswordMatch(u, password) {
			lg.Log.Printf("AuthUser MATCH")
			err := sess.PutInt(w, "user_id", u.Id)
			if err != nil {
				panic(err)
			}
			return true
		} else {
			lg.Log.Printf("AuthUser NO MATCH")
		}
	}

	lg.Log.Printf("AuthUser err:%s", err)
	sess.PopInt(w, "user_id")
	sess.PopString(w, "user_username")
	return false
}
