// Package common
// @Title  用于文章板块的工具函数
// @Description  文章板块的专用函数，包括文章字数等内容
package common

import (
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"regexp"
	"strings"
	"time"
	"xBlog/internal/app/model"
	"xBlog/internal/db"
	"xBlog/pkg/database"
	"xBlog/tools"
)

// FilterContent 过滤掉标签内容
func FilterContent(s string) string {
	re, _ := regexp.Compile(`</?[^>]*>`)
	return re.ReplaceAllString(s, "")
}

// GetDec 获取文章摘要
func GetDec(content string, num int) string {
	//先去除前后空格和换行
	content = strings.Trim(content, " \n")
	//判断内容是不是信息框
	if data := tools.FindMatch("\\[.*?](.*?)\\[/.*?]", content); data != "" {
		return data
	} else {
		if len([]rune(content)) <= num {
			return content + "...."
		} else {
			return string([]rune(content)[:num]) + "...."
		}
	}
}

// PostPostReplace 文章内容替换
func PostPostReplace(content string) string {
	content = strings.Replace(content, "\r\n\r\n", "<br>", -1)
	content = strings.Replace(content, "\r", "\\r", -1)
	content = strings.Replace(content, "\n", "\\n", -1)
	//这里是把所有的 [][/]形式的内容都匹配出来
	r, _ := regexp.Compile("(\\[([a-z|0-9A-Z=”“\"\u4e00-\u9fa5\\s]+)]).*?\\[/.*?]")
	results := r.FindAllString(content, -1)
	//逐个替换
	for _, s := range results {
		//成功提示
		var rStr string
		/*背景栏*/
		if strings.Index(s, "[success]") != -1 {
			data := tools.FindMatch("\\[.*?](.*?)\\[/.*?]", s)
			rStr = "<div class=\"alert alert-success\">" + data + "</div>"
			content = strings.Replace(content, s, rStr, 1)
		}
		if strings.Index(s, "[info]") != -1 {
			data := tools.FindMatch("\\[.*?](.*?)\\[/.*?]", s)
			rStr = "<div class=\"alert alert-info\">" + data + "</div>"
			content = strings.Replace(content, s, rStr, 1)
		}
		if strings.Index(s, "[danger]") != -1 {
			data := tools.FindMatch("\\[.*?](.*?)\\[/.*?]", s)
			rStr = "<div class=\"alert alert-danger\">" + data + "</div>"
			content = strings.Replace(content, s, rStr, 1)
		}
		/*代码块*/
		if strings.Index(s, "[block]") != -1 {
			//先去掉里面的<pre>标签
			rStr = strings.Replace(s, "<pre>", "", -1)
			rStr = strings.Replace(rStr, "</pre>", "", -1)
			data := tools.FindMatch("\\[.*?](.*?)\\[/.*?]", rStr)
			rStr = "<pre class=\"block\"><code class=\"blockjs\">" + data + "</code></pre>"
			content = strings.Replace(content, s, rStr, 1)
		}
		/*代码高亮*/
		if strings.Index(s, "highlight") != -1 {
			rStr = strings.Replace(s, "<pre>", "", -1)
			rStr = strings.Replace(rStr, "</pre>", "", -1)
			//获取语言
			lan := tools.FindMatch("lanaguage=[\"|”|“](.*?)[\"|”|“]", rStr)
			//获取代码框的内容
			data := tools.FindMatch("\\[.*?](.*?)\\[/.*?]", rStr)
			data = strings.Replace(data, "<br>", "\n", -1)
			rStr = "<pre><code class=\"line-numbers language-" + lan + "\">" + data + "</code></pre>"
			content = strings.Replace(content, s, rStr, 1)
		}
		/*面板*/
		if strings.Index(s, "dangerbox") != -1 {
			title := tools.FindMatch("title=[\"|”|“](.*?)[\"|”|“]", s)
			data := tools.FindMatch("\\[.*?](.*?)\\[/.*?]", s)
			rStr = "<div class=\"panel panel-danger\"><div class=\"panel-heading\"><h3 class=\"panel-title\">" + title + "</h3></div><div class=\"panel-body\">" + data + "</div></div>"
			content = strings.Replace(content, s, rStr, 1)
		}
		if strings.Index(s, "infobox") != -1 {
			title := tools.FindMatch("title=[\"|”|“](.*?)[\"|”|“]", s)
			data := tools.FindMatch("\\[.*?](.*?)\\[/.*?]", s)
			rStr = "<div class=\"panel panel-info\"><div class=\"panel-heading\"><h3 class=\"panel-title\">" + title + "</h3></div><div class=\"panel-body\">" + data + "</div></div>"
			content = strings.Replace(content, s, rStr, 1)
		}
		if strings.Index(s, "successbox") != -1 {
			title := tools.FindMatch("title=[\"|”|“](.*?)[\"|”|“]", s)
			data := tools.FindMatch("\\[.*?](.*?)\\[/.*?]", s)
			rStr = "<div class=\"panel panel-success\"><div class=\"panel-heading\"><h3 class=\"panel-title\">" + title + "</h3></div><div class=\"panel-body\">" + data + "</div></div>"
			content = strings.Replace(content, s, rStr, 1)
		}
		/*折叠框*/
		if strings.Index(s, "collapse") != -1 {
			title := tools.FindMatch("title=[\"|”|“](.*?)[\"|”|“]", s)
			data := tools.FindMatch("\\[.*?](.*?)\\[/.*?]", s)
			rStr = "<div class=\"xControl\"><div  class=\"xHeading\"><div class=\"xIcon\"><svg aria-hidden=\"true\" width=\"14\" height=\"16\" viewBox=\"0 0 448 512\" focusable=\"false\" class=\"fa-icon\"><g><path d=\"M416 208c17.7 0 32 14.3 32 32v32c0 17.7-14.3 32-32 32h-144v144c0 17.7-14.3 32-32 32h-32c-17.7 0-32-14.3-32-32v-144h-144c-17.7 0-32-14.3-32-32v-32c0-17.7 14.3-32 32-32h144v-144c0-17.7 14.3-32 32-32h32c17.7 0 32 14.3 32 32v144h144z\"></path></g></svg></div><h5>" + title + "</h5></div><div class=\"xContent\"><div class=\"inner\">" + data + "</div></div></div>"
			content = strings.Replace(content, s, rStr, 1)
		}
		/*标题替换*/
		if strings.Index(s, "title") != -1 {
			data := tools.FindMatch("\\[.*?](.*?)\\[/.*?]", s)
			rStr = "<h2 class=\"title-h2\">" + data + "</h2>"
			content = strings.Replace(content, s, rStr, 1)
		}
		/*下载按钮*/
		if strings.Index(s, "ypbtn") != -1 {
			src := tools.FindMatch("\\[.*?](.*?)\\[/.*?]", s)
			rStr = "<button class=\"btn-primary btn-addon\" onclick=\"window.open('" + src + "')\"><svg aria-hidden=\"true\" width=\"16\" height=\"16\" viewBox=\"0 0 512 512\" focusable=\"false\" class=\"fa-icon\"><g><path d=\"M216 0h80c13.3 0 24 10.7 24 24v168h87.7c17.8 0 26.7 21.5 14.1 34.1l-152.1 152.2c-7.5 7.5-19.8 7.5-27.3 0l-152.3-152.2c-12.6-12.6-3.7-34.1 14.1-34.1h87.8v-168c0-13.3 10.7-24 24-24zM512 376v112c0 13.3-10.7 24-24 24h-464c-13.3 0-24-10.7-24-24v-112c0-13.3 10.7-24 24-24h146.7l49 49c20.1 20.1 52.5 20.1 72.6 0l49-49h146.7c13.3 0 24 10.7 24 24zM388 464c0-11-9-20-20-20s-20 9-20 20 9 20 20 20 20-9 20-20zM452 464c0-11-9-20-20-20s-20 9-20 20 9 20 20 20 20-9 20-20z\"></path></g></svg>点击下载</button>"
			content = strings.Replace(content, s, rStr, 1)
		}
		/*哔哩哔哩*/
		if strings.Index(s, "bilibili") != -1 {
			cid := tools.FindMatch("cid=[\"|”|“](.*?)[\"|”|“]", s)
			aid := tools.FindMatch("\\[.*?](.*?)\\[/.*?]", s)
			rStr = "<iframe src=\"//player.bilibili.com/player.html?aid=" + aid + "&cid=" + cid + "\" allowtransparency=\"true\" width=\"100%\" height=\"498\" scrolling=\"no\" frameborder=\"0\"></iframe>"
			content = strings.Replace(content, s, rStr, 1)
		}

		/*网易云*/
		if strings.Index(s, "music") != -1 && strings.Index(s, "wxmusic") == -1 {
			id := tools.FindMatch("\\[.*?](.*?)\\[/.*?]", s)
			autoplay := tools.FindMatch("autoplay=[\"|”|“](.*?)[\"|”|“]", s)
			rStr = "<iframe style=\"width:100%\" frameborder=\"no\" border=\"0\" marginwidth=\"0\" marginheight=\"0\" height=\"86\" src=\"https://music.163.com/outchain/player?type=2&id=" + id + "&auto=" + autoplay + "&height=66\"></iframe>"
			content = strings.Replace(content, s, rStr, 1)
		}

	}
	/*再把回车换行换回去*/
	content = strings.Replace(content, "\\r", "\r", -1)
	content = strings.Replace(content, "\\n", "\n", -1)
	return content
}

