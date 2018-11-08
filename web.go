package hippo

import (
//	"fmt"
	"html/template"
	"gopkg.in/urfave/cli.v1"
	"github.com/gin-gonic/gin"
	"github.com/argosity/webpacking"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func InitWebpack(router *gin.Engine, config *cli.Context) *webpacking.WebPacking {
	wpConfig := &webpacking.Config{
		IsDev: IsDevMode,
	}
	packager, err := webpacking.New(wpConfig)
	CheckError(err)
	err = packager.Run()
	CheckError(err)

	router.SetFuncMap(template.FuncMap{
		"asset": packager.AssetHelper(),
	})
	return packager
}

func InitSessions(cookie_name string, r *gin.Engine, config *cli.Context) {
	secret := []byte(config.String("session_secret"))
	SessionsKeyValue = secret
	store := cookie.NewStore(secret)
	r.Use(sessions.Sessions(cookie_name, store));
}

func CreateRouter() *gin.Engine {
	return gin.New()
}

func AddGraphqlProxyRoutes(r *gin.Engine, config *cli.Context) {
	// Recovery middleware recovers from any panics and writes a 500 if there was one.

	graphql_port := config.Int("graphql_port")

	r.LoadHTMLGlob("templates/*")
//	r.Static("/public", "./public")

	r.OPTIONS("/v1/*query", allowCorsReply)
	r.OPTIONS("/v1alpha1/*graphql", allowCorsReply)
	r.POST("/v1/*query", reverseProxy(graphql_port))
	r.POST("/v1alpha1/*graphql", reverseProxy(graphql_port))
	r.POST("/apis/migrate", reverseProxy(graphql_port))
	r.GET("/v1alpha1/*graphql", reverseProxy(graphql_port))

//	r.POST("/login", UserLoginHandler)
//	r.POST("/logout", UserLogoutHandler)
//	r.POST("/signup", TenantSignupHandler)
//	r.POST("/reset-password", UserPasswordResetHandler)
//	r.GET("/forgot-password", UserDisplayPasswordResetHandler)

	// r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	// })

}
