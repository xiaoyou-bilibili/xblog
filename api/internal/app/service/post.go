// Package server @Title  文章板块相关api处理函数
// @Description  这里负责处理各种和文章相关的API接口
// @Author 小游
// @Date 2021/04/10
package server

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"time"
	"xBlog/internal/app/common"
	"xBlog/internal/app/model"
	"xBlog/internal/db"
	"xBlog/pkg/database"
	"xBlog/tools"
)

// PostGetPostList 获取文章列表
func PostGetPostList(c *gin.Context) {
	//获取必要参数(第几页以及关键词)
	ids := c.Query("page")
	key := c.Query("q")
	tag := c.Query("tag")
	category := c.Query("category")
	if ids == "" {
		ids = "1"
	}
	//转换id
	id, err := strconv.Atoi(ids)
	if err != nil {
		tools.GlobalResponse.ResponseBadRequest(c, "页码错误")
	}
	//初始化数据库连接
	dB := database.NewDb(db.CollArticle)
	//判断有无关键词
	dB.SetFilter(bson.M{"post_type": "post", "status": bson.M{"$in": []string{"publish", "encrypt"}}, "is_draft": false, "delete": false})
	if key != "" {
		dB.OR([]bson.M{{"title": db.Regex(key)}, {"content": db.Regex(key)}})
	}
	//注意这里articles直接为指针变量，避免内存占用
	var articles []db.Article
	// 文章临时页数
	var page int
	// 文章总页数
	var totalPage = 0
	//如果有标签或者分类，那么就需要进行连表查询
	if tag != "" || category != "" {
		//先查标签表判断有无标签id
		var filter string
		if tag != "" {
			filter = tag
		} else {
			filter = category
		}
		// 临时文章，用于处理置顶文章
		tempArticles := new([]db.Article)
		//  首先我们查询置顶的文章
		page, err = database.NewDb(db.CollTag).
			SetUnwind("posts").
			SetLookUp(db.CollArticle, "posts", "post_id", "post").
			SetUnwind("post").
			SetAddFields(bson.M{"post_id": "$post.post_id", "title": "$post.title", "content": "$post.content", "post_time": "$post.post_time", "view": "$post.view", "good": "$post.good", "comment": "$post.comment", "post_type": "$post.post_type", "status": "$post.status", "is_draft": "$post.is_draft", "is_top": "$post.is_top", "delete": "$post.delete"}).
			SetSort(bson.M{"is_top": -1, "post_id": -1}).
			SetFilter(dB.Filter).
			OR([]bson.M{{"name": filter}, {"chain": filter}}).
			AggregatePaginate(id, db.GetSiteOptionInt(db.KeyPostListCount), tempArticles)
		if err == nil {
			totalPage += page
			// 我们把获取到的临时的文章全部添加到文章里面
			articles = append(articles, *tempArticles...)
		}
	} else { // 不是分类查询的情况
		// 这里设置两段排序is_top是先查询置顶的
		var tempArticles []db.Article
		page, err = dB.SetSort(bson.M{"is_top": -1, "post_id": -1}).Paginate(id, db.GetSiteOptionInt(db.KeyPostListCount), &tempArticles)
		if err == nil {
			totalPage += page
			// 我们把获取到的临时的文章全部添加到文章里面
			articles = append(articles, tempArticles...)
		}
	}
	//初始化返回数据
	postList := new(model.List)
	postList.Current = id
	postList.Total = totalPage
	postList.Contents = []interface{}{}
	for _, article := range articles {
		//获取文章的属性
		post := new(model.PostListContent)
		post.ID = article.PostID
		post.Title = article.Title
		// 判断文章是否加密
		if article.Status == "encrypt" {
			post.Encryption = true
			post.Content = "文章已加密"
		} else {
			// 过滤html标签
			post.Content = common.FilterContent(article.Content)
			post.Content = common.GetDec(post.Content, db.GetSiteOptionInt(db.KeyPostDecCount))
		}
		post.Date = tools.Time2String(article.PostTime, false)
		post.Image = tools.GetRandomImg()
		post.View = article.View
		post.Good = article.Good
		post.Comment = article.Comment
		post.IsTop = article.ISTop
		postList.Contents = append(postList.Contents, *post)
	}
	tools.GlobalResponse.ResponseOk(c, postList)
}

