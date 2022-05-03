// Package middleware
// @Description 接口认证模块
// @Author 小游
// @Date 2021/05/27
package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"xBlog/internal/app/common"
	"xBlog/tools"
)

// RequestAuth 管理员接口权限认证
func RequestAuth(c *gin.Context) {
	//获取请求的各项参数
	url := c.Request.URL.RequestURI()
	//这里说明需要拦截数据
	if strings.Index(url, "/api/v3/admin") != -1 {
		// 鉴权
		if _, err := common.AccessAdminTokenV2(c); err != nil {
			tools.GlobalResponse.ResponseUnauthorized(c)
			c.Abort()
		} else {
			c.Next()
		}
	}
}
