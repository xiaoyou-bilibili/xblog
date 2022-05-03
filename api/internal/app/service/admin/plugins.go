// Package admin
// @Description 插件管理相关的包
// @Author 小游
// @Date 2021/05/03
package admin

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"xBlog/internal/app/model"
	"xBlog/internal/db"
	"xBlog/pkg/database"
	"xBlog/tools"
)

// GlobalAllPlugins 全局所有的插件信息
var GlobalAllPlugins []model.GlobalInfo

// GlobalAllPluginsSetting 全局插件设置
var GlobalAllPluginsSetting map[string][]model.AdminPluginsSetting

// Reload 为了避免循环依赖这里这使用一个全局变量替代
var Reload func()

// PluginsGetPlugins 获取所有的插件
func PluginsGetPlugins(c *gin.Context) {
	var result []model.AdminPlugins
	// 获取插件更新信息
	var pluginsUnique []string
	// 遍历全局函数
	for i, v := range GlobalAllPlugins {
		pluginsUnique = append(pluginsUnique, v.Config.Unique)
		result = append(result, model.AdminPlugins{
			Id:          i + 1,
			Name:        v.Config.Name,
			Unique:      v.Config.Unique,
			Description: v.Config.Description,
			Author:      v.Config.Author,
			Site:        v.Config.Site,
			Version:     v.Config.Version,
			NewVersion:  v.Config.Version,
		})
	}
	// 获取更新信息
	//data := tools.HttpGet(server.XBLogServer + "/blog/get/plugins/update?plugins=" + strings.Join(pluginsUnique, ","))
	// 尝试进行解析
	//updateInfo := new(model.PluginsUpdate)
	//if json.Unmarshal([]byte(data), updateInfo) == nil {
	//	// 获取更新信息
	//	info := map[string]model.AdminUpdateInfo{}
	//	for _, v := range updateInfo.Data {
	//		info[v.Unique] = model.AdminUpdateInfo{Version: v.Version, DownloadUrl: v.DownloadUrl}
	//	}
	//	// 更新一下插件信息
	//	for i, v := range result {
	//		if v, ok := info[v.Unique]; ok {
	//			result[i].NewVersion = v.Version
	//			result[i].DownloadUrl = v.DownloadUrl
	//		}
	//	}
	//}
	//返回数据
	tools.GlobalResponse.ResponseOk(c, result)
}

// PluginsRemove 删除插件
func PluginsRemove(c *gin.Context) {
	id := tools.Str2Int(c.Param("id"))
	if id == 0 || id > len(GlobalAllPlugins) {
		tools.GlobalResponse.ResponseBadRequest(c, "插件不存在")
		return
	}
	// 找到插件所在的目录
	path := GlobalAllPlugins[id-1].Path
	// 删除所有目录
	if tools.DeleteFiles(path) == nil {
		// 重新加载插件
		Reload()
		tools.GlobalResponse.ResponseNoContent(c)
	} else {
		tools.GlobalResponse.ResponseServerError(c)
	}
}

// PluginsAdd 添加插件
func PluginsAdd(c *gin.Context) {
	var filename = database.GetAppPath() + "plugins/tmp.zip"
	//接收上传的文件
	if file1, err := c.FormFile("file"); err == nil {
		if f, err := file1.Open(); err == nil {
			data := make([]byte, file1.Size)
			if _, err := f.Read(data); err == nil {
				// 关闭文件
				_ = f.Close()
				// 新建一个文件
				if ioutil.WriteFile(filename, data, 755) == nil {
					// 解压zip文件
					if err := tools.UnzipFile(filename, database.GetAppPath()+"plugins/"); err == nil {
						// 重新加载插件
						Reload()
						tools.GlobalResponse.ResponseOk(c, "安装成功")
					} else {
						tools.GlobalResponse.ResponseServerError(c, "解压文件失败")
					}
					// 删除zip文件
					_ = os.Remove(filename)
					return
				}
			}
		}
	}
	tools.GlobalResponse.ResponseBadRequest(c, "没有获取到插件信息")
}

// PluginsReload 重新加载插件
func PluginsReload(c *gin.Context) {
	Reload()
	tools.GlobalResponse.ResponseOk(c, "重载成功！")
}

// PluginsSetting 获取插件设置
func PluginsSetting(c *gin.Context) {
	// 获取插件id
	id := tools.Str2Int(c.Param("id"))
	if id == 0 || id > len(GlobalAllPlugins) {
		tools.GlobalResponse.ResponseBadRequest(c, "id错误")
		return
	}
	if data, ok := GlobalAllPluginsSetting[GlobalAllPlugins[id-1].Unique]; ok {
		// 每次都获取最近的设置
		for i, v := range data {
			if v.Type == 1 {
				// 批量获取设置
				for j, v := range v.Setting {
					value := db.GetSiteOption(tools.Interface2String(v.Key))
					if value == nil {
						value = v.Default
					}
					data[i].Setting[j].Value = value
				}
			}
		}
		tools.GlobalResponse.ResponseOk(c, data)
	} else {
		tools.GlobalResponse.ResponseNotFound(c, "该插件目前还没有设置哦！")
	}
}
