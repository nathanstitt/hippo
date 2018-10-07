package main

import (
	_ "fmt" // for adhoc printing
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Encryption-Decryption", func() {

	It("can encrypt/decrypt", func() {
		domain := "test.com"
		encrypted, _ := Encrypt(map[string]interface{}{"d": domain})
		decrypted, _ := DecryptStringProperty(encrypted, "d")
		Expect(domain).To(Equal(decrypted))
	})

	It("can encrypt/decrypt users", func() {
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
