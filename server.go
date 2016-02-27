package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	store := sessions.NewCookieStore([]byte("my-cookie-secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "account/login", gin.H{
			"title": "ログイン画面",
			"stuff": "ログイン画面",
		})
	})

	router.POST("/login", func(c *gin.Context) {
		var form Login
		// This will infer what binder to use depending on the content-type header.
		name := c.PostForm("User")

		log.Println(c.Bind(&form))
		log.Println(form)
		log.Println(name)

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

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "base", gin.H{
			"title": "title:ハロー、ワールド",
			"stuff": ":ハロー、ワールド",
		})
	})

	router.GET("/account/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "account/regiser", gin.H{
			"title": "title:ハロー、ワールド",
			"stuff": ":ハロー、ワールド",
		})
	})

	router.POST("/api/v1/register", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"status":  http.StatusOK,
		})
	})

	router.Run(":8080")
}
