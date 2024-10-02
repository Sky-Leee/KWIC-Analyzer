package main

import (
	"software-architecture-teaching-software/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
    r.LoadHTMLFiles("../static/kwic.html")

	r.GET("/kwic", handlers.RenderKWICPage)
	r.POST("/kwic/analyze", handlers.AnalyzeKWIC)
	r.Run(":8080") // 监听8080端口
}
