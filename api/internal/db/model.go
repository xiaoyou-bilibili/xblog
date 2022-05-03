// Package db @Description 数据库相关模型
// @Author 小游
// @Date 2021/04/10
package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// CollUser 集合名
const CollUser = "user"                        //用户
const CollArticle = "article"                  //文章
const CollComment = "comment"                  //评论
const CollShop = "shop_buy"                    //用户的购买信息
const CollShopResource = "shop_resource"       //资源信息
const CollTag = "tag"                          //标签信息
const CollBarrage = "barrage"                  //弹幕池
const CollFriend = "friend"                    //友人帐
const CollDonate = "donate"                    //赞助博主
const CollDouBan = "dou_ban"                   //我的豆瓣
const CollProjectTop = "project_top"           //我的项目顶部轮播图
const CollProjectCard = "project_card"         //我的项目卡片
const CollPostCollection = "post_collection"   //用户收藏或者点赞的文章
const CollEmailRegistered = "email_registered" //邮箱验证码注册
const CollSiteSetting = "site_setting"         //网站的各种设置信息
const CollOther = "other"                      //存放各种临时数据
const CollNavigation = "navigation"            //个人导航
const CollLog = "log"                          //网站日志
const CollVisitorGood = "visitor_good"         // 游客点赞记录
const CollChat = "chat"                        // 聊天信息

// MongoNoResult mongo返回的结果
const MongoNoResult = "mongo: no documents in result" //mongo没有找到数据时的报错内容

// KeyAdvice 关键词（用于临时集合）
const KeyAdvice = "user_advice"                 //意见反馈
const KeyBiliInfo = "bili_info"                 //B站个人信息卡片
const ShopBuyInfo = "shop_buy_info"             // 用户的购买信息
const WechatAccessToken = "wechat_access_token" // 微信小程序accessAccessToken

