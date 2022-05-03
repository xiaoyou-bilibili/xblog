// Package common
// @Description 和设置有关的模块
// @Author 小游
// @Date 2021/04/17
package common

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"xBlog/internal/app/model"
	"xBlog/internal/db"
	"xBlog/tools"
)

// SideInfo 侧边栏信息
var SideInfo = map[string]func() (error, model.SideInfo){}

// PluginPage 插件页面信息
var PluginPage = map[string]func() (error, model.SettingPlugin){}

// InitHeadMeta 初始化头部标签
func InitHeadMeta(title string, keyword string, description string, url string, image string) model.HeadMeta {
	return model.HeadMeta{
		Title:       title + "-" + db.GetSiteOptionString(db.KeySiteName),
		Keyword:     keyword,
		Description: description,
		Url:         db.GetSiteOptionString(db.KeySiteApiServer) + url,
		Image:       image,
		Icon:        db.GetSiteOptionString(db.KeySiteIcon),
	}
}

// InitSideInfo 初始化左右侧边栏
func InitSideInfo(index *model.SettingIndex) {
	index.LeftSide = []model.SideInfo{}
	index.RightSide = []model.SideInfo{}
	// 从设置中获取信息
	var left = db.GetSiteOption(db.KeySettingSideInfoLeft)
	var right = db.GetSiteOption(db.KeySettingSideInfoRight)
	// 获取左右侧边栏信息
	for _, v := range left.(primitive.A) {
		tmp := db.SettingAdminSideDetail{}
		if tools.Primitive2Struct(v, &tmp) == nil && SideInfo[tmp.Unique] != nil {
			if v, ok := SideInfo[tmp.Unique]; ok {
				if err, info := v(); err == nil {
					index.LeftSide = append(index.LeftSide, info)
				} else {
					fmt.Println(err)
				}
			}
		}
	}
	for _, v := range right.(primitive.A) {
		tmp := db.SettingAdminSideDetail{}
		if tools.Primitive2Struct(v, &tmp) == nil && SideInfo[tmp.Unique] != nil {
			if v, ok := SideInfo[tmp.Unique]; ok {
				if err, info := v(); err == nil {
					index.RightSide = append(index.RightSide, info)
				} else {
					fmt.Println(err)
				}
			}
		}
	}
}
