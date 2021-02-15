package internal

import (
	"github.com/gin-gonic/gin"
)

// 秘匿されたデータをシミュレートする
var accounts = gin.H{
	"admin": gin.H{"email": "abc@example.com", "phone": "090-1234-5678"},
}

func Handler() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("web/*.html")
	r.Static("/web", "./web")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})

	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))

	// localhost:8080/admin/login でアクセス
	authorized.GET("/login", func(c *gin.Context) {
		// BasicAuth ミドルウェアで設定されたユーザー名にアクセス
		user := c.MustGet(gin.AuthUserKey).(string)
		if account, ok := accounts[user]; ok {
			c.JSON(200, gin.H{"user": user, "secret": account})
		} else {
			c.JSON(200, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	r.GET("/booklist", func(ctx *gin.Context) {
		ctx.HTML(200, "booklist.html", gin.H{})
	})

	return r
}
