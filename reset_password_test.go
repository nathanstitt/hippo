package hippo

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"net/http/httptest"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/nathanstitt/webpacking"
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

func addPasswordResetRoute(
	r *gin.Engine,
	config Configuration,
	webpack *webpacking.WebPacking,
) {
	r.GET("/reset-password", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})
	r.POST("/reset-password", UserPasswordResetHandler())
	r.GET("/forgot-password", UserDisplayPasswordResetHandler)
}

var _ = Describe("Resetting Password", func() {

	Test("sends reset link", &TestFlags{WithRoutes: addPasswordResetRoute}, func(env *TestEnv) {

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

		Expect(resp.Body.String()).To(ContainSubstring(
			fmt.Sprintf("emailed a login link to you at %s", data.Email)))
		Expect(LastEmailDelivery).ToNot(BeNil())
		Expect(LastEmailDelivery.Contents).To(ContainSubstring(
			fmt.Sprintf("reset the password for email address %s", data.Email)))

		// follow link from email
		Expect(LastEmailDelivery.Contents).To(ContainSubstring("/forgot-password"))
		token, _ := EncryptStringProperty("email", data.Email)
		path := fmt.Sprintf("/forgot-password?t=%s", token)
		req, _ = http.NewRequest("GET", path, nil)
		resp = httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		Expect(resp.Body.String()).To(
			ContainSubstring("Set New Password"),
		)

		form = url.Values{}
		form.Add("password", "test-123-reset")
		req, _ = http.NewRequest("POST", "/reset-password", strings.NewReader(form.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Cookie", resp.Header().Get("Set-Cookie"))
		resp = httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		Expect(resp.Header().Get("Location")).To(Equal("/"))
		user := FindUserByEmail(data.Email, db)
		Expect(IsValidPassword(user, "test-123-reset")).To(BeTrue())
	})


})
