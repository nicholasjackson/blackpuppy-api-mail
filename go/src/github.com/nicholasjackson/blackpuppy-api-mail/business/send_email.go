package business

import (
	"net/smtp"
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

	if global.Config.SmtpServerSettings.UseAuth {
		auth := smtp.PlainAuth(
			"",
			global.Config.SmtpServerSettings.User,
			global.Config.SmtpServerSettings.Password,
			global.Config.SmtpServerSettings.Server)

		return mail.SendMail(
			global.Config.SmtpServerSettings.Server+":" +global.Config.SmtpServerSettings.Port,
			auth,
			"admin@demo.gs",
			emails,
			global.Config.ContactUsSubject,
			message)
	} else {
		return mail.SendMail(
			global.Config.SmtpServerSettings.Server+":" +global.Config.SmtpServerSettings.Port,
			 nil,
			"admin@demo.gs",
			emails,
			global.Config.ContactUsSubject,
			message)
	}

}
