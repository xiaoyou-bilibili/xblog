// Package module
// @Description 和邮件有关的工具类
// @Author 小游
// @Date 2021/04/08
package module

import (
	"github.com/dop251/goja"
	"xBlog/tools"
)

// RegisterMail 注册邮件发送服务
func RegisterMail(vm *goja.Runtime) *goja.Object {
	blog := vm.NewObject()
	// 发送邮件
	_ = blog.Set("sendMail", tools.SendMail)
	return blog
}
