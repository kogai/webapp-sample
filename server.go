package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kogai/webapp-sample/controllers"
)

func main() {
	router := gin.Default()

	// LoadHTMLGlobはテンプレートファイルのリストを上書きする
	// LoadHTMLGlobs関数を作ってもいいかも
	// https://github.com/gin-gonic/gin/issues/241
	router.LoadHTMLGlob("templates/**/*")
	router.StaticFile("/app.js", "./public/app.js")

	store := sessions.NewCookieStore([]byte("my-cookie-secret"))
	router.Use(sessions.Sessions("my-session", store))

	controllers.SetRouters(router)

	router.Run(":8080")
}
