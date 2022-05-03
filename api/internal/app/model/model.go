// Package model @Title  公开接口返回的模型
// @Description  V2接口返回的各种模型，全部采用结构体方式
package model

import (
	"xBlog/internal/db"
)

// List 带分页的数据
type List struct {
	Total    int           `json:"total"`    // 总计多少页
	Current  int           `json:"current"`  // 当前第几页
	Contents []interface{} `json:"contents"` //内容
}

// AppDownload APP下载链接
type AppDownload struct {
	Version  string `json:"version"`  //APP版本号
	Dec      string `json:"dec"`      //描述信息
	Download string `json:"download"` //下载链接
}

// PostListContent 文章内容
type PostListContent struct {
	ID         int    `json:"id"`         //文章id
	Title      string `json:"title"`      //文章标题
	Content    string `json:"content"`    //文章内容
	Date       string `json:"date"`       //文章发布时间
	Good       int    `json:"good"`       //文章点赞数
	View       int    `json:"view"`       //文章浏览量
	Image      string `json:"image"`      //文章图片
	Comment    int    `json:"comment"`    //文章评论数
	IsTop      bool   `json:"is_top"`     //文章是否置顶
	Encryption bool   `json:"encryption"` //文章是否加密
}

// PostList 文章列表
type PostList struct {
	Page     int               `json:"page"`      //文章总页数
	Now      int               `json:"now"`       //文章当前的页数
	PostList []PostListContent `json:"post_list"` //文章列表
}

// CategoryInfoParent 父节点的分类信息
type CategoryInfoParent struct {
	Count int    `json:"count"` //该分类下文章总数
	ID    int    `json:"id"`    //分类的id
	Link  string `json:"link"`  //分类的链接
	Name  string `json:"name"`  //分类的名字
}

// CategoryInfoChild 子节点的分类信息
type CategoryInfoChild struct {
	ID     int    `json:"id"`     //分类的id
	Parent int    `json:"parent"` //分类的父节点id
	Link   string `json:"link"`   //分类的地址
	Name   string `json:"name"`   //分类的名字
}

// Category 分类的结构
type Category struct {
	Parent []CategoryInfoParent `json:"parent"` //父级分类
	Child  []CategoryInfoChild  `json:"child"`  //子分类
}

// CategoryContent 分类和标签结构体
type CategoryContent struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

// PostContent 文章内容
type PostContent struct {
	ID            int               `json:"id"`             // 文章id
	Title         string            `json:"title"`          //文章标题
	Date          string            `json:"date"`           //文章日期
	View          int               `json:"view"`           //浏览数
	Comment       int               `json:"comment"`        //评论数
	Good          int               `json:"good"`           //点赞数
	CommentStatus string            `json:"comment_status"` //当前评论的状态
	Tag           []CategoryContent `json:"tag"`            //文章标签
	Category      []CategoryContent `json:"category"`       //文章的分类
	Content       string            `json:"content"`        //文章的内容
	Modify        string            `json:"modify"`         //文章修改时间
	Image         string            `json:"image"`          //文章的代表图片
	Alipay        string            `json:"alipay"`         // 支付宝
	Wechat        string            `json:"wechat"`         // 微信
	Encrypt       bool              `json:"encrypt"`        // 文章是否加密
}

// PostEncryptContent 加密文章内容
type PostEncryptContent struct {
	ID      int    `json:"id"`      // 文章id
	Content string `json:"content"` // 加密文章的内容
}

// PostComment 文章评论
type PostComment struct {
	ID       int    `json:"id"`       //评论id
	UserID   int    `json:"user_id"`  //评论者用户id(针对登录用户)
	Nickname string `json:"nickname"` //评论者昵称
	Avatar   string `json:"avatar"`   //评论者头像
	Content  string `json:"content"`  //评论的内容
	Date     string `json:"date"`     //评论的时间
	Url      string `json:"url"`      //评论者网站url
	PostID   int    `json:"post_id"`  //评论所属文章id
	Parent   int    `json:"parent"`   //父节点
	Hang     string `json:"hang"`     //评论者头像挂件
	Level    int    `json:"level"`    //评论者等级
	Uid      string `json:"uid"`      //B站用户uid
}

