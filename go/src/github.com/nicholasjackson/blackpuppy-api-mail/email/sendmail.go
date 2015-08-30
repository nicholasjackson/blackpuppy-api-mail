package email

import (
	"net/smtp"

	"github.com/hoisie/mustache"
)

type SendMail interface {
	SendMail(addr string, a smtp.Auth, from string, to []string, subject string, msg []byte) error
	BuildMessage(subject string, template string, data interface{}) []byte
}

type BaseMail struct{}

func (m *BaseMail) BuildMessage(subject string, template string, data interface{}) []byte {
	//mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	// load the template
	body := mustache.RenderFile(template, data)

	msg := []byte(body)
	return msg
}
