// Package cron
// @Description 定时任务
// @Author 小游
// @Date 2021/05/26
package cron

import (
	"xBlog/internal/db"
	"xBlog/tools"
)

// https://cron.qqe2.com/

// UpdateAccessToken 定时更新微信小程序access_token数据
func UpdateAccessToken() {
	//定时任务
	spec := "0 0 0/1 * * ?" //cron表达式，每小时执行一次
	// 任务执行的内容
	task := func() {
		// 定时更新微信token
		_, _ = tools.WechatGetAccessToken()
	}
	// 启动任务
	startNewTask(spec, task)
}

// EveryDayTask 每天定时执行一次
func EveryDayTask() {
	//定时任务
	spec := "0 0 0 1/1 * ?" //cron表达式，每天执行一次
	// 任务执行的内容
	task := func() {
		// 定时更新xml文件
		go tools.XmlCreateSiteMap()
		// 同步数据库
		if db.GetSiteOptionBool(db.KeyDatabaseBackup) {
			go db.DatabaseBackup()
		}
		// 同步豆瓣数据
		//if database.GetSiteOptionBool("time_task_sync_dou_ban"){
		//	database.SetSiteOption("dou_ban_last_update", common.Time2String(time.Now(), true))
		//	go tools.DouBanUpdateDouBan(database.GetSiteOptionString("dou_ban_cookie"), database.GetSiteOptionString("dou_ban_user"))
		//}
		//// 同步网易云歌单
		//if database.GetSiteOptionBool("time_task_sync_music163"){
		//	_, _ = tools.ToolsV3SyncMusic163()
		//}
	}
	// 启动任务
	startNewTask(spec, task)
}
