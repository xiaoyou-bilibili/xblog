// Package admin
// @Description 文件处理相关
// @Author 小游
// @Date 2021/04/14
package admin

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"strings"
	"xBlog/internal/app/common"
	"xBlog/internal/app/model"
	"xBlog/internal/db"
	"xBlog/tools"
)

// FileUploadImageFile 图片上传(文件方式)
func FileUploadImageFile(c *gin.Context) {
	// 验证用户身份
	if _, err := common.AccessGetTokenV2(c); err != nil {
		tools.GlobalResponse.ResponseUnauthorized(c)
		return
	}
	// 获取上传字段
	field := c.PostForm("field")
	if field == "" {
		field = "file"
	}
	var file *multipart.FileHeader
	var err error

	//接收上传的文件
	file, err = c.FormFile(field)
	if err != nil {
		tools.GlobalResponse.ResponseServerError(c, "无法读取文件内容")
		return
	}
	// 判断上传方式
	open := db.GetSiteOptionBool(db.KeyImgBedUpload)
	ch := db.GetSiteOptionString(db.KeyImgBedAddr)
	ls := db.GetSiteOptionString(db.KeyImgLypBedAddr)

	var url = ""
	var filename = ""
	if !open {
		url, filename = common.NormalFileUpload(file, false, "")
	} else if ch != "" {
		url, filename = common.ChFileUpload(file, false, "", ch)
	} else if ls != "" {
		url, filename = common.LsyUpload(file, false, "")
	}
	if url != "" {
		tools.GlobalResponse.ResponseOk(c, model.AdminUploadImage{Url: url, Name: filename})
	} else {
		tools.GlobalResponse.ResponseServerError(c, "上传文件失败")
	}
}

// FileUploadImageBase64 以base64的方式上传图片
func FileUploadImageBase64(c *gin.Context) {
	// 验证用户身份
	if _, err := common.AccessGetTokenV2(c); err != nil {
		tools.GlobalResponse.ResponseUnauthorized(c)
		return
	}
	// 获取上传字段
	var base64Content string
	// 获取base64格式
	base64Content = c.PostForm("data")
	//替换掉前面多余的前缀
	base64Content = strings.Replace(base64Content, "data:image/png;base64,", "", 1)
	// 判断上传方式
	open := db.GetSiteOptionBool(db.KeyImgBedUpload)
	ch := db.GetSiteOptionString(db.KeyImgBedAddr)
	ls := db.GetSiteOptionString(db.KeyImgLypBedAddr)
	var url = ""
	var filename = ""
	if !open {
		url, filename = common.NormalFileUpload(nil, true, base64Content)
	} else if ch != "" {
		url, filename = common.ChFileUpload(nil, true, base64Content, ch)
	} else if ls != "" {
		url, filename = common.LsyUpload(nil, true, base64Content)
	}
	if url != "" {
		tools.GlobalResponse.ResponseOk(c, model.AdminUploadImage{Url: url, Name: filename})
	} else {
		tools.GlobalResponse.ResponseServerError(c, "上传文件失败")
	}
}