// Smile 表情数据json数据
type Smile struct {
	Type      string `json:"type"` //表情类型
	Container []struct {
		Desc string `json:"desc"`
		Icon string `json:"icon"`
		Text string `json:"text"`
	} `json:"container"`
}

// PostCollection 用户的收藏或者点赞状态
type PostCollection struct {
	Good       bool `json:"good"`
	Collection bool `json:"collection"`
}

// BiliPersonInfo B站个人信息
type BiliPersonInfo struct {
	Nickname string `json:"nickname"`  //B站昵称
	Uid      string `json:"uid"`       //B站uid
	Avatar   string `json:"avatar"`    //B站头像
	Level    int    `json:"level"`     //B站等级
	Sign     string `json:"sign"`      //B站个性签名
	Sex      string `json:"sex"`       //性别
	IsVip    int    `json:"is_vip"`    //是否为vip 0否 1是
	Hang     string `json:"hang"`      //头像挂件
	Card     string `json:"card"`      //动态背景卡片
	View     int    `json:"view"`      //浏览数
	Good     int    `json:"good"`      //点赞数
	Watch    int    `json:"watch"`     //播放数
	Fans     int    `json:"fans"`      //粉丝数
	TopImage string `json:"top_image"` //个人背景顶部图片
	Article  int    `json:"article"`   // 文章浏览数
}

// HeadMeta 网页SEO标签
type HeadMeta struct {
	Title       string `json:"title"`       // 网站标题
	Keyword     string `json:"keyword"`     // 网站关键词
	Description string `json:"description"` // 描述信息
	Url         string `json:"url"`         // 网页链接
	Image       string `json:"image"`       // 代表性的图片
	Icon        string `json:"icon"`        // 网站图标
}

// SideInfo 侧边栏信息
type SideInfo struct {
	Title string `json:"title"` // 插件的标题
	Html  string `json:"html"`  // HTML卡片
}

// SettingIndex 主页的设置信息
type SettingIndex struct {
	HeadMeta             HeadMeta     `json:"head_meta"`             // SEO标签
	HeadNav              []db.HeadNav `json:"head_nav"`              // 顶部标签
	Background           string       `json:"background"`            //主页背景
	NavigationBackground string       `json:"navigation_background"` // 头部背景
	SiteName             string       `json:"site_name"`             //网站名字
	Description          string       `json:"description"`           //网站描述
	SiteUrl              string       `json:"site_url"`              //网站url
	BuildTime            string       `json:"build_time"`            //网站创建时间
	BeiAn                string       `json:"bei_an"`                //网站备案号
	GovBeiAn             string       `json:"gov_bei_an"`            //公安备案
	LeftSide             []SideInfo   `json:"left_side"`             // 左侧边栏内容
	RightSide            []SideInfo   `json:"right_side"`            // 右侧边栏内容
}

// SettingPost 文章界面的设置
type SettingPost struct {
	HeadMeta           HeadMeta `json:"head_meta"`           //seo标签
	Background         string   `json:"background"`          //文章界面的背景
	UrlRedirect        bool     `json:"url_redirect"`        // 开启文章跳转
	RedirectBackground string   `json:"redirect_background"` // 跳转的背景
}

// SettingDiary 日记界面的设置
type SettingDiary struct {
	HeadMeta   HeadMeta `json:"head_meta"`  //seo标签
	Background string   `json:"background"` //日记界面的背景图片
}

// SettingDonate 赞助界面的设置
type SettingDonate struct {
	HeadMeta   HeadMeta `json:"head_meta"`  //seo标签
	Alipay     string   `json:"alipay"`     //支付宝收款码
	WeChat     string   `json:"wechat"`     //微信收款码
	Background string   `json:"background"` //赞助界面背景图片
}

