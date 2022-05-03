// Package admin
// @Description 管理员设置管理
// @Author 小游
// @Date 2021/04/14
package admin

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"xBlog/internal/app/common"
	"xBlog/internal/app/model"
	"xBlog/internal/db"
	"xBlog/tools"
)

// SettingsUpdate 更新网站设置
func SettingsUpdate(c *gin.Context) {
	// 获取参数
	param := new(model.AdminUpdateOption)
	// 参数验证
	if tools.ValidatorParam(c, param) {
		return
	}
	// 判断类型并插入
	var result bool
	switch param.Type {
	case "string":
		result = db.SetSiteOption(param.Key, param.Value)
		break
	case "int":
		result = db.SetSiteOption(param.Key, tools.Str2Int(param.Value))
	case "bool":
		result = db.SetSiteOption(param.Key, tools.Str2Bool(param.Value))
	}
	if result {
		tools.GlobalResponse.ResponseCreated(c, param)
	} else {
		tools.GlobalResponse.ResponseServerError(c)
	}
}

// SettingsGetSite 获取站点设置
func SettingsGetSite(c *gin.Context) {
	// 设置信息   // title表示设置的文字 type表示设置的类型 key表示在数据库中存储的字段名字 value 表示设置的值
	options := []model.AdminSetting{
		{Title: "前台设置", Type: "divider"},
		{Title: "网站标题", Type: "input", Key: db.KeySiteName},
		{Title: "网站副标题", Type: "input", Key: db.KeyWebText},
		{Title: "网站关键词(多个关键词用,隔开)", Type: "input", Key: db.KeyKeyword},
		{Title: "网站描述", Type: "input", Key: db.KeySiteDescription},
		{Title: "备案号", Type: "input", Key: db.KeySiteBeiAn},
		{Title: "公安备案", Type: "input", Key: db.KeySiteGovBeiAn},
		{Title: "建站时间(格式:03/20/2019 00:00:00)", Type: "input", Key: db.KeyBuildTime},
		{Title: "首页一页显示文章数", Type: "number", Key: db.KeyPostListCount},
		{Title: "文章摘要字数", Type: "number", Key: db.KeyPostDecCount},
		{Title: "开启文章网址跳转显示", Type: "switch", Key: db.KeyUrlRedirectOpen, Value: false},
		{Title: "网址跳转背景", Type: "upload", Key: db.KeyUrlRedirectBackground},
		{Title: "评论审核功能", Type: "switch", Key: db.KeyCommentCheck, Value: false},

		{Title: "日记一页展示个数", Type: "number", Key: db.KeyDiaryListCount},
		{Title: "网站顶部ico", Type: "upload", Key: db.KeySiteIcon},

		{Title: "微信支付二维码", Type: "upload", Key: db.KeyWechat},
		{Title: "支付宝支付二维码", Type: "upload", Key: db.KeyAlipay},

		{Title: "安卓APP设置", Type: "divider"},
		{Title: "APP下载链接", Type: "input", Key: db.KeyAppDownloadDownload},
		{Title: "APP版本号", Type: "input", Key: db.KeyAppDownloadVersion},
		{Title: "APP更新说明", Type: "input", Key: db.KeyAppDownloadDescription},
		{Title: "顶部APP下载二维码", Type: "upload", Key: db.KeyAppTopImg},

		{Title: "登录界面设置", Type: "divider"},
		{Title: "登录界面logo", Type: "upload", Key: "login_logo"},

		{Title: "后台设置", Type: "divider"},
		{Title: "管理员邮箱", Type: "input", Key: db.KeySiteEmail},
		{Title: "日记管理摘要显示字数", Type: "number", Key: db.KeyDiaryDecNum},

		{Title: "图片上传设置(地址只能一个有值，默认选择chevereto)", Type: "divider"},
		{Title: "开启图床上传", Type: "switch", Key: db.KeyImgBedUpload, Value: false},
		{Title: "chevereto图床api地址(例:https://img.xxx.com/api/1/upload/?key=你的秘钥)", Type: "input", Key: db.KeyImgBedAddr},

		{Title: "Lsky Pro设置", Type: "divider"},
		{Title: "Lsky Pro图床地址(后面不要加/)", Type: "input", Key: db.KeyImgLypBedAddr},
		{Title: "Lsky Pro邮箱", Type: "input", Key: db.KeyImgLypBedUser},
		{Title: "Lsky Pro密码", Type: "input", Key: db.KeyImgLypBedPassword},

		{Title: "后端设置", Type: "divider", Key: ""},
		{Title: "网站地址(结尾不要加/)", Type: "input", Key: db.KeySiteApiServer},
		{Title: "数据库设置", Type: "divider", Key: ""},
		{Title: "数据库定时备份", Type: "switch", Key: db.KeyDatabaseBackup},
		{Title: "备份数据库", Type: "row", Value: "admin/settings/database/backup"},
		{Title: "恢复数据库", Type: "row", Value: "admin/settings/database/restore"},
		{Title: "站点地图", Type: "divider"},
		{Title: "更新站点地图", Type: "row", Value: "admin/settings/site/sitemap"},

		//{Title: "web服务地址(结尾不要加/)", Type: "input", Key: "site_web_server"},

		{Title: "邮件服务器设置", Type: "divider"},
		{Title: "smtp用户名", Type: "input", Key: db.KeySmtpUser},
		{Title: "smtp密码", Type: "input", Key: db.KeySmtpPass},
		{Title: "smtp地址", Type: "input", Key: db.KeySmtpServer},
		{Title: "smtp端口", Type: "input", Key: db.KeySmtpPort},
		{Title: "smtp发件人名字(只能为英文)", Type: "input", Key: db.KeySmtpName},
	}
	// 返回设置内容
	common.AdminReturnSetting(c, options)
}

