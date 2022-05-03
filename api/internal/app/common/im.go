// Package common
// @Description 聊天相关的接口
// @Author 小游
// @Date 2021/04/14
package common

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"xBlog/internal/app/model"
	"xBlog/internal/db"
	"xBlog/pkg/database"
	"xBlog/tools"
)

// WsQueue websocket队列的结构体
type WsQueue struct {
	// websocket对象
	Ws *websocket.Conn
	// 当前socket对应的用户id
	UserId int
}

// Queue 全局变量存储所有的请求
var Queue []WsQueue

// WsStruct 定义结构体保存数据
type WsStruct struct {
	// websocket对象
	Ws *websocket.Conn
	// 用户发送的内容
	Content string
}

// UserMessage 定义用户发送的数据
type UserMessage struct {
	UserId int         `json:"user_id"` // 用户id
	Token  string      `json:"token"`   // 用户token信息
	To     int         `json:"to"`      // 发送或者获取的对象
	Option string      `json:"option"`  //  用户进行的操作
	Data   interface{} `json:"data"`    //  携带的额外数据
}

// UserSend 用户发送的消息 发送的内容
type UserSend struct {
	MessageType int    `mapstructure:"message_type"` // 消息类型
	Content     string `mapstructure:"content"`      // 消息内容
}

// UserGet 用户读取消息 发送的内容
type UserGet struct {
	Date int64 `mapstructure:"date"` // 为了方便，使用unix时间戳
	Size int   `mapstructure:"size"` // 一次性获取多少条记录
}

// WsReturn websocket返回的数据类型
type WsReturn struct {
	Code    int         `json:"code"`    // 状态码 0表示错误 1 表示客户端向服务端发送请求通过 2 有新的消息  3 表示获取到了数据 4 表示获取到了历史数据
	Message string      `json:"message"` // 提示的信息
	Data    interface{} `json:"data"`    // 数据
}

// WsChain 定义一个结构体管道，用户保存结构体,给通道设置缓冲，避免过多的请求
var WsChain = make(chan WsStruct, 10)

// ImErrorProcess 错误处理函数，自动返回错误的信息
func ImErrorProcess(code int, message string, data ...interface{}) string {
	var returnData interface{}
	// 判断是否需要返回数据
	if len(data) > 0 {
		returnData = data[0]
	} else {
		returnData = nil
	}
	// 把数据转换为json字符串
	if data, err := json.Marshal(WsReturn{Code: code, Message: message, Data: returnData}); err == nil {
		return string(data)
	} else {
		// 解析出现异常时我们就返回自定义的
		return `{"result":false,"message":"解析异常","data":null}`
	}
}

// ChainProcess 管道处理函数，专门负责处理管道请求
func ChainProcess() {
	defer func() {
		// 这里recover相当于捕获异常，如果recover错误，说明出现了异常，我们这里强制重启
		if recover() != nil {
			fmt.Println(recover())
			// 如果发送错误那么我们尝试重启
			ChainProcess()
		}
	}()
	// 这里我们使用for循环来进行遍历处理
	for {
		// 从管道中获取数据
		ws := <-WsChain
		send := ImErrorProcess(0, "未知请求")
		data := UserMessage{}
		err := json.Unmarshal([]byte(ws.Content), &data)
		if err == nil {
			// 权限验证,id为0的用户不需要验证权限
			if data.UserId == 0 || AccessTokenStatusV2(data.UserId, data.Token, false) {
				// 判断用户的请求
				switch data.Option {
				case "send":
					var userSend UserSend
					// 因为我们解析到的是map类型数据，所以需要进行二次转换
					if mapstructure.Decode(data.Data, &userSend) == nil {
						send = CommitMsg(data.UserId, data.To, userSend)
					}
				case "get":
					var userGet UserGet
					// 因为我们解析到的是map类型数据，所以需要进行二次转换
					if mapstructure.Decode(data.Data, &userGet) == nil {
						send = GetMessage(data, userGet)
					}
				}
			} else {
				send = ImErrorProcess(0, "权限不足，请登录")
			}
		} else {
			send = ImErrorProcess(0, "发送的格式错误")
		}
		// 发送消息
		_ = ws.Ws.WriteMessage(websocket.TextMessage, []byte(send))
	}
}

// SendBorderCast 发送广播消息，用于通知有新消息到来
func SendBorderCast(userId int, target int, message string) {
	// 遍历队列
	for index := 0; index < len(Queue); index++ {
		//fmt.Println(index,"当前队列长度",len(Queue))
		ws := Queue[index]
		// 发送消息,如果target为0，或者用户id对应，那么我们就发送信息
		if target == 0 || target == ws.UserId || userId == ws.UserId {
			err := ws.Ws.WriteMessage(websocket.TextMessage, []byte(message))
			// 如果发送失败我们就删除指定位置的连接
			if err != nil {
				QuaProcess(&Queue, index)
				// 删除后index需要--，避免漏发
				index--
			}
		}
	}
}

// QuaProcess 删除指定位置的连接
func QuaProcess(Queue *[]WsQueue, index int) {
	defer func() {
		if recover() != nil {
			// 这里我们不处理，这样就不会抛出异常
		}
	}()
	// 如果index+1小于Queue的大小,说明index不是最后一个
	if index+1 < len(*Queue) {
		*Queue = append((*Queue)[:index], (*Queue)[index+1:]...)
	} else if index+1 == len(*Queue) {
		// index恰好为最后一个
		*Queue = (*Queue)[:index]
	}

}

