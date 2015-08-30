package business

import (
	"log"
	"github.com/nicholasjackson/blackpuppy-api-mail/email"
	"github.com/nicholasjackson/blackpuppy-api-mail/global"
)

type ContactUsEmail struct {
	Name  string
	Email string
	Body  string
}

func SendEmail(name string, email string, body string, mail email.SendMail) error {
	var emails []string = make([]string, 1)
	emails[0] = global.Config.ContactUsEmail

	emailData := ContactUsEmail{}
	emailData.Name = name
	emailData.Email = email
	emailData.Body = body

	message := mail.BuildMessage(
		global.Config.ContactUsSubject,
		global.Config.RootFolder+global.Config.ContactUsTemplate,
		emailData)

	log.Println("Body: ", string(message[:]))

	return mail.SendMail(global.Config.SmtpServerSettings.Server, nil, "admin@demo.gs", emails, global.Config.ContactUsSubject, message)
}
