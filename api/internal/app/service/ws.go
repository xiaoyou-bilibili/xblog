// Package server
// @Description websocket 接口
// @Author 小游
// @Date 2021/04/14
package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"xBlog/internal/app/common"
	"xBlog/tools"
)

var (
	grader = websocket.Upgrader{}
)

// WsChat 聊天接口请求处理
func WsChat(c *gin.Context) {
	// 用户用户id和token数据
	id := tools.Str2Int(c.Query("id"))
	token := c.Query("token")
	// 如果id不为0那么我们就验证一下用户的身份
	if id != 0 {
		if !common.AccessTokenStatusV2(id, token, false) {
			tools.GlobalResponse.ResponseUnauthorized(c, "权限不足请登录")
			return
		}
	}
	//允许websocket跨域
	grader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	// 使用grader把我们的get请求升级为websocket请求
	ws, err := grader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	// 把当前的ws放到请求队列里面去
	common.Queue = append(common.Queue, common.WsQueue{Ws: ws, UserId: id})
	// 结束时关闭websocket请求
	defer func(ws *websocket.Conn) {
		_ = ws.Close()
	}(ws)
	for {
		// 使用for循环来阅读数据
		_, msg, err := ws.ReadMessage()
		// 读取信息错误的时候我们关闭websocket链接
		if err != nil {
			_ = ws.Close()
			return
		}
		//新建一个结构体并赋值
		structs := new(common.WsStruct)
		structs.Ws = ws
		structs.Content = string(msg)
		//把结构体放入通道，等待处理
		common.WsChain <- *structs
	}
}
