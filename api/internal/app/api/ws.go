// Package api
// @Description 聊天室相关的接口
// @Author 小游
// @Date 2021/04/10
package api

import (
	"github.com/gin-gonic/gin"
	server "xBlog/internal/app/service"
)

func ChatRouter(e *gin.Engine) {
	const base = "/ws/v1/chat"
	// 聊天相关的接口
	e.GET(base, server.WsChat)
}
