package hippo

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/matcornic/hermes"
)

var _ = Describe("Sending Email", func() {

	Test("can send", &TestFlags{}, func(env *TestEnv) {
		email := NewEmailMessage(env.Tenant, env.Config)
		email.SetTo("test@test.com")
		email.SetSubject("Hello %s, you have %d things", "bob", 33)
		email.SetFrom("foo-test@test.com")
		email.Body = &hermes.Body{
			Name: "test@test.com",
			Signature: "GO AWAY!",
		}
		err := email.Deliver()
		Expect(err).To(BeNil())
		Expect(LastEmailDelivery.Subject).To(Equal("Hello bob, you have 33 things"))
		Expect(LastEmailDelivery.To).To(Equal("test@test.com"))
		Expect(LastEmailDelivery.Contents).To(
			ContainSubstring("GO AWAY!"),
		)
	})
})