// SettingsGetBackground 获取壁纸设置
func SettingsGetBackground(c *gin.Context) {
	// 设置信息
	options := []model.AdminSetting{
		{Title: "主页壁纸", Type: "upload", Key: db.KeyIndexImg},
		{Title: "顶部导航栏壁纸", Type: "upload", Key: db.KeyHeadImg},
		{Title: "文章页面壁纸", Type: "upload", Key: db.KeyPostImg},
		{Title: "日记界面壁纸", Type: "upload", Key: db.KeyDiaryImg},

		//{Title: "赞助页面壁纸", Type: "upload", Key: db.KeyDonateImg},
		//{Title: "友链页面壁纸", Type: "upload", Key: db.KeyFriendImg},
		//{Title: "追番页面壁纸", Type: "upload", Key: db.KeyAnimationImg},
		//{Title: "弹幕页面壁纸", Type: "upload", Key: db.KeyBarrageImg},
		//{Title: "豆瓣页面壁纸", Type: "upload", Key: db.KeyDouBanImg},
		//{Title: "B博页面壁纸", Type: "upload", Key: db.KeyBiBoImg},
		//{Title: "个人导航页面壁纸", Type: "upload", Key: db.KeyMoeImg},
	}
	// 返回设置内容
	common.AdminReturnSetting(c, options)
}

// SettingsGetSpider 获取壁纸设置
func SettingsGetSpider(c *gin.Context) {
	// 设置信息
	options := []model.AdminSetting{
		{Title: "B站设置", Type: "divider"},
		{Title: "B站uid", Type: "input", Key: db.KeyBiliUID},
		{Title: "B站cookie", Type: "text", Key: db.KeyBiliCookie},

		{Title: "豆瓣设置", Type: "divider"},
		{Title: "豆瓣用户id", Type: "input", Key: db.KeyDouBanUser},
		{Title: "豆瓣cookie", Type: "text", Key: db.KeyDouBanCookie},

		{Title: "网易云设置", Type: "divider"},
		{Title: "网易云用户id", Type: "input", Key: db.KeyMusicU},
		{Title: "网易云歌单id", Type: "input", Key: db.KeyMusicId},

		{Title: "自动同步设置", Type: "divider"},
		{Title: "每日定时同步豆瓣数据", Type: "switch", Key: db.KeyTimeTaskSyncDouBan},
		{Title: "每日定时同步网易云歌单", Type: "switch", Key: db.KeyTimeTaskSyncMusic163},
	}
	// 返回设置内容
	common.AdminReturnSetting(c, options)
}

// SettingsSpiderMusic163 同步网易云歌单
func SettingsSpiderMusic163(c *gin.Context) {
	if re, err := common.ToolsSyncMusic163(); err == nil {
		tools.GlobalResponse.ResponseCreated(c, re)
	} else {
		tools.GlobalResponse.ResponseServerError(c, err.Error())
	}
}

