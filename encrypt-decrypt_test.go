package hippo

import (
	_ "fmt" // for adhoc printing
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Encryption-Decryption", func() {

	Test("can encrypt/decrypt", &TestFlags{}, func(env *TestEnv) {
		domain := "test.com"
		encrypted, _ := Encrypt(map[string]interface{}{"d": domain})
		decrypted, _ := DecryptStringProperty(encrypted, "d")
		Expect(domain).To(Equal(decrypted))
	})

	Test("can encrypt/decrypt users", &TestFlags{}, func(env *TestEnv) {
		user := User{
			Name: "My Name", Email:"test@test.com",
		}
		encrypted, _ := JWTforUser(&user)
		decrypted, _ := UserforJWT(encrypted)
		Expect(user.ID).To(Equal(decrypted.ID))
		Expect(user.Name).To(Equal(decrypted.Name))
		Expect(user.Email).To(Equal(decrypted.Email))
	})
})
