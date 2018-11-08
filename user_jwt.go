package hippo

import (
//	"fmt"
	"gopkg.in/urfave/cli.v1"
	"github.com/dgrijalva/jwt-go"
)

func (u *User) JWT(config *cli.Context) string {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": u.Name,
		"admin": u.IsAdmin(),
		"graphql_claims": jwt.MapClaims{
			"x-hasura-default-role": u.RoleName(),
			"x-hasura-allowed-roles": u.AllowedRoleNames(),
			"x-hasura-user-id": u.ID,
			"x-hasura-org-id": u.Tenant.ID,
		},
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(config.String("session_secret")))
	if err != nil {
		panic(err)
	}
	return tokenString
}