// SettingFriend 友链界面的设置
type SettingFriend struct {
	HeadMeta   HeadMeta `json:"head_meta"`  //seo标签
	Name       string   `json:"name"`       //友链名字
	Dec        string   `json:"dec"`        //友链描述
	Link       string   `json:"link"`       //友链地址
	Avatar     string   `json:"avatar"`     //友链头像
	Background string   `json:"background"` //友链背景
}

// SettingAnimation 追番界面的设置
type SettingAnimation struct {
	HeadMeta   HeadMeta `json:"head_meta"`  //seo标签
	Background string   `json:"background"` //追番界面的背景图片
}

// SettingDouBan 我的豆瓣的设置
type SettingDouBan struct {
	HeadMeta   HeadMeta `json:"head_meta"`  //seo标签
	Background string   `json:"background"` //我的豆瓣界面的背景图片
	Last       string   `json:"last"`       //我的豆瓣最后更新的试卷
	Book       int      `json:"book"`       //我收藏的书的数目
	Movie      int      `json:"movie"`      //我收藏的电影的数目
	Music      int      `json:"music"`      //我收藏的音乐的数目
	Talk       string   `json:"talk"`       //一言信息
}

// SettingMoe 二次元导航设置
type SettingMoe struct {
	HeadMeta   HeadMeta `json:"head_meta"`  //seo标签
	Background string   `json:"background"` //个人导航背景图片
}

// SettingDoc 文档系统的设置
type SettingDoc struct {
	HeadMeta HeadMeta `json:"head_meta"` //seo标签
}

// SettingLogin 登录注册界面的设置
type SettingLogin struct {
	HeadMeta   HeadMeta `json:"head_meta"`  //SEO标签
	Background string   `json:"background"` //登录界面背景
	Logo       string   `json:"logo"`       //登录界面logo
	WebText    string   `json:"web_text"`   //登录界面文字
	SiteName   string   `json:"site_name"`  //网站名
}

// SettingAdmin 后台管理员设置
type SettingAdmin struct {
	Icon    string `json:"icon"`    //登录界面logo
	Title   string `json:"title"`   //登录界面文字
	Version string `json:"version"` // 当前后端的版本号
}

// SettingPlugin 插件界面设置
type SettingPlugin struct {
	HeadMeta   HeadMeta `json:"head_meta"`  // SEO标签
	Url        string   `json:"url"`        // 二级URL界面
	Background string   `json:"background"` // 背景
	Content    string   `json:"content"`    // 显示的内容
	CSS        []string `json:"css"`        // css样式
	Script     []string `json:"script"`     // JavaScript脚本
	Full       bool     `json:"full"`       // 是否全屏
	Side       bool     `json:"side"`       // 是否显示侧边栏
}

// SettingWechat 微信小程序设置
type SettingWechat struct {
	Price     string `json:"price"`      // 赞赏二维码
	HeadImage string `json:"head_image"` // 我的界面头部图片
	About     string `json:"about"`      // 关于作者界面展示内容
	Animation bool   `json:"animation"`  // 显示我的追番
	Friend    bool   `json:"friend"`     // 显示友人帐
	Donate    bool   `json:"donate"`     // 显示赞助界面
	DouBan    bool   `json:"dou_ban"`    // 显示我的豆瓣界面
	Comment   bool   `json:"comment"`    // 是否显示评论
	PostBar   bool   `json:"post_bar"`   // 是否显示评论工具条
}

// SettingApp APP设置
type SettingApp struct {
	Chat       bool   `json:"chat"`       // 是否开启聊天
	Friend     bool   `json:"friend"`     // 显示友人帐
	Animation  bool   `json:"animation"`  // 显示我的追番
	Donate     bool   `json:"donate"`     // 显示赞助界面
	DouBan     bool   `json:"dou_ban"`    // 显示我的豆瓣界面
	Music      bool   `json:"music"`      // 显示音乐盒
	Doc        bool   `json:"doc"`        // 显示文档系统
	Project    bool   `json:"project"`    // 显示我的项目
	Navigation bool   `json:"navigation"` // 显示个人导航
	Login      string `json:"login"`      // 登录注册界面背景
}

