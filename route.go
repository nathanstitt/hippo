package main

import (
	"log"
	"fmt"
	"net/http"
	"html/template"
	"gopkg.in/urfave/cli.v1"
	"github.com/gin-gonic/gin"
	"github.com/go-webpack/webpack"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func initWebpack(router *gin.Engine) {
	webpack.DevHost = "localhost:3808"
	webpack.Plugin = "manifest"
	webpack.Init(isDevMode)
}

func initSessions(r *gin.Engine, config *cli.Context) {
	secret := []byte(config.String("session_secret"))
	SessionsKeyValue = secret
	store := cookie.NewStore(secret)
	r.Use(sessions.Sessions("spendily", store))
}


func addRoutes(r *gin.Engine, config *cli.Context) { // c *cli.Context, db *gorm.DB) *gin.Engine {
	r.SetFuncMap(template.FuncMap{
		"asset": webpack.AssetHelper,
		"isSet": func(a interface{}) bool {
			if a == nil || a == "" || a == 0 {
				fmt.Println("is not set")
				return false
			}
			fmt.Println("is set")
			return false
		},
	})
	log.SetOutput(gin.DefaultWriter)

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	graphql_port := config.Int("graphql_port")

	r.LoadHTMLGlob("templates/*")
	r.Static("/public", "./public")

	r.OPTIONS("/v1/*query", allowCorsReply)
	r.OPTIONS("/v1alpha1/*graphql", allowCorsReply)
	r.POST("/v1/*query", reverseProxy(graphql_port))
	r.POST("/v1alpha1/*graphql", reverseProxy(graphql_port))
	r.GET("/v1alpha1/*graphql", reverseProxy(graphql_port))

	r.POST("/login", UserLoginHandler)
	r.POST("/logout", UserLogoutHandler)
	r.POST("/signup", TenantSignupHandler)
	r.POST("/reset-password", UserPasswordResetHandler)
	r.GET("/forgot-password", UserDisplayPasswordResetHandler)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {
		u := UserFromSession(c)
		if (u == nil) {
			c.HTML(http.StatusOK, "home.html", gin.H{})
		} else {
			renderApplication(u, c)
		}
	})
}
