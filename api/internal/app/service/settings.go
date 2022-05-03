// Package server @Description  设置板块的v3版本
// @Author 小游
// @Date 2021/01/21
package server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"xBlog/internal/app/common"
	"xBlog/internal/app/model"
	"xBlog/internal/db"
	"xBlog/pkg/database"
	"xBlog/tools"
)

// SideMusicBox 获取侧边栏音乐盒的播放数据
func SideMusicBox(c *gin.Context) {
	// 因为获取到的数据需要进行处理
	//music:=database.GetSiteOption("music_content")
	//fmt.Println(reflect.TypeOf(database.GetSiteOption("music_content")))
	//reflect.TypeOf(database.GetSiteOption("music_content"))
	if data, ok := db.GetSiteOption(db.KeyMusicContent).(primitive.A); ok {
		var returnData []interface{}
		for _, v := range data {
			if data2, ok := v.(bson.D); ok {
				returnData = append(returnData, data2.Map())
			}
		}
		tools.GlobalResponse.ResponseOk(c, returnData)
	} else {
		tools.GlobalResponse.ResponseServerError(c)
	}
}

// SideMusicIrc 获取某一首歌曲的歌词
func SideMusicIrc(c *gin.Context) {
	id := c.Param("id")
	head := tools.HttpNewHead()
	head["Host"] = "music.163.com"
	irc, err := tools.HttpGetHead("http://music.163.com/api/song/media?id="+id, head)
	// 没有获取到内容直接返回空
	if !err {
		tools.GlobalResponse.ResponseHtml(c, "")
		return
	}
	var data interface{}
	_ = json.Unmarshal([]byte(irc), &data)
	if lyric, ok := data.(map[string]interface{})["lyric"]; ok && lyric != nil {
		tools.GlobalResponse.ResponseHtml(c, tools.Interface2String(lyric))
	} else {
		tools.GlobalResponse.ResponseHtml(c, "")
	}
}

// Wechat 获取微信小程序的设置
func Wechat(c *gin.Context) {
	var setting model.SettingWechat
	setting.Price = db.GetSiteOptionString(db.KeyWechatMiniProgramPrice)
	setting.HeadImage = db.GetSiteOptionString(db.KeyWechatMiniProgramHeadImage)
	setting.About = db.GetSiteOptionString(db.KeyWechatMiniProgramAbout)
	setting.Animation = db.GetSiteOptionBool(db.KeyWechatMiniProgramAnimation)
	setting.Friend = db.GetSiteOptionBool(db.KeyWechatMiniProgramFriend)
	setting.Donate = db.GetSiteOptionBool(db.KeyWechatMiniProgramDonate)
	setting.DouBan = db.GetSiteOptionBool(db.KeyWechatMiniProgramDouBan)
	setting.Comment = db.GetSiteOptionBool(db.KeyWechatMiniProgramComment)
	setting.PostBar = db.GetSiteOptionBool(db.KeyWechatMiniProgramPostBar)
	tools.GlobalResponse.ResponseOk(c, setting)
}

// PluginsGetDonate 获取赞助界面的设置
func PluginsGetDonate(c *gin.Context) {
	var setting model.SettingDonate
	setting.Alipay = db.GetSiteOptionString(db.KeyAlipay)
	setting.WeChat = db.GetSiteOptionString(db.KeyWechat)
	setting.Background = db.GetSiteOptionString(db.KeyDonateImg)
	setting.HeadMeta.Title = "赞助博主" + "-" + db.GetSiteOptionString(db.KeySiteName)
	setting.HeadMeta.Icon = db.GetSiteOptionString(db.KeySiteIcon)
	setting.HeadMeta.Url = db.GetSiteOptionString(db.KeySiteApiServer) + "/more/pay"
	setting.HeadMeta.Description = "如果觉得博主写的不错，可以打赏博客让博主更有动力哦"
	setting.HeadMeta.Keyword = "赞助博主,打赏博主,支持博主,爱心箱,支付宝付款,微信付款"
	tools.GlobalResponse.ResponseOk(c, setting)
}

