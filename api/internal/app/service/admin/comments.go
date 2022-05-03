// Package admin
// @Description 管理员评论管理
// @Author 小游
// @Date 2021/04/14
package admin

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"xBlog/internal/app/common"
	"xBlog/internal/app/model"
	"xBlog/internal/db"
	"xBlog/pkg/database"
	"xBlog/tools"
)

// CommentsGetComments 管理员获取文章评论
func CommentsGetComments(c *gin.Context) {
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
	// 如果type和key不为空，说明设置了关键词搜索
	if search != "" && key != "" {
		// 因为可能涉及到评论状态，所以我们需要再进行判断
		if search == "agree" {
			filter[search] = tools.Str2Int(key)
		} else {
			filter[search] = db.Regex(key)
		}
	}
	// 直接使用指针变量，避免内存占用
	comments := new([]db.Comment)
	// 开始进行分页查询
	if total, page, err := database.NewDb(db.CollComment).
		SetFilter(filter).
		SetSort(bson.M{"_id": -1}).
		PaginateWithTotal(id, size, comments); err == nil {
		// 循环遍历添加
		var commentList model.AdminList
		commentList.TotalNum = total
		commentList.Total = page
		commentList.Current = id
		// 遍历数据
		for _, comment := range *comments {
			var content model.AdminCommentContent
			content.ID = comment.CommentID
			content.Author = comment.Nickname
			content.Content = comment.Content
			content.Date = tools.Time2String(comment.CommentTime, true)
			content.Email = comment.Email
			content.IP = comment.Ip
			content.Status = comment.Agree
			commentList.Contents = append(commentList.Contents, content)
		}
		// 返回数据
		tools.GlobalResponse.ResponseOk(c, commentList)
	} else {
		tools.GlobalResponse.ResponseServerError(c)
	}
}

// CommentsUpdateComments 管理员更新评论状态
func CommentsUpdateComments(c *gin.Context) {
	// 获取评论id
	id := c.Param("id")
	if id == "" {
		tools.GlobalResponse.ResponseBadRequest(c, "评论id格式错误")
		return
	}
	param := new(model.AdminUpdateComments)
	// 获取其他参数
	if tools.ValidatorParam(c, param) {
		return
	}
	// 判断内容对文章进行更新
	set := bson.M{}
	if param.Agree != "" {
		set["agree"] = tools.Str2Int(param.Agree)
	}
	// 设置评论过滤
	filter := bson.M{}
	filter["comment_id"] = bson.M{"$in": tools.String2IntArray(id, ",")}
	if database.NewDb(db.CollComment).Set(set).SetFilter(filter).UpdateMany() == nil {
		// 更新成功
		tools.GlobalResponse.ResponseCreated(c, param)
	} else {
		tools.GlobalResponse.ResponseServerError(c)
	}
}

// CommentsDeleteComments 管理员删除评论
func CommentsDeleteComments(c *gin.Context) {
	// 调用删除函数
	common.AdminDeleteData(c, db.CollComment, "comment_id")
}
