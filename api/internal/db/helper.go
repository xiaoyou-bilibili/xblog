// Package db @Description 一些常用的帮助类
// @Author 小游
// @Date 2021/04/10
package db

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os/exec"
	"runtime"
	"xBlog/pkg/database"
)

// 把所有的设置数据存储进内存中
var nowSiteSetting map[string]interface{}

// Regex 返回正则
func Regex(regex string) primitive.Regex {
	return primitive.Regex{Pattern: regex}
}

func primitive2Struct(data interface{}, result interface{}) error {
	bsonBytes, err := bson.Marshal(data)
	if err != nil {
		return err
	}
	err = bson.Unmarshal(bsonBytes, result)
	if err != nil {
		return err
	}
	return nil
}

// GetSiteOption 获取网站数据（内存存储版本,这个速度更快）
func GetSiteOption(settingName string) interface{} {
	// 先判断内存中是否有需要的数据
	if nowSiteSetting != nil {
		if value, ok := nowSiteSetting[settingName]; ok {
			return value
		}
	}
	// 再判断数据库中是否有该数据
	db := database.NewDb(CollSiteSetting)
	sites := *new([]SiteSetting)
	nowSiteSetting = make(map[string]interface{})
	//先查数据库
	if db.FindMore(&sites) == nil {
		//把所有数据遍历出来
		for _, v := range sites {
			nowSiteSetting[v.Key] = v.Value
		}
		//再次判断是否存在
		if value, ok := nowSiteSetting[settingName]; ok {
			return value
		}
	}
	//最后我们就判断默认值
	if value, ok := SiteOption[settingName]; ok {
		site := *new(SiteSetting)
		//自动插入这个值
		site.ID = primitive.NewObjectID()
		site.Key = settingName
		site.Value = value
		_ = db.InsertOne(site)
		return value
	}
	//没有这个设置
	return nil
}

// GetSiteOptionString 获取网站数据的string类型
func GetSiteOptionString(settingName string) string {
	res, ok := GetSiteOption(settingName).(string)
	if !ok {
		return ""
	} else {
		return res
	}
}

// GetSiteOptionInt 获取网站数据的int类型
func GetSiteOptionInt(settingName string) int {
	res, ok := GetSiteOption(settingName).(int32)
	if !ok {
		return 0
	} else {
		return int(res)
	}
}

// GetSiteOptionBool 获取网站数据的bool类型
func GetSiteOptionBool(settingName string) bool {
	res, ok := GetSiteOption(settingName).(bool)
	if !ok {
		return false
	} else {
		return res
	}
}

// SetSiteOption 保存站点设置
func SetSiteOption(settingName string, value interface{}) bool {
	// 直接把这个值插入数据库(不存在就新建)
	db := database.NewDb(CollSiteSetting).SetFilter(bson.M{"key": settingName}).SetUpsert(true)
	setting := new(SiteSetting)
	if db.Set(bson.M{"value": value}).UpdateOne() == nil {
		// 手动更新内容
		if nowSiteSetting != nil {
			// 为了避免直接设置和转换的时候出现问题，这里我们重新获取数据库的数据
			if db.FindOne(setting) == nil {
				nowSiteSetting[settingName] = setting.Value
			}
		}
		return true
	}
	return false
}

// DatabaseBackup 数据库备份
func DatabaseBackup() error {
	// 判断系统运行环境
	var err error
	if runtime.GOOS == "windows" {
		// 删除原来的文件，然后进行备份
		_, err = exec.Command("cmd", "/c", "rmdir /S/Q backup || mongodump -d xblog -o backup").Output()
	} else {
		// 删除原来的文件，然后进行备份
		_, err = exec.Command("/bin/sh", "-c", "rm -rf /xblog/backup/database/xblog && mongodump -d xblog -o /xblog/backup/database").Output()
	}
	return err
}

// DatabaseRestore 数据库恢复函数
func DatabaseRestore() error {
	// 判断系统运行环境
	var err error
	if runtime.GOOS == "windows" {
		// 删除原来的文件，然后进行备份
		_, err = exec.Command("cmd", "/c", "mongorestore --drop -d xblog backup\\xblog").Output()

	} else {
		// 删除原来的文件，然后进行备份
		_, err = exec.Command("/bin/sh", "-c", "mongorestore -d xblog /xblog/backup/database/xblog --drop").Output()
	}
	fmt.Println(err)
	return err
}
