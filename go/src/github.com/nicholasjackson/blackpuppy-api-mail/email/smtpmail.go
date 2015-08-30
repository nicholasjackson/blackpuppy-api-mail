package email

import (
	"log"
	"net/smtp"
	"net/mail"
	"github.com/jpoehls/gophermail"
)

type SmtpMail struct {
	BaseMail
}

func (m *SmtpMail) SendMail(addr string, a smtp.Auth, from string, to []string, subject string, msg []byte) error {
	log.Println("Sending Email:", subject)
	log.Println("Email Server:", addr)

	toAddress := make([]mail.Address, 1)
	toAddress[0].Address = to[0]

	message := gophermail.Message{}
	message.To = toAddress
	message.SetFrom(from)
	message.Subject = subject
	message.HTMLBody = string(msg)

	return gophermail.SendMail(addr, a, &message)
}
