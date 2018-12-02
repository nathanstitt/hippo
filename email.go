package hippo

import (
	"os"
	"fmt"
	"time"
	"github.com/go-mail/mail"
	"github.com/matcornic/hermes"
	"github.com/nathanstitt/hippo/models"
)


func MakeEmailMessage(tenant *hm.Tenant, config Configuration) hermes.Hermes {
	return hermes.Hermes{
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: tenant.Name,
			Link: tenant.HomepageURL.String,
			Logo: tenant.LogoURL.String,
			Copyright: fmt.Sprintf(
				"Copyright © %d %s. All rights reserved.",
				time.Now().Year(),
				config.String("product_name"),
			),
		},
	}
}

func decodeInviteToken(token string) (string, error) {
	return DecryptStringProperty(token, "email")
}



type EmailSenderInterface interface {
	SendEmail(config Configuration, to string, subject string, mailBody string) error
}

// Mail sender
type LocalhostEmailSender struct {}

func (s *LocalhostEmailSender) SendEmail(
	config Configuration, to string, subject string, mailBody string,
) error {
	m := mail.NewMessage()
	m.SetHeader("From", "contact@thescrumgame.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", mailBody)

	d := mail.Dialer{Host: "localhost", Port: 25}

	if IsDevMode {
		m.WriteTo(os.Stdout)
		return nil
	}
	return d.DialAndSend(m)
}

var EmailSender EmailSenderInterface = &LocalhostEmailSender{}

func deliverResetEmail(user *hm.User, token string, db DB, config Configuration) error {
	mailBody, err := passwordResetEmail(user, token, db, config)
	if (err != nil) {
		return err;
	}
	return EmailSender.SendEmail(
		config,
		user.Email,
		fmt.Sprintf("Password Reset for %s", config.String("product_name")),
		mailBody,
	)
}

func deliverLoginEmail(email string, tenant *hm.Tenant, config Configuration) error {
	mailBody, err := signupEmail(email, tenant, config)
	if (err != nil) {
		return err;
	}
	return EmailSender.SendEmail(
		config,
		email,
		fmt.Sprintf("Login to %s", config.String("product_name")),
		mailBody,
	)
}
