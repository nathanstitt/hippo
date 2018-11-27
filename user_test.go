package hippo

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/dgrijalva/jwt-go"
)

var _ = Describe("User", func() {

	Test("creates a JWT token", &TestFlags{WithRoutes: addLoginRoute}, func(env *TestEnv) {

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
		Expect(err).To(BeNil())
		claims, ok := token.Claims.(jwt.MapClaims);
		Expect(ok).To(BeTrue())
		Expect(token.Valid).To(BeTrue())
		Expect(claims["name"]).To(Equal(data.Name))
	})

	Test("can get/set user role", &TestFlags{WithRoutes: addLoginRoute}, func(env *TestEnv) {

		db := env.DB

		data := &SignupData{
			Name: "Nathan",
			Email: "foo@test.com",
			Password: "password",
			Tenant: "Acme",
		}

		tenant, _ := CreateTenant(data, db)
		Expect(tenant.Users).To(HaveLen(2))
		user := &tenant.Users[0]
		Expect(user.Tenant.ID).To(Equal(tenant.ID))

		Expect(user.IsAdmin()).To(BeTrue())
		Expect(user.AllowedRoleNames()).Should(ConsistOf(
			[]string{"admin", "manager", "user", "guest"},
		))
	})

});
