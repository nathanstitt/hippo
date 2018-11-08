package hippo

import (
//	"fmt"
//	"regexp"
	"testing"
//	"strings"
	//	"net/http/httptest"
	"github.com/dgrijalva/jwt-go"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUser(t *testing.T) {
	it("creates a JWT token", t, func(env *TestEnv) {
		data := &SignupData{
			Name: "Nathan",
			Email: "foo@test.com",
			Password: "password",
			Tenant: "Acme",
		}

		tenant, _ := CreateTenant(data, env.DB)
		user := &tenant.Users[0]

		tokenString := user.JWT(env.Config)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(env.Config.String("session_secret")), nil
		})
		So(err, ShouldBeEmpty)
		claims, ok := token.Claims.(jwt.MapClaims);
		So(ok, ShouldEqual, true)
		So(token.Valid, ShouldEqual, true)
		So(claims["name"], ShouldEqual, data.Name)
	})

	it("can get/set user role", t, func(env *TestEnv) {
		data := &SignupData{
			Name: "Nathan",
			Email: "foo@test.com",
			Password: "password",
			Tenant: "Acme",
		}

		tenant, _ := CreateTenant(data, env.DB)
		So(tenant.Users, ShouldHaveLength, 1)
		user := &tenant.Users[0]
		So(user.Tenant.ID, ShouldEqual, tenant.ID)
		So(user.RoleNames, ShouldHaveLength, 1)
		So(user.RoleNames, ShouldContain, "admin")
		roles := user.Roles()
		So(roles.Admin, ShouldEqual, true)
		So(roles.Manager, ShouldEqual, false)
		roles.Manager = true
		roles.Admin = false
		env.DB.Save(user)
		So(tenant.Users, ShouldHaveLength, 1)
		So(user.RoleNames, ShouldContain, "manager")
	})

}