// SettingsSpiderDouBan 同步豆瓣数据
// todo 使用JavaScript脚本重写
func SettingsSpiderDouBan(c *gin.Context) {
	// 更新同步时间
	db.SetSiteOption(db.KeyDouBanLastUpdate, tools.Time2String(time.Now(), true))
	//go tools.DouBanUpdateDouBan(db.GetSiteOptionString("dou_ban_cookie"), db.GetSiteOptionString("dou_ban_user"))
	tools.GlobalResponse.ResponseCreated(c, bson.M{"tip": "后台正在更新数据"})
}

// SettingsGetSide 获取侧边栏设置
func SettingsGetSide(c *gin.Context) {
	// 初始化一下避免前端报错
	res := db.SettingAdminSide{
		Right:  []db.SettingAdminSideDetail{},
		Left:   []db.SettingAdminSideDetail{},
		Unused: []db.SettingAdminSideDetail{},
	}
	// 左右侧边栏信息
	var left = db.GetSiteOption(db.KeySettingSideInfoLeft)
	var right = db.GetSiteOption(db.KeySettingSideInfoRight)
	// 使用map来保存侧边栏信息
	var tmpMap = map[string]bool{}
	// 获取左右侧边栏信息
	for _, v := range left.(primitive.A) {
		tmp := db.SettingAdminSideDetail{}
		if tools.Primitive2Struct(v, &tmp) == nil && common.SideInfo[tmp.Unique] != nil {
			tmpMap[tmp.Unique] = true
			res.Left = append(res.Left, tmp)
		}
	}
	for _, v := range right.(primitive.A) {
		tmp := db.SettingAdminSideDetail{}
		if tools.Primitive2Struct(v, &tmp) == nil && common.SideInfo[tmp.Unique] != nil {
			tmpMap[tmp.Unique] = true
			res.Right = append(res.Right, tmp)
		}
	}
	// 获取所有的侧边栏
	for _, v := range GlobalAllPlugins {
		if common.SideInfo[v.Unique] != nil {
			// 左右侧边栏都没有才添加
			if !tmpMap[v.Unique] {
				res.Unused = append(res.Unused, db.SettingAdminSideDetail{
					Name:   v.Config.Name,
					Unique: v.Unique,
				})
			}
		}
	}
	tools.GlobalResponse.ResponseOk(c, res)
}

// SettingsUpdateSide 更新侧边栏设置
func SettingsUpdateSide(c *gin.Context) {
	// 获取参数
	var param model.AdminPutSideParam
	if tools.ValidatorParam(c, &param) {
		return
	}
	// 保存到数据库
	db.SetSiteOption(db.KeySettingSideInfoLeft, param.Left)
	db.SetSiteOption(db.KeySettingSideInfoRight, param.Right)
	tools.GlobalResponse.ResponseOk(c, param)
}

// SettingsGetMiniProgram 获取微信小程序设置
func SettingsGetMiniProgram(c *gin.Context) {
	// 设置信息
	options := []model.AdminSetting{
		{Title: "小程序appID", Type: "input", Key: db.KeyWechatMiniProgramId},
		{Title: "小程序秘钥", Type: "input", Key: db.KeyWechatMiniProgramSecret},
		{Title: "赞赏码", Type: "upload", Key: db.KeyWechatMiniProgramPrice},
		{Title: "小程序头部图片", Type: "upload", Key: db.KeyWechatMiniProgramHeadImage},
		{Title: "小程序关于博主内容", Type: "input", Key: db.KeyWechatMiniProgramAbout},

		{Title: "功能界面设置", Type: "divider"},
		{Title: "显示我的追番", Type: "switch", Key: db.KeyWechatMiniProgramAnimation},
		{Title: "显示友链", Type: "switch", Key: db.KeyWechatMiniProgramFriend},
		{Title: "显示赞助", Type: "switch", Key: db.KeyWechatMiniProgramDonate},
		{Title: "显示豆瓣", Type: "switch", Key: db.KeyWechatMiniProgramDouBan},
		{Title: "显示评论", Type: "switch", Key: db.KeyWechatMiniProgramComment},
		{Title: "显示文章底部工具栏", Type: "switch", Key: db.KeyWechatMiniProgramPostBar},
	}
	// 返回设置内容
	common.AdminReturnSetting(c, options)
}

