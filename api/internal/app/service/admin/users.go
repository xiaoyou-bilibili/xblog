// Package admin
// @Description 管理员用户管理
// @Author 小游
// @Date 2021/04/14
package admin

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
	"xBlog/internal/app/common"
	"xBlog/internal/app/model"
	"xBlog/internal/db"
	"xBlog/pkg/database"
	"xBlog/tools"
)

// UsersGetUser 管理员获取用户列表
func UsersGetUser(c *gin.Context) {
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
		// 这里如果是用户id或者用户状态，我们需要额外转换为int类型
		if search == "user_id" || search == "status" || search == "identity" {
			filter[search] = tools.Str2Int(key)
		} else if search == "subscription" {
			filter[search] = tools.Str2Bool(key)
		} else {
			filter[search] = db.Regex(key)
		}
	}
	// 直接使用指针变量，避免内存占用
	users := new([]db.User)
	// 开始进行分页查询
	if total, page, err := database.NewDb(db.CollUser).
		SetFilter(filter).
		SetSort(bson.M{"_id": -1}).
		PaginateWithTotal(id, size, users); err == nil {
		// 循环遍历添加
		var userList model.AdminList
		userList.Total = page
		userList.TotalNum = total
		userList.Current = id
		// 遍历数据
		for _, user := range *users {
			var content model.AdminUserContent
			content.ID = user.UserID
			content.Username = user.Username
			content.Nickname = user.Nickname
			content.Email = user.Email
			content.Status = user.Status
			content.Subscription = user.Subscription
			content.Registered = tools.Time2String(user.Registered, true)
			content.LastLogin = tools.Time2String(user.LastTime, true)
			content.LoginIp = user.LastIp
			content.Identity = user.Identity
			// 用户列表添加用户
			userList.Contents = append(userList.Contents, content)
		}
		// 返回数据
		tools.GlobalResponse.ResponseOk(c, userList)
	} else {
		tools.GlobalResponse.ResponseServerError(c)
	}
}

// UsersUpdateUser 管理员更新用户数据
func UsersUpdateUser(c *gin.Context) {
	// 获取用户id
	id := c.Param("id")
	if id == "" {
		tools.GlobalResponse.ResponseBadRequest(c, "用户id格式错误")
		return
	}
	param := new(model.AdminUpdateUser)
	// 获取其他参数
	if tools.ValidatorParam(c, param) {
		return
	}
	// 判断内容对文章进行更新
	set := bson.M{}
	if param.Username != "" {
		set["username"] = param.Username
	}
	if param.Nickname != "" {
		set["nickname"] = param.Nickname
	}
	if param.Status != "" {
		set["status"] = tools.Str2Int(param.Status)
	}
	if param.Subscription != "" {
		set["subscription"] = tools.Str2Bool(param.Subscription)
	}
	if param.Email != "" {
		set["email"] = param.Email
	}
	if param.Identity != "" {
		set["identity"] = tools.Str2Int(param.Identity)
	}
	// 设置文章过滤
	filter := bson.M{}
	filter["user_id"] = bson.M{"$in": tools.String2IntArray(id, ",")}
	if database.NewDb(db.CollUser).Set(set).SetFilter(filter).UpdateMany() == nil {
		// 更新成功
		tools.GlobalResponse.ResponseCreated(c, param)
	} else {
		tools.GlobalResponse.ResponseServerError(c)
	}
}

// UsersDeleteUsers 管理员删除用户
func UsersDeleteUsers(c *gin.Context) {
	id := c.Param("id")
	if tools.IsInStringArray(strings.Split(id, ","), "1") {
		tools.GlobalResponse.ResponseServerError(c, "id为1的用户禁止删除")
	} else {
		// 调用删除函数
		common.AdminDeleteData(c, db.CollUser, "user_id")
	}
}