// PostGetCategory 获取文章分类
func PostGetCategory(c *gin.Context) {
	//先获取所有的分类信息
	category := new([]db.Tag)
	result := new(model.Category)
	result.Parent = []model.CategoryInfoParent{}
	result.Child = []model.CategoryInfoChild{}
	if database.NewDb(db.CollTag).SetFilter(bson.M{"item_type": "category"}).FindMore(category) != nil {
		tools.GlobalResponse.ResponseServerError(c, "获取分类数据失败")
		return
	}
	//遍历分类
	for _, v := range *category {
		//先获取父分类(自动过滤掉默认分类)
		if v.Parent == 0 && v.ItemID != 1 {
			parent := model.CategoryInfoParent{}
			parent.ID = v.ItemID
			parent.Link = v.Chain
			parent.Name = v.Name
			parent.Count = 0
			//计算子分类的数目
			for _, v2 := range *category {
				if v2.Parent == v.ItemID {
					parent.Count += len(v2.Posts)
				}
			}
			result.Parent = append(result.Parent, parent)
		} else if v.ItemID != 1 {
			//获取子分类
			child := model.CategoryInfoChild{}
			child.ID = v.ItemID
			child.Link = v.Chain
			child.Parent = v.Parent
			child.Name = v.Name
			result.Child = append(result.Child, child)
		}
	}
	tools.GlobalResponse.ResponseOk(c, result)
}

// PostGetPostContent 获取文章内容
func PostGetPostContent(c *gin.Context) {
	//获取id信息
	id := tools.Str2Int(c.Param("id"))
	if id == 0 {
		tools.GlobalResponse.ResponseBadRequest(c, "id错误")
		return
	}
	content := new(model.PostContent)
	article := new(db.Article)
	if database.NewDb(db.CollArticle).SetFilter(bson.M{"post_id": id, "is_draft": false, "delete": false, "status": bson.M{"$in": []string{"publish", "encrypt"}}}).FindOne(article) == nil {
		// 文章阅读数+1
		_ = database.NewDb(db.CollArticle).SetFilter(bson.M{"post_id": id}).Inc(bson.M{"view": 1}).UpdateOne()
		// 设置文章基本信息
		content.ID = article.PostID
		content.Title = article.Title
		content.Date = tools.Time2String(article.PostTime, false)
		content.View = article.View
		content.Comment = article.Comment
		content.Good = article.Good
		content.CommentStatus = article.CommentStatus
		// 初始化分类和标签
		content.Tag = []model.CategoryContent{}
		content.Category = []model.CategoryContent{}
		// 判断文章是否加密
		if article.Status == "encrypt" {
			content.Encrypt = true
			content.Content = "文章已加密，请输入访问密码"
		} else {
			content.Content = common.PostPostReplace(article.Content)
		}
		content.Modify = tools.Time2String(article.Modify, false)
		//图片暂时以随机获取为主
		content.Image = tools.GetRandomImg()
		//获取支付宝和微信的二维码
		content.Alipay = db.GetSiteOptionString(db.KeyAlipay)
		content.Wechat = db.GetSiteOptionString(db.KeyWechat)
		tags := new([]db.Tag)
		err := database.NewDb(db.CollTag).SetFilter(bson.M{"posts": bson.M{"$in": []int{id}}}).FindMore(tags)
		if err != nil {
			tools.GlobalResponse.ResponseServerError(c, "获取标签信息失败")
			return
		}
		//遍历文章或者标签
		for _, v := range *tags {
			//根据不同的类型添加数据
			if v.ItemType == "tag" {
				content.Tag = append(content.Tag, model.CategoryContent{Name: v.Name, Link: v.Chain})
			} else if v.ItemType == "category" {
				content.Category = append(content.Category, model.CategoryContent{Name: v.Name, Link: v.Chain})
			}
		}
		tools.GlobalResponse.ResponseOk(c, content)
	} else {
		tools.GlobalResponse.ResponseNotFound(c, "没有找到这篇文章")
	}
}

