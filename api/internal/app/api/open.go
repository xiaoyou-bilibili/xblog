// Package api @Description
// @Author 小游
// @Date 2021/04/10
package api

import (
	"github.com/gin-gonic/gin"
	server "xBlog/internal/app/service"
	"xBlog/internal/app/service/admin"
)

// PostRouter 文章相关的路由函数
func PostRouter(e *gin.Engine) {
	const base = Version + "posts"
	//获取文章列表
	e.GET(base, server.PostGetPostList)
	//获取文章的分类
	e.GET(base+"/category", server.PostGetCategory)
	//获取文章的内容
	e.GET(base+"/:id", server.PostGetPostContent)
	// 获取加密文章的内容
	e.GET(base+"/:id/encryption", server.PostGetPostEncryptContent)
	//获取文章的评论
	e.GET(base+"/:id/comments", server.PostGetPostComment)
	//文章评论
	e.POST(base+"/:id/comments", server.PostCommitComment)
	//文章点赞或者收藏（登录用户点赞）
	e.PUT(base+"/:id/status", server.PostUpdateStatus)
	//文章点赞（非登录用户点赞）
	e.PUT(base+"/:id/liked", server.PostUpdateVisitorGood)
	//获取某个用户某个文章的点赞和评论情况
	e.GET(base+"/:id/status", server.PostGetCollection)
	//小程序判断文章的收藏状态
	e.GET(base+"/:id/wechat_status/:openid", server.PostWechatGetCollection)
	//文章点赞或者收藏（小程序用户点赞）
	e.PUT(base+"/:id/wechat_status/:openid", server.PostWechatUpdateStatus)
	// 微信小程序用户发布评论
	e.POST(base+"/:id/wechat_comments", server.PostWechatCommitComment)
	// 获取小程序码
	e.GET(base+"/:id/mini_program_code", server.PostsWechatGetCode)
}

// UserRouter 用户相关的路由
func UserRouter(e *gin.Engine) {
	const base = Version + "user"

	// 获取用户信息
	e.GET(base, server.UserGetUser)

	// 更新用户个人信息
	e.PUT(base, server.UserPutUser)
	// 用户注册
	e.POST(base, server.UserAddUser)
	// 激活账户(因为链接限制，所以只能用get请求)
	e.GET(base+"/:id/status", server.UserActiveUser)

	// 用户登录
	e.POST(base+"/token", server.UserToken)
	// 判断用户是否登录
	e.GET(base+"/token", server.UserGetToken)

	// 发送重置密码邮件（这里还需要加上get请求,方便网页端操作）
	e.POST(base+"/password/email", server.UserV3PasswordEmail)
	e.GET(base+"/password/email", server.UserV3PasswordEmail)
	// 用户通过token来重置密码(用于重置密码界面)
	e.PUT(base+"/password/token", server.UserResetPasswordToken)
	// 兼容就版本的重置密码
	e.POST("/api/v2/access/set/password", server.UserResetPasswordToken)

	// 判断用户名是否存在
	e.GET(base+"/username", server.UserUserGetUserName)

	// 手机端获取验证码
	e.POST(base+"/code", server.UserPostCode)
	// 手机端用户注册
	e.POST(base+"/app", server.UserPostApp)
	// 手机端重置密码
	e.PUT(base+"/app/password", server.UserPutAppPassword)

	// 获取用户收藏的所有文章
	e.GET(base+"/collections", server.UserCollections)
	// 获取微信小程序用户的收藏的文章
	e.GET(base+"/wechat_collections/:openid", server.UserWeChatCollections)
}

// SettingRouter 设置相关的路由
func SettingRouter(e *gin.Engine) {
	const base = Version + "settings"
	const wechat = base + "/wechat"
	const plugins = base + "/plugins"
	const app = base + "/app"
	const site = base + "/site"
	const theme = base + "/themes"

	/* 微信小程序设置 */
	e.GET(wechat, server.Wechat)

	/* 插件设置,遗留设置 */
	//获取赞助界面设置
	e.GET(plugins+"/sponsor", server.PluginsGetDonate)
	//获取友链界面设置
	e.GET(plugins+"/friends", server.PluginsGetFriend)

	/* APP设置 */
	e.GET(app, server.SettingAPP)
	// 获取网页顶部和底部的设置
	e.GET(site+"/index", server.SettingGetIndex)
	//获取登录注册界面的设置
	e.GET(site+"/login", server.SettingGetLogin)
	//获取文章界面设置
	e.GET(site+"/post", server.SettingGetPost)

	//获取日记界面设置
	e.GET(site+"/diary", server.SettingGetDiary)
	//获取文档界面设置
	e.GET(site+"/doc", server.SettingGetDoc)
	//获取后台管理系统的设置
	e.GET(site+"/admin", server.SettingGetAdmin)
	// 更新用户id
	e.PUT(site+"/auth/user/:id", server.SettingUpdateUser)

	// 获取插件的设置
	e.GET(site+"/plugins/:name", server.SettingGetPlugins)

	// 获取主题设置
	e.GET(theme+"/:name", server.SettingGetThemes)
}

// ToolsRouter 一些通用的工具路由
func ToolsRouter(e *gin.Engine) {
	const base = Version + "tools"
	//获取表情的json数据
	e.GET(base+"/smiles", server.ToolsGetPostSmile)
	//小程序获取openid
	e.GET(base+"/openid/:code", server.ToolsGetOpenid)
	//提交意见反馈
	e.POST(base+"/advice", server.ToolsSubmitAdvice)
	//根据uid快速获取B站的头像和昵称
	e.GET(base+"/bili_info/:uid", server.ToolsGetBiliInfo)
	//获取网站的站点地图
	e.GET(base+"/sitemap", server.ToolsGetSiteMap)
	//获取手机APP的下载链接
	e.GET(base+"/app/version", server.ToolsGetAppVersion)
	// 图片上传接口 base64 格式
	e.POST(base+"/file/images/base64", admin.FileUploadImageBase64)
	// 普通格式
	e.POST(base+"/file/images", admin.FileUploadImageFile)
}

// PluginRouter 插件功能板块的路由
func PluginRouter(e *gin.Engine) {
	const base = Version + "plugins"
	//获取所有日记
	e.GET(base+"/diary", server.PluginsDiary)
	//获取文档列表
	e.GET(base+"/docs", server.PluginsDocList)
	//获取文档内容
	e.GET(base+"/docs/:id", server.PluginsDocContent)
	// 获取聊天室消息
	e.GET(base+"/chatRoom", server.PluginsChatroom)
}
