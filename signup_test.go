package hippo

import (
	"fmt"
	"strings"
	"net/url"
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
)

func formData(email string) url.Values {
	data := url.Values{}
	data.Add("name", "Bob")
	data.Add("email", email)
	data.Add("password", "password1234")
	data.Add("tenant", "Acme Inc")
	return data
}

func makeRequest(data url.Values, router *gin.Engine) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("POST", "/signup",
		strings.NewReader(data.Encode()))
	req.PostForm = data
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	return resp
}

func TestSignupHandler(t *testing.T) {

	it("it emails a login link", t, func(env *TestEnv) {
		r := env.Router
		db := env.DB
		email := "test1234@test.com"

		data := formData(email)
		resp := makeRequest(data, r)
		user := FindUserByEmail(email, db)
		So(user.ID, ShouldNotEqual, 0)
		So(resp.Code, ShouldEqual, http.StatusOK)
		if user.ID == 0 {
			fmt.Printf("BODY: %s", resp.Body.String())
		}
		So(resp.Header().Get("Set-Cookie"), ShouldNotBeEmpty)
		So(resp.Header().Get("Location"), ShouldEqual, "/")
		So(resp.Body.String(), ShouldContainSubstring, email)
	})

	it("errors when email is duplicate", t, func(env *TestEnv) {
		email := "nathan1234@stitt.org"
		form := &SignupData{
			Name: "Nathan",
			Email: email,
			Password: "password",
			Tenant: "Acme",
		}
		CreateTenant(form, env.DB)
		data := formData(email)
		resp := makeRequest(data, env.Router)
		So(resp.Code, ShouldEqual, http.StatusOK)
		So(resp.Body.String(), ShouldContainSubstring, "email is in use")
	})
}
