package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	store := sessions.NewCookieStore([]byte("my-cookie-secret"))
	router.Use(sessions.Sessions("mysession", store))

	// authorized := router.Group("/mypage")
	/*
		store := sessions.NewCookieStore([]byte("my-cookie-secret"))
		router.Use(sessions.Sessions("mysession", store))

		router.GET("/incr", func(c *gin.Context) {
			session := sessions.Default(c)
			var count int
			v := session.Get("count")
			if v == nil {
				count = 0
			} else {
				count = v.(int)
				count += 1
			}
			session.Set("count", count)
			session.Save()
			c.JSON(200, gin.H{
				"count":   count,
				"session": session,
			})
		})
	*/
	// type HandlerFunc func(*Context)

	authorized := router.Group("/mypage", func(c *gin.Context) {
		session := sessions.Default(c)
		log.Println(session)
	})

	authorized.GET("/secrets", func(c *gin.Context) {
		// user := c.MustGet(gin.AuthUserKey).(string)
		// log.Println(user)

		c.JSON(http.StatusOK, gin.H{"user": "user"})
	})

	/*
		authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
			"foo":    "bar",
			"austin": "1234",
			"lena":   "hello2",
			"manu":   "4321",
		}))
		authorized.GET("/secrets", func(c *gin.Context) {
			// get user, it was setted by the BasicAuth middleware
			user := c.MustGet(gin.AuthUserKey).(string)
			if secret, ok := secrets[user]; ok {
				c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
			} else {
				c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
			}
		})
	*/

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