// SettingsGetAPP 获取APP全局设置
func SettingsGetAPP(c *gin.Context) {
	// 设置信息
	options := []model.AdminSetting{
		{Title: "开启友链功能", Type: "switch", Key: db.KeySettingAppFriend},
		{Title: "开启追番功能", Type: "switch", Key: db.KeySettingAppAnimation},
		{Title: "开启赞助功能", Type: "switch", Key: db.KeySettingAppDonate},
		{Title: "开启豆瓣功能", Type: "switch", Key: db.KeySettingAppDouBan},
		{Title: "开启音乐盒功能", Type: "switch", Key: db.KeySettingAppMusic},
		{Title: "开启文档功能", Type: "switch", Key: db.KeySettingAppDoc},
		{Title: "开启我的项目功能", Type: "switch", Key: db.KeySettingAppProject},
		{Title: "开启个人导航功能", Type: "switch", Key: db.KeySettingAppNavigation},
		{Title: "登录界面背景", Type: "upload", Key: db.KeySettingAppLogin},
	}
	// 返回设置内容
	common.AdminReturnSetting(c, options)
}

// SettingsGetNav 获取导航栏设置
func SettingsGetNav(c *gin.Context) {
	// 初始化一下避免前端报错
	var res []db.HeadNav
	// 获取导航栏设置
	var nav = db.GetSiteOption(db.KeySettingNavInfo)
	// 获取左右侧边栏信息
	for _, v := range nav.(primitive.A) {
		tmp := db.HeadNav{}
		if tools.Primitive2Struct(v, &tmp) == nil {
			res = append(res, tmp)
		}
	}
	tools.GlobalResponse.ResponseOk(c, res)
}

// SettingsUpdateNav 更新导航栏设置
func SettingsUpdateNav(c *gin.Context) {
	// 获取参数
	var param []db.HeadNav
	if c.Bind(&param) != nil {
		tools.GlobalResponse.ResponseBadRequest(c, "参数格式错误")
		return
	}
	// 保存到数据库
	db.SetSiteOption(db.KeySettingNavInfo, param)
	tools.GlobalResponse.ResponseOk(c, param)
}

// SettingsDatabaseBackup 数据库备份功能
func SettingsDatabaseBackup(c *gin.Context) {
	if db.DatabaseBackup() == nil {
		tools.GlobalResponse.ResponseOk(c, map[string]string{})
	} else {
		tools.GlobalResponse.ResponseServerError(c, "备份失败")
	}
}

// SettingsDatabaseRestore 数据库恢复功能
func SettingsDatabaseRestore(c *gin.Context) {
	if db.DatabaseRestore() == nil {
		tools.GlobalResponse.ResponseOk(c, map[string]string{})
	} else {
		tools.GlobalResponse.ResponseServerError(c, "恢复失败")
	}
}

// SettingsGetSitemap 同步站点地图
func SettingsGetSitemap(c *gin.Context) {
	if tools.XmlCreateSiteMap() {
		tools.GlobalResponse.ResponseCreated(c, gin.H{"message": "同步成功"})
	} else {
		tools.GlobalResponse.ResponseServerError(c)
	}
}

type ThemeSetting struct {
	Title string               `json:"title"`
	Items []model.AdminSetting `json:"items"`
}

// SettingGetTheme 获取主题设置
func SettingGetTheme(c *gin.Context) {
	// 获取主题名字
	name := c.Param("name")
	server := db.GetSiteOptionString(db.KeySiteApiServer)
	// 发送请求获取数据
	res := tools.HttpGet(server + "/php/api/v1/themes/setting?name=" + name)
	// 解析数据
	data := new([]ThemeSetting)
	//fmt.Println(server + "/php/api/v1/themes/setting?name=" + name)
	//fmt.Println(res)
	if json.Unmarshal([]byte(res), data) == nil {
		// 解析所有数据
		for i := range *data {
			for k, v := range (*data)[i].Items {
				if tmp := db.GetSiteOption(v.Key); tmp == nil {
					(*data)[i].Items[k].Value = v.Default
				} else {
					(*data)[i].Items[k].Value = tmp
				}
			}
		}
		tools.GlobalResponse.ResponseOk(c, data)
		return
	}
	tools.GlobalResponse.ResponseServerError(c, "无法获取设置，可能该主题没有设置!")
}
