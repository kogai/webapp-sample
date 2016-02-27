package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register", gin.H{
		"title": "アカウント登録",
		"stuff": "アカウント登録ページです",
	})
}

func postRegister(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"status":  http.StatusOK,
	})
}

var registerHandlers handlers = handlers{
	"GET":  getRegister,
	"POST": postRegister,
}
