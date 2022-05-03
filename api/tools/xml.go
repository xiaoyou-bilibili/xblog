// Package tools @Description siteMap生成工具
// @Author 小游
// @Date 2021/04/10
package tools

import (
	"encoding/xml"
	"io/ioutil"
	"time"
	"xBlog/internal/db"
	"xBlog/pkg/database"
)

// Url 网站地图的xml格式
type Url struct {
	Loc        string `xml:"loc"`
	LastMod    string `xml:"lastmod"`
	ChangeFreq string `xml:"changefreq"`
	Priority   string `xml:"priority"`
}

type XMLServers struct {
	XMLName xml.Name `xml:"urlset"`
	Version string   `xml:"version,attr"`
	Urls    []Url    `xml:"url"`
	Xml     string   `xml:"xmlns,attr"`
}

// XmlCreateSiteMap 创建网站地图
func XmlCreateSiteMap() bool {
	web := db.GetSiteOptionString(db.KeySiteApiServer)
	//指定版本
	v := &XMLServers{
		Version: "1",
		Xml:     "https://www.sitemaps.org/schemas/sitemap/0.9",
	}
	//先添加主站
	v.Urls = append(v.Urls, Url{
		Loc:        web,
		LastMod:    time.Now().Format("2006-01-02T15:04:05") + "+00:00",
		ChangeFreq: "daily",
		Priority:   "1.0",
	})
	var url string
	//获取所有的文章和文档
	article := new([]db.Article)
	if database.NewDb(db.CollArticle).FindMore(article) == nil {
		for _, item := range *article {
			//判断文章类型
			if item.PostType == "doc" {
				url = web + "/doc/" + Int2Str(item.PostID)
			} else {
				url = web + "/archives/" + Int2Str(item.PostID)
			}
			//解析为xml文件
			v.Urls = append(v.Urls, Url{
				Loc:        url,
				LastMod:    item.Modify.Format("2006-01-02T15:04:05") + "+00:00",
				ChangeFreq: "monthly",
				Priority:   "0.6",
			})
		}
	}
	// 读取文章分类
	tags := new([]db.Tag)
	if database.NewDb(db.CollTag).FindMore(tags) == nil {
		for _, item := range *tags {
			if item.ItemType == "category" {
				url = web + "/?category=" + item.Chain
			} else if item.ItemType == "tag" {
				url = web + "/?tag=" + item.Chain
			} else {
				// 如果啥都没有就跳过
				continue
			}
			// 判断文章类型
			v.Urls = append(v.Urls, Url{
				Loc:        web + "/?category=" + item.Chain,
				LastMod:    time.Now().Format("2006-01-02T15:04:05") + "+00:00",
				ChangeFreq: "monthly",
				Priority:   "0.6",
			})
		}
	}
	//生成xml文件
	out, err := xml.Marshal(v)
	if err == nil {
		// 自己写文件
		b := []byte(xml.Header + string(out))
		path := database.GetAppPath()
		_ = ioutil.WriteFile(path+"assets/sitemap.xml", b, 755)
	}
	return true
}
