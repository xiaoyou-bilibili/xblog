// Package plugin @Description 解析插件的核心部分
// @Author 小游
// @Date 2021/04/08
package plugin

import (
	"github.com/dop251/goja"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"io/ioutil"
	"xBlog/internal/app"
	"xBlog/internal/app/model"
	"xBlog/internal/app/service/admin"
	"xBlog/internal/plugin/module"
	"xBlog/pkg/database"
)

// 当前注册的对象
var current []*goja.Runtime

// Global 全局对象一些信息
type Global struct {
	PluginName string // 插件的名字
}

// 加载插件
func loadPlugin() {
	// 加载的时候先把当前插件列表置空
	current = []*goja.Runtime{}
	// 清空当前所有插件
	admin.GlobalAllPlugins = []model.GlobalInfo{}
	// 清空所有插件设置
	admin.GlobalAllPluginsSetting = map[string][]model.AdminPluginsSetting{}
	// 清空全局map对象，因为gin路由重复注册会报错，所有这里只清空函数
	for k := range module.GlobalRouter {
		module.GlobalRouter[k] = module.Router{}
	}
	// 关闭所有的定时任务
	for _, c := range module.GlobalCronQueue {
		if c != nil {
			c.Stop()
		}
	}
	module.GlobalCronQueue = []*cron.Cron{}
	// 扫描所有的插件
	plugins := ScanPath()
	// 把所有的插件都加载进去
	for _, v := range plugins {
		// 读取一下js文件和package.json文件
		js, err1 := ioutil.ReadFile(v + "/index.js")
		json, err2 := ioutil.ReadFile(v + "/plugins.json")
		if err1 == nil && err2 == nil {
			current = append(current, ParseJs(string(js), string(json), v))
		}
	}
}

// ScanAllPlugins 扫描所有的插件并加载
func ScanAllPlugins() {
	// 注册widget调试
	module.InitWidget()
	// 注册web界面调试
	module.InitPage()
	// 先加载插件
	loadPlugin()
	// 获取路径
	path := database.GetAppPath()
	// 配置静态文件映射
	app.GinEngine.Static("/plugins/static", path+"plugins")
	// 配置回调函数
	admin.Reload = loadPlugin
	// 注册一个重新加载插件的路由
	app.GinEngine.GET("/restart", func(c *gin.Context) {
		// 重新加载插件
		loadPlugin()
		c.PureJSON(200, "重启完毕")
	})
}
