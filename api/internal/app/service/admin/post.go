// Package admin
// @Description 文章板块接口
// @Author 小游
// @Date 2021/04/14
package admin

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/url"
	"strings"
	"time"
	"xBlog/internal/app/common"
	"xBlog/internal/app/model"
	"xBlog/internal/db"
	"xBlog/pkg/database"
	"xBlog/tools"
)

// PostsGetArticles 管理员获取所有文章
func PostsGetArticles(c *gin.Context) {
	// 获取参数
	id := tools.Str2Int(c.Query("page"))
	size := tools.Str2Int(c.Query("page_size"))
	search := c.Query("search_type")
	key := c.Query("search_key")
	if id == 0 {
		id = 1
	}
	if size == 0 {
		size = 10
	}
	//设置关键词
	filter := bson.M{}
	filter["post_type"] = "post"
	filter["delete"] = false
	// 如果type和key不为空，说明设置了关键词搜索
	if search != "" && key != "" {
		// 这里如果是文章id，我们还需要额外判断
		if search == "post_id" {
			filter[search] = tools.Str2Int(key)
		} else {
			filter[search] = db.Regex(key)
		}
	}
	// 直接使用指针变量，避免内存占用
	articles := new([]db.Article)
	// 开始进行分页查询
	if total, page, err := database.NewDb(db.CollArticle).SetSort(bson.M{"_id": -1}).SetFilter(filter).PaginateWithTotal(id, size, articles); err == nil {
		// 循环遍历添加
		var postList model.AdminList
		postList.TotalNum = total
		postList.Total = page
		postList.Current = id
		postList.Contents = []interface{}{}
		// 初始化标签数据库，避免后面重复初始化
		tagDb := database.NewDb(db.CollTag)
		for _, article := range *articles {
			var postContent model.AdminPostContent
			postContent.Date = tools.Time2String(article.PostTime, true)
			postContent.ID = article.PostID
			postContent.Title = article.Title
			postContent.Content = common.GetPostDec(article.Content)
			postContent.Good = article.Good
			postContent.View = article.View
			postContent.Comment = article.Comment
			postContent.Status = article.Status
			postContent.IsDraft = article.IsDraft
			// 分类和tag初始化
			postContent.Category = []string{}
			postContent.Tags = []string{}
			// 开始获取文章目录
			tags := new([]db.Tag)
			if tagDb.SetFilter(bson.M{"posts": bson.M{"$in": []int{article.PostID}}}).FindMore(tags) == nil {
				for _, v := range *tags {
					//根据不同的类型添加数据
					if v.ItemType == "tag" {
						postContent.Tags = append(postContent.Tags, v.Name)
					} else if v.ItemType == "category" {
						postContent.Category = append(postContent.Category, v.Name)
					}
				}
			}
			postList.Contents = append(postList.Contents, postContent)
		}
		// 返回数据
		tools.GlobalResponse.ResponseOk(c, postList)
	} else {
		tools.GlobalResponse.ResponseServerError(c)
	}
}

// PostsUpdateArticle 管理员更新文章内容
func PostsUpdateArticle(c *gin.Context) {
	// 获取文章id
	id := c.Param("id")
	if id == "" {
		tools.GlobalResponse.ResponseBadRequest(c, "文章id格式错误")
		return
	}
	param := new(model.AdminUpdatePost)
	// 参数验证
	if tools.ValidatorParam(c, param) {
		return
	}
	// 判断内容对文章进行更新
	set := bson.M{}
	if param.Title != "" {
		set["title"] = param.Title
	}
	if param.Html != "" {
		set["content"] = strings.Replace(param.Html, "&quot;", "\"", -1)
	}
	if param.Status != "" {
		set["status"] = param.Status
	}
	if param.Password != "" {
		set["password"] = param.Password
	}
	if param.Md != "" {
		set["markdown"] = param.Md
	}
	if param.IsTop != "" {
		set["is_top"] = tools.Str2Bool(param.IsTop)
	}
	if param.IsDraft != "" {
		set["is_draft"] = tools.Str2Bool(param.IsDraft)
	}
	if param.Delete != "" {
		set["delete"] = tools.Str2Bool(param.Delete)
	}
	if param.Parent != "" {
		set["parent"] = tools.Str2Int(param.Parent)
	}
	// 更新文章的标签信息(这里因为id可能有多个的情况，我们这里默认只取第一个)
	common.AdminPostUpdateTag(tools.String2IntArray(id, ",")[0], param.Category, param.Tags)
	// 设置更新时间
	set["modify"] = time.Now()
	//time.Sleep(2*time.Second)
	// 设置文章过滤
	filter := bson.M{}
	filter["post_id"] = bson.M{"$in": tools.String2IntArray(id, ",")}
	if database.NewDb(db.CollArticle).Set(set).SetFilter(filter).UpdateMany() == nil {
		// 更新成功
		tools.GlobalResponse.ResponseCreated(c, param)
	} else {
		tools.GlobalResponse.ResponseServerError(c)
	}
}

