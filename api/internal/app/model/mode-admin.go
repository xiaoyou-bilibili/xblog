// Package model @Title  管理员接口模型
// @Description  管理员接口返回的模型
package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// AdminList 管理员带分页的数据
type AdminList struct {
	TotalNum int           `json:"total_num"` // 多少条数据
	Total    int           `json:"total"`     // 多少条数据
	Current  int           `json:"current"`   // 当前第几页
	Contents []interface{} `json:"contents"`  //内容
}

// AdminPostContent 文章内容
type AdminPostContent struct {
	ID       int      `json:"id"`       //文章id
	Title    string   `json:"title"`    //文章标题
	Content  string   `json:"content"`  //文章内容
	Date     string   `json:"date"`     //文章发布时间
	Good     int      `json:"good"`     //文章点赞数
	View     int      `json:"view"`     //文章浏览量
	Comment  int      `json:"comment"`  //文章评论数
	Status   string   `json:"status"`   //文章状态
	Category []string `json:"category"` // 文章分类信息
	Tags     []string `json:"tags"`     // 文章标签
	IsDraft  bool     `json:"is_draft"` // 是否为草稿
}

// AdminPostList 管理员文章列表
type AdminPostList struct {
	Page     int                `json:"page"`      //文章总页数
	Now      int                `json:"now"`       //文章当前的页数
	PostList []AdminPostContent `json:"post_list"` //文章列表
}

// AdminDiaryContent 日记内容
type AdminDiaryContent struct {
	ID      int    `json:"id"`       //文章id
	Content string `json:"content"`  //日记内容
	Date    string `json:"date"`     //日记发布时间
	Good    int    `json:"good"`     //日记点赞数
	Comment int    `json:"comment"`  //日记评论数
	Status  string `json:"status"`   //日记状态
	IsDraft bool   `json:"is_draft"` // 是否为草稿
}

// AdminPostDelete 管理员删除的文章列表
type AdminPostDelete struct {
	ID      int    `json:"id"`      // 文章id
	Title   string `json:"title"`   // 文章标题
	Type    string `json:"type"`    // 文章类型
	Content string `json:"content"` // 文章内容
}

// AdminDiaryList 日记列表
type AdminDiaryList struct {
	Page      int                 `json:"page"`       //日记总页数
	Now       int                 `json:"now"`        //日记当前的页数
	DiaryList []AdminDiaryContent `json:"diary_list"` //日记列表
}

// AdminCommentContent 评论内容
type AdminCommentContent struct {
	ID      int    `json:"id"`      //评论id
	Author  string `json:"author"`  //评论作者\
	Content string `json:"content"` // 评论内容
	//Title  string `json:"title"`  //评论文章
	Date   string `json:"date"`   //评论发布时间
	Email  string `json:"email"`  //评论者邮箱
	IP     string `json:"ip"`     //评论者ip
	Status int    `json:"status"` // 评论状态
}

// AdminCommentList 评论列表
type AdminCommentList struct {
	Page        int                   `json:"page"`         //总页数
	Now         int                   `json:"now"`          //当前的页数
	CommentList []AdminCommentContent `json:"comment_list"` //列表
}

// AdminUserContent 用户内容
type AdminUserContent struct {
	ID           int    `json:"id"`           //用户id
	Username     string `json:"username"`     //用户名
	Nickname     string `json:"nickname"`     //用户昵称
	Email        string `json:"email"`        //用户邮箱
	Registered   string `json:"registered"`   //用户注册时间
	LastLogin    string `json:"last_login"`   //用户上次登录时间
	LoginIp      string `json:"login_ip"`     // 用户上次登录ip
	Status       int    `json:"status"`       // 用户状态
	Subscription bool   `json:"subscription"` // 用户是否订阅
	Identity     int    `json:"identity"`     // 用户权限
}

// AdminUserList 用户列表
type AdminUserList struct {
	Page     int                `json:"page"`      //总页数
	Now      int                `json:"now"`       //当前的页数
	UserList []AdminUserContent `json:"user_list"` //列表
}

