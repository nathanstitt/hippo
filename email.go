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
	Message *mail.Message
	Body *hermes.Body
	Tenant *hm.Tenant
	Product hermes.Product
	Configuration Configuration
}

func NewEmailMessage(tenant *hm.Tenant, config Configuration) *Email {
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
	email := &Email{
		Message: mail.NewMessage(),
		Configuration: config,
		Tenant: tenant,
		Product: product,
	}
	email.SetFrom(tenant.Email, tenant.Name)
	return email
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
	d.StartTLSPolicy = mail.NoStartTLS
	if IsDevMode {
		m.WriteTo(os.Stdout)
		return nil
	} else {
		return d.DialAndSend(m)
	}
}

var EmailSender EmailSenderInterface = &LocalhostEmailSender{}

func setAddress(mail *mail.Message, header string, address string, names []string) {
	if len(names) > 0 {
		mail.SetAddressHeader(header, address, names[0])
	} else {
		mail.SetHeader(header, address)
	}
}

func (email *Email) SetSubject(subject string, a ...interface{}) {
	email.Message.SetHeader("Subject", fmt.Sprintf(subject, a...))
}
func (email *Email) SetFrom(address string, name ...string) {
	setAddress(email.Message, "From", address, name)
}

func (email *Email) SetTo(address string, name ...string) {
	setAddress(email.Message, "To", address, name)
}
func (email *Email) SetReplyTo(address string, name ...string) {
	setAddress(email.Message, "ReplyTo", address, name)
}
func (email *Email) FormatAddress(address, name string) string {
	return email.Message.FormatAddress(address, name)
}

func (email *Email) BuildMessage() error {
	h := hermes.Hermes{
		Product: email.Product,
	}
	if email.Body == nil {
		return fmt.Errorf("Unable to send email without body")
	}
	contents := hermes.Email{ Body: *email.Body}
	htmlEmailBody, err := h.GenerateHTML(contents)
	if err != nil {
		return err
	}

	textEmailBody, err := h.GeneratePlainText(contents)
	if err != nil {
		return err
	}
	email.Message.SetBody("text/plain", textEmailBody)
	email.Message.AddAlternative("text/html", htmlEmailBody)
	return nil
}

func (email *Email) Deliver() error {
	err := email.BuildMessage()
	if err != nil {
		return err
	}
	return EmailSender.SendEmail(email.Configuration, email.Message)
}