// GetMessage 获取消息
func GetMessage(send UserMessage, data UserGet) string {
	var code = 3
	// 先判断是获取历史消息还是最新的消息
	// 我们发送消息时一个双向过程，包括你发给我和我发给你
	var or = []bson.M{
		{"user_id": send.To, "target": send.UserId},
		{"user_id": send.UserId, "target": send.To},
	}
	// 如果是公共聊天室,那么就只需要获取所有target为0就可以了
	if send.To == 0 {
		or = []bson.M{{"target": send.To}}
	}
	// 是否是查询历史记录
	if data.Date != 0 {
		// 小于("$lt")、小于等于("$lte")、大于("$gt")、大于等于("$gte")、不等于("$ne")
		or[0]["date"] = bson.M{"$lt": data.Date}
		if send.To != 0 {
			or[1]["date"] = bson.M{"$lt": data.Date}
		}
		code = 4
	}
	var response []model.ChatMessage
	// 初始化为空数组
	response = []model.ChatMessage{}
	// 我们获取最新的消息
	var result []db.ChatInfo
	// 获取数据
	collection := database.NewDb(db.CollChat)
	if collection.
		SetLimit(int64(data.Size)).
		SetSort(bson.M{"_id": -1}).
		OR(or).
		FindMore(&result) == nil {
		// 因为顺序反了，所以我们还需要调整让数组倒过来
		ArrayReverse(&result)
		// todo 可能是因为驱动限制，没找到update的时候进行limit操作，所以只能换一种方法了
		var ids []primitive.ObjectID
		var userInfo db.User
		userDb := database.NewDb(db.CollUser)
		for _, v := range result {
			var chatMessage model.ChatMessage
			// 把我们的消息id放入，当消息时自己发送的时候，不能标记已读
			if userInfo.UserID == send.UserId {
				ids = append(ids, v.ID)
			}
			chatMessage.ID = v.ID.Hex()
			chatMessage.UserId = v.UserId
			chatMessage.Date = v.Date
			chatMessage.Target = v.Target
			chatMessage.MessageType = v.MessageType
			chatMessage.Read = v.Read
			// 获取用户信息
			if userDb.SetFilter(bson.M{"user_id": v.UserId}).FindOne(&userInfo) == nil {
				chatMessage.Avatar = userInfo.Avatar
				chatMessage.Nickname = userInfo.Nickname
			}
			// 这里我们替换一下表情
			chatMessage.Content = CommentChangeSmile(v.Content, "100")
			response = append(response, chatMessage)
		}
		// 获取数据的同时把数据置为已读
		_ = collection.SetFilter(bson.M{"_id": bson.M{"$in": ids}}).Set(bson.M{"read": true}).UpdateMany()
		//fmt.Println(response)
		return ImErrorProcess(code, "获取数据成功", response)
	}
	return ImErrorProcess(0, "获取数据失败")
}

// CommitMsg 提交数据
func CommitMsg(id int, to int, message UserSend) string {
	//获取一些数据
	var user db.User
	if database.NewDb(db.CollUser).SetFilter(bson.M{"user_id": id}).FindOne(&user) == nil {
		//避免xss攻击
		message.Content = tools.ReplaceXss(message.Content)
		// 插入数据
		var info db.ChatInfo
		info.ID = primitive.NewObjectID()
		info.UserId = user.UserID
		info.Content = message.Content
		// 这里我们使用毫秒来作为时间戳
		// https://blog.csdn.net/mirage003/article/details/80822608
		info.Date = time.Now().UnixNano() / 1e6
		info.Target = to
		info.MessageType = message.MessageType
		// 如果是0那么我们就直接置为已读
		if to == 0 {
			info.Read = true
		}
		// 数据库插入数据
		if database.NewDb(db.CollChat).InsertOne(info) == nil {
			// 查询用户信息
			var user db.User
			var message model.ChatMessage
			if database.NewDb(db.CollUser).SetFilter(bson.M{"user_id": info.UserId}).FindOne(&user) == nil {
				message.Avatar = user.Avatar
				message.Nickname = user.Nickname
			}
			message.ID = info.ID.Hex()
			message.UserId = info.UserId
			message.Content = CommentChangeSmile(info.Content, "100")
			message.Date = info.Date
			message.Target = info.Target
			message.MessageType = info.MessageType
			message.Read = info.Read
			// 这里我们发送广播来进行通知
			SendBorderCast(user.UserID, to, ImErrorProcess(2, "有新消息", message))
			// 返回正确数据
			return ImErrorProcess(1, "提交成功！")
		}
	}
	return ImErrorProcess(0, "提交失败！")
}

// ArrayReverse 数组倒序
func ArrayReverse(arr *[]db.ChatInfo) {
	var temp db.ChatInfo
	length := len(*arr)
	// 这里我们直接让第一个和最后一个交换，直接换到中间
	for i := 0; i < length/2; i++ {
		temp = (*arr)[i]
		(*arr)[i] = (*arr)[length-1-i]
		(*arr)[length-1-i] = temp
	}
}
