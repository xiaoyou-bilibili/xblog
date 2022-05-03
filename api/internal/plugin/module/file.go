// Package module
// @Description 文件操作相关的函数
// @Author 小游
// @Date 2021/04/08
package module

import (
	"github.com/dop251/goja"
	"io/ioutil"
)

// RegisterFile 注册文件读取对象
func RegisterFile(vm *goja.Runtime) *goja.Object {
	blog := vm.NewObject()
	_ = blog.Set("read", FileReadFile)
	return blog
}

// FileReadFile 读取文件
func FileReadFile(filename string, call func(err interface{}, data interface{})) {
	if data, err := ioutil.ReadFile(Global.Path + "/" + filename); err != nil {
		call(err.Error(), nil)
	} else {
		call(nil, string(data))
	}
}
