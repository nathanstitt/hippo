package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"github.com/matcornic/hermes"
)

func passwordResetEmail(email string, token string, config *cli.Context) (string, error) {
	productName := config.String("product_name")
	domain := config.String("domain")
	mail := makeEmailMessage(config)
	body := hermes.Body{
		Name: email,
		Intros: []string{
			fmt.Sprintf("You have received this email because someone requested to reset the password for email address %s at %s", email, productName),
		},
		Actions: []hermes.Action{
			{
				Instructions: fmt.Sprintf("Click the button below to reset your password for %s", productName),
				Button: hermes.Button{
					Color: "#DC4D2F",
					Text:  "Reset Password",
					Link:  fmt.Sprintf("https://%s/forgot-password?t=%s",
						domain, token),
				},
			},
		},
		Outros: []string{
			fmt.Sprintf("If you did not request a password reset for %s, please ignore this email. No further action is required on your part.", productName),
		},
		Signature: "Thanks!",
	}

	return mail.GenerateHTML(hermes.Email{ Body: body })
}
