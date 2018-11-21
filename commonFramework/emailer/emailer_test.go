package emailer

import (
	"testing"
)

func TestEmailSend(t *testing.T) {
	email := Email{
		SMTPServer: "smtp-relay",
		Port:       25,
		Sender:     "emailer_test@crateandbarrel.com",
		Recipients: []string{
			"sktripathy@crateandbarrel.com",
		},
		CcRecipients: []string{
			"sktripathy@crateandbarrel.com",
		},
		BccRecipients: []string{
			"sktripathy@crateandbarrel.com",
		},
		Subject:     "Testing Go Emailer",
		Body:        "Hello! <br /><br />This is a test email sent from Go email tester. Please ignore and do not reply.<br /><br />Thank You<br />Go Email Tester",
		ContentType: "text/html",
	}

	err := email.Send()

	if err != nil {
		t.Errorf("Error sending test email %v", err)
	}
}

func TestEmailSendMissingRequiredParams(t *testing.T) {
	email := Email{
		SMTPServer:  "",
		Port:        25,
		Sender:      "emailer_test@crateandbarrel.com",
		Subject:     "Testing Go Emailer",
		Body:        "Hello! <br /><br />This is a test email sent from Go email tester. Please ignore and do not reply.<br /><br />Thank You<br />Go Email Tester",
		ContentType: "text/html",
	}

	err := email.Send()

	if err == nil {
		t.Errorf("Error testing email validation")
	} else {
		if err.Error() != paramMissing {
			t.Errorf("Error testing email validation. Expected %s, Received %s", paramMissing, err.Error())
		}
	}
}
