package hippo

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/matcornic/hermes"
)

var _ = Describe("Sending Email", func() {

	Test("can send", &TestFlags{}, func(env *TestEnv) {
		email := NewEmailMessage(env.Tenant, env.Config)
		email.To = "test@test.com"
		email.Body = &hermes.Body{
			Name: "test@test.com",
			Signature: "GO AWAY!",
		}
		err := email.deliver()
		Expect(err).To(BeNil())
		Expect(LastEmailDelivery.To).To(Equal("test@test.com"))
		Expect(LastEmailDelivery.Contents).To(
			ContainSubstring("GO AWAY!"),
		)
	})
})
