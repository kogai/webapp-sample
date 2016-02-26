package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		log.Println("Hello, world")
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "ハロー、ワールド",
		})
	})

	router.Run(":8080")
}
