// Package server @Description 工具类V3版本
// @Author 小游
// @Date 2021/01/18
package server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"time"
	"xBlog/internal/app/common"
	"xBlog/internal/app/model"
	"xBlog/internal/db"
	"xBlog/pkg/database"
	"xBlog/tools"
)

// ToolsGetPostSmile 获取表情的json文件
func ToolsGetPostSmile(c *gin.Context) {
	smile := new(map[string]model.Smile)
	bytes, err := ioutil.ReadFile(database.GetAppPath() + "configs/owo.json")
	if err != nil {
		tools.GlobalResponse.ResponseServerError(c, "读取表情数据失败")
		return
	}
	err = json.Unmarshal(bytes, smile)
	if err != nil {
		tools.GlobalResponse.ResponseServerError(c, "表情数据转换失败")
		return
	}
	//遍历表情数据并转换
	for k, v := range *smile {
		for k1, v1 := range v.Container {
			if v.Type == "images" {
				(*smile)[k].Container[k1].Icon = `<img title="` + v1.Desc[1:len(v1.Desc)-1] + `" class="x-owo-icon" src="` + db.GetSiteOptionString(db.KeySiteApiServer) + "/" + v1.Icon + `"/>`
			} else if v.Type == "emoticon" {
				(*smile)[k].Container[k1].Icon = `<span title="` + v1.Icon + `" class="x-owo-icon">` + v1.Icon + `</span>`
			}
		}
	}
	// 返回成功数据
	tools.GlobalResponse.ResponseOk(c, smile)
}

// ToolsGetOpenid 微信小程序获取openid
func ToolsGetOpenid(c *gin.Context) {
	code := c.Param("code")
	if tools.JudgeParams(code) {
		tools.GlobalResponse.ResponseBadRequest(c)
	}
	appId := db.GetSiteOptionString(db.KeyWechatMiniProgramId)
	secret := db.GetSiteOptionString(db.KeyWechatMiniProgramSecret)
	//发送请求
	result := tools.HttpGet("https://api.weixin.qq.com/sns/jscode2session?appid=" + appId + "&secret=" + secret + "&js_code=" + code + "&grant_type=authorization_code")
	//解析获取到的json数据
	var v interface{}
	err := json.Unmarshal([]byte(result), &v)
	if err != nil {
		tools.GlobalResponse.ResponseServerError(c, "获取用户数据失败")
	} else {
		tools.GlobalResponse.ResponseOk(c, v)
	}
}

// ToolsSubmitAdvice 用户提交意见反馈
func ToolsSubmitAdvice(c *gin.Context) {
	// 获取参数
	param := new(model.Advice)
	if tools.ValidatorParam(c, param) {
		return
	}
	//数据库中插入数据
	var other db.Other
	var advice model.UserAdvice
	advice.Content = param.Content
	advice.Concat = param.Concat
	other.ID = primitive.NewObjectID()
	other.Key = db.KeyAdvice
	other.Value = advice
	other.Create = time.Now()
	if database.NewDb(db.CollOther).InsertOne(other) == nil {
		body := "联系方式：" + param.Concat + "<br>反馈内容：" + param.Content
		_ = tools.SendMail([]string{db.GetSiteOptionString(db.KeySiteEmail)}, "意见反馈", body)
		tools.GlobalResponse.ResponseCreated(c, param)
	} else {
		tools.GlobalResponse.ResponseServerError(c, "反馈失败")
	}
}

// ToolsGetBiliInfo 根据BID快速获取个人信息
func ToolsGetBiliInfo(c *gin.Context) {
	uid := c.Param("uid")
	if uid == "" {
		tools.GlobalResponse.ResponseBadRequest(c)
		return
	}
	if info, err := common.BiliGetBaseInfo(uid); err == nil {
		tools.GlobalResponse.ResponseOk(c, info)
	} else {
		tools.GlobalResponse.ResponseServerError(c, "获取数据失败")
	}
}

// ToolsGetSiteMap 获取站点地图
func ToolsGetSiteMap(c *gin.Context) {
	//先获取博客的配置信息
	var sitemap = model.Sitemap{}
	sitemap.Post = []model.SiteMapContent{}
	sitemap.Doc = []model.SiteMapContent{}
	sitemap.Tag = []model.SiteMapContent{}
	sitemap.Category = []model.SiteMapContent{}
	server := db.GetSiteOptionString(db.KeySiteApiServer)
	sitemap.Map = server + "/assets/sitemap.xml"
	sitemap.Site = db.GetSiteOptionString(db.KeySiteName)
	//获取文章数据
	articles := new([]db.Article)
	if database.NewDb(db.CollArticle).SetFilter(bson.M{"is_draft": false, "status": bson.M{"$in": []string{"publish", "encrypt"}}}).FindMore(articles) == nil {
		//遍历文章和文档
		for _, v := range *articles {
			if v.PostType == "post" {
				sitemap.Post = append(sitemap.Post, model.SiteMapContent{Title: v.Title, Url: server + "/archives/" + tools.Int2Str(v.PostID)})
			} else if v.PostType == "doc" {
				sitemap.Doc = append(sitemap.Doc, model.SiteMapContent{Title: v.Title, Url: server + "/doc/" + tools.Int2Str(v.PostID)})
			}
		}
	}
	//遍历分类和列表
	tags := new([]db.Tag)
	if database.NewDb(db.CollTag).FindMore(tags) == nil {
		//遍历文章和分类
		for _, v := range *tags {
			if v.ItemType == "category" {
				sitemap.Category = append(sitemap.Category, model.SiteMapContent{Title: v.Name, Url: server + "/?category=" + v.Chain})
			} else if v.ItemType == "tag" {
				sitemap.Category = append(sitemap.Category, model.SiteMapContent{Title: v.Name, Url: server + "/?tag=" + v.Chain})
			}
		}
	}
	tools.GlobalResponse.ResponseOk(c, sitemap)
}

// ToolsGetAppVersion 获取APP的版本信息
func ToolsGetAppVersion(c *gin.Context) {
	app := new(model.AppDownload)
	app.Download = db.GetSiteOptionString(db.KeyAppDownloadDownload)
	app.Version = db.GetSiteOptionString(db.KeyAppDownloadVersion)
	app.Dec = db.GetSiteOptionString(db.KeyAppDownloadDescription)
	tools.GlobalResponse.ResponseOk(c, app)
}
