package main

import (
	_ "fmt" // for adhoc printing
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestEncryptDecrypt(t *testing.T) {

	Convey("it can encrypt/decrypt", t, func() {
		domain := "test.com"
		encrypted, err := Encrypt(map[string]interface{}{"d": domain})
		if err != nil {
			t.Fatal(err)
		}
		decrypted, err := DecryptStringProperty(encrypted, "d")
		if err != nil {
			t.Fatal(err)
		}
		So(domain, ShouldEqual, decrypted)
	})

	Convey("it can encrypt/decrypt users", t, func() {
		user := User{

			Name: "My Name", Email:"test@test.com",
		}
		encrypted, err := JWTforUser(&user)
		if err != nil {
			t.Fatal(err)
		}

		decrypted, err := UserforJWT(encrypted)
		if err != nil {
			t.Fatal(err)
		}
		So(user.ID, ShouldEqual, decrypted.ID)
		So(user.Name, ShouldEqual, decrypted.Name)
		So(user.Email, ShouldEqual, decrypted.Email)
	})
}