// SiteOption 网站的默认配置信息
var SiteOption = map[string]interface{}{
	KeyAppImg:                     "https://cdn.jsdelivr.net/gh/xiaoyou66/imgbed@1.0/xblog/RYKe.png",
	KeyH5Img:                      "https://cdn.jsdelivr.net/gh/xiaoyou66/imgbed@1.0/xblog/RTTQ.png",
	KeyMiniProgramImg:             "https://cdn.jsdelivr.net/gh/xiaoyou66/imgbed@1.0/xblog/Rgmt.jpg",
	KeyAppDownloadVersion:         "3.0",
	KeyAppDownloadDescription:     "1.接口适配，已经更新到最新的接口%%2.添加文章点赞，收藏功能%%3.新添加我的豆瓣功能%%4.添加个人中心%%5.修复安卓10图片选择闪退问题",
	KeyAppDownloadDownload:        "http://cdn.xiaoyou66.com/xiaoyoublog3.0.apk",
	KeyBiBoImg:                    "https://cdn.jsdelivr.net/gh/xiaoyou66/imgbed@1.0/xblog/tKb5.jpg",
	KeyBiliCookie:                 "",
	KeyBiliShow:                   "bilibili个人认证:一个永远都火不了的可怜up主",
	KeyBiliUID:                    "343147393",
	KeyBuildTime:                  "03/20/2019 00:00:00",
	KeyCommentListCount:           5,
	KeyCommentNum:                 7,
	KeyBarrageImg:                 "https://cdn.jsdelivr.net/gh/xiaoyou66/imgbed@1.0/xblog/tKb5.jpg",
	KeyDiaryAdminList:             10,
	KeyDiaryDecNum:                20,
	KeyDiaryImg:                   "https://cdn.jsdelivr.net/gh/xiaoyou66/imgbed@1.0/xblog/tKb5.jpg",
	KeyDiaryListCount:             10,
	KeyDouBanCookie:               "",
	KeyDouBanImg:                  "https://cdn.jsdelivr.net/gh/xiaoyou66/imgbed@1.0/xblog/tKb5.jpg",
	KeyDouBanLastUpdate:           "2020-07-14 23:49:06",
	KeyDouBanUser:                 "199424307",
	KeyDouBanListCount:            20,
	KeyAnimationImg:               "https://cdn.jsdelivr.net/gh/xiaoyou66/imgbed@1.0/xblog/tKb5.jpg",
	KeyFriendAvatar:               "https://cdn.jsdelivr.net/gh/xiaoyou66/imgbed@1.0/xblog/tTSY.jpg",
	KeyFriendDec:                  "二次元技术宅",
	KeyFriendLink:                 "https://xiaoyou66.com",
	KeyFriendName:                 "小游",
	KeyFriendImg:                  "https://cdn.jsdelivr.net/gh/xiaoyou66/imgbed@1.0/xblog/tKb5.jpg",
	KeyHeadImg:                    "https://cdn.jsdelivr.net/gh/xiaoyou66/imgbed@1.0/xblog/npBl.png",
	KeyAppTopImg:                  "https://cdn.jsdelivr.net/gh/xiaoyou66/imgbed@1.0/xblog/Rzrc.png",
	KeyAppDownloadUrl:             "http://api.xiaoyou66.com/blog/xiaoyoublog.html",
	KeyIndexImg:                   "https://cdn.jsdelivr.net/gh/xiaoyou66/imgbed@1.0/xblog/tKb5.jpg",
	KeyKeyword:                    "小游网,二次元博客,个人网站,萌萌的网站",
	KeyLoginLogo:                  "https://cdn.jsdelivr.net/gh/xiaoyou66/imgbed@1.0/xblog/yw9w.png",
	KeyMusicId:                    "2817376656",
	KeyMusicU:                     "47167ace8e856c0810dae1d3c903b68489b4f4718eabb5cc43ff0d2b53cf797533a649814e309366",
	KeyMusicContent:               nil,
	KeyMyAvatar:                   "https://cdn.jsdelivr.net/gh/xiaoyou66/imgbed@1.0/xblog/tTSY.jpg",
	KeyMyBackground:               "https://cdn.jsdelivr.net/gh/xiaoyou66/imgbed@1.0/xblog/y0qR.png",
	KeyMyIntroduce:                "二次元技术宅",
	KeyMyBiliBili:                 "https://space.bilibili.com/343147393",
	KeyMyGithub:                   "https://github.com/xiaoyou66",
	KeyMySteam:                    "https://steamcommunity.com/id/xiaoyou66",
	KeyMyTelegram:                 "https://t.me/xiaoyou625",
	KeyMyTwitter:                  "https://twitter.com/xiaoyou625",
	KeyMyZhiHu:                    "https://www.zhihu.com/people/xiao-you-89-91",
	KeyDonateImg:                  "https://cdn.jsdelivr.net/gh/xiaoyou66/imgbed@1.0/xblog/tKb5.jpg",
	KeyPostDecCount:               50,
	KeyPostImg:                    "https://cdn.jsdelivr.net/gh/xiaoyou66/imgbed@1.0/xblog/tKb5.jpg",
	KeyPostListCount:              5,
	KeySiteBeiAn:                  "赣ICP备19003009号",
	KeySiteGovBeiAn:               "",
	KeySiteDescription:            "一个分享个人笔记,个人生活经历,分享各种有趣资源,实用技能的个人二次元博客网站",
	KeySiteEmail:                  "xiaoyou2333@foxmail.com",
	KeySiteIcon:                   "https://cdn.jsdelivr.net/gh/xiaoyou66/imgbed@1.0/xblog/tTSY.jpg",
	KeySiteName:                   "小游网",
	KeyCommentCheck:               false,
	KeyToolNotice:                 "目前系统还在开发中~并不是稳定版本，还希望大家积极反馈bug",
	KeyUserListCount:              5,
	KeyWebText:                    "我的个人小站",
	KeyWechat:                     "https://cdn.jsdelivr.net/gh/xiaoyou66/imgbed@1.0/xblog/ySH4.png",
	KeyAlipay:                     "https://cdn.jsdelivr.net/gh/xiaoyou66/imgbed@1.0/xblog/yJWT.jpg",
	KeyNoFace:                     "//static.hdslb.com/images/member/noface.gif",
	KeyMoeImg:                     "https://cdn.jsdelivr.net/gh/xiaoyou66/imgbed@1.0/xblog/tKb5.jpg",
	KeyWechatMiniProgramId:        "wx144674b7cc62e08d",
	KeyWechatMiniProgramSecret:    "",
	KeyImgBedAddr:                 "",
	KeyImgLypBedAddr:              "",
	KeyImgLypBedUser:              "",
	KeyImgLypBedPassword:          "",
	KeyImgLypBedToken:             "",
	KeyImgBedUpload:               false,
	KeySiteApiServer:              "http://127.0.0.1:2333",
	KeyDatabaseBackup:             false,
	KeySmtpUser:                   "",
	KeySmtpPass:                   "",
	KeySmtpServer:                 "",
	KeySmtpPort:                   "",
	KeySmtpName:                   "XIAOYOU",
	KeyUrlRedirectOpen:            false,
	KeyUrlRedirectBackground:      "https://cdn.jsdelivr.net/gh/xiaoyou66/imgbed@1.0/xblog/tKb5.jpg",
	KeyWechatMiniProgramPrice:     "https://cdn.jsdelivr.net/gh/xiaoyou66/imgbed@1.0/xblog/price.png",
	KeyWechatMiniProgramHeadImage: "https://cdn.jsdelivr.net/gh/xiaoyou66/imgbed@1.0/xblog/head.png",
	KeyWechatMiniProgramAbout:     "作者:小游\r\n作者博客:xiaoyou66.com\r\n作者B站uid:343147393",
	KeyWechatMiniProgramAnimation: true,
	KeyWechatMiniProgramFriend:    true,
	KeyWechatMiniProgramDonate:    true,
	KeyWechatMiniProgramDouBan:    true,
	KeyWechatMiniProgramComment:   true,
	KeyWechatMiniProgramPostBar:   true,
	KeyTimeTaskSyncDouBan:         false,
	KeyTimeTaskSyncMusic163:       false,
	KeyChatRoomAvatar:             "https://cdn.jsdelivr.net/gh/xiaoyou66/imgbed@1.0/xblog/tTSY.jpg",
	KeySettingAppChat:             true,
	KeySettingAppFriend:           true,
	KeySettingAppAnimation:        true,
	KeySettingAppDonate:           true,
	KeySettingAppDouBan:           true,
	KeySettingAppMusic:            true,
	KeySettingAppDoc:              true,
	KeySettingAppProject:          true,
	KeySettingAppNavigation:       true,
	KeySettingAppLogin:            "",
	KeySettingSideMoreDiary:       true,
	KeySettingSideMoreDonate:      true,
	KeySettingSideMoreFriend:      true,
	KeySettingSideMoreAnimation:   true,
	KeySettingSideMoreBarrage:     true,
	KeySettingSideMoreDouBan:      true,
	KeySettingSideMoreProject:     true,
	KeySettingSideInfoLeft:        []SettingAdminSideDetail{},
	KeySettingSideInfoRight:       []SettingAdminSideDetail{},
	KeySettingNavInfo:             []HeadNav{{Title: "主站", Link: "/", Children: []HeadNav{}}},
	KeySettingUserDomainID:        "1",
}

