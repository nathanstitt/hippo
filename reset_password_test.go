package hippo

import (
	"fmt"
	"regexp"
	"testing"
	"net/http"
	"net/url"
	"strings"
	"net/http/httptest"
	. "github.com/smartystreets/goconvey/convey"
)

func prepareResetRequest(db DB) *SignupData {
	data := SignupData{
		Name: "Bob",
		Email: "test-invite-123@test.com",
		Password: "password1234",
		Tenant: "Acme Inc",
	}
	CreateTenant(&data, db)
	return &data
}

func TestResetPassword(t *testing.T) {

	Test("sends reset link", &TestFlags{WithRoutes: addLoginRoute}, func(env *TestEnv) {

		r := env.Router
		db := env.DB

		// initiate reset
		data := prepareResetRequest(db)
		form := url.Values{}
		form.Add("email", data.Email)

		req, _ := http.NewRequest( "POST", "/reset-password", strings.NewReader(form.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)


		So(resp.Body.String(), ShouldContainSubstring,
			fmt.Sprintf("emailed a login link to you at %s", data.Email))
		So(testEmail.body, ShouldContainSubstring,
			fmt.Sprintf("reset the password for email address %s", data.Email))


		// follow link from email
		re := regexp.MustCompile("(/forgot-password\\?t=.*?)\"")
		match := re.FindStringSubmatch(testEmail.body)
		req, _ = http.NewRequest("GET", match[1], nil)
		resp = httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		So(resp.Body.String(), ShouldContainSubstring, "Set New Password")

		form = url.Values{}
		form.Add("password", "test-123-reset")
		req, _ = http.NewRequest("POST", "/reset-password", strings.NewReader(form.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Cookie", resp.Header().Get("Set-Cookie"))

		resp = httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		So(resp.Header().Get("Location"), ShouldEqual, "/")

		user := FindUserByEmail(data.Email, db)
		So(user.ValidatePassword("test-123-reset"), ShouldEqual, true)
	})


}