// GetPostDec 获取文章摘要
func GetPostDec(content string) string {
	//先去除前后空格和换行
	content = strings.Trim(content, " \n")
	//替换所有的 html标签
	content = FilterContent(content)
	//判断内容是不是信息框
	//判断文章字数是否足够
	if len([]rune(content)) < 300 {
		return content
	}
	return string([]rune(content)[:290]) + "...."
}

// PostGetUserBaseInfo 更据用户id获取头像和名字V2
func PostGetUserBaseInfo(id int) (model.UserBaseInfo, bool) {
	user := new(db.User)
	var info model.UserBaseInfo
	if database.NewDb(db.CollUser).SetFilter(bson.M{"user_id": id}).FindOne(user) == nil {
		info.Nickname = user.Nickname
		info.Avatar = user.Avatar
		return info, true
	}
	return info, false
}

// PostNoticeUser 自动通知用户博主文章更新
func PostNoticeUser(postId int) error {
	//读取所有订阅用户的邮箱
	users := new([]db.User)
	var emails []string
	if database.NewDb(db.CollUser).SetFilter(bson.M{"subscription": true}).FindMore(users) == nil {
		for _, v := range *users {
			emails = append(emails, v.Email)
		}
	}
	if len(emails) == 0 {
		return nil
	}
	//读取文章的数据
	article := new(db.Article)
	if database.NewDb(db.CollArticle).SetFilter(bson.M{"post_id": postId}).FindOne(article) == nil {
		site := db.GetSiteOptionString(db.KeySiteApiServer)
		href := site + "/archives/" + tools.Int2Str(postId)
		now := tools.Time2String(time.Now(), true)
		title := article.Title
		dec := GetDec(article.Content, 80)
		body := `
		<style>.qmbox img.wp-smiley{width:auto!important;height:auto!important;max-height:8em!important;margin-top:-4px;display:inline}</style>
		<div style="background:#ececec;width:100%;padding:50px 0;text-align:center">
			<div style="background:#fff;width:750px;text-align:left;position:relative;margin:0 auto;font-size:14px;line-height:1.5">
				<div style="zoom:1;padding:5px 40px;background:#518bcb; border-bottom:1px solid #467ec3;">
					<h3 style="color:#FFF;text-align:center">你关注的博主有文章更新啦( • ̀ω•́ )✧</h3>
				</div>
				<div style="padding:10px 40px 30px">
					<h2 style="font-size:18px;border-bottom:1px solid #ccc">` + title + `</h2>
					<p style="color:#a0a0a0;line-height:20px;font-size:15px;margin:20px 0">` + dec + `</p>
					<h4><a target="_blank" href="` + href + `" style="text-decoration:none;color:#518bcb;">点击查看原文</a></h4>
					<div style="font-size:13px;color:#a0a0a0;padding-top:px">该邮件由系统自动发出，如果不是您本人操作，请忽略此邮件。</div>
					<div class="qmSysSign" style="padding-top:5px;font-size:12px;color:#a0a0a0">
						<p style="color:#a0a0a0;line-height:18px;font-size:12px;margin:5px 0"><span style="border-bottom:1px dashed #ccc" t="5" times="">` + now + `</span></p>
					</div>
				</div>
			</div>
		</div>
	`
		// 我们直接遍历通知用户
		for _, v := range emails {
		_:
			tools.SendMail([]string{v}, "你关注的博主有文章更新啦!("+title+")", body)
		}
		return nil
	}
	return errors.New("获取文章信息失败")
}

// CommentChangeSmile 替换评论表情
func CommentChangeSmile(content string, width string) string {
	server := db.GetSiteOptionString(db.KeySiteApiServer)
	// todo 这个函数耗时较长，需要优化
	smile := new(map[string]model.Smile)
	bytes, err := ioutil.ReadFile(database.GetAppPath() + "configs/owo.json")
	if err != nil {
		return ""
	}
	err = json.Unmarshal(bytes, smile)
	if err != nil {
		return ""
	}
	for _, v := range *smile {
		if v.Type == "images" {
			for _, v1 := range v.Container {
				var img string
				if width != "" {
					img = `<img class="comment-img" style="width:` + width + `px;height:` + width + `px" src="` + server + "/" + v1.Icon + `"/>`
				} else {
					img = `<img class="comment-img" src="` + server + "/" + v1.Icon + `"/>`
				}
				content = strings.Replace(content, v1.Desc[1:len(v1.Desc)-1], img, -1)
			}
		}
	}
	return content
}
