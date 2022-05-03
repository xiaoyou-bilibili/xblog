// Package module
// @Description 静态资源映射
// @Author 小游
// @Date 2021/04/08
package module

import (
	"github.com/dop251/goja"
	"xBlog/internal/app"
)

// 静态路径也搞一个全局管理对象
var globalStatic = make(map[string]bool)

// RegisterStatic 注册静态路径
func RegisterStatic(vm *goja.Runtime) *goja.Object {
	blog := vm.NewObject()
	_ = blog.Set("staticPath", StaticRegister)
	_ = blog.Set("staticFile", StaticFileRegister)
	return blog
}

// StaticFileRegister 注册静态文件
func StaticFileRegister(filename string) {
	url := "/static/" + Global.Unique + "/" + filename
	filename = Global.Path + "/" + filename
	// 先判断路径是否存在
	if !globalStatic[filename] {
		app.GinEngine.StaticFile(url, filename)
		globalStatic[filename] = true
	}
}

// StaticRegister 注册静态路径
func StaticRegister(name string, path string) {
	path = Global.Path + "/" + path
	url := "/static/" + Global.Unique + "/" + name
	// 先判断路径是否存在
	if !globalStatic[path] {
		app.GinEngine.Static(url, path)
		globalStatic[path] = true
	}
}
