package main

import (
	"gopkg.in/urfave/cli.v1"
	"github.com/urfave/cli/altsrc"
)

func setupConfiguration() *cli.App {
	app := cli.NewApp()
	app.Name = "spendily"
	app.Version = "0.0.1"
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
			Name:   "Spendily",
			Value:  "product_name",
			Usage:  "The name of the product",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:   "",
			Value:  "logo_url",
			Usage:  "url for logo",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:   "spendily.com",
			Value:  "domain",
			Usage:  "The domain to use for URLs",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:   "bind_address",
			Value:  "localhost",
			Usage:  "ip address to bind to",
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
