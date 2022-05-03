// Package main @Description
// @Author 小游
// @Date 2021/04/11
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kardianos/service"
	"log"
	"os"
	"xBlog/internal/app"
	"xBlog/internal/app/middleware"
	"xBlog/internal/cron"
	"xBlog/internal/manage"
	"xBlog/internal/plugin"
)

func SerStart() {
	// 关闭debug模式
	gin.SetMode(gin.ReleaseMode)
	//初始化gin框架
	app.GinEngine = gin.Default()
	//初始化中间件
	middleware.MiddleInit(app.GinEngine)
	// 注册插件
	plugin.ScanAllPlugins()
	// 启动定时任务
	cron.Init()
	// 启动项目
	app.Init()
}

// 当前应用程序的结构体
type program struct{}

// Start 服务启动
func (p *program) Start(s service.Service) error {
	log.Println("开始服务")
	go p.run()
	return nil
}

// Stop 停止程序
func (p *program) Stop(s service.Service) error {
	log.Println("停止服务")
	return nil
}

// 启动程序
func (p *program) run() {
	//API初始化
	SerStart()
}

func main() {
	// 是否需要可视化面板
	manage.Init()
	//服务的配置信息
	cfg := &service.Config{
		Name:        "xblog",
		DisplayName: "XBlog博客系统",
		Description: "一款基于前后端分离模式开发的二次元多端博客系统",
	}
	// 程序的接口
	prg := &program{}
	// 构建服务对象
	s, err := service.New(prg, cfg)
	if err != nil {
		log.Fatal(err)
	}
	// logger 用于记录系统日志
	logger, err := s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	// 如果有参数的话我们按照对应的参数执行
	if len(os.Args) == 2 {
		err = service.Control(s, os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// 如果没有那么我们就直接启动
		err = s.Run()
		if err != nil {
			logger.Error(err)
		}
	}
	if err != nil {
		logger.Error(err)
	}
}