// PostGetPostEncryptContent 获取加密文章内容
func PostGetPostEncryptContent(c *gin.Context) {
	//获取id信息
	id := tools.Str2Int(c.Param("id"))
	password := c.Query("password")
	//fmt.Println(password)
	if id == 0 || password == "" {
		tools.GlobalResponse.ResponseBadRequest(c)
		return
	}
	content := new(model.PostEncryptContent)
	article := new(db.Article)
	// 开始查找文章内容
	if database.NewDb(db.CollArticle).SetFilter(bson.M{"post_id": id}).FindOne(article) == nil && article.Password == password {
		content.Content = common.PostPostReplace(article.Content)
		content.ID = article.PostID
		tools.GlobalResponse.ResponseOk(c, content)
	} else {
		tools.GlobalResponse.ResponseForbidden(c, "密码错误")
	}
}

// PostGetPostComment 获取文章评论
func PostGetPostComment(c *gin.Context) {
	//获取文章id
	id := tools.Str2Int(c.Param("id"))
	if id == 0 {
		tools.GlobalResponse.ResponseBadRequest(c, "文章id格式错误")
		return
	}
	//开始获取评论
	comments := new([]db.Comment)
	if database.NewDb(db.CollComment).SetSort(bson.M{"_id": -1}).SetFilter(bson.M{"post_id": id, "agree": 1}).FindMore(comments) == nil {
		//遍历数据
		var result []model.PostComment
		result = []model.PostComment{}
		for _, v := range *comments {
			comment := model.PostComment{
				ID:       v.CommentID,
				Nickname: v.Nickname,
				Url:      v.Site,
				PostID:   v.PostID,
				UserID:   v.UserID,
				Content:  common.CommentChangeSmile(v.Content, ""),
				Date:     tools.Time2String(v.CommentTime, true),
				Parent:   v.Parent,
				Hang:     v.Hang,
				Level:    v.Level,
				Uid:      v.Uid,
			}
			if v.Avatar != "" {
				comment.Avatar = v.Avatar
			} else {
				comment.Avatar = tools.GetRandomAvatar()
			}
			result = append(result, comment)
		}
		tools.GlobalResponse.ResponseOk(c, result)
	} else {
		tools.GlobalResponse.ResponseNotFound(c, "没有找到数据")
	}
}

// PostCommitComment 文章发布评论
func PostCommitComment(c *gin.Context) {
	param := new(model.Comment)
	if tools.ValidatorParam(c, param) {
		return
	}
	// 获取评论id
	postID := tools.Str2Int(c.Param("id"))
	if postID == 0 {
		tools.GlobalResponse.ResponseBadRequest(c, "文章id非法")
		return
	}
	var comment db.Comment
	comment.ID = primitive.NewObjectID()
	comment.PostID = postID
	comment.UserID = param.UserID
	comment.Parent = param.Parent
	comment.Nickname = param.Name
	comment.Site = param.Site
	comment.Email = param.Email
	comment.CommentTime = time.Now()
	comment.Ip = c.ClientIP()
	comment.Content = tools.ReplaceXss(param.Content)
	comment.Agent = c.GetHeader("User-Agent")
	//评论是否同意
	if db.GetSiteOptionBool(db.KeyCommentCheck) {
		comment.Agree = 0
	} else {
		comment.Agree = 1
	}
	comment.Avatar = param.Avatar
	comment.Level = tools.Str2Int(param.Level)
	comment.Hang = param.Hang
	comment.Uid = param.Uid
	// fmt.Println(comment.UserID)
	// 如果用户id不为空那么头像就直接查找个人信息表
	if comment.UserID != 0 {
		// 我们还需要验证一下用户权限
		if _, err := common.AccessGetTokenV2(c); err != nil {
			tools.GlobalResponse.ResponseUnauthorized(c)
			return
		}
		user := new(db.User)
		if database.NewDb(db.CollUser).SetFilter(bson.M{"user_id": comment.UserID}).FindOne(user) == nil {
			comment.Avatar = user.Avatar
		}
	}
	//插入评论数据
	commentId, err := database.NewDb(db.CollComment).InsertOneIncrease(comment, "comment_id")
	if err != nil {
		tools.GlobalResponse.ResponseServerError(c, "插入评论数据失败")
		return
	}
	// 设置评论id
	comment.CommentID = commentId
	//文章评论数+1（不考虑是否执行成功）
	_ = database.NewDb(db.CollArticle).SetFilter(bson.M{"post_id": comment.PostID}).Inc(bson.M{"comment": 1}).UpdateOne()
	//发送消息给站长
	server := db.GetSiteOptionString(db.KeySiteApiServer)
	body := "昵称:" + comment.Nickname + "<br>内容:" + comment.Content + "<br>文章地址:<a href='" + server + "/archives/" + strconv.Itoa(comment.PostID) + "'>" + server + "/archives/" + strconv.Itoa(comment.PostID) + "</a><br>邮箱:" + comment.Email
	err = tools.SendMail([]string{db.GetSiteOptionString(db.KeySiteEmail)}, "网站新评论", body)
	//发送邮件通知对方回复
	if comment.Parent != 0 {
		//获取对方的邮箱地址
		data := new(db.Comment)
		if database.NewDb(db.CollComment).SetFilter(bson.M{"comment_id": comment.Parent}).FindOne(data) == nil {
			body = `
				<div style="border-left:3px solid #d3d3d3;padding: 10px 20px;margin: 1.4em 0;color: #646464;">` + data.Content + `</div>
				昵称:` + comment.Nickname + `
				<br>内容: ` + comment.Content + `
				<br>文章地址:<a href="` + server + `/archives/` + strconv.Itoa(comment.PostID) + `">` + server + `/archives/` + strconv.Itoa(comment.PostID) + `</a>
				<br>本邮件由系统自动发出，请勿直接回复`
			_ = tools.SendMail([]string{data.Email}, "你的评论有了新的回复", body)
		}
	}
	//返回结果（结果里面包含评论发布信息）
	tools.GlobalResponse.ResponseCreated(c, comment)
}

