// Package module
// @Description colly爬虫包
// @Author 小游
// @Date 2021/04/08
package module

import (
	"github.com/dop251/goja"
	"github.com/gocolly/colly"
)

// RegisterColly 注册colly爬虫请求对象
func RegisterColly(vm *goja.Runtime) *goja.Object {
	blog := vm.NewObject()
	_ = blog.Set("init", CollyInit)
	return blog
}

// CollyInit 初始化colly爬虫对象
func CollyInit(head map[string]string) *colly.Collector {
	c := colly.NewCollector()
	//设置头部避免被被识别为爬虫
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36"
	//设置头部信息
	c.OnRequest(func(r *colly.Request) {
		for k, v := range head {
			r.Headers.Set(k, v)
		}
	})
	//设置多线程
	_ = c.Limit(&colly.LimitRule{
		DomainRegexp: "",
		DomainGlob:   "",
		Delay:        0,
		RandomDelay:  0,
		Parallelism:  50, //这里是限制50个线程同时执行（其实可以调高一点）
	})
	return c
}
