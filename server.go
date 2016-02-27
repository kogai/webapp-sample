package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kogai/dotspace/controllers"
	"log"
	"net/http"
)

type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")

	store := sessions.NewCookieStore([]byte("my-cookie-secret"))
	router.Use(sessions.Sessions("my-session", store))

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login", gin.H{
			"title": "ログイン画面",
			"stuff": "ログイン画面",
		})
	})

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

	controllers.SetRouters(router)

	router.Run(":8080")
}
