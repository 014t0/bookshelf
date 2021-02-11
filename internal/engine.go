package internal

import "github.com/gin-gonic/gin"

func Handler() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("web/*.html")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})

	r.GET("/booklist", func(ctx *gin.Context) {
		ctx.HTML(200, "booklist.html", gin.H{})
	})

	return r
}
