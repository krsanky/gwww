package account

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"oldcode.org/gow/db"
	"oldcode.org/gow/lg"
	"oldcode.org/gow/session"
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
	return fmt.Sprintf("[id:%d username:%s email:%s pw:%s act:%v]",
		u.Id, u.Username, u.Email, u.Password, u.Is_active)
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
	sess := session.Manager.Load(r)
	err := sess.PutInt(w, UserIdString, u.Id)
	if err != nil {
		panic(err)
	}
}

// user is implied as current logged in user
func Logout(w http.ResponseWriter, r *http.Request) {
	session.Manager.Load(r).Clear(w)
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
	sess := session.Manager.Load(r)
	if err == nil {
		lg.Log.Printf("user email:%s", u.Email)
		if PasswordMatch(u, password) {
			lg.Log.Printf("AuthUser MATCH")
			//LoginUser(w, r, u)
			err := sess.PutInt(w, UserIdString, u.Id)
			if err != nil {
				panic(err)
			}
			return true
		} else {
			lg.Log.Printf("AuthUser NO MATCH")
		}
	}

	lg.Log.Printf("AuthUser err:%s", err)
	sess.PopInt(w, UserIdString)
	sess.PopString(w, "user_username")
	return false
}

// This implies a user can just be an email address.
func (u *User) Register() {

}
