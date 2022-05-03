// Package middleware
// @Description 中间件配置
// @Author 小游
// @Date 2021/04/16
package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// MiddleInit 初始化中间件
func MiddleInit(g *gin.Engine) {
	// 允许跨域
	g.Use(CORS)
	// 认证中间件
	g.Use(RequestAuth)
}

// CORS 设置跨域
// https://segmentfault.com/a/1190000022781975
func CORS(c *gin.Context) {
	method := c.Request.Method
	//  设置头部信息
	// 可将将* 替换为指定的域名
	c.Header("Access-Control-Allow-Origin", "*")
	//服务器支持的所有跨域请求的方法
	c.Header("Access-Control-Allow-Methods", "*")
	//允许跨域设置可以返回其他子段，可以自定义字段
	c.Header("Access-Control-Allow-Headers", "*")
	// 允许浏览器（客户端）可以解析的头部 （重要）
	c.Header("Access-Control-Expose-Headers", "user_id,token,Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
	//允许客户端传递校验信息比如 cookie (重要)
	c.Header("Access-Control-Allow-Credentials", "true")
	// 允许options请求
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	// 执行下一步操作（gin使用的就是类似于责任链的模式）
	c.Next()
}
