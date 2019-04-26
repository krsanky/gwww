package email

import (
	"fmt"
	"net/smtp"
)

var (
	host     string
	port     int
	user     string
	password string
	use_ssl  bool
)

func init() {
	host = "smtp.fastmail.com"
	port = 587
	user = "wisehart@fastmail.fm"
	password = "n5jkbctpygfqy3pv"
	use_ssl = true
}

func Send_test() {
	auth := smtp.PlainAuth("", user, password, host)
	hostp := fmt.Sprintf("%s:%d", host, port)
	//fmt.Printf("hostp:%s\n", hostp)
	to := []string{"paul@oldcode.org"}
	msg := []byte("To: paul@oldcode.org\r\n" +
		"Subject: 123 Gophers\r\n" +
		"\r\n" +
		"123 This is the email body.\r\n")
	err := smtp.SendMail(hostp, auth, "asd123@oldcode.org", to, msg)
	if err != nil {
		fmt.Println("error sending email\n")
	}
}
