// Package module
// @Description 路由注册相关
// @Author 小游
// @Date 2021/04/06
package module

import (
	"github.com/dop251/goja"
	"github.com/gin-gonic/gin"
	"xBlog/internal/app"
	"xBlog/internal/app/api"
	"xBlog/tools"
)

// Router 这里维护一个全局的映射函数
type Router struct {
	Fun func(c *gin.Context)
}

// GlobalRouter 全局router对象，避免路由重复注册造成报错的bug
var GlobalRouter = make(map[string]Router)

// RegisterRouter 注册所有路由,返回object对象
func RegisterRouter(vm *goja.Runtime) *goja.Object {
	blog := vm.NewObject()
	_ = blog.Set("registerRouter", routerRegisterRouter)
	_ = blog.Set("registerAdminRouter", routerRegisterAdminRouter)
	// 传入一个response对象
	_ = blog.Set("response", &tools.GlobalResponse)
	// 获取post的值
	_ = blog.Set("getPostJson", routerGetPostJson)
	return blog
}

// 路由管理类来管理路由分发(类似于代理模式)
func routerManage(c *gin.Context) {
	// 获取路径
	path := c.FullPath()
	// 获取请求方式
	method := c.Request.Method
	// 拼接一下
	path = method + path
	// 判断当前路径是否存在，存在我们就通过代理来执行对应操作
	if _, ok := GlobalRouter[path]; ok && GlobalRouter[path].Fun != nil {
		GlobalRouter[path].Fun(c)
	} else {
		tools.GlobalResponse.ResponseNotFound(c, "路由不存在")
	}
}

// 注册管理路由
func routerRegisterAdminRouter(method string, url string, function func(*gin.Context)) {
	// 注册路由的时候我们需要自动加上插件的路径信息
	url = api.Version + "admin/plugins/" + Global.Unique + url
	addRouter(method, url, function)
}

// 注册路由
func routerRegisterRouter(method string, url string, function func(*gin.Context)) {
	// 注册路由的时候我们需要自动加上插件的路径信息
	url = api.Version + "plugins/" + Global.Unique + url
	addRouter(method, url, function)
}

func addRouter(method string, url string, function func(*gin.Context)) {
	// 注册路由的回调函数
	fun := func(c *gin.Context) {
		// 异常处理
		defer func() {
			if err := recover(); err != nil {
				tools.GlobalResponse.ResponseServerError(c, "插件执行错误!错误信息:"+err.(*goja.Exception).Error())
			}
		}()
		// 调用js的回调函数
		function(c)
	}
	// 如果路由不存在那么就注册,否则就不管
	if _, ok := GlobalRouter[method+url]; !ok {
		// 判断请求方式
		switch method {
		case "GET":
			app.GinEngine.GET(url, routerManage)
		case "POST":
			app.GinEngine.POST(url, routerManage)
		case "PUT":
			app.GinEngine.PUT(url, routerManage)
		case "DELETE":
			app.GinEngine.DELETE(url, routerManage)
		}
	}
	// 更新处理函数到map上
	GlobalRouter[method+url] = Router{Fun: fun}
}

// 获取json数据
func routerGetPostJson(c *gin.Context) interface{} {
	data := map[string]interface{}{}
	err := c.BindJSON(&data)
	if err != nil {
		return nil
	}
	return data
}
