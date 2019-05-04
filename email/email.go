package email

import (
	"fmt"
	"net/smtp"

	"oldcode.org/gow/settings"
)

func Send_test() {
	port := settings.GetInt("email.port")
	host := settings.GetString("email.host")
	user := settings.GetString("email.user")
	password := settings.GetString("email.password")

	auth := smtp.PlainAuth("", user, password, host)
	hostp := fmt.Sprintf("%s:%d", host, port)
	//fmt.Printf("hostp:%s\n", hostp)
	to := []string{"paul@oldcode.org"}
	msg := []byte("To: paul@oldcode.org\r\n" +
		"Subject: 123 Gophers\r\n" +
		"\r\n" +
		"123 This is the email body.\r\n")
	//asd123 is the real from and sender
	err := smtp.SendMail(hostp, auth, "asd123@oldcode.org", to, msg)
	if err != nil {
		fmt.Println("error sending email\n")
	}
}
