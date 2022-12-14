package utils

import (
	"github.com/davidsolisdev/portafolioAPI/config"

	mail "github.com/xhit/go-simple-mail/v2"
)

type NewEmail struct {
	From    string
	To      string
	Subject string
}

func SendEmail(email *NewEmail, bodyHtml string) (bool, error) {
	// conecting mail server
	smtp, err := config.SmtpClient()
	if err != nil {
		return false, err
	}
	defer smtp.Close()

	// create new mail
	var eMail *mail.Email = mail.NewMSG()
	eMail.SetFrom(email.From)
	eMail.AddTo(email.To)
	eMail.SetSubject(email.Subject)
	eMail.SetBody(mail.TextHTML, bodyHtml)

	// comprobate mail
	if eMail.Error != nil {
		return false, eMail.Error
	}

	// send mail
	err = eMail.Send(smtp)
	if err != nil {
		return false, err
	}

	return true, err
}
