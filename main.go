package main

import (
	"os"
	"strings"
	"gopkg.in/urfave/cli.v1"
	"github.com/gin-gonic/gin"
)

var SessionsKeyValue = []byte("32-byte-long-auth-key-123-45-712")
var isDevMode = false

func startServer(c *cli.Context) error {
	isDevMode = 0 != strings.Compare(gin.Mode(), "release")

	db := ConnectDB(c)
	router := gin.New()
	router.Use(contextMiddleware(c, db))
	initSessions(router, c)
	addRoutes(router, c)
	initWebpack(router)
	startGraphql(c)

	router.Run(":1281") // listen and serve on 0.0.0.0:8080
	return nil
}

func main() {
	app := setupConfiguration();
	app.Action = startServer
	err := app.Run(os.Args)
	checkError(err)
}
