package mail

import (
	gomail "gopkg.in/gomail.v2"
)

type EmailPayload struct {
	Subject string
	Message string
	From    string
	To      []string
}

// SendMail send an email to recipients
func SendMail(
	e *EmailPayload,
	authUser,
	authPassword,
	host string,
	port int,
) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", e.From)
	msg.SetHeader("To", e.To...)
	msg.SetHeader("Subject", e.Subject)
	msg.SetBody("text/html", e.Message)

	n := gomail.NewDialer(host, port, authUser, authPassword)
	if err := n.DialAndSend(msg); err != nil {
		return err
	}
	return nil
}
