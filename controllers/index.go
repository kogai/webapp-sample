package controllers

import (
	"github.com/gin-gonic/gin"
	// "log"
	"net/http"
)

type handlers map[string]gin.HandlerFunc
type account struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func getRoute(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "トップページ",
		"stuff": "トップページです",
	})
}

var routeHandlers handlers = handlers{
	"GET": getRoute,
}

func SetRouters(router *gin.Engine) {
	router.GET("/", routeHandlers["GET"])

	router.GET("/login", loginHandlers["GET"])
	router.POST("/login", loginHandlers["POST"])

	router.GET("/register", registerHandlers["GET"])
	router.POST("/register", registerHandlers["POST"])

	/*
		authorized := router.Group("/", func(c *gin.Context) {
			session := sessions.Default(c)
			sessionId := session.Get("sessionId")
			if sessionId == nil {
				log.Println("This session not authorized.")
				return
			}
			log.Println(sessionId, " is authorized.")
		})

		authorized.GET("/secrets", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"user": "user"})
		})

	*/
}
