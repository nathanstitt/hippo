package hippo

import (
	"fmt"
//	"regexp"
	"testing"
	"net/http"
	"net/url"
	"strings"
	"net/http/httptest"
	"github.com/jinzhu/gorm"
	. "github.com/smartystreets/goconvey/convey"
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

func TestLoginHandler(t *testing.T) {

	it("can log in", t, func(env *TestEnv) {
		r := env.Router
		db := env.DB

		form := prepareLoginRequest(db);
		req, _ := http.NewRequest( "POST", "/login", strings.NewReader(form.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		So(resp.Header().Get("Set-Cookie"), ShouldNotBeEmpty)
		So(resp.Header().Get("Location"), ShouldEqual, "/")
	})

	it("it rejects invalid logins", t, func(env *TestEnv) {
		r := env.Router
		db := env.DB
		form := prepareLoginRequest(db);
		form.Set("password", "foo")
		req, _ := http.NewRequest( "POST", "/login", strings.NewReader(form.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		So(resp.Header().Get("Location"), ShouldBeEmpty)
		So(resp.Body.String(), ShouldContainSubstring, "tab login active")
		So(resp.Body.String(), ShouldContainSubstring, "email or password is incorrect")
	})

}
