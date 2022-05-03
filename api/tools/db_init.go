// Package tools 数据库初始化工具类
// @Description
// @Author 小游
// @Date 2021/05/25
package tools

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"xBlog/internal/db"
	"xBlog/pkg/database"
)

// DbInit 初始化博客系统数据库
func DbInit() {
	// 判断是否有用户数据
	user := new(db.User)
	if err := database.NewDb(db.CollUser).FindOne(user); err != nil && err.Error() == db.MongoNoResult {
		// 初始化用户信息
		_ = database.NewDb(db.CollUser).DropCollection()
		// 初始化用户信息
		user.ID = primitive.NewObjectID()
		user.UserID = 1
		user.Username = "admin"
		user.Password = Encrypt("123456")
		user.Nickname = "小游"
		user.Registered = time.Now()
		user.Status = 1
		user.Identity = 1
		user.LoginInfo = []db.LoginInfo{}
		if database.NewDb(db.CollUser).InsertOne(&user) != nil {
			fmt.Println("初始化用户信息失败")
		} else {
			fmt.Println("初始化用户信息成功")
		}
	}

	// 判断是否有文章数据
	post := new(db.Article)
	if err := database.NewDb(db.CollArticle).FindOne(post); err != nil && err.Error() == db.MongoNoResult {
		// 文章信息初始化
		post.ID = primitive.NewObjectID()
		post.PostID = 1
		post.AuthorID = 1
		post.PostTime = time.Now()
		post.Md = "[success]CSS样式测试，如果你看到显示的是一个绿色的框，说明css工作正常！[/success]\n\n## 说明\n如果你看到**这段文字**，就说明你的博客安装成功！"
		post.Content = "<p>[success]CSS样式测试，如果你看到显示的是一个绿色的框，说明css工作正常！[/success]</p>\n<h2 id=\"h2-u8BF4u660E\"><a name=\"说明\" class=\"reference-link\"></a><span class=\"header-link octicon octicon-link\"></span>说明</h2><p>如果你看到<strong>这段文字</strong>，就说明你的博客安装成功！</p>\n"
		post.Title = "博客系统安装成功"
		post.Status = "publish"
		post.CommentStatus = "open"
		post.Modify = time.Now()
		post.Parent = 0
		post.PostType = "post"
		dbc := database.NewDb(db.CollArticle)
		if dbc.InsertOne(&post) != nil {
			fmt.Println("初始化文章信息失败")
		} else {
			fmt.Println("初始化文章信息成功")
		}
		// 初始化日记
		post.ID = primitive.NewObjectID()
		post.PostID = 2
		post.Md = "这个是你的第一篇日记，如果你看到了这条信息说明日记功能正常！"
		post.Content = "<p>这个是你的第一篇日记，如果你看到了这条信息说明日记功能正常！</p>"
		post.Title = ""
		post.PostType = "diary"
		if dbc.InsertOne(&post) != nil {
			fmt.Println("初始化日记信息失败")
		} else {
			fmt.Println("初始化日记信息成功")
		}
		// 初始化文档
		post.ID = primitive.NewObjectID()
		post.PostID = 3
		post.Md = "这个是你的第一篇文档，如果你看到了这条信息说明文档功能正常！"
		post.Content = "<p>这个是你的第一篇文档，如果你看到了这条信息说明文档功能正常！</p>"
		post.Title = "文档系统测试"
		post.PostType = "doc"
		if dbc.InsertOne(&post) != nil {
			fmt.Println("初始化文档信息失败")
		} else {
			fmt.Println("初始化文档信息成功")
		}
	}

	// 判断是否有评论数据
	comment := new(db.Comment)
	if err := database.NewDb(db.CollComment).FindOne(comment); err != nil && err.Error() == db.MongoNoResult {
		comment.ID = primitive.NewObjectID()
		comment.CommentID = 1
		comment.PostID = 1
		comment.UserID = 1
		comment.Nickname = "小游"
		comment.CommentTime = time.Now()
		comment.Content = "文章评论测试"
		comment.Agree = 1
		if database.NewDb(db.CollComment).InsertOne(&comment) != nil {
			fmt.Println("初始化评论失败")
		} else {
			fmt.Println("初始化评论成功")
		}
	}

	// 初始化友链数据
	friend := new(db.Friend)
	if err := database.NewDb(db.CollFriend).FindOne(user); err != nil && err.Error() == db.MongoNoResult {
		friend.ID = primitive.NewObjectID()
		friend.Url = "https://xiaoyou66.com"
		friend.Name = "小游"
		friend.Avatar = "https://img.xiaoyou66.com/images/2020/02/20/tTSY.jpg"
		friend.Description = "二次元技术宅"
		friend.Email = "xiaoyou2333@foxmail.com"
		friend.Status = 1
		friend.ApplicationTime = time.Now()
		if database.NewDb(db.CollFriend).InsertOne(&friend) != nil {
			fmt.Println("初始化友链失败")
		} else {
			fmt.Println("初始化友链成功")
		}
	}

	// 初始化分类和标签
	tag := new(db.Tag)
	if err := database.NewDb(db.CollTag).FindOne(user); err != nil && err.Error() == db.MongoNoResult {
		tag.ID = primitive.NewObjectID()
		tag.ItemID = 1
		tag.Name = "默认"
		tag.Chain = "%e9%bb%98%e8%ae%a4"
		tag.ItemType = "category"
		tag.Posts = []int{1}
		dbc := database.NewDb(db.CollTag)
		if dbc.InsertOne(&tag) != nil {
			fmt.Println("初始化分类失败")
		} else {
			fmt.Println("初始化分类成功")
		}
		// 初始化标签
		tag.ID = primitive.NewObjectID()
		tag.ItemID = 2
		tag.ItemType = "tag"
		if dbc.InsertOne(&tag) != nil {
			fmt.Println("初始化标签失败")
		} else {
			fmt.Println("初始化标签成功")
		}
	}

	// 初始化网站数据
	setting := new(db.SiteSetting)
	if err := database.NewDb(db.CollSiteSetting).FindOne(setting); err != nil && err.Error() == db.MongoNoResult {
		// 初始化网站设置数据
		dbc := database.NewDb(db.CollSiteSetting)
		for k, v := range db.SiteOption {
			site := *new(db.SiteSetting)
			//自动插入这个值
			site.ID = primitive.NewObjectID()
			site.Key = k
			site.Value = v
			_ = dbc.InsertOne(site)
		}
	}

	// 初始化聊天数据
	chat := new(db.ChatInfo)
	if err := database.NewDb(db.CollChat).FindOne(chat); err != nil && err.Error() == db.MongoNoResult {
		chat.ID = primitive.NewObjectID()
		chat.UserId = 1
		chat.Date = time.Now().UnixNano()
		chat.Content = "测试"
		dbc := database.NewDb(db.CollChat)
		if dbc.InsertOne(chat) == nil {
			fmt.Println("初始化聊天数据成功")
		} else {
			fmt.Println("初始化聊天数据失败")
		}
	}

	// 初始化分类和导航
	navigation := new(db.Navigation)
	if err := database.NewDb(db.CollNavigation).FindOne(navigation); err != nil && err.Error() == db.MongoNoResult {
		navigation.ID = primitive.NewObjectID()
		navigation.NavigationID = 1
		navigation.Name = "测试"
		navigation.Value = "#000000D"
		dbc := database.NewDb(db.CollNavigation)
		if dbc.InsertOne(navigation) == nil {
			fmt.Println("初始化分类和导航成功")
		} else {
			fmt.Println("初始化分类和导航失败")
		}
	}
}
