package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getAccountRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register", gin.H{
		"title": "アカウント登録",
		"stuff": "アカウント登録ページです",
	})
}

var accountRegisterHandlers handlers = handlers{
	"GET": getAccountRegister,
}
