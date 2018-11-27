package hippo

import (
	"io"

//	"fmt"
	"flag"
	"bytes"
	"net/http"
	"io/ioutil"
	"html/template"
	"net/http/httptest"
	"gopkg.in/urfave/cli.v1"
	"github.com/onsi/ginkgo"
	"github.com/gin-gonic/gin"
	"github.com/nathanstitt/webpacking"
)

type FakeEmailSender struct {
	to string
	subject string
	body string
}

func (f *FakeEmailSender) SendEmail(config Configuration, to string, subject string, body string) error {
	f.to, f.subject, f.body = to, subject, body
	return nil
}

var testEmail *FakeEmailSender

func testingContextMiddleware(config Configuration, tx DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("dbTx", tx)
		c.Set("config", config)
		c.Next()
	}
}

type TestEnv struct {
	Router *gin.Engine
	DB DB
	Config Configuration
	Webpack *webpacking.WebPacking
	Tenant *Tenant
}

type RequestOptions struct {
	Body *string
	SessionCookie string
	User *User
}

func (env *TestEnv) MakeRequest(
	method string,
	path string,
	options *RequestOptions,
) *httptest.ResponseRecorder {
	var body io.Reader
	if options != nil {
		if options.Body != nil {
			body = bytes.NewReader([]byte(*options.Body))
		}
	}
	req, _ := http.NewRequest(method, path, body)
	if options != nil {
		if options.User != nil {
			req.Header.Set("Cookie",
				TestingCookieForUser(
					options.User, env.Config,
				),
			)
		}
	}
	resp := httptest.NewRecorder()
	env.Router.ServeHTTP(resp, req)
	return resp
}

type TestFlags struct {
	WithRoutes func(
		*gin.Engine,
		Configuration,
		*webpacking.WebPacking,
	)
}

type TestSetupEnv struct {
	SessionSecret string
	DBConnectionUrl string
}

func TestingCookieForUser(u *User, config Configuration) string {
	r := gin.Default()
	InitSessions("test", r, config)
	r.GET("/", func(c *gin.Context) {
		LoginUser(u, c)
		c.String(200, "")
	})
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(res, req)
	return res.Header().Get("Set-Cookie")
}

var TestingEnvironment = &TestSetupEnv{
	SessionSecret: "32-byte-long-auth-key-123-45-712",
	DBConnectionUrl: "postgres://nas@localhost",
}


func RunSpec(flags *TestFlags, testFunc func(*TestEnv)) {
	testEmail = &FakeEmailSender{}
	EmailSender = testEmail;
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	set := flag.NewFlagSet("test", 0)
	set.String(
		"session_secret", TestingEnvironment.SessionSecret, "doc",
	)
	set.String(
		"db_conn_url", TestingEnvironment.DBConnectionUrl, "doc",
	)

	var config Configuration
	config = cli.NewContext(nil, set, nil)
	config.String("foo")
	db := ConnectDB(config)
	tx := db.Begin()

	var router *gin.Engine
	var webpack *webpacking.WebPacking

	tenant, _ := CreateTenant(
		&SignupData{
			Name: "Tester Testing",
			Email: "test@test.com",
			Password: "password",
			Tenant: "testing",
		}, tx,
	)

	if flags.WithRoutes != nil {
		router = gin.New()


		router.Use(testingContextMiddleware(config, tx))
		InitSessions("test", router, config)
		IsDevMode = true
		fake := webpacking.InstallFakeAssetReader()
		defer fake.Restore()
		router.SetFuncMap(template.FuncMap{
			"asset": func(asset string) (template.HTML, error) {
				return template.HTML(asset), nil
			},
		})
		router.LoadHTMLGlob("templates/*")

		flags.WithRoutes(router, config, webpack)
	}

	//	ginkgo.It(description, func() {
	defer func() {
		tx.Rollback()
	}()

	testFunc(&TestEnv{
		Router: router,
		DB: tx,
		Config: config,
		Webpack: webpack,
		Tenant: tenant,
	});
}


func Test(description string, flags *TestFlags, testFunc func(*TestEnv)) {
	ginkgo.It(description, func() {
		RunSpec(flags, testFunc)
	})
}

func XTest(description string, flags *TestFlags, testFunc func(*TestEnv)) {
	ginkgo.XIt(description, func() {})

}

func FTest(description string, flags *TestFlags, testFunc func(*TestEnv)) {
	ginkgo.FIt(description, func() {
		RunSpec(flags, testFunc)
	})
}