// PluginsGetFriend 获取友链界面的设置
func PluginsGetFriend(c *gin.Context) {
	var setting model.SettingFriend
	setting.Name = db.GetSiteOptionString(db.KeyFriendName)
	setting.Dec = db.GetSiteOptionString(db.KeyFriendDec)
	setting.Link = db.GetSiteOptionString(db.KeyFriendLink)
	setting.Avatar = db.GetSiteOptionString(db.KeyFriendAvatar)
	setting.Background = db.GetSiteOptionString(db.KeyFriendImg)
	setting.HeadMeta.Title = "友人帐" + "-" + db.GetSiteOptionString(db.KeySiteName)
	setting.HeadMeta.Icon = db.GetSiteOptionString(db.KeySiteIcon)
	setting.HeadMeta.Url = db.GetSiteOptionString(db.KeySiteApiServer) + "/more/friend"
	setting.HeadMeta.Description = "欢迎各位大佬来交换友链"
	setting.HeadMeta.Keyword = "友情链接,友人帐,友链,我的伙伴"
	tools.GlobalResponse.ResponseOk(c, setting)
}

// SettingAPP 获取APP的设置
func SettingAPP(c *gin.Context) {
	var setting model.SettingApp
	setting.Chat = db.GetSiteOptionBool(db.KeySettingAppChat)
	setting.Friend = db.GetSiteOptionBool(db.KeySettingAppFriend)
	setting.Animation = db.GetSiteOptionBool(db.KeySettingAppAnimation)
	setting.Donate = db.GetSiteOptionBool(db.KeySettingAppDonate)
	setting.DouBan = db.GetSiteOptionBool(db.KeySettingAppDouBan)
	setting.Music = db.GetSiteOptionBool(db.KeySettingAppMusic)
	setting.Doc = db.GetSiteOptionBool(db.KeySettingAppDoc)
	setting.Project = db.GetSiteOptionBool(db.KeySettingAppProject)
	setting.Navigation = db.GetSiteOptionBool(db.KeySettingAppNavigation)
	setting.Login = db.GetSiteOptionString(db.KeySettingAppLogin)
	tools.GlobalResponse.ResponseOk(c, setting)
}

// SettingGetIndex 获取主页所有的设置
func SettingGetIndex(c *gin.Context) {
	// 获取页面信息
	page := c.Query("page")
	var index model.SettingIndex
	// 获取导航栏设置
	var nav = db.GetSiteOption(db.KeySettingNavInfo)
	// 获取左右侧边栏信息
	for _, v := range nav.(primitive.A) {
		tmp := db.HeadNav{}
		if tools.Primitive2Struct(v, &tmp) == nil {
			index.HeadNav = append(index.HeadNav, tmp)
		}
	}
	// 初始化headMeta信息
	index.HeadMeta = common.InitHeadMeta(
		db.GetSiteOptionString(db.KeyWebText),
		db.GetSiteOptionString(db.KeyKeyword),
		db.GetSiteOptionString(db.KeySiteDescription),
		"",
		"",
	)
	// 获取navigation
	index.NavigationBackground = db.GetSiteOptionString(db.KeyHeadImg)
	// 获取顶部设置
	index.SiteName = db.GetSiteOptionString(db.KeySiteName)
	index.Description = db.GetSiteOptionString(db.KeyWebText)
	index.SiteUrl = db.GetSiteOptionString(db.KeySiteApiServer)
	index.BuildTime = db.GetSiteOptionString(db.KeyBuildTime)
	index.BeiAn = db.GetSiteOptionString(db.KeySiteBeiAn)
	index.GovBeiAn = db.GetSiteOptionString(db.KeySiteGovBeiAn)
	// 根据不同的页面获取主页信息
	switch page {
	case "post":
		index.Background = db.GetSiteOptionString(db.KeyPostImg)
	case "diary":
		index.Background = db.GetSiteOptionString(db.KeyDiaryImg)
	default:
		index.Background = db.GetSiteOptionString(db.KeyIndexImg)
	}
	// 左右侧边栏内容
	common.InitSideInfo(&index)
	tools.GlobalResponse.ResponseOk(c, index)
}