// PostUpdateStatus 更新文章状态(更新文章的收藏或点赞)
func PostUpdateStatus(c *gin.Context) {
	//获取参数
	parm := new(model.PostStatus)
	// 验证参数
	if tools.ValidatorParam(c, parm) {
		return
	}
	// 验证身份
	user, err := common.AccessGetTokenV2(c)
	if err != nil {
		tools.GlobalResponse.ResponseUnauthorized(c, err.Error())
		return
	}
	// 获取文章id
	postId := tools.Str2Int(c.Param("id"))
	// 打印参数
	// 使用upsert。如果记录存在那么就自动更新，如果不存在就插入
	postCollection := database.NewDb(db.CollPostCollection).SetFilter(bson.M{"user_id": user.UserID}).SetUpsert(true)
	postCollection.Set(bson.M{"user_id": user.UserID})
	add := bson.M{}
	unset := bson.M{}
	// 判断用户是否收藏
	if parm.Collection {
		add["collection"] = postId
	} else {
		unset["collection"] = postId
	}
	// 判断用户是否点赞
	if parm.Good {
		add["good"] = postId
	} else {
		unset["good"] = postId
	}
	// 设置点赞或收藏
	if len(add) > 0 {
		postCollection.AddToSet(add)
	}
	if len(unset) > 0 {
		postCollection.Unset(unset)
	}
	// 先判断用户是否收藏过这篇文章
	collection := database.NewDb(db.CollPostCollection).
		SetFilter(bson.M{"user_id": user.UserID, "good": bson.M{"$in": []int{postId}}}).
		FindOne(new(db.PostCollection)) == nil
	// 更新数据库
	if postCollection.UpdateOne() == nil {
		inc := -1
		// 更新点赞记录
		if collection {
			if parm.Good {
				inc = 0
			} else {
				inc = -1
			}
		} else {
			if parm.Good {
				inc = 1
			} else {
				inc = 0
			}
		}
		if inc != 0 {
			// 更新点赞数据
			_ = database.NewDb(db.CollArticle).
				SetFilter(bson.M{"post_id": postId}).
				Inc(bson.M{"good": inc}).
				UpdateOne()
		}
		// 更新成功
		tools.GlobalResponse.ResponseCreated(c, parm)
	} else {
		tools.GlobalResponse.ResponseServerError(c, "更新失败")
	}
}

