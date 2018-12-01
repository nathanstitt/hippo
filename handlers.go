package hippo

import (
	"log"
	"fmt"
	"net/http"
	"database/sql"
	"html/template"
	"encoding/json"
	"net/http/httputil"
	"github.com/gin-gonic/gin"
)

func GetConfig(c *gin.Context) Configuration {
	config, ok := c.MustGet("config").(Configuration)
	if ok {
		return config
	}
	panic("config isn't the correct type")
}

func GetDB(c *gin.Context) DB {
	tx, ok := c.MustGet("dbTx").(DB)
	if ok {
		return tx
	}
	panic("config isn't the correct type")
}

func RoutingMiddleware(config Configuration, db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tx, err := db.Begin()
		if err != nil {
			panic(err)
		}
		c.Set("dbTx", tx)
		c.Set("config", config)
		defer func() {
			status := c.Writer.Status()
			if status >= 400 {
				log.Printf("Transaction is being rolled back; status = %d\n", status)
				tx.Rollback();
			}
			return
		}()
		c.Next()
		if (c.Writer.Status() < 400) {
			tx.Commit();
		}
	}
}

func RenderErrorPage(message string, c *gin.Context, err *error) {
	if err != nil {
		log.Printf("Error occured: %s", *err)
	}
	c.HTML(http.StatusInternalServerError, "error.html", gin.H{
		"message": message,
	})
}

func allowCorsReply(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Max-Age", "86400")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max, X-HASURA-ACCESS-KEY")
	c.Header("Access-Control-Allow-Credentials", "true")
}

func reverseProxy(port int) gin.HandlerFunc {
	target := fmt.Sprintf("localhost:%d", port)
	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = target
	}
	return func(c *gin.Context) {
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func RenderHomepage(signup *SignupData, err *error, c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"signup": signup,
		"error": err,
	})
}

func RenderApplication(user *User, c *gin.Context) {
	cfg := GetConfig(c)
	c.HTML(http.StatusOK, "application.html", gin.H{
		"webDomain": cfg.String("web_domain"),
		"bootstrapData": BootstrapData(user, cfg),
	})
}

func BootstrapData(user *User, cfg Configuration) template.JS {
	type BootstrapDataT map[string]interface{}
	bootstrapData, err := json.Marshal(
		BootstrapDataT{
			"user": user,
			"graphql" : BootstrapDataT{
				"token": user.JWT(cfg),
				"endpoint": cfg.String("web_domain"),
			},
		})
	if err != nil {
		panic(err)
	}
	return template.JS(string(bootstrapData))
}
