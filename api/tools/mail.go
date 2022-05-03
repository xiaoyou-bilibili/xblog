// Package tools @Description 邮件相关
// @Author 小游
// @Date 2021/04/10
package tools

import (
	"gopkg.in/gomail.v2"
	"strconv"
	"xBlog/internal/db"
)

func SendMail(mailTo []string, subject string, body string) error {
	user := db.GetSiteOptionString(db.KeySmtpUser)
	pass := db.GetSiteOptionString(db.KeySmtpPass)
	server := db.GetSiteOptionString(db.KeySmtpServer)
	portS := db.GetSiteOptionString(db.KeySmtpPort)
	name := db.GetSiteOptionString(db.KeySmtpName)
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string{
		"user": user,
		"pass": pass,
		"host": server,
		"port": portS,
	}
	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int
	m := gomail.NewMessage()
	m.SetHeader("From", name+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailTo...)                       //发送给多个用户
	m.SetHeader("Subject", subject)                    //设置邮件主题
	m.SetBody("text/html", body)                       //设置邮件正文
	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	err := d.DialAndSend(m)
	return err
}