// PostUpdateVisitorGood 非登录用户点赞
func PostUpdateVisitorGood(c *gin.Context) {
	// 获取文章id
	postId := tools.Str2Int(c.Param("id"))
	if postId == 0 {
		tools.GlobalResponse.ResponseBadRequest(c, "文章id格式错误")
		return
	}
	// 获取用户ip地址
	ip := c.ClientIP()
	// 先判断用户是否点赞了文章
	collection := database.NewDb(db.CollVisitorGood)
	if collection.SetFilter(bson.M{"ip": ip, "posts": bson.M{"$in": []int{postId}}}).
		FindOne(new(db.PostCollection)) != nil {
		// 插入一条点赞数据
		if collection.SetFilter(bson.M{"ip": ip}).SetUpsert(true).AddToSet(bson.M{"posts": postId}).UpdateOne() == nil {
			// 更新点赞数据
			_ = database.NewDb(db.CollArticle).
				SetFilter(bson.M{"post_id": postId}).
				Inc(bson.M{"good": 1}).
				UpdateOne()
			tools.GlobalResponse.ResponseCreated(c, bson.M{"id": postId})
		} else {
			tools.GlobalResponse.ResponseServerError(c, "点赞失败")
		}
	} else {
		tools.GlobalResponse.ResponseUnProcessEntity(c, "你已点赞")
	}
}

// PostGetCollection 获取文章的收藏状态
func PostGetCollection(c *gin.Context) {
	//获取参数
	postID := tools.Str2Int(c.Param("id"))
	if postID == 0 {
		tools.GlobalResponse.ResponseBadRequest(c)
		return
	}
	// 验证用户
	user, err := common.AccessGetTokenV2(c)
	if err != nil {
		tools.GlobalResponse.ResponseUnauthorized(c)
		return
	}
	//获取用户的收藏数据
	var result model.PostCollection
	collection := new(db.PostCollection)
	if database.NewDb(db.CollPostCollection).SetFilter(bson.M{"user_id": user.UserID}).FindOne(collection) == nil {
		//找到数据了
		result.Collection = tools.IsInIntArray(collection.Collection, postID)
		result.Good = tools.IsInIntArray(collection.Good, postID)
	} else {
		//没找到数据
		result.Collection = false
		result.Good = false
	}
	tools.GlobalResponse.ResponseOk(c, result)
}

// PostWechatGetCollection 微信小程序获取文章的收藏状态
func PostWechatGetCollection(c *gin.Context) {
	//获取参数
	postID := tools.Str2Int(c.Param("id"))
	openID := c.Param("openid")
	if postID == 0 || openID == "" {
		tools.GlobalResponse.ResponseBadRequest(c)
		return
	}
	//获取用户的收藏数据
	var result model.PostCollection
	collection := new(db.PostCollection)
	if database.NewDb(db.CollPostCollection).SetFilter(bson.M{"openid": openID}).FindOne(collection) == nil {
		//找到数据了
		result.Collection = tools.IsInIntArray(collection.Collection, postID)
		result.Good = tools.IsInIntArray(collection.Good, postID)
	} else {
		//没找到数据
		result.Collection = false
		result.Good = false
	}
	tools.GlobalResponse.ResponseOk(c, result)
}

// PostWechatUpdateStatus 微信小程序用户更新文章状态(更新文章的收藏或点赞)
func PostWechatUpdateStatus(c *gin.Context) {
	// 验证参数
	parm := new(model.PostStatus)
	if tools.ValidatorParam(c, parm) {
		return
	}
	// 获取文章id
	postID := tools.Str2Int(c.Param("id"))
	openID := c.Param("openid")
	// 打印参数
	// 使用upsert。如果记录存在那么就自动更新，如果不存在就插入
	postCollection := database.NewDb(db.CollPostCollection).SetFilter(bson.M{"openid": openID}).
		SetUpsert(true)
	add := bson.M{}
	unset := bson.M{}
	// 判断用户是否收藏
	if parm.Collection {
		add["collection"] = postID
	} else {
		unset["collection"] = postID
	}
	// 判断用户是否点赞
	if parm.Good {
		add["good"] = postID
	} else {
		unset["good"] = postID
	}
	// 设置点赞或收藏
	if len(add) > 0 {
		postCollection.AddToSet(add)
	}
	if len(unset) > 0 {
		postCollection.Unset(unset)
	}
	// 先判断用户是否收藏过这篇文章
	collection := database.NewDb(db.CollPostCollection).
		SetFilter(bson.M{"openid": openID, "good": bson.M{"$in": []int{postID}}}).
		FindOne(new(db.PostCollection)) == nil
	// 更新数据库
	if postCollection.UpdateOne() == nil {
		inc := -1
		// 更新点赞记录
		if collection {
			if parm.Good {
				inc = 0
			} else {
				inc = -1
			}
		} else {
			if parm.Good {
				inc = 1
			} else {
				inc = 0
			}
		}
		if inc != 0 {
			// 更新点赞数据
		_:
			database.NewDb(db.CollArticle).
				SetFilter(bson.M{"post_id": postID}).
				Inc(bson.M{"good": inc}).
				UpdateOne()
		}
		// 更新成功
		tools.GlobalResponse.ResponseCreated(c, parm)
	} else {
		tools.GlobalResponse.ResponseServerError(c, "更新失败")
	}
}

