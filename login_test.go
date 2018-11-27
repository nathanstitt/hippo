package hippo

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"net/http/httptest"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"github.com/nathanstitt/webpacking"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func prepareLoginRequest(db *gorm.DB) url.Values {
	data := SignupData{
		Name: "Bob",
		Email: "test123@test.com",
		Password: "password1234",
		Tenant: "Acme Inc",
	}

	_, err := CreateTenant(&data, db)
	if err != nil {
		panic(fmt.Sprintf("add tenant failed: %s", err))
	}
	form := url.Values{}
	form.Add("email", data.Email)
	form.Add("tenant", "acme-inc")
	form.Add("password", data.Password)
	return form;
}

func addLoginRoute(
	r *gin.Engine,
	config Configuration,
	webpack *webpacking.WebPacking,
) {
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})
	r.POST("/login", UserLoginHandler("/"))
}

var _ = Describe("Login", func() {

	Test("can log in", &TestFlags{WithRoutes: addLoginRoute}, func(env *TestEnv) {
		r := env.Router
		db := env.DB

		form := prepareLoginRequest(db);
		req, _ := http.NewRequest( "POST", "/login", strings.NewReader(form.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		Expect(resp.Header().Get("Set-Cookie")).To(Not(BeEmpty()))
		Expect(resp.Header().Get("Location")).To(Equal("/"))
	})

	Test("it rejects invalid logins", &TestFlags{WithRoutes: addLoginRoute}, func(env *TestEnv) {


		r := env.Router
		db := env.DB
		form := prepareLoginRequest(db);
		form.Set("password", "foo")
		req, _ := http.NewRequest( "POST", "/login", strings.NewReader(form.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		Expect(resp.Header().Get("Location")).To(BeEmpty())
		Expect(resp.Body.String()).To(ContainSubstring("email or password is incorrect"))
	})

});
