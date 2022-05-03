// Package tools
// @Description 微信小程序相关的工具类
// @Author 小游
// @Date 2021/05/26
package tools

import (
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"xBlog/internal/app/model"
	"xBlog/internal/db"
	"xBlog/pkg/database"
)

// WechatGetAccessToken 微信小程序获取access_token
func WechatGetAccessToken() (model.WechatAccessToken, error) {
	appID := db.GetSiteOptionString(db.KeyWechatMiniProgramId)
	secret := db.GetSiteOptionString(db.KeyWechatMiniProgramSecret)
	access := model.WechatAccessToken{}
	// appid获取secret为空就直接退出
	if appID == "" || secret == "" {
		return access, nil
	}
	// 发送请求获取accessToken
	result := HttpGet("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + appID + "&secret=" + secret)
	if json.Unmarshal([]byte(result), &access) == nil {
		// 更新到数据库（这里我们可以实现数据不存在的时候自动插入，存在自动更新）
		if database.NewDb(db.CollOther).
			SetUpsert(true).
			SetFilter(bson.M{"key": db.WechatAccessToken}).
			Set(bson.M{"value": access, "create": time.Now()}).
			UpdateOne() == nil {
			return access, nil
		}
		return access, errors.New("保存到数据库失败")
	}
	return access, errors.New("获取token失败")
}
