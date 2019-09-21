package email

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"text/template"

	"oldcode.org/home/wise/repo/go/oldcode.org/gow/lg"
	"oldcode.org/home/wise/repo/go/oldcode.org/gow/settings"
)

var (
	port     int64
	host     string
	user     string
	password string
	pages    map[string]*template.Template
)

func Init() {
	port = settings.GetInt("email.port")
	host = settings.GetString("email.host")
	user = settings.GetString("email.user")
	password = settings.GetString("email.password")
}

func Send_test() {
	lg.Log.Printf("Send_test() start...")
	if port == 0 {
		Init()
	}

	auth := smtp.PlainAuth("", user, password, host)
	hostp := fmt.Sprintf("%s:%d", host, port)
	//fmt.Printf("hostp:%s\n", hostp)
	to := []string{"paul@oldcode.org"}
	msg := []byte("To: paul@oldcode.org\r\n" +
		"Subject: Gow Gophers\r\n" +
		"\r\n" +
		"123 This is the email body.\r\n")

	//asd123 is the real from and sender
	err := smtp.SendMail(hostp, auth, "gow+test@oldcode.org", to, msg)
	if err != nil {
		lg.Log.Printf("Send_test(): error sending email\n")
	}
}

func registerPage(page string, tmpls []string) error {
	dir, err := os.Getwd()
	if err != nil {
		return (err)
	}
	if e := os.Chdir("tmpl/"); e != nil {
		return (e)
	}

	t := template.New(page)

	_, err = t.ParseFiles(tmpls...)
	if err != nil {
		return (err)
	}
	pages[page] = t

	if e := os.Chdir(dir); e != nil {
		return (e)
	}
	return nil
}

// Send is like web.Render but with email...
//func Render(w http.ResponseWriter, data interface{}, tmpls ...string) {
func Send(to []string, data interface{}, tmpls ...string) {

	var err error
	page := tmpls[len(tmpls)-1]
	if _, tmpl_exists := pages[page]; !tmpl_exists {
		if err = registerPage(tmpls[len(tmpls)-1], tmpls); err != nil {
			panic(err)
		}
	}

	var buf bytes.Buffer
	if err = pages[page].Execute(&buf, data); err != nil {
		panic(err)
	}
	//lg.Log.Printf("BUF:%s", buf) wrong formt specifier
}
