package email

import (
	"log"
	"net/smtp"
)

type SmtpMail struct {
	BaseMail
}

func (m *SmtpMail) SendMail(addr string, a smtp.Auth, from string, to []string, subject string, msg []byte) error {
	//body := m.BuildMessage(subject, msg)
	log.Println("Sending Email:", subject)
	log.Println("Email Server:", addr)
	subject = "Subject: " + subject + "\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	messagebody := []byte(subject + mime + string(msg))
	return smtp.SendMail(addr, a, from, to, messagebody)
}
