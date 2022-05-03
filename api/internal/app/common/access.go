// Package common
// @Description 权限相关
// @Author 小游
// @Date 2021/04/13
package common

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
	"time"
	"xBlog/internal/app/model"
	"xBlog/internal/db"
	"xBlog/pkg/database"
	"xBlog/tools"
)

// AccessGetTokenV2 普通用户权限验证V2版本
func AccessGetTokenV2(c *gin.Context) (model.UserLogin, error) {
	var login model.UserLogin
	// 为了兼容某些请求无法在头部发送，这里的权限信息可以在头部发送或者参数里面发送
	if tools.Str2Int(c.GetHeader("user_id")) == 0 {
		login.UserID = tools.Str2Int(c.PostForm("user_id"))
		login.Token = c.PostForm("token")
	} else {
		login.UserID = tools.Str2Int(c.GetHeader("user_id"))
		login.Token = c.GetHeader("token")
	}
	if login.UserID == 0 || login.Token == "" {
		return model.UserLogin{}, errors.New("参数为空")
	}
	if AccessTokenStatusV2(login.UserID, login.Token, false) {
		return login, nil
	} else {
		return model.UserLogin{}, errors.New("token无效")
	}
}

// AccessAdminTokenV2 管理员权限验证V2版本
func AccessAdminTokenV2(c *gin.Context) (model.UserLogin, error) {
	var login model.UserLogin
	//fmt.Println(c.GetHeader("user_id"))
	// todo 这里因为nginx转发会把下划线进行转义，所以这里后续需要改成中划线
	// 或者加上这句话   underscores_in_headers on;
	login.UserID = tools.Str2Int(c.GetHeader("user_id"))
	login.Token = c.GetHeader("token")
	//fmt.Println(c.Request.Header)
	//fmt.Println(login)
	if login.UserID == 0 || login.Token == "" {
		return model.UserLogin{}, errors.New("参数为空")
	}
	if AccessTokenStatusV2(login.UserID, login.Token, true) {
		return login, nil
	} else {
		return model.UserLogin{}, errors.New("token无效")
	}
}

// AccessAdminGetId 管理员获取用户id
func AccessAdminGetId(c *gin.Context) int {
	// 这里为了避免
	return tools.Str2Int(c.GetHeader("user_id"))
}

// AccessTokenStatusV2 判断token是否有效V2版本
func AccessTokenStatusV2(id int, token string, isAdmin bool) bool {
	//先查询该用户的登录数据
	user := new(db.User)
	collection := database.NewDb(db.CollUser).SetFilter(bson.M{"user_id": id})
	if collection.FindOne(user) == nil {
		// 判断是否为管理员验证
		if isAdmin && user.Identity != 1 {
			return false
		}
		//遍历数据
		for _, v := range user.LoginInfo {
			//判断token是否过期(token一般保存30天)
			if (time.Since(v.LoginTime).Hours() / 24) >= 30 {
				_ = collection.Pull(bson.M{"login_info": v}).UpdateOne()
			} else if v.Token == token {
				//token验证
				return true
			}
		}
	}
	return false
}

// AccessResetPasswordV2 重置密码发送邮件
func AccessResetPasswordV2(email string) bool {
	//随机生成token
	token := tools.GetRandomNum(20)
	sever := db.GetSiteOptionString(db.KeySiteApiServer)
	site := db.GetSiteOptionString(db.KeySiteName)
	//判断这个用户是否存在
	user := new(db.User)
	collection := database.NewDb(db.CollUser).SetFilter(bson.M{"email": email})
	if collection.FindOne(user) == nil {
		urls := sever + "/api/v3/user/password/email?id=" + strconv.Itoa(user.UserID) + "&token=" + token + "&option=reset&site=" + sever
		//更新token
		if collection.Set(bson.M{"token": token}).UpdateOne() != nil {
			return false
		}
		//发送邮件
		body := `
			<div class="content" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; max-width: 600px; display: block; margin: 0 auto; padding: 20px;">
				<table class="main" width="100%" cellpadding="0" cellspacing="0" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; border-radius: 3px; background-color: #fff; margin: 0; border: 1px solid #e9e9e9;" bgcolor="#fff"><tbody><tr style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;"><td class="alert alert-warning" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 16px; vertical-align: top; color: #fff; font-weight: 500; text-align: center; border-radius: 3px 3px 0 0; background-color: #009688; margin: 0; padding: 20px;" align="center" bgcolor="#FF9F00" valign="top">
							` + site + `密码重置
						</td>
					</tr><tr style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;"><td class="content-wrap" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; vertical-align: top; margin: 0; padding: 20px;" valign="top">
							<table width="100%" cellpadding="0" cellspacing="0" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;"><tbody><tr style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;"><td class="content-block" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; vertical-align: top; margin: 0; padding: 0 0 20px;" valign="top">
										亲爱的 <strong style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;"><span style="border-bottom: 1px dashed rgb(204, 204, 204); z-index: 1; position: static;" t="7" onclick="return false;" data="1589294503" isout="1">` + user.Nickname + `</span></strong> ：
									</td>
								</tr><tr style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;"><td class="content-block" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; vertical-align: top; margin: 0; padding: 0 0 20px;" valign="top">
										点击下面的按钮进行密码重置
									</td>
								</tr><tr style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;"><td class="content-block" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; vertical-align: top; margin: 0; padding: 0 0 20px;" valign="top">
										<a href="` + urls + `" class="btn-primary" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; color: #FFF; text-decoration: none; line-height: 2em; font-weight: bold; text-align: center; cursor: pointer; display: inline-block; border-radius: 5px; text-transform: capitalize; background-color: #009688; margin: 0; border-color: #009688; border-style: solid; border-width: 10px 20px;" rel="noopener" target="_blank">重置密码</a>
									</td>
								</tr><tr style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;"><td class="content-block" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; vertical-align: top; margin: 0; padding: 0 0 20px;" valign="top">
										
									</td>
								</tr></tbody></table></td>
					</tr></tbody></table><div class="footer" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; width: 100%; clear: both; color: #999; margin: 0; padding: 20px;">
					<table width="100%" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;"><tbody><tr style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;"><td class="aligncenter content-block" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 12px; vertical-align: top; color: #999; text-align: center; margin: 0; padding: 0 0 20px;" align="center" valign="top">此邮件由系统自动发送，请不要直接回复。</td>
						</tr></tbody></table></div></div>
 		`
		if tools.SendMail([]string{email}, "重置你的账户", body) == nil {
			return true
		}
	}
	return false
}