// User 用户信息
type User struct {
	ID           primitive.ObjectID `bson:"_id"`          //文档id
	UserID       int                `bson:"user_id"`      //用户id
	Username     string             `bson:"username"`     //用户名
	Password     string             `bson:"password"`     //密码
	Nickname     string             `bson:"nickname"`     //昵称
	Email        string             `bson:"email"`        //用户邮箱
	Registered   time.Time          `bson:"registered"`   //注册时间
	Status       int                `bson:"status"`       //用户状态（0未激活 1已激活）
	Token        string             `bson:"token"`        //注册验证token
	LastTime     time.Time          `bson:"last_time"`    //上次登录时间
	LastIp       string             `bson:"last_ip"`      //上次登录ip
	Sign         string             `bson:"sign"`         //签名
	Hang         string             `bson:"hang"`         //头像挂件
	Level        int                `bson:"level"`        //用户等级
	Avatar       string             `bson:"avatar"`       //头像地址
	Subscription bool               `bson:"subscription"` //邮件订阅是否开启
	LoginFail    int                `bson:"login_fail"`   //用户登录失败的次数
	Identity     int                `bson:"identity"`     //用户身份1 是管理员 2普通用户
	LoginInfo    []LoginInfo        `bson:"login_info"`   //登录信息（用于身份验证）
}

