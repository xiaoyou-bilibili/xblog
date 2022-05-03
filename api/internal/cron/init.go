// Package cron
// @Description 定时任务初始化函数
// @Author 小游
// @Date 2021/05/26
package cron

import "github.com/robfig/cron/v3"

func startNewTask(spec string, task func()) {
	// 新建一个定时任务对象
	// 根据cron表达式进行时间调度，cron可以精确到秒，大部分表达式格式也是从秒开始。
	crontab := cron.New(cron.WithSeconds()) //默认从分开始进行时间调度
	//crontab := cron.New(cron.WithSeconds()) //精确到秒
	// 添加定时任务
	if _, err := crontab.AddFunc(spec, task); err != nil {
		return
	}
	crontab.Start()
}

// Init 定时任务模块初始化
func Init() {
	UpdateAccessToken()
	EveryDayTask()
}