// UserLogin 用户的登录数据
type UserLogin struct {
	UserID int    `json:"user_id"` //用户id
	Token  string `json:"token"`   //用户token数据
}

// BiliBaseInfo 用户B站的基本信息
type BiliBaseInfo struct {
	Nickname string `json:"nickname"` //用户昵称
	Avatar   string `json:"avatar"`   //用户头像
	Hang     string `json:"hang"`     //头像挂件
	Level    int    `json:"level"`    //用户等级
}

// Sitemap 网站的站点地图
type Sitemap struct {
	Site     string           `json:"site"`     //网站名字
	Map      string           `json:"map"`      //站点地图地址
	Post     []SiteMapContent `json:"post"`     //网站所有文章
	Doc      []SiteMapContent `json:"doc"`      //网站所有文档
	Category []SiteMapContent `json:"category"` //网站所有分类
	Tag      []SiteMapContent `json:"tag"`      //网站所有标签
}

// SiteMapContent 站点地图的内容
type SiteMapContent struct {
	Title string `json:"title"` //标题
	Url   string `json:"url"`   //地址
}

// UserAdvice 用户的意见反馈
type UserAdvice struct {
	Concat  string `bson:"concat" json:"concat"`   //联系方式
	Content string `bson:"content" json:"content"` //用户反馈内容
}

// DiaryContent 日记内容
type DiaryContent struct {
	DiaryID  int    `json:"diary_id"` //日记id
	Content  string `json:"content"`  //日记内容
	Date     string `json:"date"`     //日记发布时间
	Comment  int    `json:"comment"`  //评论数
	Good     int    `json:"good"`     //点赞数
	Avatar   string `json:"avatar"`   //头像
	Nickname string `json:"nickname"` //昵称
	Encrypt  bool   `json:"encrypt"`  // 日记是否加密
}

// UserBaseInfo 获取作者的基本信息
type UserBaseInfo struct {
	Nickname string `json:"nickname"` //昵称
	Avatar   string `json:"avatar"`   //头像
}

// DocListContent 文档列表（包括文档的基本信息）
type DocListContent struct {
	ID     int    `json:"id"`     //文档id
	Title  string `json:"title"`  //文档标题
	Parent int    `json:"parent"` //文档父节点
}

// DocContent 文档内容
type DocContent struct {
	ID      int    `json:"id"`      // 文档id
	Title   string `json:"title"`   // 文档标题
	Content string `json:"content"` // 文档内容
}

// BiBoDynamic B站动态
type BiBoDynamic struct {
	View     int          `json:"view"`     //浏览量
	RePost   int          `json:"re_post"`  //转发量
	Comment  int          `json:"comment"`  //评论数
	Like     int          `json:"like"`     //点赞数
	Time     string       `json:"time"`     //发布时间
	Id       string       `json:"id"`       //动态id
	Bid      string       `json:"bid"`      //视频bid
	Types    string       `json:"types"`    //动态类型(video 视频 dynamic 普通动态 share分享数据)
	Pendant  string       `json:"pendant"`  //头像挂件
	Decorate BiBoDecorate `json:"decorate"` //动态挂件
	Content  interface{}  `json:"content"`  //动态内容
	Images   []string     `json:"images"`   //动态图片
	OImages  []string     `json:"o_images"` //转发的内容的图片
}

// BiBoDecorate 卡片挂件数据
type BiBoDecorate struct {
	Card   string `json:"card"`   //卡片
	Number string `json:"number"` //专属粉丝卡片号码
	Color  string `json:"color"`  //专属粉丝卡片号码颜色
	Name   string `json:"name"`   //专属粉丝卡片名字
}

