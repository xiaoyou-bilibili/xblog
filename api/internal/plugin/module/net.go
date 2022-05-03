// Package module
// @Description 网络请求包
// @Author 小游
// @Date 2021/04/07
package module

import (
	"github.com/dop251/goja"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// RegisterNet 注册网络请求对象
func RegisterNet(vm *goja.Runtime) *goja.Object {
	blog := vm.NewObject()
	_ = blog.Set("get", NetGet)
	_ = blog.Set("post", NetPost)
	return blog
}

// NetGet 发送get请求
func NetGet(url string, head map[string]string, callback func(interface{}, interface{})) {
	client := &http.Client{}
	//新建一个请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		callback(err.Error(), nil)
		return
	}
	// 默认添加agent信息
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36 Edg/92.0.902.73")
	//添加头部信息
	for k, v := range head {
		request.Header.Set(k, v)
	}
	//执行请求
	resp, err := client.Do(request)
	if err != nil {
		callback(err.Error(), nil)
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	// 把请求转换为字节数据
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		callback(err.Error(), nil)
		return
	}
	// 如果没啥问题的话，我们就直接回调
	callback(nil, string(s))
}

// NetPost 发送POST请求
func NetPost(target string, head map[string]string, param map[string]string, callback func(interface{}, interface{})) {
	client := &http.Client{}
	// 封装请求
	DataUrlVal := url.Values{}
	for key, val := range param {
		DataUrlVal.Add(key, val)
	}
	req, err := http.NewRequest("POST", target, strings.NewReader(DataUrlVal.Encode()))
	if err != nil {
		callback(err.Error(), nil)
		return
	}
	// 添加头部
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36 Edg/92.0.902.73")
	for k, v := range head {
		req.Header.Set(k, v)
	}
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		callback(err.Error(), nil)
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	// 读取返回值
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		callback(err.Error(), nil)
		return
	}
	callback(nil, string(result))
}
