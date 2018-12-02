package hippo

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/nathanstitt/hippo/models"
)

func JWTForUser(u *hm.User, config Configuration) string {
	// Create a new token object, specifying signing method
	// and the claims you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": u.Name,
		"admin": UserIsAdmin(u),
		"graphql_claims": jwt.MapClaims{
			"x-hasura-default-role": UserRoleName(u),
			"x-hasura-allowed-roles": UserAllowedRoleNames(u),
			"x-hasura-user-id": u.ID,
			"x-hasura-org-id": u.TenantID,
		},
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(config.String("session_secret")))
	if err != nil {
		panic(err)
	}
	return tokenString
}
