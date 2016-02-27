package controllers

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type account struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func getLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login", gin.H{
		"title": "ログイン画面",
		"stuff": "ログイン画面",
	})
}

func postLogin(c *gin.Context) {
	// Authentication

	var form account
	err := c.Bind(&form)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	// 仮にuser=manu password=123のアカウントのみ認可している
	if form.User != "manu" && form.Password != "123" {
		c.JSON(http.StatusOK, gin.H{
			"status":  401,
			"message": "un authorized",
		})
		return
	}

	session := sessions.Default(c)
	session.Set("sessionId", form.User)
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "ok",
	})
}

var loginHandlers handlers = handlers{
	"GET":  getLogin,
	"POST": postLogin,
}
