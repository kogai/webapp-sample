package controllers

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login", gin.H{
		"title": "ログイン画面",
		"stuff": "ログイン画面",
	})
}

func postLogin(c *gin.Context) {
	var form account
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
}

var loginHandlers handlers = handlers{
	"GET":  getLogin,
	"POST": postLogin,
}
