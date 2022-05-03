// Package module
// @Description 提供的一些实用工具包
// @Author 小游
// @Date 2021/04/08
package module

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dop251/goja"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"regexp"
	"strings"
	"time"
	"xBlog/internal/app/common"
	"xBlog/internal/app/model"
	"xBlog/internal/db"
	"xBlog/pkg/database"
	"xBlog/tools"
)

// RegisterTools 注册工具包
func RegisterTools(vm *goja.Runtime) *goja.Object {
	blog := vm.NewObject()
	// 打印日志
	_ = blog.Set("log", fmt.Println)
	_ = blog.Set("getSetting", ToolsGetSetting)
	// 设置设置信息
	_ = blog.Set("setSetting", db.SetSiteOption)
	_ = blog.Set("getAdminPluginSetting", ToolsGetSettingInfo)
	_ = blog.Set("setKey", ToolsSetKey)
	_ = blog.Set("getKey", ToolsGetKey)
	// 正则匹配
	_ = blog.Set("findMatch", tools.FindMatch)
	// 替换html标签，避免xss攻击
	_ = blog.Set("replaceXSS", tools.ReplaceXss)
	_ = blog.Set("verifyField", ToolsVerifyField)
	_ = blog.Set("verifyEmail", ToolsVerifyEmail)
	_ = blog.Set("changeCommentSmile", ToolsReplaceCommentSmile)
	// 获取随机头像
	_ = blog.Set("getRandomAvatar", tools.GetRandomAvatar)
	// 时间转字符串
	_ = blog.Set("time2String", tools.Time2String)
	// 字符串传时间
	_ = blog.Set("string2time", tools.Str2Time)
	// 字符串转int
	_ = blog.Set("str2int", tools.Str2Int)
	// 字符串转bool
	_ = blog.Set("str2bool", tools.Str2Bool)
	// 字符串批量转objectID
	_ = blog.Set("string2objetIdArray", tools.String2ObjectIDArray)
	// 字符串转objectID
	_ = blog.Set("str2objectId", tools.Str2ObjectID)
	// 字符串替换
	_ = blog.Set("strReplace", strings.Replace)
	_ = blog.Set("getBiliPersonInfo", ToolsGetBiliPersonInfo)
	_ = blog.Set("coverGBKToUTF8", tools.CoverGBKToUTF8)
	return blog
}

// ToolsGetSetting 获取网站设置信息
func ToolsGetSetting(key string) interface{} {
	row := db.GetSiteOption(key)
	if data, ok := row.(primitive.A); ok {
		var returnData []interface{}
		for _, v := range data {
			if data2, ok := v.(bson.D); ok {
				returnData = append(returnData, data2.Map())
			}
		}
		return returnData
	} else {
		return row
	}
}

// ToolsGetSettingInfo 获取设置信息
func ToolsGetSettingInfo(data []interface{}) []model.AdminSetting {
	res := make([]model.AdminSetting, len(data))
	for i, v := range data {
		tmp := v.(map[string]interface{})
		value := db.GetSiteOption(tools.Interface2String(tmp["key"]))
		if value == nil {
			value = tmp["default"]
		}
		res[i] = model.AdminSetting{
			Title: tools.Interface2String(tmp["title"]),
			Type:  tools.Interface2String(tmp["type"]),
			Key:   tools.Interface2String(tmp["key"]),
			Dec:   tools.Interface2String(tmp["dec"]),
			Value: value,
		}
	}
	return res
}

// ToolsSetKey 存储key
func ToolsSetKey(key string, value interface{}) interface{} {
	//把数据存起来
	if err := database.NewDb(db.CollOther).
		SetFilter(bson.M{"key": key}).
		SetUpsert(true).
		Set(bson.M{"key": key, "value": value, "create": time.Now()}).UpdateOne(); err == nil {
		return nil
	} else {
		return err.Error()
	}
}

// ToolsGetKey 获取key
func ToolsGetKey(key string) interface{} {
	data := new(db.Other)
	if database.NewDb(db.CollOther).SetFilter(bson.M{"key": key}).FindOne(data) == nil {
		// 判断类型primitive.D说明是对象，需要进行二次转换
		if reflect.TypeOf(data.Value) == reflect.TypeOf(primitive.D{}) {
			var result map[string]interface{}
			// 转换为map对象
			if tools.Primitive2Struct(data.Value, &result) == nil {
				return result
			} else {
				return nil
			}
		} else {
			return data.Value
		}
	} else {
		return nil
	}
}