// AdminPostEditContent 文章编辑器的详细数据
type AdminPostEditContent struct {
	PostID   int      `json:"post_id"`  // 文章id
	Title    string   `json:"title"`    // 文章标题
	Html     string   `json:"html"`     // 文章html内容
	Md       string   `json:"md"`       // 文章md格式的内容
	Status   string   `json:"status"`   // 文章当前的状态
	Password string   `json:"password"` // 文章密码
	IsTop    bool     `json:"is_top"`   // 是否置顶
	Category []int    `json:"category"` // 文章分类
	Tags     []string `json:"tags"`     // 文章标签
	IsDraft  bool     `json:"is_draft"` // 是否为草稿
}

// AdminDiaryEditContent 日记编辑器的数据
type AdminDiaryEditContent struct {
	PostID   int    `json:"post_id"`  // 日记id
	Html     string `json:"html"`     // 日记html内容
	Md       string `json:"md"`       // 日记md格式的内容
	Status   string `json:"status"`   // 日记当前的状态
	Password string `json:"password"` // 日记密码
	IsDraft  bool   `json:"is_draft"` // 是否为草稿
}

// AdminCategoryList 分类列表
type AdminCategoryList struct {
	ID     int    `json:"id"`     //分类id
	Title  string `json:"title"`  //分类标题
	Parent int    `json:"parent"` //分类父节点
}

// AdminAddCategory 管理员添加新分类
type AdminAddCategory struct {
	ID int `json:"id"` // 文章id
}

// AdminPublishPost 发布文章返回的操作
type AdminPublishPost struct {
	PostID int `json:"post_id"` // 发布成功后的文章id
}

// AdminUploadImage 管理员上传图片的操作
type AdminUploadImage struct {
	Url  string `json:"url"`  // 图片上传成功的路径
	Name string `json:"name"` // 上传成功的文件名
}

// AdminFriends 管理员友链信息
type AdminFriends struct {
	ID          primitive.ObjectID `json:"id"`          // 友链id
	Url         string             `json:"url"`         // 友链地址
	Name        string             `json:"name"`        // 友链名字
	Avatar      string             `json:"avatar"`      // 友链头像
	Description string             `json:"description"` // 友链描述
	Email       string             `json:"email"`       // 申请者邮箱
	Status      int                `json:"status"`      // 友链状态
}

// AdminDonate 管理员赞助信息
type AdminDonate struct {
	ID         primitive.ObjectID `json:"id"`          // 赞助id
	Nickname   string             `json:"nickname"`    // 赞助者名字
	Amount     string             `json:"amount"`      // 赞助额
	Comment    string             `json:"comment"`     // 留言
	DonateTime string             `json:"donate_time"` // 赞助时间
}

// AdminNavigationCategory 管理员分类内容
type AdminNavigationCategory struct {
	ID    int    `json:"id"`    // 分类id
	Name  string `json:"name"`  // 分类名字
	Color string `json:"color"` // 分类颜色
}

// AdminNavigation 管理员导航网址
type AdminNavigation struct {
	ID     int    `json:"id"`     // 网址id
	Name   string `json:"name"`   // 网址的名字
	Url    string `json:"url"`    // 网址地址
	Parent int    `json:"parent"` // 网址所属的分类id
}

// AdminProjectTop 管理员我的项目的顶部链接
type AdminProjectTop struct {
	ID    primitive.ObjectID `json:"id"`    // 顶部链接id
	Image string             `json:"image"` //展示图片
	Title string             `json:"title"` //标题
	Url   string             `json:"url"`   //跳转地址
}

// AdminProjectCard 管理员我的项目的项目卡片
type AdminProjectCard struct {
	ID    primitive.ObjectID `json:"id"`    // 项目id
	Name  string             `json:"name"`  // 项目名字
	Image string             `json:"image"` // 项目图片
	Dec   string             `json:"dec"`   // 项目描述
	Time  string             `json:"time"`  // 项目发布时间
	Video string             `json:"video"` // 视频介绍地址
	Blog  string             `json:"blog"`  // 博客介绍地址
	Code  string             `json:"code"`  // 源码地址
}

// AdminDocAddChapter 管理员插入新的文档后返回插入的信息
type AdminDocAddChapter struct {
	ID int `json:"id"` // 文章id
}

