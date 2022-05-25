// Package module
// @Description
// @Author 小游
// @Date 2021/04/17
package module

import (
	"bytes"
	"fmt"
	"github.com/dop251/goja"
	"github.com/gin-gonic/gin"
	"html/template"
	"xBlog/internal/app"
	"xBlog/internal/app/common"
	"xBlog/internal/app/model"
	"xBlog/internal/app/service/admin"
	"xBlog/internal/db"
	"xBlog/tools"
)

// 侧边栏信息
type widgetInfo struct {
	file    string
	content func() interface{}
	path    string
}

// 全局CSS变量
var globalCSS = map[string]string{
	"element": "https://unpkg.com/element-ui@2.15.1/lib/theme-chalk/index.css",
}

// 全局JavaScript变量
var globalScript = map[string]string{
	"xiaoyou": "/static/js/xiaoyou.js",
	"element": "https://unpkg.com/element-ui/lib/index.js",
	"vue":     "/static/js/vue.js",
	"jquery":  "https://cdn.bootcdn.net/ajax/libs/jquery/3.6.0/jquery.min.js",
}

// 全局html插件对象
var globalWidget = map[string]widgetInfo{}

// InitWidget 注册一个路由用于读取插件信息
func InitWidget() {
	// 设置加载函数
	app.GinEngine.SetFuncMap(template.FuncMap{"unescaped": unescaped})
	app.GinEngine.GET("/plugins/widget", func(c *gin.Context) {
		// 获取插件的名字
		name := c.Query("name")
		widget := globalWidget[name]
		// 加载html文件
		app.GinEngine.LoadHTMLFiles(widget.path)
		// 输出内容
		c.HTML(200, widget.file, widget.content())
	})
}

// InitPage 页面信息调试
func InitPage() {
	app.GinEngine.GET("/plugins/page/:name", func(c *gin.Context) {
		// 获取插件的名字
		name := c.Param("name")
		url := c.Query("url")
		if info, ok := common.PluginPage[name+url]; ok {
			// 执行函数
			if err, data := info(); err == nil {
				css := ""
				// 加载CSS和JS文件
				for _, v := range data.CSS {
					css += "<link rel='stylesheet' href='" + v + "' />"
				}
				for _, v := range data.Script {
					css += "<script src='" + v + "'></script>"
				}
				tools.GlobalResponse.ResponseHtml(c, css+data.Content)
			} else {
				tools.GlobalResponse.ResponseServerError(c, "插件报错:"+err.Error())
			}
		} else {
			tools.GlobalResponse.ResponseNotFound(c, "页面不存在")
		}
	})
}

// RegisterWidget 注册页面部件包
func RegisterWidget(vm *goja.Runtime) *goja.Object {
	blog := vm.NewObject()
	_ = blog.Set("addSide", WidgetAddSide)
	_ = blog.Set("addPage", WidgetAddPage)
	_ = blog.Set("addSetting", WidgetAddSetting)
	return blog
}

// HTML标签不转义
func unescaped(x string) interface{} { return template.HTML(x) }

// 加载html内容
func loadTemplate(path string, file string, call func() interface{}) (error, string) {
	t := template.New("")
	// 加载函数
	t.Funcs(template.FuncMap{"unescaped": unescaped})
	// 初始化模板引擎
	t, err := t.ParseFiles(path + "/" + file)
	// 判断是否加载错误
	if err != nil {
		//fmt.Println(err)
		return err, ""
	}
	// 因为我们需要io.Writer对象，我们这里使用bytes.Buffer替代
	// https://stackoverflow.com/questions/13765797/the-best-way-to-get-a-string-from-a-writer
	buf := new(bytes.Buffer)
	// 解析模板
	if t.ExecuteTemplate(buf, file, call()) != nil {
		return err, ""
	}
	// 返回侧边栏数据
	return nil, fmt.Sprint(buf.String())
}

