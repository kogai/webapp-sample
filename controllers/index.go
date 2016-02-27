package controllers

import (
	"github.com/gin-gonic/gin"
	// "log"
	"net/http"
)

type handlers map[string]gin.HandlerFunc

func getRoute(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "タイトルで〜す",
		"stuff": "トップページで〜す",
	})
}

var routeHandlers handlers = handlers{
	"GET": getRoute,
}

func SetRouters(r *gin.Engine) {
	r.GET("/", routeHandlers["GET"])
	r.GET("/register", accountRegisterHandlers["GET"])

	r.POST("/api/v1/register", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"status":  http.StatusOK,
		})
	})

	/*
		r.POST("/api/v1/login", func(c *gin.Context) {
			var form Login
			if c.Bind(&form) == nil {
				if form.User == "manu" && form.Password == "123" {
					session := sessions.Default(c)
					session.Set("sessionId", "my-session-id")
					session.Save()
					c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
				} else {
					c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
				}
			}
		})
	*/
}
