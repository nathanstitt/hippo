package hippo

import (
	_ "fmt" // for adhoc printing
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	q "github.com/volatiletech/sqlboiler/queries/qm"
)

var _ = Describe("Tenant methods", func() {

	Test("can query", &TestFlags{}, func(env *TestEnv) {
		Expect(
			Tenants(q.Where("id=?", env.Tenant.ID)).ExistsP(env.DB),
		).To(Equal(true))

	})

})
