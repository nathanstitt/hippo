package main

import (
	"fmt"
	"strings"
	"gopkg.in/urfave/cli.v1"
	"github.com/dgrijalva/jwt-go"
)

func (u *User) JWT(config *cli.Context) string {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": u.Name,
		"admin": u.Roles().Admin,
		"graphql_claims": jwt.MapClaims{
			"x-hasura-allowed-roles": strings.Join(u.RoleNames, ","),
			"x-hasura-default-role": u.RoleNames[0],
			"x-hasura-user-id": fmt.Sprintf("%d", u.ID),
			"x-hasura-org-id": fmt.Sprintf("%d", u.Tenant.ID),
		},
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(config.String("session_secret")))
	if err != nil {
		panic(err)
	}
	return tokenString
}
