// Package tools @Description 和图片处理有关
// @Author 小游
// @Date 2021/04/10
package tools

import (
	"math/rand"
	"path/filepath"
	"runtime"
	"strings"
	"xBlog/internal/db"
	"xBlog/pkg/database"
)

// ReplaceFile 替换文件名字，使用绝对路径，避免相对路径的问题
func ReplaceFile(file string) (filename string) {
	path := database.GetAppPath()
	server := db.GetSiteOptionString(db.KeySiteApiServer)
	if runtime.GOOS == "windows" {
		filename = strings.Replace(file, strings.Replace(path, "/", "\\", -1), "", -1)
	} else {
		filename = strings.Replace(file, path, "", -1)
	}
	filename = server + "/" + strings.Replace(filename, "\\", "/", -1)
	return
}

// GetRandomImg 生成随机文章图片
func GetRandomImg() string {
	//获取文件夹下所有图片
	path := database.GetAppPath()
	files, err := filepath.Glob(path + "/assets/images/background/*")
	if err != nil {
		return ""
	}
	filename := ""
	if len(files) > 0 {
		filename = ReplaceFile(files[rand.Intn(len(files))])
	} else {
		filename = ""
	}
	return filename
}

// GetRandomLoginImg 生成随机登录界面背景
func GetRandomLoginImg() string {
	//获取文件夹下所有图片
	path := database.GetAppPath()
	files, err := filepath.Glob(path + "/assets/images/login/*")
	if err != nil {
		return ""
	}
	filename := ""
	if len(files) > 0 {
		filename = ReplaceFile(files[rand.Intn(len(files))])
	} else {
		filename = ""
	}
	return filename
}

// GetRandomAvatar 生成随机头像
func GetRandomAvatar() string {
	//获取模板
	path := database.GetAppPath()
	files, err := filepath.Glob(path + "/assets/images/avatar/*")
	if err != nil {
		return ""
	}
	return ReplaceFile(files[rand.Intn(len(files))])
}
