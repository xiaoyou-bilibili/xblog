// Package app @Description
// @Author 小游
// @Date 2021/04/11
package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
	"xBlog/internal/app/api"
	"xBlog/internal/app/common"
	"xBlog/pkg/database"
	"xBlog/tools"
)

var GinEngine *gin.Engine

// Init 初始化
func Init() {
	//设置随机数种子
	rand.Seed(time.Now().Unix())
	//读取网站的配置
	configs := tools.GetConfig("site")
	// 新建一个线程来处理webSocket数据
	go common.ChainProcess()
	//mongo数据库连接池初始化
	database.DbInit()
	// 数据库初始化
	tools.DbInit()
	defer database.DbClose()
	//初始化路由
	api.Router(GinEngine)
	// 启动项目,这里我们使用http2.0版本
	if err := GinEngine.Run(":" + configs["apiPort"]); err != nil {
		fmt.Println(err)
	}
}
