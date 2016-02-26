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

	router.Run(":8080")
}
