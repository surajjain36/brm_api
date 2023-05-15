package mailservice

import (
	"brm_api/utils/apperror"
	"crypto/tls"
	"os"

	"gopkg.in/gomail.v2"
)

func SendMail(toMailId string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("MAIL_FROM_EMAILID"))
	m.SetHeader("To", toMailId)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(os.Getenv("MAIL_HOST"), 587, os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: false,
		ServerName: os.Getenv("MAIL_HOST")}

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return &apperror.AppError{
			Status: 500,
			Msg:    err.Error(),
			Errors: nil,
		}
	}
	return nil
}
