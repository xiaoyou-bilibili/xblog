// Package plugin
// @Description 这个包负责扫描插件目录
// @Author 小游
// @Date 2021/04/04
package plugin

import (
	"io/ioutil"
	"os"
	"xBlog/pkg/database"
)

// ScanPath 扫描路径，获取所有插件的路径
func ScanPath() []string {
	path := database.GetAppPath() + "plugins"
	var paths []string
	if data, err := ioutil.ReadDir(path); err == nil {
		for k := range data {
			dir := data[k]
			if dir.IsDir() {
				// 如果是目录的话，判断是否存在plugin.json和index.js文件
				if _, err := os.Stat(path + "/" + dir.Name() + "/index.js"); err == nil {
					if _, err := os.Stat(path + "/" + dir.Name() + "/plugins.json"); err == nil {
						paths = append(paths, path+"/"+dir.Name())
					}
				}
			}
		}
	}
	return paths
}
