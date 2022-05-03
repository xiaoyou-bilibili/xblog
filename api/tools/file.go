// Package tools @Description 文件读取相关的函数
// @Author 小游
// @Date 2021/04/10
package tools

import (
	"archive/zip"
	"fmt"
	"github.com/Unknwon/goconfig"
	"io"
	"os"
	"path/filepath"
	"strings"
	"xBlog/pkg/database"
)

// GetConfig 读取配置文件
func GetConfig(module string) map[string]string {
	config, err := goconfig.LoadConfigFile(database.GetAppPath() + "configs/app.ini") //加载配置文件
	if err != nil {
		return nil
	}
	glob, _ := config.GetSection(module) //读取全部mysql配置
	return glob
}

// SetConfig 设置配置文件
func SetConfig(module string, key string, value string) bool {
	var file = database.GetAppPath() + "configs/app.ini"
	config, err := goconfig.LoadConfigFile(file) //加载配置文件
	if err != nil {
		return false
	}
	if !config.SetValue(module, key, value) {
		if err := goconfig.SaveConfigFile(config, file); err == nil {
			return true
		} else {
			fmt.Println(err.Error())
		}
	}
	return false
}

// IsExist 文件夹是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		} else {
			return false
		}
	}
	return true
}

// CreatePath 创建文件夹
func CreatePath(filepath string) error {
	if !IsExist(filepath) {
		err := os.MkdirAll(filepath, os.ModePerm)
		return err
	}
	return nil
}

// GetConfigString 获取配置文件
func GetConfigString(module string, key string) string {
	config, err := goconfig.LoadConfigFile(database.GetAppPath() + "configs/app.ini") //加载配置文件
	if err != nil {
		return ""
	}
	glob, err := config.GetSection(module) //读取全部mysql配置
	if err == nil {
		return glob[key]
	}
	return ""
}

// DeleteFiles 删除整个目录
func DeleteFiles(path string) error {
	return os.RemoveAll(path)
}

// UnzipFile 解压文件
func UnzipFile(filename, target string) error {
	// 打开压缩文件
	reader, err := zip.OpenReader(filename)
	if err != nil {
		return err
	}
	defer reader.Close()
	// 首先获取第一个文件的根路径，然后删除
	if len(reader.File) > 0 && len(strings.Split(reader.File[0].Name, "/")) > 0 {
		path := target + strings.Split(reader.File[0].Name, "/")[0]
		_ = os.RemoveAll(path)
	}
	// 遍历压缩包里面所有的文件夹
	for _, file := range reader.File {
		// 设置我们的绝对值
		absPath := filepath.Join(target, file.Name)
		// 如果是目录，就创建目录
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(absPath, file.Mode()); err != nil {
				return err
			}
			// 因为是目录，跳过当前循环，因为后面都是文件的处理
			continue
		}
		// 打开文件
		fr, err := file.Open()
		if err != nil {
			return err
		}
		fw, err := os.OpenFile(absPath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		_, err = io.Copy(fw, fr)
		if err != nil {
			return err
		}
		// 关闭文件流
		_ = fw.Close()
		_ = fr.Close()
	}
	return nil
}
