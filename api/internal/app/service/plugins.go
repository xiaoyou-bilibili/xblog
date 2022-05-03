// Package server @Description 自定义插件的服务
// @Author 小游
// @Date 2021/01/21
package server

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"xBlog/internal/app/common"
	"xBlog/internal/app/model"
	"xBlog/internal/db"
	"xBlog/pkg/database"
	"xBlog/tools"
)

// PluginsDiary 获取所有日记
func PluginsDiary(c *gin.Context) {
	id := tools.Str2Int(c.Query("page"))
	if id == 0 {
		id = 1
	}
	filter := c.Query("filter")
	article := new([]db.Article)
	if page, err := database.NewDb(db.CollArticle).SetSort(bson.M{"_id": -1}).SetFilter(bson.M{"status": bson.M{"$in": []string{"publish", "encrypt"}}, "post_type": "diary", "is_draft": false, "delete": false}).Paginate(id, db.GetSiteOptionInt(db.KeyDiaryListCount), article); err == nil {
		var diaryList model.List
		diaryList.Total = page
		diaryList.Current = id
		diaryList.Contents = []interface{}{}
		//遍历日记
		for _, v := range *article {
			var diary model.DiaryContent
			diary.DiaryID = v.PostID
			diary.Date = tools.Time2String(v.PostTime, true)
			diary.Comment = v.Comment
			diary.Good = v.Good
			if info, ok := common.PostGetUserBaseInfo(v.AuthorID); ok {
				diary.Nickname = info.Nickname
				diary.Avatar = info.Avatar
			}
			if v.Status == "encrypt" {
				diary.Encrypt = true
				diary.Content = "这篇日记已被加密"
			} else {
				if filter != "" {
					diary.Content = common.FilterContent(common.PostPostReplace(v.Content))
				} else {
					diary.Content = common.PostPostReplace(v.Content)
				}
			}
			diaryList.Contents = append(diaryList.Contents, diary)
		}
		tools.GlobalResponse.ResponseOk(c, diaryList)
	} else {
		tools.GlobalResponse.ResponseServerError(c)
	}
}

// PluginsChatroom 获取聊天室内容
func PluginsChatroom(c *gin.Context) {
	var user model.UserLogin
	user, _ = common.AccessGetTokenV2(c)
	// 默认为空数组
	var result []model.ChatDialog
	result = []model.ChatDialog{}
	collection := database.NewDb(db.CollChat).SetSort(bson.M{"_id": -1})
	// 先获取公共频道的数据
	var commonMessage db.ChatInfo
	if err := collection.SetFilter(bson.M{"target": 0}).FindOne(&commonMessage); err == nil {
		result = append(result, model.ChatDialog{Id: 0, Name: "公共频道", Avatar: db.GetSiteOptionString(db.KeyChatRoomAvatar), Message: commonMessage, Count: 0})
	}
	// 然后获取和我有关的消息,搜索的时候排除公共频道的数据,同时设置最多获取1000条数据，这样一些很久远的记录就不会显示了
	var meMessage []db.ChatInfo
	if user.UserID != 0 && collection.SetLimit(1000).SetFilter(bson.M{}).OR([]bson.M{{"user_id": user.UserID, "target": bson.M{"$ne": 0}}, {"target": user.UserID}}).FindMore(&meMessage) == nil {
		// 下面我们通过一个map来实现聊天室的定位
		rooms := make(map[int]db.ChatInfo)
		// 未读数据数目
		unRead := make(map[int]int)
		// 用户id
		userId := 0
		// 注意，map数据可能会乱，所以我们需要额外数组来保存用户id
		var order []int
		// 循环遍历获取数据
		for _, v := range meMessage {
			// 我发给别人,这个时候直接存储别人的id
			if v.UserId == user.UserID {
				// 不存在的时候才插入
				userId = v.Target
				if _, ok := rooms[userId]; !ok {
					order = append(order, userId)
					rooms[userId] = v
					unRead[userId] = 0
				}
				// 注意，我发送给别人是不需要已读的
			}
			// 别人发给我的时候
			if v.Target == user.UserID {
				// 不存在的时候才插入
				userId = v.UserId
				if _, ok := rooms[userId]; !ok {
					order = append(order, userId)
					rooms[userId] = v
					unRead[userId] = 0
				}
				// 判断数据是否已读
				if !v.Read {
					unRead[userId]++
				}
			}
		}
		// 经过上面循环后，我们就获取到了聊天室的数据,下面我们来依次添加
		var userInfo db.User
		userDb := database.NewDb(db.CollUser)
		for _, v := range order {
			// 首先获取用户信息
			if userDb.SetFilter(bson.M{"user_id": v}).FindOne(&userInfo) == nil {
				result = append(result, model.ChatDialog{Id: v, Name: userInfo.Nickname, Avatar: userInfo.Avatar, Message: rooms[v], Count: unRead[v]})
			}
		}
	}
	tools.GlobalResponse.ResponseOk(c, result)
}

// PluginsDocList 获取文档列表
func PluginsDocList(c *gin.Context) {
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
		tools.GlobalResponse.ResponseNotFound(c)
	}
}

// PluginsDocContent 获取文档内容
func PluginsDocContent(c *gin.Context) {
	id := tools.Str2Int(c.Param("id"))
	if id == 0 {
		tools.GlobalResponse.ResponseBadRequest(c, "文档id不符合格式")
		return
	}
	article := new(db.Article)
	if database.NewDb(db.CollArticle).SetFilter(bson.M{"post_id": id}).FindOne(article) == nil {
		var doc model.DocContent
		doc.ID = id
		doc.Title = article.Title
		doc.Content = common.PostPostReplace(article.Content)
		tools.GlobalResponse.ResponseOk(c, doc)
	} else {
		tools.GlobalResponse.ResponseNotFound(c)
	}
}
