package emailer

import (
	"errors"

	"github.com/zalora_icecream/commonFramework/external/gopkg.in/gomail.v2"
)

const (
	paramMissing string = "Either SMTPServer or Sender or Recipients is missing"
)

//Email models an Email
type Email struct {
	SMTPServer    string
	Port          int
	Sender        string
	Recipients    []string
	CcRecipients  []string
	BccRecipients []string
	Subject       string
	Body          string
	ContentType   string
}

//Send sends an email
func (email Email) Send() error {

	if len(email.SMTPServer) == 0 ||
		len(email.Sender) == 0 ||
		len(email.Recipients) == 0 {
		return errors.New(paramMissing)
	}

	if len(email.ContentType) == 0 {
		email.ContentType = "text/plain"
	}

	d := gomail.Dialer{Host: email.SMTPServer, Port: email.Port}

	m := gomail.NewMessage()
	m.SetHeader("From", email.Sender)

	m.SetHeader("To", email.Recipients...)

	if len(email.CcRecipients) > 0 {
		m.SetHeader("Cc", email.CcRecipients...)
	}

	if len(email.BccRecipients) > 0 {
		m.SetHeader("Bcc", email.BccRecipients...)
	}

	m.SetHeader("Subject", email.Subject)
	m.SetBody(email.ContentType, email.Body)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
