package email

import (
	"fmt"
	"net/smtp"

	"oldcode.org/gow/lg"
	"oldcode.org/gow/settings"
)

func Send_test() {
	lg.Log.Printf("Send_test() start...")
	port := settings.GetInt("email.port")
	host := settings.GetString("email.host")
	user := settings.GetString("email.user")
	password := settings.GetString("email.password")
	//sender := settings.GetString("email.default_sender")

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