// SettingGetLogin 获取登录界面设置
func SettingGetLogin(c *gin.Context) {
	var setting model.SettingLogin
	// 先获取SEO标签
	// 初始化headMeta信息
	setting.HeadMeta = common.InitHeadMeta(
		"",
		"登录注册界面，用于用户登录和注册",
		db.GetSiteOptionString(db.KeySiteDescription),
		"/login",
		"",
	)
	// 获取设置
	setting.Logo = db.GetSiteOptionString(db.KeyLoginLogo)
	setting.WebText = db.GetSiteOptionString(db.KeyWebText)
	setting.SiteName = db.GetSiteOptionString(db.KeySiteName)
	setting.Background = tools.GetRandomLoginImg()
	tools.GlobalResponse.ResponseOk(c, setting)
}

// SettingGetPost 获取文章界面设置
func SettingGetPost(c *gin.Context) {
	//获取文章id
	id := tools.Str2Int(c.Query("id"))
	if id == 0 {
		tools.GlobalResponse.ResponseBadRequest(c, "id格式非法")
		return
	}
	var setting model.SettingPost
	setting.Background = db.GetSiteOptionString(db.KeyPostImg)
	setting.RedirectBackground = db.GetSiteOptionString(db.KeyUrlRedirectBackground)
	setting.UrlRedirect = db.GetSiteOptionBool(db.KeyUrlRedirectOpen)
	setting.HeadMeta = common.InitHeadMeta(
		"",
		"",
		"",
		"/archives/"+tools.Int2Str(id),
		"",
	)
	//获取文章数据
	post := new(db.Article)
	if database.NewDb(db.CollArticle).SetFilter(bson.M{"post_id": id}).FindOne(post) == nil {
		setting.HeadMeta.Title = post.Title + "-" + setting.HeadMeta.Title
		setting.HeadMeta.Description = common.GetPostDec(post.Content)
		tags := new([]db.Tag)
		if database.NewDb(db.CollTag).SetFilter(bson.M{"posts": bson.M{"$in": []int{id}}}).FindMore(tags) == nil {
			//把文章标签和分类全部加到关键词里去
			for _, v := range *tags {
				//根据不同的类型添加数据
				setting.HeadMeta.Keyword += v.Name + ","
			}
		}
	}
	tools.GlobalResponse.ResponseOk(c, setting)
}

// SettingGetDiary 获取日记界面设置
func SettingGetDiary(c *gin.Context) {
	var setting model.SettingDiary
	setting.Background = db.GetSiteOptionString(db.KeyDiaryImg)
	setting.HeadMeta = common.InitHeadMeta(
		"日记板块",
		"博客日记,二次元日记,我的日记,个人日记",
		"日记板块是专门用来记录自己的日记，随笔等内容，过的再忙，也不要忘记写写日记记录时间",
		"/more/diary",
		"",
	)
	tools.GlobalResponse.ResponseOk(c, setting)
}

// SettingGetDoc 获取文档界面设置
func SettingGetDoc(c *gin.Context) {
	id := tools.Str2Int(c.Query("id"))
	var setting model.SettingDoc
	doc := new(db.Article)
	title := ""
	description := ""
	if database.NewDb(db.CollArticle).SetFilter(bson.M{"post_id": id, "post_type": "doc", "status": "publish"}).FindOne(doc) == nil {
		title = doc.Title
		description = common.GetPostDec(doc.Content)
	}
	setting.HeadMeta = common.InitHeadMeta(
		title,
		"文档系统,文档,个人文档系统,个人文档",
		description,
		"/doc",
		"",
	)
	tools.GlobalResponse.ResponseOk(c, setting)
}

// SettingGetAdmin 获取后台管理界面设置
func SettingGetAdmin(c *gin.Context) {
	var setting model.SettingAdmin
	setting.Title = "后台管理系统"
	setting.Icon = db.GetSiteOptionString(db.KeySiteIcon)
	setting.Version = Version
	tools.GlobalResponse.ResponseOk(c, setting)
}