// LoginInfo 用户登录的一些信息
type LoginInfo struct {
	Agent     string    `bson:"agent"`      //浏览器UA
	Ip        string    `bson:"ip"`         //登录ip地址
	Token     string    `bson:"token"`      //用户的token数据
	LoginTime time.Time `bson:"login_time"` //用户登录的时间
}

// Article 文章基本信息
type Article struct {
	ID            primitive.ObjectID `bson:"_id"`            //文档id
	PostID        int                `bson:"post_id"`        //文章id
	AuthorID      int                `bson:"author_id"`      //作者id
	PostTime      time.Time          `bson:"post_time"`      //文章发布时间
	Content       string             `bson:"content"`        //文章内容
	Title         string             `bson:"title"`          //文章标题
	Status        string             `bson:"status"`         //状态(private私有 publish公开 draft草稿)
	CommentStatus string             `bson:"comment_status"` //评论状态(open允许评论 close关闭评论 login仅允许登录评论)
	Password      string             `bson:"password"`       //文章密码
	Modify        time.Time          `bson:"modify"`         //文章修改时间
	Parent        int                `bson:"parent"`         //文章的父文档(一般用于文档的分级)
	PostType      string             `bson:"post_type"`      //文章类型(post普通文章 doc文档 diary日记)
	Md            string             `bson:"markdown"`       //文章markdown格式
	Comment       int                `bson:"comment"`        //评论数
	Good          int                `bson:"good"`           //点赞数
	View          int                `bson:"view"`           //浏览数
	Guid          string             `bson:"guid"`           //文章代表图
	ISTop         bool               `bson:"is_top"`         //文章是否置顶
	IsDraft       bool               `bson:"is_draft"`       //是否为草稿
	Delete        bool               `json:"delete"`         //文章是否删除
}

// Comment 评论基本信息
type Comment struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`                    //文档id
	CommentID   int                `bson:"comment_id" json:"comment_id"`     //评论id
	PostID      int                `bson:"post_id" json:"post_id"`           //文章id
	UserID      int                `bson:"user_id" json:"user_id"`           //用户id
	Parent      int                `bson:"parent" json:"parent"`             //父评论id（如果没有那么就为0）
	Nickname    string             `bson:"nickname" json:"nickname"`         //昵称
	Site        string             `bson:"site" json:"site"`                 //站点url
	Email       string             `bson:"email" json:"email"`               //邮箱地址
	CommentTime time.Time          `bson:"comment_time" json:"comment_time"` //评论时间
	Ip          string             `bson:"ip" json:"ip"`                     //评论ip
	Content     string             `bson:"content" json:"content"`           //评论内容
	Agree       int                `bson:"agree" json:"agree"`               //评论状态(0拒绝 1同意)
	Agent       string             `bson:"agent" json:"agent"`               //用户浏览器UA
	Avatar      string             `bson:"avatar" json:"avatar"`             //评论者头像
	Level       int                `bson:"level" json:"level"`               //评论者B站等级
	Hang        string             `bson:"hang" json:"hang"`                 //评论者B站头像挂件
	Uid         string             `bson:"uid" json:"uid"`                   //评论者B站uid
	OpenID      string             `bson:"openid" json:"openid"`             //评论者微信openid
}

// BuyInfo 用户的购买信息
type BuyInfo struct {
	BuyTime time.Time `bson:"buy_time"` //购买时间
	ShopID  int       `bson:"shop_id"`  //商品id
	Agree   int       `bson:"agree"`    //是否同意购买(0拒绝 1同意)
}

// ShopBuy 用户购买的商品数据(仅限小游网站长使用)
type ShopBuy struct {
	ID      primitive.ObjectID `bson:"_id"`      //文档id
	UserID  int                `bson:"user_id"`  //用户id
	BuyInfo []BuyInfo          `bson:"buy_info"` //用户的购买信息
}

// ShopResource 资源数据表(仅限小游网站长使用)
type ShopResource struct {
	ID           primitive.ObjectID `bson:"_id"`           //文档id
	ShopID       int                `bson:"shop_id"`       //商品id
	Name         string             `bson:"name"`          //商品名字
	Version      string             `bson:"version"`       //版本号
	HelpName     string             `bson:"help_name"`     //帮助文档名字
	HelpUrl      string             `bson:"help_url"`      //帮助文档地址
	ResourceName string             `bson:"resource_name"` //资源名字
	ResourceUrl  string             `bson:"resource_url"`  //资源下载地址
}

// Tag 标签分类信息
type Tag struct {
	ID          primitive.ObjectID `bson:"_id"`         //文档id
	ItemID      int                `bson:"item_id"`     //项目id
	Name        string             `bson:"name"`        //名字
	Chain       string             `bson:"chain"`       //自定义链接
	Description string             `bson:"description"` //项目描述
	Parent      int                `bson:"parent"`      //父节点
	ItemType    string             `bson:"item_type"`   //项目类型(tag标签 category分类)
	Posts       []int              `bson:"posts"`       //该标签下所有文章
}

// Barrage 弹幕池
type Barrage struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`            //文档id
	Avatar   string             `bson:"avatar" json:"avatar"`     //头像地址
	Content  string             `bson:"content" json:"content"`   //弹幕内容
	Nickname string             `bson:"nickname" json:"nickname"` //用户昵称
	Color    string             `bson:"color" json:"color"`       //弹幕颜色
	Send     time.Time          `bson:"send_time" json:"send"`    //弹幕发送时间
}