// AdminDocContent 管理员获取文档内容
type AdminDocContent struct {
	PostID  int    `json:"post_id"` // 文章id
	Content string `json:"content"` // 文档内容
	Md      string `json:"md"`      // markdown格式内容
}

// AdminIndexTotal 主页文章，用户，浏览量和评论数
type AdminIndexTotal struct {
	Post    int `json:"post"`    // 文章总数
	User    int `json:"user"`    // 用户总数
	View    int `json:"view"`    // 浏览量
	Comment int `json:"comment"` // 评论数
}

// AdminPostDistributed 文章类型分布
type AdminPostDistributed struct {
	Post  int `json:"post"`  // 文章数目
	Doc   int `json:"doc"`   // 文档数目
	Diary int `json:"diary"` // 日记数目
}

// AdminVisualPostDetail 文章详细数据说明
type AdminVisualPostDetail struct {
	Title   []string `json:"title"`   // 文章标题列表
	Good    []int    `json:"good"`    // 文章点赞数
	Comment []int    `json:"comment"` // 文章评论数
	View    []int    `json:"view"`    // 文章浏览数
}

// AdminMusic163 网易云音乐数据结构
type AdminMusic163 struct {
	Name   string `bson:"name" json:"name"`     // 名字
	Artist string `bson:"artist" json:"artist"` // 作曲家
	Url    string `bson:"url" json:"url"`       // 音乐地址
	Cover  string `bson:"cover" json:"cover"`   // 封面
	Irc    string `bson:"lrc" json:"irc"`       // 歌词
}

// AdminSetting 管理员设置
type AdminSetting struct {
	Title   string      `json:"title"`   // 设置标题
	Type    string      `json:"type"`    // 设置的类型 input 输入框 switch 开关 upload 图片 divider 分割线 number 步进器
	Key     string      `json:"key"`     // 设置所在数据库字段
	Value   interface{} `json:"value"`   // 设置的值
	Dec     string      `json:"dec"`     // 设置的描述
	Default interface{} `json:"default"` // 默认的值
}

// AdminPluginsSetting 管理员界面插件的设置信息
type AdminPluginsSetting struct {
	Type    int            `json:"type"`    // 设置类型(1是直接返回对应的设置，2是使用vue的配置)
	Name    string         `json:"name"`    // 设置标签名字
	Extra   string         `json:"extra"`   // 额外的信息，用于选项选择
	Setting []AdminSetting `json:"setting"` // 设置值
}

// AdminPlugins 管理员插件设置
type AdminPlugins struct {
	Id          int    `json:"id"`
	Unique      string `json:"unique"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Site        string `json:"site"`
	Version     string `json:"version"`
	NewVersion  string `json:"new_version"`  // 新版本号
	DownloadUrl string `json:"download_url"` // 下载URL
}

// PluginsUpdate 插件更新信息
type PluginsUpdate struct {
	Code int `json:"code"`
	Data []struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Unique      string `json:"unique"`
		Description string `json:"description"`
		Version     string `json:"version"`
		Author      string `json:"author"`
		Site        string `json:"site"`
		DownloadUrl string `json:"download_url"`
	} `json:"data"`
	Msg string `json:"msg"`
}

// AdminUpdateInfo 更新信息
type AdminUpdateInfo struct {
	Version     string `json:"version"`
	DownloadUrl string `json:"download_url"`
}

// AdminPluginsShop 插件市场信息
type AdminPluginsShop struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Author      string `json:"author"`
	Site        string `json:"site"`
	DownloadUrl string `json:"download_url"`
}

// GlobalInfo 插件的全局信息
type GlobalInfo struct {
	Unique string  // 插件的唯一标识
	Path   string  // 插件路径
	Config Plugins // 插件的配置文件信息
}

// Plugins 插件的配置信息
type Plugins struct {
	Name        string   `json:"name"`
	Unique      string   `json:"unique"`
	Description string   `json:"description"`
	Version     string   `json:"version"`
	Author      string   `json:"author"`
	Site        string   `json:"site"`
	Auth        []string `json:"auth"`
}