// SettingUpdateUser 更新用户id
func SettingUpdateUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		tools.GlobalResponse.ResponseBadRequest(c)
		return
	}
	if db.SetSiteOption(db.KeySettingUserDomainID, id) {
		tools.GlobalResponse.ResponseOk(c, gin.H{"id": id})
	} else {
		tools.GlobalResponse.ResponseServerError(c)
	}
}

// SettingGetPlugins 获取插件界面的设置
func SettingGetPlugins(c *gin.Context) {
	// 首先判断这个页面是否存在
	if info, ok := common.PluginPage[c.Param("name")]; ok {
		// 执行函数获取返回结果
		if err, data := info(); err == nil {
			// 直接返回插件信息
			tools.GlobalResponse.ResponseOk(c, data)
		} else {
			tools.GlobalResponse.ResponseServerError(c, "插件执行错误，错误信息："+err.Error())
		}
	} else {
		tools.GlobalResponse.ResponseNotFound(c, "该页面不存在")
	}
}

// SettingGetDonate 赞助界面设置
func SettingGetDonate(c *gin.Context) {
	var setting model.SettingDonate
	setting.HeadMeta = common.InitHeadMeta(
		"赞助博主",
		"赞助博主,打赏博主,支持博主,爱心箱,支付宝付款,微信付",
		"如果觉得博主写的不错，可以打赏博客让博主更有动力哦",
		"/more/pay",
		"款",
	)
	setting.Alipay = db.GetSiteOptionString(db.KeyAlipay)
	setting.WeChat = db.GetSiteOptionString(db.KeyWechat)
	setting.Background = db.GetSiteOptionString(db.KeyDonateImg)
	tools.GlobalResponse.ResponseOk(c, setting)
}

// SettingGetAnimation 追番界面设置
func SettingGetAnimation(c *gin.Context) {
	var setting model.SettingAnimation
	setting.Background = db.GetSiteOptionString(db.KeyAnimationImg)
	setting.HeadMeta = common.InitHeadMeta(
		"我的追番",
		"B站追番,我的追番,追番,番剧,B站动漫,B站",
		"我的追番，记录自己的番剧记录",
		"/more/pay",
		"",
	)
	tools.GlobalResponse.ResponseOk(c, setting)
}

// SettingGetDouBan 获取豆瓣界面设置
func SettingGetDouBan(c *gin.Context) {
	var setting model.SettingDouBan
	setting.Background = db.GetSiteOptionString(db.KeyDouBanImg)
	setting.Last = db.GetSiteOptionString(db.KeyDouBanLastUpdate)
	//获取我收藏的电影，音乐，书籍的数目
	collection := database.NewDb(db.CollDouBan)
	if count, err := collection.SetFilter(bson.M{"item_type": "book"}).GetCount(); err == nil {
		setting.Book = int(count)
	}
	if count, err := collection.SetFilter(bson.M{"item_type": "movie"}).GetCount(); err == nil {
		setting.Movie = int(count)
	}
	if count, err := collection.SetFilter(bson.M{"item_type": "music"}).GetCount(); err == nil {
		setting.Music = int(count)
	}
	//获取一言
	setting.Talk = tools.HttpGet("https://v1.hitokoto.cn?c=a&c=b&c=c&c=d&c=e&c=f&c=g&c=h&c=i&c=j&c=k&c=l&encode=text")
	// 获取SEO标签
	setting.HeadMeta = common.InitHeadMeta(
		"我的豆瓣",
		"我的豆瓣,豆瓣,我收藏的豆瓣,博客豆瓣模板,豆瓣读书,豆瓣电影",
		"我的豆瓣可以显示我的读书记录，观影记录和听歌记录",
		"/more/dou-ban",
		"",
	)
	tools.GlobalResponse.ResponseOk(c, setting)
}

// SettingGetThemes 获取主题设置
func SettingGetThemes(c *gin.Context) {
	name := c.Param("name")
	// 为了避免用户获取到网站设置，这里需要对name进行过滤
	if !strings.HasPrefix(name, "theme_") {
		tools.GlobalResponse.ResponseBadRequest(c, "name非法！")
		return
	}
	tools.GlobalResponse.ResponseOk(c, gin.H{
		"key":   name,
		"value": db.GetSiteOption(name),
	})
}
