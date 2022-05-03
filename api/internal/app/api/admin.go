// Package api
// @Description 管理员相关的接口
// @Author 小游
// @Date 2021/04/10
package api

import (
	"github.com/gin-gonic/gin"
	"xBlog/internal/app/service/admin"
)

// AdminVisualRouter 数据可视化
func AdminVisualRouter(e *gin.Engine) {
	const base = AdminVersion + "visualization"
	// 获取文章数，用户数，浏览数，评论数
	e.GET(base, admin.Visual)
	// 获取文章，日记，文档的数目分布情况
	e.GET(base+"/posts/distributed", admin.VisualPostDistributed)
	// 获取文章的详细分布数据
	e.GET(base+"/posts/detail", admin.VisualV3GetPostDetail)
}

// AdminPostsRouter 文章管理
func AdminPostsRouter(e *gin.Engine) {
	const base = AdminVersion + "posts"
	// 获取文章列表
	e.GET(base+"/articles", admin.PostsGetArticles)
	// 获取文章详细数据
	e.GET(base+"/articles/:id", admin.PostsGetContent)
	// 发布新文章
	e.POST(base+"/articles", admin.PostsAddArticle)
	// 更新文章内容
	e.PUT(base+"/articles/:id", admin.PostsUpdateArticle)
	// 删除某篇文章
	e.DELETE(base+"/articles/:id", admin.PostsDeleteArticle)

	// 获取日记列表
	e.GET(base+"/diary", admin.PostsGetDiary)
	// 发布日记/
	e.POST(base+"/diary", admin.PostsAddDiary)
	// 获取日记内容
	e.GET(base+"/diary/:id", admin.PostsGetDiaryContent)

	// 获取文章回收站数据
	e.GET(base+"/trash", admin.PostsGetTrash)

	// 获取所有文档
	e.GET(base+"/docs", admin.PostsGetDocs)
	// 添加文档
	e.POST(base+"/docs", admin.PostsAddDocs)
	// 获取文档内容
	e.GET(base+"/docs/:id", admin.PostsGetDocsContent)

	// 邮件通知功能,发送邮件通知用户
	e.POST(base+"/:id/subscription", admin.PostsNoticeUser)
	// 管理员获取文章分类
	e.GET(base+"/category", admin.PostsGetCategory)
	// 管理员新增文章分类
	e.POST(base+"/category", admin.PostsAddCategory)
	// 管理员更新分类
	e.PUT(base+"/category/:id", admin.PostsUpdateCategory)
	// 管理员删除分类
	e.DELETE(base+"/category/:id", admin.PostsDeleteCategory)
}

// AdminUsersRouter 用户管理
func AdminUsersRouter(e *gin.Engine) {
	const base = AdminVersion + "users"
	// 获取用户数据
	e.GET(base, admin.UsersGetUser)
	// 更新用户数据
	e.PUT(base+"/:id", admin.UsersUpdateUser)
	// 删除用户数据
	e.DELETE(base+"/:id", admin.UsersDeleteUsers)
}

// AdminCommentsRouter 评论管理
func AdminCommentsRouter(e *gin.Engine) {
	const base = AdminVersion + "comments"
	// 获取评论内容
	e.GET(base, admin.CommentsGetComments)
	// 更新评论内容
	e.PUT(base+"/:id", admin.CommentsUpdateComments)
	// 删除评论
	e.DELETE(base+"/:id", admin.CommentsDeleteComments)
}

// AdminSettingsRouter 设置管理
func AdminSettingsRouter(e *gin.Engine) {
	const base = "/api/v3/admin/settings"
	// 更新网站的设置
	e.PUT(base, admin.SettingsUpdate)
	// 获取站点设置
	e.GET(base+"/site", admin.SettingsGetSite)

	// 自动同步站点地图
	e.GET(base+"/site/sitemap", admin.SettingsGetSitemap)

	// 获取壁纸设置
	e.GET(base+"/background", admin.SettingsGetBackground)
	// 获取爬虫设置
	//e.GET(base+"/spider", admin.SettingsGetSpider)
	// 同步网易云音乐歌单
	e.PUT(base+"/spider/music163", admin.SettingsSpiderMusic163)
	// 同步豆瓣数据
	e.PUT(base+"/spider/dou_ban", admin.SettingsSpiderDouBan)

	// 获取侧边栏设置
	e.GET(base+"/side", admin.SettingsGetSide)
	// 更新侧边栏设置
	e.PUT(base+"/side", admin.SettingsUpdateSide)

	// 获取导航栏设置
	e.GET(base+"/nav", admin.SettingsGetNav)
	// 更新导航栏设置
	e.PUT(base+"/nav", admin.SettingsUpdateNav)

	// 获取小程序设置
	e.GET(base+"/other/mini_program", admin.SettingsGetMiniProgram)
	// 获取APP设置
	e.GET(base+"/other/app", admin.SettingsGetAPP)

	// 数据库备份
	e.GET(base+"/database/backup", admin.SettingsDatabaseBackup)
	// 数据库恢复
	e.GET(base+"/database/restore", admin.SettingsDatabaseRestore)

	// 获取主题设置
	e.GET(base+"/theme/:name", admin.SettingGetTheme)

}

// AdminPluginsRouter 插件管理
func AdminPluginsRouter(e *gin.Engine) {
	const base = "/api/v3/admin/plugins"
	const manage = base + "/manage"
	// 获取所有的插件信息
	e.GET(manage, admin.PluginsGetPlugins)
	// 删除插件
	e.DELETE(manage+"/:id", admin.PluginsRemove)
	// 添加插件
	e.POST(manage, admin.PluginsAdd)
	// 重载插件
	e.PUT(manage, admin.PluginsReload)
	// 获取插件设置
	e.GET(manage+"/setting/:id", admin.PluginsSetting)
}
