package hippo

import (
//	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/dgrijalva/jwt-go"
)

var _ = Describe("User", func() {

	Test("creates a JWT token", &TestFlags{WithRoutes: addLoginRoute}, func(env *TestEnv) {
		user := env.Tenant.R.Users[0]
		tokenString := user.JWT(env.Config)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(env.Config.String("session_secret")), nil
		})
		Expect(err).To(BeNil())
		claims, ok := token.Claims.(jwt.MapClaims);
		Expect(ok).To(BeTrue())
		Expect(token.Valid).To(BeTrue())
		Expect(claims["name"]).To(Equal(user.Name))
	})

	Test("can get/set user role", &TestFlags{WithRoutes: addLoginRoute}, func(env *TestEnv) {
		tenant := env.Tenant
		Expect(tenant.R.Users).To(HaveLen(2))
		user := tenant.R.Users[0]
		Expect(user.TenantID).To(Equal(tenant.ID))

		Expect(user.IsAdmin()).To(BeTrue())
		Expect(user.AllowedRoleNames()).Should(ConsistOf(
			[]string{"admin", "manager", "user", "guest"},
		))
	})

});