// ToolsVerifyField 验证某个字段是否有效
func ToolsVerifyField(data interface{}) bool {
	if data == nil {
		return false
	}
	switch reflect.TypeOf(data).Name() {
	case "string":
		return data.(string) != ""
	}
	return true
}

// ToolsVerifyEmail 验证邮箱是否符合格式
func ToolsVerifyEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// ToolsReplaceCommentSmile 替换评论表情
func ToolsReplaceCommentSmile(content string) string {
	return common.CommentChangeSmile(content, "")
}

// ToolsGetBiliPersonInfo 获取B站个人信息
func ToolsGetBiliPersonInfo(uid, cookie string) (model.BiliPersonInfo, error) {
	//避免被B站识别为爬虫
	head := tools.HttpNewHead()
	head["origin"] = "https://space.bilibili.com"
	head["Referer"] = "https://space.bilibili.com"
	head["Cookie"] = cookie
	var info model.BiliPersonInfo
	info.Uid = uid
	//获取用户基本信息
	content, _ := tools.HttpGetHead("https://api.bilibili.com/x/space/acc/info?mid="+uid, head)
	var row interface{}
	if json.Unmarshal([]byte(content), &row) != nil {
		return model.BiliPersonInfo{}, errors.New("获取个人信息失败")
	}
	infos, ok := row.(map[string]interface{})["data"].(map[string]interface{})
	if ok {
		info.Nickname = tools.Interface2String(infos["name"])
		info.Avatar = tools.Interface2String(infos["face"])
		info.Level = tools.InterfaceFloat2Int(infos["level"])
		info.Sign = tools.Interface2String(infos["sign"])
		info.Sex = tools.Interface2String(infos["sex"])
		info.TopImage = strings.Replace(tools.Interface2String(infos["top_photo"]), "http://", "//", 1)
		if vip, ok := infos["vip"].(map[string]interface{})["status"]; ok {
			info.IsVip = tools.InterfaceFloat2Int(vip)
		}
		//替换头像https链接
		info.Avatar = tools.HttpReplaceHttp(info.Avatar)
	}
	//获取头像挂件，背景卡片等信息
	content, _ = tools.HttpGetHead("https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/space_history?host_uid="+uid, head)
	//解析用户数据
	if json.Unmarshal([]byte(content), &row) != nil {
		return model.BiliPersonInfo{}, errors.New("获取背景卡片失败")
	}
	infos, ok = row.(map[string]interface{})["data"].(map[string]interface{})
	if ok && tools.InterfaceFloat2Int(infos["has_more"]) == 1 { //发布过动态，可以获取卡片信息
		infos, ok = infos["cards"].([]interface{})[0].(map[string]interface{})["desc"].(map[string]interface{})["user_profile"].(map[string]interface{})
		if ok {
			info.Hang = tools.Interface2String(infos["pendant"].(map[string]interface{})["image"])
			//装饰卡片可能为空
			if infos["decorate_card"] != nil {
				info.Card = tools.Interface2String(infos["decorate_card"].(map[string]interface{})["card_url"])
			}
			//替换https链接
			info.Hang = tools.HttpReplaceHttp(info.Hang)
			info.Card = tools.HttpReplaceHttp(info.Card)
		}
	}
	//获取浏览量等内容
	content, _ = tools.HttpGetHead("https://api.bilibili.com/x/space/upstat?mid="+uid, head)
	if json.Unmarshal([]byte(content), &row) != nil {
		return model.BiliPersonInfo{}, errors.New("获取浏览量等数据失败")
	}
	infos, ok = row.(map[string]interface{})["data"].(map[string]interface{})
	if ok {
		if infos["archive"] != nil {
			info.View = tools.InterfaceFloat2Int(infos["archive"].(map[string]interface{})["view"])
		}
		if infos["article"] != nil {
			info.Article = tools.InterfaceFloat2Int(infos["article"].(map[string]interface{})["view"])
		}
		info.Good = tools.InterfaceFloat2Int(infos["likes"])
	}
	//获取播放数和粉丝数
	content, _ = tools.HttpGetHead("https://api.bilibili.com/x/relation/stat?vmid="+uid, head)
	if json.Unmarshal([]byte(content), &row) != nil {
		return model.BiliPersonInfo{}, errors.New("获取播放数粉丝数失败")
	}
	if infos, ok = row.(map[string]interface{})["data"].(map[string]interface{}); ok {
		info.Fans = tools.InterfaceFloat2Int(infos["follower"])
		info.Watch = tools.InterfaceFloat2Int(infos["following"])
	}
	return info, nil
}
