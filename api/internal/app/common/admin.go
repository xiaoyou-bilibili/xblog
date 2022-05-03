// Package common
// @Description 管理员板块工具类
// @Author 小游
// @Date 2021/04/14
package common

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/url"
	"xBlog/internal/app/model"
	"xBlog/internal/db"
	"xBlog/pkg/database"
	"xBlog/tools"
)

// AdminPostUpdateTag 文章更新分类和标签
func AdminPostUpdateTag(id int, category []int, tag []string) {
	// 先删除这个文章的所有分类和标签(这里不能复用，后面更新会复用pull导致删除) 这个pull会把数组里面的数据删除
	_ = database.NewDb(db.CollTag).SetFilter(bson.M{"posts": bson.M{"$in": []int{id}}}).Pull(bson.M{"posts": bson.M{"$in": []int{id}}}).UpdateMany()
	// 获取数据库对象
	Db := database.NewDb(db.CollTag)
	// 开始遍历分类
	for _, v := range category {
		// 因为addToSet可以自动判断数组里面是否有这个值，所以我们就不需要主动查找(这里我们设置upsert可以确保没有数据的时候自动添加)
		_ = Db.SetUpsert(true).SetFilter(bson.M{"item_id": v, "item_type": "category"}).AddToSet(bson.M{"posts": id}).UpdateOne()
	}
	// 开始遍历标签
	for _, v := range tag {
		// 避免赋值空标签
		if v == "" {
			break
		}
		// 因为标签并不是固定的，所以我们需要先判断这个标签是否存在
		if err := Db.SetUpsert(false).SetFilter(bson.M{"name": v}).FindOne(new(db.Tag)); err != nil && err.Error() == database.MongoNoResult {
			var data db.Tag
			data.ID = primitive.NewObjectID()
			data.Name = v
			data.Chain = url.QueryEscape(v)
			data.Posts = []int{id}
			data.ItemType = "tag"
			// 直接插入一个新的标签，不管是否成功
			_, _ = Db.InsertOneIncrease(data, "item_id")
		} else if err == nil {
			// 找到了就主动添加
			_ = Db.SetFilter(bson.M{"name": v}).AddToSet(bson.M{"posts": id}).UpdateOne()
		}
	}
}

// AdminDeleteData 数据库删除数据
func AdminDeleteData(c *gin.Context, dbName string, field string) {
	// 获取分类id
	id := c.Param("id")
	if id == "" {
		tools.GlobalResponse.ResponseBadRequest(c, "id格式错误")
		return
	}
	// 直接删除对应的文章
	if err := database.NewDb(dbName).SetFilter(bson.M{field: bson.M{"$in": tools.String2IntArray(id, ",")}}).DeleteMore(); err == nil {
		tools.GlobalResponse.ResponseNoContent(c)
	} else {
		tools.GlobalResponse.ResponseServerError(c)
	}
}

// AdminDeleteObjectData 数据库删除数据（objectID的版本）
func AdminDeleteObjectData(c *gin.Context, dbName string, field string) {
	// 获取分类id
	id := c.Param("id")
	if id == "" {
		tools.GlobalResponse.ResponseBadRequest(c, "id格式错误")
		return
	}
	// 直接删除对应的文章
	if database.NewDb(dbName).SetFilter(bson.M{field: bson.M{"$in": tools.String2ObjectIDArray(id, ",")}}).DeleteMore() == nil {
		tools.GlobalResponse.ResponseNoContent(c)
	} else {
		tools.GlobalResponse.ResponseServerError(c)
	}
}

// AdminReturnSetting 返回界面设置
func AdminReturnSetting(c *gin.Context, options []model.AdminSetting) {
	for k, v := range options {
		if v.Key != "" {
			// 从设置中查询数据
			options[k].Value = db.GetSiteOption(v.Key)
		}
	}
	tools.GlobalResponse.ResponseOk(c, options)
}

// AdminGetSetting 获取设置信息
func AdminGetSetting(options []model.AdminSetting) {
	for k, v := range options {
		if v.Key != "" {
			// 从设置中查询数据
			options[k].Value = db.GetSiteOption(v.Key)
		}
	}
}
