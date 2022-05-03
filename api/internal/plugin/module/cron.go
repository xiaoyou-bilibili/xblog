// Package module
// @Description 定时任务包
// @Author 小游
// @Date 2021/04/08
package module

import (
	"github.com/dop251/goja"
	"github.com/robfig/cron/v3"
)

// GlobalCronQueue 全局定时任务队列
var GlobalCronQueue []*cron.Cron

// RegisterCron 注册定时任务对象
func RegisterCron(vm *goja.Runtime) *goja.Object {
	blog := vm.NewObject()
	_ = blog.Set("start", CronStart)
	return blog
}

// CronStart 启动一个定时任务
func CronStart(spec string, task func()) {
	// 新建一个定时任务对象
	// 根据cron表达式进行时间调度，cron可以精确到秒，大部分表达式格式也是从秒开始。
	crontab := cron.New(cron.WithSeconds())
	//crontab := cron.New(cron.WithSeconds()) //精确到秒
	// 添加定时任务
	if _, err := crontab.AddFunc(spec, task); err != nil {
		return
	}
	// 启动定时任务
	crontab.Start()
	// 加入全局队列中
	GlobalCronQueue = append(GlobalCronQueue, crontab)
}