// WidgetAddSide 添加侧边栏卡片 （是否在左边，卡片标题，侧边栏文件，模板渲染的内容，是否开启debug模式）
func WidgetAddSide(title string, file string, call func() interface{}, debug bool) {
	// 是否开启debug模式
	if debug {
		globalWidget[Global.Unique] = widgetInfo{file: file, content: call, path: Global.Path + "/" + file}
	}
	// 因为路径会变化，所以我们临时存储数据
	path := Global.Path
	//fmt.Println(path)
	// 注册一个插件函数
	common.SideInfo[Global.Unique] = func() (error, model.SideInfo) {
		if err, data := loadTemplate(path, file, call); err == nil {
			return nil, model.SideInfo{Title: title, Html: data}
		} else {
			return err, model.SideInfo{}
		}
	}
}

// WidgetAddPage 添加页面
func WidgetAddPage(
	option map[string]interface{},
	call func() interface{}, // 数据回调
) {
	// 为了避免用户某些地方不填造成程序错误，这里我手动处理一下
	if option["headMeta"] == nil {
		option["headMeta"] = map[string]interface{}{}
	}
	if option["css"] == nil {
		option["css"] = []interface{}{}
	}
	if option["script"] == nil {
		option["script"] = []interface{}{}
	}
	// 因为路径会变化，所以我们临时存储数据
	path := Global.Path
	// 初始化相关信息
	info := model.SettingPlugin{}
	// 初始化headMeta
	head := option["headMeta"].(map[string]interface{})
	info.HeadMeta = model.HeadMeta{
		Title:       db.GetSiteOptionString(db.KeySiteName) + "-" + tools.Interface2String(head["title"]),
		Keyword:     tools.Interface2String(head["keyword"]),
		Description: tools.Interface2String(head["description"]),
		Image:       tools.Interface2String(head["image"]),
		Icon:        db.GetSiteOptionString(db.KeySiteIcon),
	}
	info.Url = tools.Interface2String(option["url"])
	info.Background = tools.Interface2String(option["background"])
	// 这里遍历CSS和JS
	for _, v := range option["css"].([]interface{}) {
		if c, ok := globalCSS[v.(string)]; ok {
			info.CSS = append(info.CSS, c)
		} else {
			info.CSS = append(info.CSS, v.(string))
		}
	}
	for _, v := range option["script"].([]interface{}) {
		if s, ok := globalScript[v.(string)]; ok {
			info.Script = append(info.Script, s)
		} else {
			info.Script = append(info.Script, v.(string))
		}
	}
	info.Full = tools.Interface2Bool(option["full"])
	info.Side = tools.Interface2Bool(option["side"])
	// 注册一个插件函数
	common.PluginPage[Global.Unique+tools.Interface2String(option["url"])] = func() (error, model.SettingPlugin) {
		// 加载页面样式
		if err, data := loadTemplate(path, tools.Interface2String(option["file"]), call); err == nil {
			info.Content = data
			return nil, info
		} else {
			return err, model.SettingPlugin{}
		}
	}
}

// WidgetAddSetting 添加设置界面
func WidgetAddSetting(title string, optionType int, option interface{}) {
	setting := model.AdminPluginsSetting{
		Name: title,
		Type: optionType,
	}
	if settings, ok := option.([]interface{}); ok {
		// 我们需要把option转换为[]model.AdminSetting的形式
		data := make([]model.AdminSetting, len(settings))
		for i, tmp := range settings {
			if setting, ok := tmp.(map[string]interface{}); ok {
				// 判断一下数据库中有没有值，没有的话就新建一条记录
				if db.GetSiteOption(tools.Interface2String(setting["key"])) == nil {
					db.SetSiteOption(tools.Interface2String(setting["key"]), setting["default"])
				}
				data[i] = model.AdminSetting{
					Title:   tools.Interface2String(setting["title"]),
					Type:    tools.Interface2String(setting["type"]),
					Key:     tools.Interface2String(setting["key"]),
					Dec:     tools.Interface2String(setting["dec"]),
					Default: setting["default"],
				}
			}
		}
		setting.Setting = data
	} else if optionType == 2 {
		setting.Extra = tools.Interface2String(option)
	}
	if setting.Name != "" {
		admin.GlobalAllPluginsSetting[Global.Unique] = append(admin.GlobalAllPluginsSetting[Global.Unique], setting)
	}
}
