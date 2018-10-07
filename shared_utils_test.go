package main

import (
//	"fmt"
	"flag"
	"testing"
	"io/ioutil"
	"github.com/jinzhu/gorm"
	"gopkg.in/urfave/cli.v1"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type FakeEmailSender struct {
	to string
	subject string
	body string
}

func (f *FakeEmailSender) SendEmail(config *cli.Context, to string, subject string, body string) error {
	f.to, f.subject, f.body = to, subject, body
	return nil
}

var testEmail *FakeEmailSender

func testingContextMiddleware(config *cli.Context, tx *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("dbTx", tx)
		c.Set("config", config)
		c.Next()
	}
}

type TestEnv struct {
	Router *gin.Engine
	DB *gorm.DB
	Config *cli.Context
}

func xit(description string, t *testing.T, testFunc func(*TestEnv)) {}

func it(description string, t *testing.T, testFunc func(*TestEnv)) {
	testEmail = &FakeEmailSender{}
	EmailSender = testEmail;
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	set := flag.NewFlagSet("test", 0)
	set.String("session_secret", "32-byte-long-auth-key-123-45-712", "doc")
	set.String(
		"db_conn_url",
		"postgres://localhost/spendily_dev?sslmode=disable",
		"doc",
	)
	config := cli.NewContext(nil, set, nil)
	db := ConnectDB(config)
	tx := db.Begin()
	router := gin.New()
	router.Use(testingContextMiddleware(config, tx))
	initSessions(router, config)
	addRoutes(router, config)
	initWebpack(router)

	Convey(description, t, func() {
		defer func() {
			tx.Rollback()
		}()

		testFunc(&TestEnv{
			Router: router,
			DB: tx,
			Config: config,
		});
	})
}
