package controllers

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type handlers map[string]gin.HandlerFunc

func getRoute(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "トップページ",
		"stuff": "トップページです",
	})
}

func authorization(c *gin.Context) {
	session := sessions.Default(c)
	sessionId := session.Get("sessionId")
	if sessionId == nil {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
	log.Println(sessionId, " is authorized.")
}

var routeHandlers handlers = handlers{
	"GET": getRoute,
}

func SetRouters(router *gin.Engine) {
	// 認証が必要なページ
	authorized := router.Group("/", authorization)
	authorized.GET("/", routeHandlers["GET"])

	// 認証が不要なページ
	router.GET("/login", loginHandlers["GET"])
	router.POST("/login", loginHandlers["POST"])
	router.GET("/register", registerHandlers["GET"])
	router.POST("/register", registerHandlers["POST"])
}