// Friend 友人帐
type Friend struct {
	ID              primitive.ObjectID `bson:"_id" json:"id"`                            //文档id
	Url             string             `bson:"url" json:"url"`                           //友链地址
	Name            string             `bson:"name" json:"name"`                         //友链名字
	Avatar          string             `bson:"avatar" json:"avatar"`                     //友链头像
	Description     string             `bson:"description" json:"description"`           //友链描述
	Email           string             `bson:"email" json:"email"`                       //友链邮箱
	Status          int                `bson:"status" json:"status"`                     //友链状态
	ApplicationTime time.Time          `bson:"application_time" json:"application_time"` //申请时间
}

// Donate 赞助
type Donate struct {
	ID         primitive.ObjectID `bson:"_id"`         //文档id
	Nickname   string             `bson:"nickname"`    //用户昵称
	Amount     string             `bson:"amount"`      //赞助额
	Comment    string             `bson:"comment"`     //留言
	DonateTime time.Time          `bson:"donate_time"` //赞助时间
}

// DouBan 我的豆瓣
type DouBan struct {
	ID       primitive.ObjectID `bson:"_id"`       //文档id
	Name     string             `bson:"name"`      //名字
	Image    string             `bson:"image"`     //图片
	Score    int                `bson:"score"`     //评分
	PubInfo  string             `bson:"pub_info"`  //出版信息
	Comment  string             `bson:"comment"`   //我的评论
	Status   string             `bson:"status"`    //状态(collect已看 do在看 wish想看)
	ItemType string             `bson:"item_type"` //类型(book图书 movie电影 music音乐)
	Url      string             `bson:"url"`       //链接地址
}

// ProjectTop 我的项目（顶部轮播图）--- v3版本废除
type ProjectTop struct {
	ID   primitive.ObjectID `bson:"_id"`  //文档id
	Name string             `bson:"name"` //项目名字
	Img  string             `bson:"img"`  //图片地址
	Url  string             `bson:"url"`  //链接
}

// ProjectCard 我的项目(卡片)
type ProjectCard struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`                  //文档id
	Name        string             `bson:"name" json:"name"`               //项目名字
	Img         string             `bson:"img" json:"img"`                 //项目图片
	MakeTime    string             `bson:"make_time" json:"make_time"`     //制作时间
	Description string             `bson:"description" json:"description"` //项目描述
	VideoUrl    string             `bson:"video_url" json:"video_url"`     //视频地址
	BlogUrl     string             `bson:"blog_url" json:"blog_url"`       //博客文章地址
	CodeUrl     string             `bson:"code_url" json:"code_url"`       //github地址
}

// Project 我的项目 V3版本
type Project struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`                  //文档id
	Name        string             `bson:"name" json:"name"`               //项目名字
	Img         string             `bson:"img" json:"img"`                 //项目图片
	MakeTime    string             `bson:"make_time" json:"make_time"`     //制作时间
	Description string             `bson:"description" json:"description"` //项目描述
	VideoUrl    string             `bson:"video_url" json:"video_url"`     //视频地址
	BlogUrl     string             `bson:"blog_url" json:"blog_url"`       //博客文章地址
	CodeUrl     string             `bson:"code_url" json:"code_url"`       //github地址
	IsTop       bool               `bson:"is_top" json:"is_top"`           //是否作为轮播图
	Link        string             `bson:"link" json:"link"`               // 轮播图链接
}

