package hippo

import (
//	"fmt"
	"html/template"
	"github.com/gin-gonic/gin"
	"github.com/nathanstitt/webpacking"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func InitWebpack(router *gin.Engine, config Configuration) *webpacking.WebPacking {
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

func InitSessions(cookie_name string, r *gin.Engine, config Configuration) {
	secret := []byte(config.String("session_secret"))
	SessionsKeyValue = secret
	store := cookie.NewStore(secret)
	r.Use(sessions.Sessions(cookie_name, store));
}

func CreateRouter() *gin.Engine {
	return gin.New()
}

func AddGraphqlProxyRoutes(r *gin.Engine, config Configuration) {
	graphql_port := config.Int("graphql_port")
	r.LoadHTMLGlob("views/*")
	r.OPTIONS("/v1/*query", allowCorsReply)
	r.OPTIONS("/v1alpha1/*graphql", allowCorsReply)
	r.POST("/v1/*query", reverseProxy(graphql_port))
	r.POST("/v1alpha1/*graphql", reverseProxy(graphql_port))
	r.POST("/apis/migrate", reverseProxy(graphql_port))
	r.GET("/v1alpha1/*graphql", reverseProxy(graphql_port))
}
