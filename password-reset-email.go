package hippo

import (
	"fmt"
	"github.com/matcornic/hermes/v2"
	"github.com/nathanstitt/hippo/models"
)

func passwordResetEmail(user *hm.User, token string, db DB, config Configuration) *hermes.Body {
	// email string
	productName := config.String("product_name")
	domain := config.String("domain")

	return &hermes.Body{
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
}