// PostCollection 文章的点赞与收藏
type PostCollection struct {
	ID         primitive.ObjectID `bson:"_id"`        //文档id
	UserID     int                `bson:"user_id"`    //用户id,现在只支持int类型
	OpenID     string             `bson:"openid"`     // 微信小程序openid
	Collection []int              `bson:"collection"` //收藏的文章号
	Good       []int              `bson:"good"`       //点赞的文章号
}

// EmailRegistered 邮箱验证码注册
type EmailRegistered struct {
	ID    primitive.ObjectID `bson:"_id"`       //文档id
	Email string             `bson:"email"`     //注册用户的邮箱地址
	Token string             `bson:"token"`     //注册用户邮箱验证码
	Send  time.Time          `bson:"send_time"` //邮件的发送时间
}

// SiteSetting 网站配置信息
type SiteSetting struct {
	ID    primitive.ObjectID `bson:"_id"`   //文档id
	Key   string             `bson:"key"`   //设置名字
	Value interface{}        `bson:"value"` //设置值
}

// Other 临时数据存放
type Other struct {
	ID     primitive.ObjectID `bson:"_id" json:"id"`        //文档id
	Key    string             `bson:"key" json:"key"`       //临时数据的key
	Value  interface{}        `bson:"value" json:"value"`   //临时数据的值
	Create time.Time          `bson:"create" json:"create"` //临时数据创建的时间
}

// Navigation 个人导航
type Navigation struct {
	ID           primitive.ObjectID `bson:"_id"`           //文档id
	NavigationID int                `bson:"navigation_id"` //导航id
	Name         string             `bson:"name"`          //链接名字
	Value        string             `bson:"value"`         //值
	Parent       int                `bson:"parent"`        //父节点
}

// WebLog 网站运行日志
type WebLog struct {
	ID       primitive.ObjectID `bson:"_id"`      // 唯一ID标识
	Duration time.Duration      `bson:"duration"` // 请求持续时间
	Url      string             `bson:"url"`      // 请求url
	Ip       string             `bson:"ip"`       // 请求的ip
	Time     time.Time          `bson:"time"`     // 请求的试卷
	UA       string             `bson:"ua"`       // 用户浏览器UA
}

// VisitorGood 游客的点赞记录
type VisitorGood struct {
	ID    primitive.ObjectID `bson:"_id"`   // 唯一ID标识
	IP    string             `bson:"ip"`    // 游客的ip地址
	Posts []int              `bson:"posts"` // 点赞文章列表
}

// ChatInfo 用户聊天的接口
type ChatInfo struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`                    // 唯一ID标识
	UserId      int                `bson:"user_id" json:"user_id"`           // 用户id
	Content     string             `bson:"content" json:"content"`           // 聊天内容
	Date        int64              `bson:"date" json:"date"`                 // 发送时间
	Target      int                `bson:"target" json:"target"`             // 发送的目标(这个是用户id，如果为0就说明是公共频道)
	MessageType int                `bson:"message_type" json:"message_type"` // 消息类型，目前一般为0,0表示文本消息
	Read        bool               `bson:"read" json:"read"`                 // 用户是否已读，公共频道一般不管
	Delete      bool               `bson:"delete" json:"delete"`             // 是否删除
}

// SettingAdminSideDetail 插件的详细信息
type SettingAdminSideDetail struct {
	Name   string `json:"name"`   // 插件名字
	Unique string `json:"unique"` // 插件unique信息
}

// SettingAdminSide 管理员侧边栏设置
type SettingAdminSide struct {
	Left   []SettingAdminSideDetail `json:"left"`   // 左边插件信息
	Right  []SettingAdminSideDetail `json:"right"`  // 右边插件信息
	Unused []SettingAdminSideDetail `json:"unused"` // 可使用的插件信息
}

// HeadNav 头部导航栏
type HeadNav struct {
	Title    string    `json:"title"` // 功能标题
	Link     string    `json:"link"`  // 链接
	Children []HeadNav `json:"children"`
}