// PostsDeleteArticle 管理员删除文章
func PostsDeleteArticle(c *gin.Context) {
	// 直接调用工具函数来删除数据
	common.AdminDeleteData(c, db.CollArticle, "post_id")
}

// PostsGetContent 获取文章详细数据
func PostsGetContent(c *gin.Context) {
	id := tools.Str2Int(c.Param("id"))
	article := new(db.Article)
	if database.NewDb(db.CollArticle).SetFilter(bson.M{"post_id": id}).FindOne(article) == nil {
		var post model.AdminPostEditContent
		post.Title = article.Title
		post.Html = article.Content
		post.Md = article.Md
		post.Status = article.Status
		post.Password = article.Password
		post.IsTop = article.ISTop
		post.IsDraft = article.IsDraft
		post.PostID = article.PostID
		// 初始化分类和标签
		post.Category = []int{}
		post.Tags = []string{}
		// 获取文章标签
		tags := new([]db.Tag)
		if database.NewDb(db.CollTag).SetFilter(bson.M{"posts": bson.M{"$in": []int{id}}}).FindMore(tags) == nil {
			for _, v := range *tags {
				if v.ItemType == "category" {
					post.Category = append(post.Category, v.ItemID)
				} else if v.ItemType == "tag" {
					post.Tags = append(post.Tags, v.Name)
				}
			}
		} else {
			tools.GlobalResponse.ResponseServerError(c, "获取标签信息失败")
			return
		}
		tools.GlobalResponse.ResponseOk(c, post)
	} else {
		tools.GlobalResponse.ResponseServerError(c, "获取数据失败")
	}
}

// PostsAddArticle 管理员新增文章
func PostsAddArticle(c *gin.Context) {
	// 获取用户id
	user := common.AccessAdminGetId(c)
	// 获取参数
	param := new(model.AdminAddPost)
	// 参数验证
	if tools.ValidatorParam(c, param) {
		return
	}
	// 新建文章
	article := new(db.Article)
	article.ID = primitive.NewObjectID()
	article.AuthorID = user
	article.PostTime = time.Now()
	article.Content = param.Html
	article.Title = param.Title
	article.Status = param.Status
	article.CommentStatus = "open"
	article.Password = param.Password
	article.Modify = time.Now()
	article.PostType = "post"
	article.Md = param.Md
	article.IsDraft = tools.Str2Bool(param.IsDraft)
	article.ISTop = tools.Str2Bool(param.IsTop)
	// 插入数据
	if id, err := database.NewDb(db.CollArticle).InsertOneIncrease(article, "post_id"); err == nil {
		// 更新文章的标签信息
		common.AdminPostUpdateTag(id, param.Category, param.Tags)
		// 返回成功信息
		tools.GlobalResponse.ResponseCreated(c, model.AdminPublishPost{PostID: id})
	} else {
		tools.GlobalResponse.ResponseServerError(c)
	}
}

// PostsGetDiary 管理员获取日记列表
func PostsGetDiary(c *gin.Context) {
	// 获取参数
	id := tools.Str2Int(c.Query("page"))
	size := tools.Str2Int(c.Query("page_size"))
	search := c.Query("search_type")
	key := c.Query("search_key")
	if id == 0 {
		id = 1
	}
	if size == 0 {
		size = 10
	}
	//设置关键词
	filter := bson.M{}
	filter["post_type"] = "diary"
	filter["delete"] = false
	// 如果type和key不为空，说明设置了关键词搜索
	if search != "" && key != "" {
		// 这里如果是文章id，我们还需要额外判断
		if search == "post_id" {
			filter[search] = tools.Str2Int(key)
		} else {
			filter[search] = db.Regex(key)
		}
	}
	// 直接使用指针变量，避免内存占用
	articles := new([]db.Article)
	// 开始进行分页查询
	if total, page, err := database.NewDb(db.CollArticle).SetSort(bson.M{"_id": -1}).SetFilter(filter).PaginateWithTotal(id, size, articles); err == nil {
		// 循环遍历添加
		var postList model.AdminList
		postList.TotalNum = total
		postList.Total = page
		postList.Current = id
		postList.Contents = []interface{}{}
		// 获取文章数据
		for _, article := range *articles {
			var postContent model.AdminDiaryContent
			postContent.Date = tools.Time2String(article.PostTime, true)
			postContent.ID = article.PostID
			postContent.Content = common.GetDec(article.Content, 20)
			postContent.Good = article.Good
			postContent.Comment = article.Comment
			postContent.Status = article.Status
			postContent.IsDraft = article.IsDraft
			// 日记列表添加数据
			postList.Contents = append(postList.Contents, postContent)
		}
		// 返回数据
		tools.GlobalResponse.ResponseOk(c, postList)
	} else {
		tools.GlobalResponse.ResponseServerError(c)
	}
}

