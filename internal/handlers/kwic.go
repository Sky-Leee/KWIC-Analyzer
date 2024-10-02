package handlers

import (
	"net/http"
	"software-architecture-teaching-software/internal/api"
	"software-architecture-teaching-software/internal/logic"
	"software-architecture-teaching-software/pkg/consts"

	"github.com/gin-gonic/gin"
)

// RenderKWICPage 渲染 KWIC 分析页面
func RenderKWICPage(c *gin.Context) {
	c.HTML(http.StatusOK, "kwic.html", nil)
}

// 解析请求
func AnalyzeKWIC(c *gin.Context) {
	var request struct {
		Type    int    `json:"type"`
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var result string
	switch request.Type {
	case consts.TypeMainSub:
		result = logic.ProcessMainSub(request.Content)
	case consts.TypeObjectOriented:
		result = logic.ProcessObjectOriented(request.Content)
	case consts.TypeEventSystem:
		result = logic.ProcessEventSystem(request.Content)
	case consts.TypePipeFilter:
		result = logic.ProcessPipeFilter(request.Content)
	default:
		c.JSON(http.StatusBadRequest, api.KWICResult{
			Msg: "Invalid type",
		})
		return
	}

	c.JSON(http.StatusOK, api.KWICResult{Data: result})
}
