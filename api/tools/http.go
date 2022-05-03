// Package tools @Description http请求有关的包
// @Author 小游
// @Date 2021/04/10
package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

type WriteCounter struct {
	Total uint64
}

func (w WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	w.Total += uint64(n)
	return n, nil
}

// HttpGet 发送get请求
func HttpGet(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(s)
}

// HttpPost 发送post请求
func HttpPost(url string, data map[string]string) string {
	bytesData, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	resp, err := http.Post(url, "application/json;charset=utf-8", bytes.NewBuffer(bytesData))
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(s)
}

// HttpPostByte 发送post请求,返回字节流类型
func HttpPostByte(url string, data map[string]string) []byte {
	if bytesData, err := json.Marshal(data); err == nil {
		if resp, err := http.Post(url, "application/json;charset=utf-8", bytes.NewBuffer(bytesData)); err == nil {
			defer resp.Body.Close()
			if s, err := ioutil.ReadAll(resp.Body); err == nil {
				return s
			}
		}
	}
	return nil
}

// HttpPostFile 发送post上传文件请求（这个只用于图床上传）
func HttpPostFile(urls string, files *multipart.FileHeader, fileField string, data map[string]string, local bool) string {
	// 判断是否是上传本地文件
	var file multipart.File
	var err error
	if local {
		//打开文件句柄操作
		file, err = os.Open("temp.png")
	} else {
		//打开文件句柄操作
		file, err = files.Open()
	}
	if err != nil {
		fmt.Println("error opening file")
		return ""
	}
	defer file.Close()

	//创建一个模拟的form中的一个选项,这个form项现在是空的
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作, 设置文件的上传参数叫uploadfile, 文件名是filename,
	//相当于现在还没选择文件, form项里选择文件的选项
	var filename string
	if local || (files != nil && files.Filename == "blob") {
		filename = "temp.png"
	} else {
		filename = files.Filename
	}
	fileWriter, err := bodyWriter.CreateFormFile(fileField, filename)
	if err != nil {
		fmt.Println("error writing to buffer", err.Error())
		return ""
	}
	//iocopy 这里相当于选择了文件,将文件放到form中
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		//fmt.Println("copy error", err.Error())
		return ""
	}
	//获取上传文件的类型,multipart/form-data; boundary=...
	contentType := bodyWriter.FormDataContentType()
	//这个很关键,必须这样写关闭,不能使用defer关闭,不然会导致错误
	bodyWriter.Close()
	//这里就是上传的其他参数设置,可以使用 bodyWriter.WriteField(key, val) 方法
	//这种设置值得仿佛 和下面再从新创建一个的一样
	for key, val := range data {
		_ = bodyWriter.WriteField(key, val)
	}

	//发送post请求到服务端
	resp, err := http.Post(urls, contentType, bodyBuf)
	if err != nil {
		//fmt.Println("发送请求失败", err.Error())
		return ""
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(respBody)
}

// HttpGetHead 发送带头部信息的get的请求
func HttpGetHead(url string, head map[string]string) (string, bool) {
	client := &http.Client{}

	//新建一个请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", false
	}

	//添加头部信息
	for k, v := range head {
		request.Header.Add(k, v)
	}

	//执行请求
	resp, err := client.Do(request)
	if err != nil {
		return "", false
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", false
	}
	return string(s), true
}

// HttpReplaceHttp 替换链接中的http或http
func HttpReplaceHttp(data string) string {
	return FindMatch(`:(.*?)$`, data)
}

// HttpNewHead 初始化一个默认的head,默认带有谷歌浏览器的agent
func HttpNewHead() map[string]string {
	return map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36 Edg/87.0.664.75"}
}

// HttpDownloadFile 下载文件
func HttpDownloadFile(url string, filename string) error {
	out, err := os.Create(filename + ".tmp")
	if err != nil {
		return err
	}
	resp, err := http.Get(url)
	if err != nil {
		_ = out.Close()
		return err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	counter := &WriteCounter{}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		_ = out.Close()
		return err
	}
	_ = out.Close()
	if err = os.Rename(filename+".tmp", filename); err != nil {
		return err
	}
	return nil
}
