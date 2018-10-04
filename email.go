package main

import (
	"os"
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"github.com/go-mail/mail"
	"github.com/matcornic/hermes"
)


func makeEmailMessage(config *cli.Context) hermes.Hermes {
	return hermes.Hermes{
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: config.String("product_name"),
			Link: fmt.Sprintf("https://%s/", config.String("domain")),
			Logo: fmt.Sprintf("https://%s%s", config.String("domain"),
				config.String("logo_url")),
			Copyright: "Copyright © 2018 Argosity. All rights reserved.",
		},
	}
}

func decodeInviteToken(token string) (string, error) {
	return DecryptStringProperty(token, "email")
}



type EmailSenderInterface interface {
	SendEmail(config *cli.Context, to string, subject string, mailBody string) error
}

// Mail sender
type LocalhostEmailSender struct {}

func (s *LocalhostEmailSender) SendEmail(
	config *cli.Context, to string, subject string, mailBody string,
) error {
	m := mail.NewMessage()
	m.SetHeader("From", "contact@thescrumgame.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", mailBody)

	d := mail.Dialer{Host: "localhost", Port: 25}

	if isDevMode {
		m.WriteTo(os.Stdout)
		return nil
	}
	return d.DialAndSend(m)
}

var EmailSender EmailSenderInterface = &LocalhostEmailSender{}

func deliverResetEmail(email string, token string, config *cli.Context) error {
	mailBody, err := passwordResetEmail(email, token, config)
	if (err != nil) {
		return err;
	}
	return EmailSender.SendEmail(
		config,
		email,
		fmt.Sprintf("Password Reset for %s", config.String("product_name")),
		mailBody,
	)
}

func deliverLoginEmail(email string, config *cli.Context) error {
	mailBody, err := signupEmail(email, config)
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
