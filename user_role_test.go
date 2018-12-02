package hippo

import (
	_ "fmt" // for adhoc printing
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	q "github.com/volatiletech/sqlboiler/queries/qm"
)

var _ = Describe("User Role", func() {

	Test("prevents removing the last admin/guest", nil, func(env *TestEnv) {
		admin := env.Tenant.Users(
			q.Where("role_id = ?", UserAdminRoleID),
		).OneP(env.DB)
		var err error
		_, err = admin.Delete(env.DB)
		Expect(err).To(HaveOccurred())

		env.Tenant.AddUsersP(env.DB, true, &User{
			Name: "Bob",
			Email: "bob@test.com",
			RoleID: UserAdminRoleID,
		})
		_, err = admin.Delete(env.DB)
		Expect(err).ToNot(HaveOccurred())

	})

})
