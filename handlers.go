package main

import (
	"log"
	"fmt"
	"net/http"
	"html/template"
	"encoding/json"
	"net/http/httputil"
	"github.com/jinzhu/gorm"
	"gopkg.in/urfave/cli.v1"
	"github.com/gin-gonic/gin"
)

func getConfig(c *gin.Context) *cli.Context {
	config, ok := c.MustGet("config").(*cli.Context)
	if ok {
		return config
	}
	panic("config isn't the correct type")
}

func getDB(c *gin.Context) *gorm.DB {
	tx, ok := c.MustGet("dbTx").(*gorm.DB)
	if ok {
		return tx
	}
	panic("config isn't the correct type")
}

func contextMiddleware(config *cli.Context, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tx := db.Begin()
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
			fmt.Printf("COMMIT\n")
			tx.Commit();
		}
	}
}

func renderErrorPage(message string, c *gin.Context, err *error) {
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

func renderHomepage(signup *SignupData, err *error, c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"signup": signup,
		"error": err,
	})
}

func renderApplication(user *User, c *gin.Context) {
	bootstrapData, err := json.Marshal(ApplicationBootstrapData{
		User: *user,
		JWT: user.JWT(getConfig(c)),
	})
	if err != nil {
		renderErrorPage("Failed to generate page data", c, nil)
		return
	}
	c.HTML(http.StatusOK, "application.html", gin.H{
		"bootstrapData": template.HTML(string(bootstrapData)),
	})
}