// BiBoVideoContent 视频动态内容数据
type BiBoVideoContent struct {
	Aid      string `json:"aid"`      //视频aid
	Cid      string `json:"cid"`      //视频cid
	Dec      string `json:"dec"`      //视频描述
	Dynamic  string `json:"dynamic"`  //动态说明
	Pic      string `json:"pic"`      //视频封面
	Title    string `json:"title"`    //视频标题
	Coin     int    `json:"coin"`     //硬币数
	Barrage  int    `json:"barrage"`  //弹幕数
	View     int    `json:"view"`     //播放数
	Share    int    `json:"share"`    //分享数
	Comment  int    `json:"comment"`  //评论数
	Favorite int    `json:"favorite"` //收藏数
	Like     int    `json:"like"`     //点赞数
}

// BiBoDynamicContent 普通动态内容数据
type BiBoDynamicContent struct {
	Content string `json:"content"` //动态内容
}

// BiBoShareContent 转发动态
type BiBoShareContent struct {
	Content string `json:"content"` //转发说明（就是自己的评论）
	Type    string `json:"type"`    //转发的动态类型
	Uid     string `json:"uid"`     //原动态用户的uid
	Avatar  string `json:"avatar"`  //原动态用户的头像
	Name    string `json:"name"`    //原动态用户的名字
	Origin  string `json:"origin"`  //原动态哟门户的内容
	Aid     string `json:"aid"`     //视频aid
	Cid     string `json:"cid"`     //视频cid
	Desc    string `json:"desc"`    //视频描述
	Dynamic string `json:"dynamic"` //视频简介
	Pic     string `json:"pic"`     //视频封面
	Title   string `json:"title"`   //视频标题
	View    int    `json:"view"`    //播放数
	Like    int    `json:"like"`    //点赞数
}

// UserInfo ------管理员相关的返回结果
type UserInfo struct {
	Avatar       string `json:"avatar"`       // 头像
	Sign         string `json:"sign"`         // 个性签名
	Level        int    `json:"level"`        // 等级
	Hang         string `json:"hang"`         // 头像挂件
	Username     string `json:"username"`     // 用户名
	Nickname     string `json:"nickname"`     // 昵称
	Email        string `json:"email"`        // 邮箱地址
	UserID       int    `json:"user_id"`      // 用户id
	Identity     int    `json:"identity"`     // 用户身份
	Subscription bool   `json:"subscription"` // 用户是否订阅邮件
}

// WechatAccessToken 小程序获取access_token
type WechatAccessToken struct {
	AccessToken string `json:"access_token"` // token数据
	ExpiresIn   int    `json:"expires_in"`   // 过期时间
}

// ChatDialog 聊天室，列表
type ChatDialog struct {
	Id      int         `json:"id"`      // 聊天室id
	Avatar  string      `json:"avatar"`  // 头像
	Name    string      `json:"name"`    // 聊天室名字
	Message db.ChatInfo `json:"message"` // 消息
	Count   int         `json:"count"`   // 未读消息数目
}

// ChatMessage 消息内容
type ChatMessage struct {
	ID          string `json:"id"`           // 唯一ID标识
	UserId      int    `json:"user_id"`      // 用户id
	Avatar      string `json:"avatar"`       // 头像
	Nickname    string `json:"nickname"`     // 用户昵称
	Content     string `json:"content"`      // 聊天内容
	Date        int64  `json:"date"`         // 发送时间
	Target      int    `json:"target"`       // 发送的目标(这个是用户id，如果为0就说明是公共频道)
	MessageType int    `json:"message_type"` // 消息类型，目前一般为0,0表示文本消息
	Read        bool   `json:"read"`         // 用户是否已读，公共频道一般不管
}

// UserAuthDomain 用户的授权域名信息
type UserAuthDomain struct {
	Code int `json:"code"`
	Data struct {
		UserId int    `json:"user_id"`
		Domain string `json:"domain"`
	} `json:"data"`
	Msg string `json:"msg"`
}
