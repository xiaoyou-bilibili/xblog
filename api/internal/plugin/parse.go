// Package plugin
// @Description 解析JavaScript脚本
// @Author 小游
// @Date 2021/04/04
package plugin

import (
	"encoding/json"
	"fmt"
	"github.com/dop251/goja"
	"xBlog/internal/app/model"
	"xBlog/internal/app/service/admin"
	"xBlog/internal/plugin/module"
)

// ParseJs 解析js文件
func ParseJs(js string, row string, path string) *goja.Runtime {
	config := model.Plugins{}
	vm := goja.New()
	// 尝试解析JSON文件,解析无误后进行脚本相关的配置
	if json.Unmarshal([]byte(row), &config) == nil {
		blog := vm.NewObject()
		module.Global = model.GlobalInfo{
			Unique: config.Unique,
			Config: config,
			Path:   path,
		}
		// 把当前插件放入全局变量中
		admin.GlobalAllPlugins = append(admin.GlobalAllPlugins, module.Global)
		// 先配置全局信息
		_ = blog.Set("global", module.Global)
		// 根据用户的权限来分配函数
		for _, v := range config.Auth {
			switch v {
			case "router":
				_ = blog.Set("router", module.RegisterRouter(vm))
			case "database":
				_ = blog.Set("database", module.RegisterDatabase(vm))
			case "net":
				_ = blog.Set("net", module.RegisterNet(vm))
			case "spider":
				_ = blog.Set("spider", module.RegisterColly(vm))
			case "cron":
				_ = blog.Set("cron", module.RegisterCron(vm))
			case "tools":
				_ = blog.Set("tools", module.RegisterTools(vm))
			case "file":
				_ = blog.Set("file", module.RegisterFile(vm))
			case "static":
				//_ = blog.Set("static", module.RegisterStatic(vm))
			case "mail":
				_ = blog.Set("mail", module.RegisterMail(vm))
			case "widget":
				_ = blog.Set("widget", module.RegisterWidget(vm))
			}
		}
		// 挂载对象
		_ = vm.Set("xBlog", blog)
		// 运行脚本
		_, err := vm.RunString(js)
		if err != nil {
			fmt.Println(err)
		}
	}
	return vm
}