// PostWechatCommitComment 微信小程序发布评论
func PostWechatCommitComment(c *gin.Context) {
	param := new(model.WechatComment)
	if tools.ValidatorParam(c, param) {
		return
	}
	// 获取评论id
	postID := tools.Str2Int(c.Param("id"))
	if postID == 0 {
		tools.GlobalResponse.ResponseBadRequest(c, "评论id非法")
		return
	}
	var comment db.Comment
	comment.ID = primitive.NewObjectID()
	comment.PostID = postID
	comment.Parent = param.Parent
	comment.Nickname = param.Name
	comment.OpenID = param.OpenID
	comment.CommentTime = time.Now()
	comment.Ip = c.ClientIP()
	comment.Content = tools.ReplaceXss(param.Content)
	comment.Agent = c.GetHeader("User-Agent")
	//评论是否同意
	comment.Agree = 0
	comment.Avatar = param.Avatar
	//插入评论数据
	commentId, err := database.NewDb(db.CollComment).InsertOneIncrease(comment, "comment_id")
	if err != nil {
		tools.GlobalResponse.ResponseServerError(c, "插入评论数据失败")
		return
	}
	// 设置评论id
	comment.CommentID = commentId
	//文章评论数+1（不考虑是否执行成功）
	_ = database.NewDb(db.CollArticle).SetFilter(bson.M{"post_id": comment.PostID}).Inc(bson.M{"comment": 1}).UpdateOne()
	//发送消息给站长
	server := db.GetSiteOptionString(db.KeySiteApiServer)
	body := "昵称:" + comment.Nickname + "<br>内容:" + comment.Content + "<br>文章地址:<a href='" + server + "/archives/" + strconv.Itoa(comment.PostID) + "'>" + server + "/archives/" + strconv.Itoa(comment.PostID) + "</a><br>邮箱:" + comment.Email
	err = tools.SendMail([]string{db.GetSiteOptionString(db.KeySiteEmail)}, "微信小程序评论", body)
	//发送邮件通知对方回复
	if comment.Parent != 0 {
		//获取对方的邮箱地址
		data := new(db.Comment)
		if database.NewDb(db.CollComment).SetFilter(bson.M{"comment_id": comment.Parent}).FindOne(data) == nil {
			body = `
				<div style="border-left:3px solid #d3d3d3;padding: 10px 20px;margin: 1.4em 0;color: #646464;">` + data.Content + `</div>
				昵称:` + comment.Nickname + `
				<br>内容: ` + comment.Content + `
				<br>文章地址:<a href="` + server + `/archives/` + strconv.Itoa(comment.PostID) + `">` + server + `/archives/` + strconv.Itoa(comment.PostID) + `</a>
				<br>本邮件由系统自动发出，请勿直接回复`
			_ = tools.SendMail([]string{data.Email}, "你的评论有了新的回复", body)
		}
	}
	//返回结果（结果里面包含评论发布信息）
	tools.GlobalResponse.ResponseCreated(c, comment)
}

// PostsWechatGetCode 微信小程序获取小程序码
func PostsWechatGetCode(c *gin.Context) {
	//获取参数
	postID := c.Param("id")
	if postID == "" {
		tools.GlobalResponse.ResponseBadRequest(c)
		return
	}
	data := new(db.Other)
	if database.NewDb(db.CollOther).
		SetFilter(bson.M{"key": db.WechatAccessToken}).
		FindOne(data) == nil {
		access := model.WechatAccessToken{}
		if tools.Primitive2Struct(data.Value, &access) == nil {
			// 这里我们获取到了access_token
			img := tools.HttpPostByte(
				"https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token="+access.AccessToken,
				map[string]string{"scene": postID, "page": "pages/index/post/post"})
			tools.GlobalResponse.ResponseByte(c, "image/jpeg", img)
		}
	}
}
