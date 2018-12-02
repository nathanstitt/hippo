package hippo

import (
	"fmt"
	"github.com/matcornic/hermes"
	"github.com/nathanstitt/hippo/models"
)

func passwordResetEmail(user *hm.User, token string, db DB, config Configuration) (string, error) {
	// email string
	productName := config.String("product_name")
	domain := config.String("domain")
	mail := MakeEmailMessage(user.Tenant().OneP(db), config)
	body := hermes.Body{
		Name: user.Email,
		Intros: []string{
			fmt.Sprintf("You have received this email because someone requested to reset the password for email address %s at %s", user.Email, productName),
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
