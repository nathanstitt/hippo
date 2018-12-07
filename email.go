package hippo

import (
	"os"
	"fmt"
	"time"
	"github.com/go-mail/mail"
	"github.com/matcornic/hermes"
	"github.com/nathanstitt/hippo/models"
)


type Email struct {
	To string
	From string
	ReplyTo string
	Subject string
	Body *hermes.Body
	Tenant *hm.Tenant
	Product hermes.Product
	Configuration Configuration
}

func MakeEmailMessage(tenant *hm.Tenant, config Configuration) *Email {
	product := hermes.Product{
		// Appears in header & footer of e-mails
		Name: tenant.Name,
		Link: tenant.HomepageURL.String,
		Logo: tenant.LogoURL.String,
		Copyright: fmt.Sprintf(
			"Copyright © %d %s. All rights reserved.",
			time.Now().Year(),
			config.String("product_name"),
		),
	}
	return &Email{
		From: config.String("product_email"),
		Configuration: config,
		Tenant: tenant,
		Product: product,
	}
}

func decodeInviteToken(token string) (string, error) {
	return DecryptStringProperty(token, "email")
}



type EmailSenderInterface interface {
	SendEmail(Configuration, *mail.Message) error
}

// Mail sender
type LocalhostEmailSender struct {}

func (s *LocalhostEmailSender) SendEmail(config Configuration, m *mail.Message) error {
	host := config.String("email_server")
	d := mail.Dialer{Host: host, Port: 25}
	if IsDevMode {
		m.WriteTo(os.Stdout)
		return nil
	} else {
		return d.DialAndSend(m)
	}
}

var EmailSender EmailSenderInterface = &LocalhostEmailSender{}

func (email *Email) BuildMessage() (*mail.Message, error) {
	h := hermes.Hermes{
		Product: email.Product,
	}
	if email.Body == nil {
		return nil, fmt.Errorf("Unable to send email without body")
	}
	contents := hermes.Email{ Body: *email.Body}
	htmlEmailBody, err := h.GenerateHTML(contents)
	if err != nil {
		return nil, err
	}

	textEmailBody, err := h.GeneratePlainText(contents)
	if err != nil {
		return nil, err
	}
	m := mail.NewMessage()
	if email.ReplyTo != "" {
		m.SetHeader("ReplyTo", email.From)
	}
	m.SetHeader("From", email.From)
	m.SetHeader("To", email.To)
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/plain", textEmailBody)
	m.AddAlternative("text/html", htmlEmailBody)
	return m, nil
}

func (email *Email) deliver() error {
	m, err := email.BuildMessage()
	if err != nil {
		return err
	}
	return EmailSender.SendEmail(email.Configuration, m)
}


func deliverResetEmail(user *hm.User, token string, db DB, config Configuration) error {
	email := MakeEmailMessage(user.Tenant().OneP(db), config)
	email.Body = passwordResetEmail(user, token, db, config)
	email.To = user.Email
	email.Subject = fmt.Sprintf("Password Reset for %s", config.String("product_name"))
	return email.deliver()
}

func deliverLoginEmail(emailAddress string, tenant *hm.Tenant, config Configuration) error {
	email := MakeEmailMessage(tenant, config)
	email.Body = signupEmail(emailAddress, tenant, config)
	email.To = emailAddress
	email.Subject = fmt.Sprintf("Login to %s", config.String("product_name"))
	return email.deliver()
}