// PostsAddDiary 管理员新增日记
func PostsAddDiary(c *gin.Context) {
	// 获取用户id
	user := common.AccessAdminGetId(c)
	// 获取参数
	param := new(model.AdminAddDiary)
	// 参数验证
	if tools.ValidatorParam(c, param) {
		return
	}
	article := new(db.Article)
	article.ID = primitive.NewObjectID()
	article.AuthorID = user
	article.PostTime = time.Now()
	article.Content = param.Html
	article.Status = param.Status
	article.CommentStatus = "open"
	article.Password = param.Password
	article.Modify = time.Now()
	article.PostType = "diary"
	article.Md = param.Md
	article.IsDraft = param.IsDraft
	// 插入数据
	if id, err := database.NewDb(db.CollArticle).InsertOneIncrease(article, "post_id"); err == nil {
		// 返回成功信息
		tools.GlobalResponse.ResponseCreated(c, model.AdminPublishPost{PostID: id})
	} else {
		tools.GlobalResponse.ResponseServerError(c)
	}
}

// PostsGetDiaryContent 管理员获取日记内容
func PostsGetDiaryContent(c *gin.Context) {
	id := tools.Str2Int(c.Param("id"))
	article := new(db.Article)
	if database.NewDb(db.CollArticle).SetFilter(bson.M{"post_id": id}).FindOne(article) == nil {
		var post model.AdminDiaryEditContent
		post.Html = article.Content
		post.Md = article.Md
		post.Status = article.Status
		post.Password = article.Password
		post.IsDraft = article.IsDraft
		post.PostID = article.PostID
		tools.GlobalResponse.ResponseOk(c, post)
	} else {
		tools.GlobalResponse.ResponseServerError(c)
	}
}

// PostsNoticeUser 发送邮件通知
func PostsNoticeUser(c *gin.Context) {
	// 权限认证
	id := tools.Str2Int(c.Param("id"))
	if common.PostNoticeUser(id) == nil {
		tools.GlobalResponse.ResponseCreated(c, bson.M{"id": id})
	} else {
		tools.GlobalResponse.ResponseServerError(c)
	}
}

// PostsGetDocs 获取所有文档
func PostsGetDocs(c *gin.Context) {
	articles := new([]db.Article)
	if database.NewDb(db.CollArticle).SetFilter(bson.M{"post_type": "doc", "status": "publish", "delete": false}).FindMore(articles) == nil {
		var docs []model.DocListContent
		for _, v := range *articles {
			var doc model.DocListContent
			doc.ID = v.PostID
			doc.Title = v.Title
			doc.Parent = v.Parent
			docs = append(docs, doc)
		}
		tools.GlobalResponse.ResponseOk(c, docs)
	} else {
		tools.GlobalResponse.ResponseServerError(c)
	}
}

// PostsAddDocs 添加文档
func PostsAddDocs(c *gin.Context) {
	user := common.AccessAdminGetId(c)
	// 获取参数
	param := new(model.AdminAddDoc)
	// 参数验证
	if tools.ValidatorParam(c, param) {
		return
	}
	// 初始化文档
	article := new(db.Article)
	article.ID = primitive.NewObjectID()
	article.AuthorID = user
	article.Title = param.Title
	article.Parent = param.Parent
	article.PostTime = time.Now()
	article.Modify = time.Now()
	article.Status = "publish"
	article.CommentStatus = "open"
	article.Modify = time.Now()
	article.PostType = "doc"
	article.IsDraft = false
	// 插入文档
	if id, err := database.NewDb(db.CollArticle).InsertOneIncrease(article, "post_id"); err == nil {
		tools.GlobalResponse.ResponseCreated(c, model.AdminDocAddChapter{ID: id})
	} else {
		tools.GlobalResponse.ResponseServerError(c)
	}
}

