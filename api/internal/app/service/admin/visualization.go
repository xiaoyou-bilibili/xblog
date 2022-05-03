// Package admin
// @Description 数据可视化板块
// @Author 小游
// @Date 2021/01/26
package admin

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"xBlog/internal/app/model"
	"xBlog/internal/db"
	"xBlog/pkg/database"
	"xBlog/tools"
)

// Visual 获取总体信息
func Visual(c *gin.Context) {
	var data model.AdminIndexTotal
	// 获取文章总数
	if num, err := database.NewDb(db.CollArticle).GetCount(); err == nil {
		data.Post = int(num)
	}
	// 获取用户总数
	if num, err := database.NewDb(db.CollUser).GetCount(); err == nil {
		data.User = int(num)
	}
	// 获取评论总数
	if num, err := database.NewDb(db.CollComment).GetCount(); err == nil {
		data.Comment = int(num)
	}
	// 获取浏览量
	if num, err := database.NewDb(db.CollArticle).GetFieldSum(nil, "view"); err == nil {
		data.View = int(num[0].Total)
	}
	tools.GlobalResponse.ResponseOk(c, data)
}

// VisualPostDistributed 获取不同类型文章的数目
func VisualPostDistributed(c *gin.Context) {
	var data model.AdminPostDistributed
	collection := database.NewDb(db.CollArticle)
	if num, err := collection.SetFilter(bson.M{"post_type": "post"}).GetCount(); err == nil {
		data.Post = int(num)
	}
	if num, err := collection.SetFilter(bson.M{"post_type": "doc"}).GetCount(); err == nil {
		data.Doc = int(num)
	}
	if num, err := collection.SetFilter(bson.M{"post_type": "diary"}).GetCount(); err == nil {
		data.Diary = int(num)
	}
	tools.GlobalResponse.ResponseOk(c, data)
}

// VisualV3GetPostDetail 获取每个文章的详细数据
func VisualV3GetPostDetail(c *gin.Context) {
	// 开始获取所有数据
	articles := new([]db.Article)
	if database.NewDb(db.CollArticle).SetFilter(bson.M{"post_type": "post", "status": "publish", "is_draft": false, "delete": false}).FindMore(articles) == nil {
		var data model.AdminVisualPostDetail
		// 开始遍历
		for _, v := range *articles {
			data.Title = append(data.Title, v.Title)
			data.Comment = append(data.Comment, v.Comment)
			data.Good = append(data.Good, v.Good)
			data.View = append(data.View, v.View)
		}
		tools.GlobalResponse.ResponseOk(c, data)
	} else {
		tools.GlobalResponse.ResponseServerError(c)
	}
}
