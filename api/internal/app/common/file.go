// Package common 文件上传相关
// @Description
// @Author 小游
// @Date 2021/04/14
package common

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
	"xBlog/internal/db"
	"xBlog/pkg/database"
	"xBlog/tools"
)

// File2Base64 file转base64 前面是base64的内容 后面是文件的后缀
func File2Base64(file *multipart.FileHeader) (string, string) {
	//使用图床进行上传
	files, err := file.Open()
	defer func(files multipart.File) {
		_ = files.Close()
	}(files)
	if err != nil {
		return "", ""
	}
	buffer := make([]byte, file.Size)
	_, err = files.Read(buffer)
	if err != nil {
		return "", ""
	}
	source := base64.StdEncoding.EncodeToString(buffer)
	suffix := tools.FindMatch("(\\.[a-z]+)$", file.Filename)
	//如果没有就自己设置一个
	if suffix == "" {
		suffix = ".png"
	}
	return source, suffix
}

// NormalFileUpload 普通方式的图片上传
// 前面一个返回是文件url后面那个是文件名
func NormalFileUpload(file *multipart.FileHeader, IsBase64 bool, base64Content string) (string, string) {

	year := database.GetAppPath() + "upload/images/" + strconv.Itoa(time.Now().Year())
	month := strconv.Itoa(int(time.Now().Month()))
	filename := time.Now().Format("150405") + tools.GetRandomNum(5)
	suffix := ".png"
	//判断是否是base64编码
	if !IsBase64 {
		base64Content, suffix = File2Base64(file)
	}
	filename += suffix
	filePath := year + "/" + month + "/" + filename
	err := tools.CreatePath(year + "/" + month)
	if err == nil && base64Content != "" {
		bas, err := base64.StdEncoding.DecodeString(base64Content)
		if err != nil {
			return "", ""
		}
		//根据base64转换为文件
		if ioutil.WriteFile(filePath, bas, 755) == nil {
			// 替换一下路径。要不然会报错
			return tools.ReplaceFile(filePath), filename
		}
	}
	return "", ""
}

// ChFileUpload cheered 图床上传
func ChFileUpload(file *multipart.FileHeader, IsBase64 bool, base64Content string, addr string) (string, string) {
	if !IsBase64 {
		base64Content, _ = File2Base64(file)
	}
	urlValues := url.Values{}
	urlValues.Add("source", base64Content)
	resp, err := http.PostForm(addr, urlValues)
	if err != nil {
		return "", ""
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", ""
	}
	//通过正则来获取图片url
	re, _ := regexp.Compile(`"url":"(.*?)"`)
	urls := re.FindStringSubmatch(string(body))
	if urls != nil {
		return strings.Replace(urls[1], "\\/", "/", -1), "1" + tools.FindMatch("(\\.[a-z]+)$", urls[1])
	}
	return "", ""
}

//===================== lsky pro相关
type returnData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Time int         `json:"time"`
}

func LsyUpdateToken() string {
	user := db.GetSiteOptionString(db.KeyImgLypBedUser)
	password := db.GetSiteOptionString(db.KeyImgLypBedPassword)
	tokenUpload := db.GetSiteOptionString(db.KeyImgLypBedAddr)
	// 获取token信息
	row := tools.HttpPost(tokenUpload+"/api/token", map[string]string{"email": user, "password": password})
	data := new(returnData)
	if json.Unmarshal([]byte(row), data) == nil {
		// 解析map
		if data.Code != 200 {
			return ""
		}
		if token1, ok := data.Data.(map[string]interface{}); ok {
			if token2, ok := token1["token"].(string); ok {
				db.SetSiteOption(db.KeyImgLypBedToken, token2)
				return token2
			}
		}
	}
	return ""
}

// LsyUploadImage liskov的上传处理 前面是是否上传成功，后面两个是文件名和上传地址
func LsyUploadImage(file *multipart.FileHeader, local bool) (bool, string, string) {
	upload := db.GetSiteOptionString(db.KeyImgLypBedAddr)
	token := map[string]string{"token": db.GetSiteOptionString(db.KeyImgLypBedToken)}
	row := tools.HttpPostFile(upload+"/api/upload", file, "image", token, local)
	data := new(returnData)
	if json.Unmarshal([]byte(row), data) == nil {
		// 解析map
		if data.Code != 200 {
			return false, "", ""
		}
		if m, ok := data.Data.(map[string]interface{}); ok {
			if urls, ok := m["url"].(string); ok {
				if name, ok := m["name"].(string); ok {
					return true, name, urls
				}
			}
		}
	}
	return false, "", ""
}

// LsyUpload Lsky Pro 图床上传
func LsyUpload(file *multipart.FileHeader, IsBase64 bool, base64Content string) (string, string) {
	// 先进行一次上传
	var urls string
	var name string
	var ok bool
	// 判断是否为图床上传
	if IsBase64 {
		bas, err := base64.StdEncoding.DecodeString(base64Content)
		if err != nil {
			return "", ""
		}
		// 根据base64转换为文件
		if ioutil.WriteFile("temp.png", bas, 755) != nil {
			// 读取文件并上传
			return "", ""
		}
	}
	if ok, name, urls = LsyUploadImage(file, IsBase64); !ok {
		// 获取token重新试一次
		if LsyUpdateToken() != "" {
			if ok, name, urls = LsyUploadImage(file, IsBase64); ok {
				return urls, name
			}
		}
		return "", ""
	} else {
		return urls, name
	}
}
