// Package api @Description
// @Author 小游
// @Date 2021/04/10
package api

import (
	"github.com/gin-gonic/gin"
	"xBlog/pkg/database"
)

// Version API版本号
const Version = "/api/v3/"
const AdminVersion = "/api/v3/admin/"

func Router(g *gin.Engine) {
	// 获取系统路径
	path := database.GetAppPath()
	//fmt.Println("系统路径",path)
	// 注册静态文件
	g.Static("/assets", path+"assets")
	g.Static("/upload", path+"upload")
	// 注册各个板块路由
	PostRouter(g)
	UserRouter(g)
	SettingRouter(g)
	ToolsRouter(g)
	PluginRouter(g)
	ChatRouter(g)
	// 管理员路由
	AdminVisualRouter(g)
	AdminPostsRouter(g)
	AdminUsersRouter(g)
	AdminCommentsRouter(g)
	AdminSettingsRouter(g)
	AdminPluginsRouter(g)
}