// PostsGetDocsContent 获取文档内容
func PostsGetDocsContent(c *gin.Context) {
	// 获取参数
	id := tools.Str2Int(c.Param("id"))
	if id == 0 {
		tools.GlobalResponse.ResponseBadRequest(c, "id格式错误")
		return
	}
	doc := new(db.Article)
	if database.NewDb(db.CollArticle).SetFilter(bson.M{"post_id": id}).FindOne(doc) == nil {
		var data model.AdminDocContent
		data.PostID = doc.PostID
		data.Content = doc.Content
		data.Md = doc.Md
		tools.GlobalResponse.ResponseOk(c, data)
	} else {
		tools.GlobalResponse.ResponseServerError(c)
	}
}

// PostsGetTrash 获取文章回收站
func PostsGetTrash(c *gin.Context) {
	// 获取参数
	id := tools.Str2Int(c.Query("page"))
	size := tools.Str2Int(c.Query("page_size"))
	search := c.Query("search_type")
	key := c.Query("search_key")
	if id == 0 {
		id = 1
	}
	if size == 0 {
		size = 10
	}
	//设置关键词
	filter := bson.M{}
	filter["delete"] = true
	// 如果type和key不为空，说明设置了关键词搜索
	if search != "" && key != "" {
		// 这里如果是文章id，我们还需要额外判断
		if search == "post_id" {
			filter[search] = tools.Str2Int(key)
		} else {
			filter[search] = db.Regex(key)
		}
	}
	// 获取文章列表
	articles := new([]db.Article)
	if total, page, err := database.NewDb(db.CollArticle).SetSort(bson.M{"_id": -1}).SetFilter(filter).PaginateWithTotal(id, size, articles); err == nil {
		// 循环遍历添加
		var postList model.AdminList
		postList.TotalNum = total
		postList.Total = page
		postList.Current = id
		postList.Contents = []interface{}{}
		// 遍历
		for _, v := range *articles {
			var article model.AdminPostDelete
			article.ID = v.PostID
			article.Content = common.GetDec(v.Content, 50)
			article.Title = v.Title
			article.Type = v.PostType
			postList.Contents = append(postList.Contents, article)
		}
		tools.GlobalResponse.ResponseOk(c, postList)
	} else {
		tools.GlobalResponse.ResponseServerError(c)
	}
}

// PostsGetCategory 获取文章分类
func PostsGetCategory(c *gin.Context) {
	// 获取所有的分类信息
	categoryS := new([]db.Tag)
	if database.NewDb(db.CollTag).SetFilter(bson.M{"item_type": "category"}).FindMore(categoryS) == nil {
		var data []model.AdminCategoryList
		// 遍历分类节点
		for _, v := range *categoryS {
			var item model.AdminCategoryList
			item.ID = v.ItemID
			item.Parent = v.Parent
			item.Title = v.Name
			data = append(data, item)
		}
		tools.GlobalResponse.ResponseOk(c, data)
	} else {
		tools.GlobalResponse.ResponseServerError(c)
	}
}

// PostsAddCategory 添加分类
func PostsAddCategory(c *gin.Context) {
	// 获取参数
	param := new(model.AdminAddCategoryParm)
	// 获取其他参数
	if tools.ValidatorParam(c, param) {
		return
	}
	// 添加分类信息
	var category db.Tag
	category.ID = primitive.NewObjectID()
	category.Name = param.Name
	category.Parent = param.Parent
	category.ItemType = param.Type
	category.Chain = url.QueryEscape(param.Name)
	category.Posts = []int{}
	// 插入分类信息
	if id, err := database.NewDb(db.CollTag).InsertOneIncrease(category, "item_id"); err == nil {
		tools.GlobalResponse.ResponseOk(c, bson.M{"item_id": id})
	} else {
		tools.GlobalResponse.ResponseServerError(c)
	}
}

// PostsUpdateCategory 更新分类
func PostsUpdateCategory(c *gin.Context) {
	// 获取分类id
	id := tools.Str2Int(c.Param("id"))
	// 获取参数
	param := new(model.AdminUpdateCategoryParm)
	// 获取其他参数
	if tools.ValidatorParam(c, param) {
		return
	}
	set := bson.M{}
	// 获取参数
	if param.Name != "" {
		set["name"] = param.Name
		set["chain"] = url.QueryEscape(param.Name)
	}
	if param.Type != "" {
		set["item_type"] = param.Type
	}
	if param.Parent != "" {
		set["parent"] = tools.Str2Int(param.Parent)
	}
	if database.NewDb(db.CollTag).SetFilter(bson.M{"item_id": id}).Set(set).UpdateOne() == nil {
		tools.GlobalResponse.ResponseCreated(c, param)
	} else {
		tools.GlobalResponse.ResponseServerError(c)
	}
}

// PostsDeleteCategory 删除分类
func PostsDeleteCategory(c *gin.Context) {
	common.AdminDeleteData(c, db.CollTag, "item_id")
}
