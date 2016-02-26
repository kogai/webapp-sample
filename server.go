package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

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
		c.JSON(200, gin.H{
			"message": "ok",
			"status":  200,
		})
	})

	router.Run(":8080")
}
