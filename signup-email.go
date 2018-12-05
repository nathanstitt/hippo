package hippo

import (
	"fmt"
	"github.com/matcornic/hermes"
	"github.com/nathanstitt/hippo/models"
)

func signupEmail(email string, tenant *hm.Tenant, config Configuration) *hermes.Body {
	productName := config.String("product_name")
	domain := config.String("domain")
	return &hermes.Body{
		Name: email,
		Intros: []string{
			fmt.Sprintf("You have received this email because %s was used to sign up for TheScrumGame.com", email),
		},
		Actions: []hermes.Action{
			{
				Instructions: fmt.Sprintf("Click the button below to access %s",
					productName),
				Button: hermes.Button{
					Color: "#DC4D2F",
					Text:  "Log me in",
					Link:  fmt.Sprintf("https://%s/login", domain),
				},
			},
		},
		Outros: []string{
			fmt.Sprintf("If you did not request an account with %s, please ignore this email. No further action is required on your part.", productName),
		},
		Signature: "Thanks!",
	}
}
