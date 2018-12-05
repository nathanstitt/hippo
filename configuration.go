package hippo

import (
	"strings"
	"github.com/gin-gonic/gin"
	"gopkg.in/urfave/cli.v1"
	"github.com/urfave/cli/altsrc"
)

type Configuration interface {
	String(string) string
	Bool(name string) bool
	Int(name string) int
}

var IsDevMode = false
var SessionsKeyValue = []byte("32-byte-long-auth-key-123-45-712")

func Initialize() *cli.App {
	IsDevMode = 0 != strings.Compare(gin.Mode(), "release")

	app := cli.NewApp()
	app.Flags = []cli.Flag {
		altsrc.NewIntFlag(cli.IntFlag{
			Name:   "port",
			Value:  8080,
			Usage:  "port to listen to",
		}),
		altsrc.NewIntFlag(cli.IntFlag{
			Name:   "graphql_port",
			Value:  8081,
			Usage:  "port to listen to",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:   "product_name",
			Value:  "Hippo Fun Time!",
			Usage:  "The name of the product",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:   "product_email",
			Usage:  "The email address to use for transactional email for the product",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:   "logo_url",
			Value:  "",
			Usage:  "url for logo",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:   "web_domain",
			Value:  "localhost",
			Usage:  "The domain to use for URLs",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:   "bind_address",
			Value:  "localhost",
			Usage:  "ip address to bind to",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:   "session_cookie_name",
			Value:  "hippo",
			Usage:  "name of session cookie",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:   "db_conn_url",
			Value:  "",
			Usage:  "PG database connection string",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:   "session_secret",
			Value:  "",
			Usage:  "32 char long string to use for encrypting session",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:   "graphql_access_key",
			Value:  "mysecretkey",
			Usage:  "32 char long string to use for encrypting session",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:   "email_server",
			Value:  "localhost",
			Usage:  "address of email server",
		}),
		cli.StringFlag{
			Name:   "config",
			Value: "./config.yaml",
			Usage:  "read configuration from file",
		},
	}

	app.Before = altsrc.InitInputSourceWithContext(
		app.Flags,
		altsrc.NewYamlSourceFromFlagFunc("config"),
	)

	return app;
}
